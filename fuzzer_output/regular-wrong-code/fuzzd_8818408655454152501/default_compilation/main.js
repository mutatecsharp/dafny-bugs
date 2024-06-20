// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, globalState) {
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("gmyjs"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(743)), function (_0_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }));
    };
    static fm1(p0, globalState) {
      return _dafny.Seq.of((_dafny.Set.fromElements(function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-470), new BigNumber(70))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(-470)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(70)))) {
            _coll0.add(_module.__default.safeModuloInt(_1_v0, new BigNumber(-241)));
          }
        }
        return _coll0;
      }(), _dafny.Set.fromElements(new BigNumber(616)), _dafny.Set.fromElements(new BigNumber(44), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(766), new BigNumber(916), new BigNumber((_dafny.MultiSet.fromElements(!(!(true)))).cardinality()), new BigNumber(948)))).cardinality()), new BigNumber(194)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-194)), function (_2_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length),new BigNumber((_dafny.Seq.of(false)).length))).length))))).IsSubsetOf(_dafny.Set.fromElements(function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-550),false)).Keys.Elements) {
          let _3_v1 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-550),false)).contains(_3_v1)) {
            _coll1.add(_module.__default.safeDivisionInt(_3_v1, new BigNumber((_dafny.Seq.UnicodeFromString("owyrbapmv")).length)));
          }
        }
        return _coll1;
      }())), (_dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D1.create_DC5()), _dafny.Seq.of(_module.D1.create_DC5()), _dafny.Seq.Create(_module.__default.abs(new BigNumber(643)), function (_4_i1) {
        return _module.D1.create_DC5();
      }), _dafny.Seq.of(_module.D1.create_DC5()))).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D1.create_DC5(), _module.D1.create_DC5()))), (_dafny.MultiSet.fromElements(false)).IsSubsetOf(_dafny.MultiSet.fromElements(false)), !(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("magrhjlxm")).length))).length)).isEqualTo(new BigNumber(-600)), !(!(false)) || (true));
    };
    static fm2(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(927)), function (_5_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_6_i1) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("o"));
    };
    static fm4(p0, globalState) {
      if (true) {
        return (_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)));
      } else {
        return _dafny.MultiSet.fromElements(!(!(!(false))), !(!(true)), !(true), false);
      }
    };
    static fm5(p0, globalState) {
      return new BigNumber(563);
    };
    static fm6(globalState) {
      return _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("x"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(292)), function (_7_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("dutfa"), _dafny.Seq.UnicodeFromString("qorshhqg"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), function (_8_i1) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("sqdlj")), _dafny.Seq.UnicodeFromString("dveovbd"));
    };
    static fm15(globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(745), new BigNumber((_dafny.Seq.UnicodeFromString("xmo")).length)));
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(244))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("virko")).length)))).Union(_dafny.Set.fromElements(new BigNumber(62)));
    };
    static fm18(globalState) {
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(191),(_module.D4.create_DC13(new BigNumber(-484), true, false, false)).dtor_cf19)).length), new BigNumber(890));
    };
    static fm19(p0, globalState) {
      return (((true) ? (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(166), new BigNumber(716), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dkqc")).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), function (_9_i0) {
        return new BigNumber(79);
      }))));
    };
    static fm20(p0, globalState) {
      let _source0 = _module.D13.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(!(false))), true, new BigNumber(963));
      if (_source0.is_DC36) {
        let _10___mcc_h0 = (_source0).cf54;
        let _11___mcc_h1 = (_source0).cf55;
        let _12___mcc_h2 = (_source0).cf56;
        let _13_cf56 = _12___mcc_h2;
        let _14_cf55 = _11___mcc_h1;
        let _15_cf54 = _10___mcc_h0;
        if ((_module.D0.create_DC3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_13_cf56,false)).length), _14_cf55)).dtor_cf5) {
          return _dafny.Set.fromElements(!(_14_cf55), true, true, _14_cf55);
        } else {
          return _dafny.Set.fromElements(_14_cf55, _14_cf55);
        }
      } else {
        let _16___mcc_h3 = (_source0).cf53;
        let _17_cf53 = _16___mcc_h3;
        return (_dafny.Set.fromElements(false, true)).Difference(_dafny.Set.fromElements(true, !(false)));
      }
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return new BigNumber(330);
    };
    static fm22(p0, p1, globalState) {
      return _module.D4.create_DC13(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(954)), function (_18_i0) {
  return new _dafny.CodePoint('g'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("cvhr"))).length), !(true), true, true);
    };
    static fm25(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D14.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(200)), function (_19_i0) {
  return new _dafny.CodePoint('c'.codePointAt(0));
})).length)));
      if (_source1.is_DC38) {
        let _20___mcc_h0 = (_source1).cf58;
        let _21_cf58 = _20___mcc_h0;
        return _module.D1.create_DC5();
      } else {
        let _22___mcc_h1 = (_source1).cf57;
        let _23_cf57 = _22___mcc_h1;
        return _module.D1.create_DC5();
      }
    };
    static fm27(p0, p1, globalState) {
      return _dafny.Set.fromElements((new BigNumber(938)).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), function (_24_i0) {
        return new BigNumber(499);
      }))).cardinality())), new BigNumber(542), ((!(true)) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, true)).length))) : (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true, false))).cardinality()))));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("xfdlh")).length)))).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber(402))) : (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(227)))));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(484)), _dafny.Seq.of(new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(993), new BigNumber(142))) {
          let _25_v0 = _compr_2;
          if (((new BigNumber(993)).isLessThanOrEqualTo(_25_v0)) && ((_25_v0).isLessThan(new BigNumber(142)))) {
            _coll2.add(_module.__default.safeModuloInt(_25_v0, new BigNumber((function () {
              let _coll3 = new _dafny.Map();
              for (const _compr_3 of (function () {
                let _coll4 = new _dafny.Set();
                for (const _compr_4 of _dafny.IntegerRange(new BigNumber(372), new BigNumber(544))) {
                  let _26_v2 = _compr_4;
                  if (((new BigNumber(372)).isLessThanOrEqualTo(_26_v2)) && ((_26_v2).isLessThan(new BigNumber(544)))) {
                    _coll4.add(_module.__default.safeModuloInt(_26_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)));
                  }
                }
                return _coll4;
              }()).Elements) {
                let _27_v1 = _compr_3;
                if ((function () {
                  let _coll5 = new _dafny.Set();
                  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(372), new BigNumber(544))) {
                    let _28_v2 = _compr_5;
                    if (((new BigNumber(372)).isLessThanOrEqualTo(_28_v2)) && ((_28_v2).isLessThan(new BigNumber(544)))) {
                      _coll5.add(_module.__default.safeModuloInt(_28_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)));
                    }
                  }
                  return _coll5;
                }()).contains(_27_v1)) {
                  _coll3.push([(_27_v1).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("crkbo"))).length)),false]);
                }
              }
              return _coll3;
            }()).length)));
          }
        }
        return _coll2;
      }()).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(407)), function (_29_i0) {
        return new BigNumber(293);
      })).length)), _dafny.Seq.of(new BigNumber(-13), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality())), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), function (_30_i1) {
        return _dafny.Seq.UnicodeFromString("bjkrgyb");
      })).length)))));
    };
    static fm30(p0, globalState) {
      if ((((_module.D18.create_DC44(new BigNumber((_dafny.Seq.of(true)).length), false, _dafny.Seq.UnicodeFromString("uyswmpk"), true, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-836)), (_module.D4.create_DC11(function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-525),!(true))).Keys.Elements) {
    let _31_v0 = _compr_6;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-525),!(true))).contains(_31_v0)) {
      _coll6.add((_31_v0).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('q'.codePointAt(0)))).length))));
    }
  }
  return _coll6;
}())).dtor_cf18, _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))))).dtor_cf66) ? (false) : (false))) {
        return _module.D1.create_DC6(new _dafny.CodePoint('v'.codePointAt(0)), true, _dafny.Seq.UnicodeFromString("bsxoyrdy"));
      } else {
        return _module.D1.create_DC6(new _dafny.CodePoint('m'.codePointAt(0)), false, _dafny.Seq.UnicodeFromString("nv"));
      }
    };
    static fm31(p0, globalState) {
      return new BigNumber(-258);
    };
    static fm32(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(-755))).Union(_dafny.MultiSet.fromElements(new BigNumber(537)));
    };
    static fm33(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('o'.codePointAt(0));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(95)), function (_32_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(-843))).length);
      }), _dafny.Seq.of((_module.D18.create_DC44(new BigNumber(404), true, _dafny.Seq.UnicodeFromString("sdtgfltt"), !(true), _dafny.Seq.of(function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-998), new BigNumber(189))) {
    let _33_v0 = _compr_7;
    if (((new BigNumber(-998)).isLessThanOrEqualTo(_33_v0)) && ((_33_v0).isLessThan(new BigNumber(189)))) {
      _coll7.add((_33_v0).plus(new BigNumber(-224)));
    }
  }
  return _coll7;
}(), function () {
  let _coll8 = new _dafny.Set();
  for (const _compr_8 of (_dafny.MultiSet.fromElements(new BigNumber(428), new BigNumber((_dafny.Seq.of(new BigNumber(343))).length))).Elements) {
    let _34_v1 = _compr_8;
    if ((_dafny.MultiSet.fromElements(new BigNumber(428), new BigNumber((_dafny.Seq.of(new BigNumber(343))).length))).contains(_34_v1)) {
      _coll8.add((_34_v1).minus(new BigNumber(-983)));
    }
  }
  return _coll8;
}()))).dtor_cf63, new BigNumber(373), new BigNumber(252), new BigNumber(-399), new BigNumber(-973)))) {
        return _module.D4.create_DC11(function () {
  let _coll9 = new _dafny.Set();
  for (const _compr_9 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kft")).length), new BigNumber((_dafny.Seq.UnicodeFromString("vdsnhhe")).length))).Elements) {
    let _35_v2 = _compr_9;
    if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kft")).length), new BigNumber((_dafny.Seq.UnicodeFromString("vdsnhhe")).length))).contains(_35_v2)) {
      _coll9.add(_module.__default.safeModuloInt(_35_v2, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(268)), function (_36_i1) {
        return new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality());
      })).length))));
    }
  }
  return _coll9;
}());
      } else {
        return _module.D4.create_DC11(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false, !(true))).cardinality())));
      }
    };
    static fm35(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(true, true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(true)))));
    };
    static fm36(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(_module.D20.create_DC50(_module.D5.create_DC17(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), true), true, _dafny.Seq.UnicodeFromString("jydkg")))).cardinality()), new BigNumber(365)),true);
    };
    static fm37(p0, globalState) {
      let _source2 = _module.D0.create_DC0(false);
      if (_source2.is_DC1) {
        let _37___mcc_h0 = (_source2).cf1;
        let _38___mcc_h1 = (_source2).cf2;
        let _39_cf2 = _38___mcc_h1;
        let _40_cf1 = _37___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_39_cf2,_39_cf2);
      } else if (_source2.is_DC2) {
        let _41___mcc_h2 = (_source2).cf3;
        let _42_cf3 = _41___mcc_h2;
        return (_dafny.Map.Empty.slice().updateUnsafe(_42_cf3,_42_cf3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_42_cf3,_42_cf3));
      } else if (_source2.is_DC3) {
        let _43___mcc_h3 = (_source2).cf4;
        let _44___mcc_h4 = (_source2).cf5;
        let _45_cf5 = _44___mcc_h4;
        let _46_cf4 = _43___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe(_45_cf5,_45_cf5)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_45_cf5,_45_cf5));
      } else {
        let _47___mcc_h5 = (_source2).cf0;
        let _48_cf0 = _47___mcc_h5;
        return _dafny.Map.Empty.slice().updateUnsafe(_48_cf0,_48_cf0);
      }
    };
    static fm39(p0, globalState) {
      if (!(false)) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      } else if (false) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }
    };
    static fm40(globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("ifrnalu")).length)))).Union(function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(518)))).Elements) {
          let _49_v0 = _compr_10;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(518)))).contains(_49_v0)) {
            _coll10.add(_49_v0);
          }
        }
        return _coll10;
      }());
    };
    static fm41(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(297), new BigNumber(125))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-699)));
    };
    static fm42(p0, globalState) {
      return _module.D0.create_DC3(_module.__default.safeDivisionInt(new BigNumber((function () {
  let _coll11 = new _dafny.Set();
  for (const _compr_11 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(false,true))).Elements) {
    let _50_v0 = _compr_11;
    if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(false,true))).contains(_50_v0)) {
      _coll11.add(_50_v0);
    }
  }
  return _coll11;
}()).length), new BigNumber(439)), (_module.D18.create_DC44(new BigNumber((_dafny.Seq.UnicodeFromString("vbu")).length), false, _dafny.Seq.UnicodeFromString("newukpfn"), true, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(508)), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-353)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false)).length))))).dtor_cf64);
    };
    static fm43(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true)));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      return (new BigNumber(483)).multipliedBy(new BigNumber(((_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements(false))).cardinality()));
    };
    static fm45(p0, p1, globalState) {
      return _module.D5.create_DC16(false);
    };
    static fm46(p0, p1, p2, globalState) {
      return _module.D10.create_DC29((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false, !(false))).cardinality()))).cardinality())).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), function (_51_i0) {
  return new _dafny.CodePoint('o'.codePointAt(0));
})).length)), !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("aif"), new _dafny.CodePoint('e'.codePointAt(0))), new BigNumber((_dafny.Seq.UnicodeFromString("djacvwmb")).length));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      if (((true) ? (true) : (true))) {
        if (false) {
          return function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(319),true),new BigNumber(-97))).length), new BigNumber(216))).Elements) {
              let _52_v0 = _compr_12;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(319),true),new BigNumber(-97))).length), new BigNumber(216)), _52_v0)) {
                _coll12.add((_52_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(476)), function (_53_i0) {
                  return new _dafny.CodePoint('i'.codePointAt(0));
                })).length)));
              }
            }
            return _coll12;
          }();
        } else {
          return _dafny.Set.fromElements(new BigNumber(151), new BigNumber(219));
        }
      } else {
        return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("mgaketdgp")).length))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber(-13))).length)));
      }
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-109)), function (_54_i0) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(132),_dafny.Set.fromElements(false, false))).length);
        })).length))).Elements) {
          let _55_v0 = _compr_13;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-109)), function (_54_i0) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(132),_dafny.Set.fromElements(false, false))).length);
          })).length))).contains(_55_v0)) {
            _coll13.push([(_55_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false, false)).length),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))).cardinality())))).length)),new BigNumber((_dafny.Seq.UnicodeFromString("wdm")).length)]);
          }
        }
        return _coll13;
      }()).length)).isLessThan(new BigNumber(-844)),_dafny.MultiSet.FromArray(((false) ? (_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))) : (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))))));
    };
    static fm49(p0, p1, p2, p3, globalState) {
      return _module.D5.create_DC15((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements((_module.D18.create_DC44(new BigNumber(478), true, _dafny.Seq.UnicodeFromString("rilg"), false, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(340))))).dtor_cf66))).length), new BigNumber(148))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(-832))));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      return (_module.D15.create_DC39(_dafny.MultiSet.fromElements(true))).dtor_cf59;
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(!(true))), _dafny.MultiSet.fromElements(true));
    };
    static fm52(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(208)),new BigNumber(869));
    };
    static fm53(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(!(!((!(false)) && (true))));
    };
    static fm54(p0, globalState) {
      return _module.D0.create_DC2(false);
    };
    static fm55(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).Difference((_module.D24.create_DC61(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)))))).dtor_cf94)).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))));
    };
    static fm56(p0, globalState) {
      if ((_dafny.MultiSet.fromElements(new BigNumber(-290))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qd")).length))).length), new BigNumber((_dafny.Seq.of(new BigNumber(586), new BigNumber(394))).length)))) {
        if (!(false)) {
          return _module.D0.create_DC1(!(true), false);
        } else {
          return _module.D0.create_DC1(!(false), !(!(true)));
        }
      } else {
        return _module.D0.create_DC1(false, true);
      }
    };
    static fm57(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((!(false)) ? (new BigNumber(874)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-241)), function (_56_i0) {
        return _module.D4.create_DC14(_module.D4.create_DC14(_module.D4.create_DC11(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kffsxfa")).length)))));
      })).length))),(_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(598), new BigNumber(843))));
    };
    static fm58(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(778)), function (_57_i0) {
        return _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), function (_58_i1) {
        return _dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)));
      })), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_59_i2) {
        return _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)));
      })));
    };
    static fm59(p0, p1, globalState) {
      return function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of ((function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_60_i0) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }))).Elements) {
            let _61_v1 = _compr_15;
            if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_60_i0) {
              return new _dafny.CodePoint('f'.codePointAt(0));
            }))).contains(_61_v1)) {
              _coll15.add(_61_v1);
            }
          }
          return _coll15;
        }()).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("n"), _dafny.Seq.UnicodeFromString("no"), _dafny.Seq.UnicodeFromString("cplx")))).Elements) {
          let _62_v0 = _compr_14;
          if (((function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_60_i0) {
              return new _dafny.CodePoint('f'.codePointAt(0));
            }))).Elements) {
              let _63_v1 = _compr_16;
              if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("i"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_60_i0) {
                return new _dafny.CodePoint('f'.codePointAt(0));
              }))).contains(_63_v1)) {
                _coll16.add(_63_v1);
              }
            }
            return _coll16;
          }()).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("n"), _dafny.Seq.UnicodeFromString("no"), _dafny.Seq.UnicodeFromString("cplx")))).contains(_62_v0)) {
            _coll14.push([_62_v0,new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(true)))).length)]);
          }
        }
        return _coll14;
      }();
    };
    static fm60(globalState) {
      let _source3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fij")).length)),!(!(true)));
      let _64___mcc_h0 = _source3;
      let _65_cf60 = _64___mcc_h0;
      return _module.D6.create_DC19(_module.D0.create_DC3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), true), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(509)), function (_66_i0) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length));
    };
    static fm61(p0, p1, p2, globalState) {
      return function () {
        let _coll17 = new _dafny.Set();
        for (const _compr_17 of (_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-25)), _dafny.Set.fromElements(new BigNumber(-605), new BigNumber((_dafny.Seq.UnicodeFromString("ahqnoab")).length), new BigNumber(839)), _dafny.Set.fromElements(new BigNumber(-105)), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(476))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-585))).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)))).length), new BigNumber(870), new BigNumber(-522))), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('c'.codePointAt(0)))).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dlcnkyi")).length)), _dafny.Set.fromElements(new BigNumber(-248))))).Elements) {
          let _67_v0 = _compr_17;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-25)), _dafny.Set.fromElements(new BigNumber(-605), new BigNumber((_dafny.Seq.UnicodeFromString("ahqnoab")).length), new BigNumber(839)), _dafny.Set.fromElements(new BigNumber(-105)), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(476))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-585))).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)))).length), new BigNumber(870), new BigNumber(-522))), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('c'.codePointAt(0)))).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("dlcnkyi")).length)), _dafny.Set.fromElements(new BigNumber(-248)))), _67_v0)) {
            _coll17.add(_67_v0);
          }
        }
        return _coll17;
      }();
    };
    static fm62(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D10.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D10.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true))))).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(!(true),_module.D10.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false)))) : (_dafny.Map.Empty.slice().updateUnsafe(true,_module.D10.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),true))))));
    };
    static fm63(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(490), new BigNumber(278)), _dafny.Set.fromElements(new BigNumber(-498), new BigNumber(752)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("j")).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(956)),true)).length)))), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), function (_68_i0) {
        return new BigNumber(-118);
      })).length)))).Union(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-822)), _dafny.Set.fromElements((_module.D0.create_DC3(new BigNumber((function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), function (_69_i1) {
    return new _dafny.CodePoint('r'.codePointAt(0));
  })).length), new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-862))).length))).Elements) {
    let _70_v0 = _compr_18;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(772)), function (_69_i1) {
      return new _dafny.CodePoint('r'.codePointAt(0));
    })).length), new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-862))).length)), _70_v0)) {
      _coll18.push([(_70_v0).plus(new BigNumber((_dafny.Set.fromElements(true, false)).length)),!(true)]);
    }
  }
  return _coll18;
}()).length), true)).dtor_cf4), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kyskdkmkh")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-82),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("faso")).length),true))).length))))).length),_dafny.Seq.UnicodeFromString("huu"))).length))));
    };
    static fm64(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)));
    };
    static fm65(p0, globalState) {
      if (false) {
        if (false) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(384),!(false));
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(414), new BigNumber(657))).cardinality()))).length),!(true));
        }
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),true);
      }
    };
    static fm66(globalState) {
      let _source4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_dafny.Seq.of(new BigNumber(526))).Elements) {
          let _71_v0 = _compr_19;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(526)), _71_v0)) {
            _coll19.push([(_71_v0).minus(new BigNumber(991)),new _dafny.CodePoint('f'.codePointAt(0))]);
          }
        }
        return _coll19;
      }()).length),false);
      let _72___mcc_h0 = _source4;
      let _73_cf60 = _72___mcc_h0;
      return _module.D19.create_DC48((_dafny.ZERO).minus(new BigNumber(-238)));
    };
    static fm67(p0, p1, p2, globalState) {
      let _source5 = _module.D24.create_DC63(new BigNumber(562), _dafny.Seq.UnicodeFromString("bikakvgqv"), new BigNumber(847), new BigNumber(174));
      if (_source5.is_DC62) {
        let _74___mcc_h0 = (_source5).cf95;
        let _75_cf95 = _74___mcc_h0;
        return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jqrsg"));
      } else if (_source5.is_DC63) {
        let _76___mcc_h1 = (_source5).cf96;
        let _77___mcc_h2 = (_source5).cf97;
        let _78___mcc_h3 = (_source5).cf98;
        let _79___mcc_h4 = (_source5).cf99;
        let _80_cf99 = _79___mcc_h4;
        let _81_cf98 = _78___mcc_h3;
        let _82_cf97 = _77___mcc_h2;
        let _83_cf96 = _76___mcc_h1;
        return function () {
          let _coll20 = new _dafny.Set();
          for (const _compr_20 of (_dafny.Seq.of(_82_cf97)).Elements) {
            let _84_v0 = _compr_20;
            if (_dafny.Seq.contains(_dafny.Seq.of(_82_cf97), _84_v0)) {
              _coll20.add(_84_v0);
            }
          }
          return _coll20;
        }();
      } else if (_source5.is_DC64) {
        let _85___mcc_h5 = (_source5).cf100;
        let _86_cf100 = _85___mcc_h5;
        return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("kgkdbmw"), _dafny.Seq.UnicodeFromString("coc"))).Difference(function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(825)), function (_87_i0) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }))).Elements) {
            let _88_v1 = _compr_21;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(825)), function (_87_i0) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })), _88_v1)) {
              _coll21.add(_88_v1);
            }
          }
          return _coll21;
        }());
      } else {
        let _89___mcc_h6 = (_source5).cf94;
        let _90_cf94 = _89___mcc_h6;
        return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), function (_91_i1) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("isxpgybe"));
      }
    };
    static fm68(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), function (_92_i0) {
        return _module.D1.create_DC4(_dafny.Seq.of(true, false, true));
      }), _dafny.Seq.of(_module.D1.create_DC4(_dafny.Seq.of(true, true)))), _dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC4(_dafny.Seq.of(true))), _dafny.Seq.of(_module.D1.create_DC4(_dafny.Seq.of(false, true)))));
    };
    static fm69(p0, globalState) {
      return (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-685)), function (_93_i0) {
        return _dafny.MultiSet.fromElements(false, !(!(true)));
      }));
    };
    static fm70(p0, p1, p2, globalState) {
      let _source6 = _module.D4.create_DC14(_module.D4.create_DC11(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(528))).length), new BigNumber(191))).cardinality()))));
      if (_source6.is_DC12) {
        return _module.D14.create_DC38(new BigNumber((_dafny.Set.fromElements(new BigNumber(60))).length));
      } else if (_source6.is_DC13) {
        let _94___mcc_h0 = (_source6).cf19;
        let _95___mcc_h1 = (_source6).cf20;
        let _96___mcc_h2 = (_source6).cf21;
        let _97___mcc_h3 = (_source6).cf22;
        let _98_cf22 = _97___mcc_h3;
        let _99_cf21 = _96___mcc_h2;
        let _100_cf20 = _95___mcc_h1;
        let _101_cf19 = _94___mcc_h0;
        return _module.D14.create_DC38(_101_cf19);
      } else if (_source6.is_DC11) {
        let _102___mcc_h4 = (_source6).cf18;
        let _103_cf18 = _102___mcc_h4;
        return _module.D14.create_DC38(new BigNumber((_dafny.Seq.UnicodeFromString("ggt")).length));
      } else {
        let _104___mcc_h5 = (_source6).cf23;
        let _105_cf23 = _104___mcc_h5;
        return _module.D14.create_DC38(new BigNumber(147));
      }
    };
    static fm71(p0, globalState) {
      if ((new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(!(false), false), _dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), !(!(false)), true, true, false)), _dafny.MultiSet.FromArray(_dafny.Seq.of(true)))).length)).isEqualTo((_module.D13.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(!(true),false), !(false), new BigNumber(749))).dtor_cf56)) {
        return _module.D1.create_DC4(_dafny.Seq.of(true, true, true, !(true), false));
      } else {
        return _module.D1.create_DC4(_dafny.Seq.of(false));
      }
    };
    static fm72(p0, p1, p2, p3, globalState) {
      let _source7 = _module.D15.create_DC40();
      if (_source7.is_DC40) {
        return _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(677))).length)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), function (_106_i0) {
          return new BigNumber((_dafny.Set.fromElements(true)).length);
        })).length)), new BigNumber(175)));
      } else {
        let _107___mcc_h0 = (_source7).cf59;
        let _108_cf59 = _107___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("j")).length))), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(606)), function (_109_i1) {
          return new BigNumber((_dafny.Seq.of(_module.D10.create_DC29(true, true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(276))).length),new _dafny.CodePoint('h'.codePointAt(0)))).length)))).length);
        }), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-534)), function (_110_i2) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })).length))));
      }
    };
    static fm73(p0, globalState) {
      return _module.D12.create_DC32(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(310), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(131))).length))).length),!(false))).length)))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("vowu")).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), new BigNumber((_dafny.Seq.of(false)).length)), (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)))));
    };
    static fm74(globalState) {
      return _dafny.Seq.of(_module.D22.create_DC56(new _dafny.CodePoint('k'.codePointAt(0)), true));
    };
    static m0(globalState) {
      let r0 = [];
      let r1 = [];
      let r2 = false;
      let _111_v0;
      _111_v0 = true;
      let _112_v1;
      _112_v1 = new BigNumber(740);
      let _113_v2;
      _113_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(_111_v0) || (_111_v0),((_111_v0) ? (_112_v1) : ((_dafny.ZERO).minus(_112_v1))));
      let _114_v3;
      let _nw0 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _114_v3 = _nw0;
      let _115_v4;
      _115_v4 = _dafny.Seq.UnicodeFromString("wwvaps");
      let _index0 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_114_v3).length));
      (_114_v3)[_index0] = _115_v4;
      let _116_v5;
      _116_v5 = new _dafny.CodePoint('e'.codePointAt(0));
      let _index1 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_114_v3).length));
      let _rhs0 = _111_v0;
      let _rhs1 = _115_v4;
      let _rhs2 = (_113_v2).Merge(_113_v2);
      let _rhs3 = _dafny.Seq.update(_dafny.Seq.Concat(_115_v4, _115_v4), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_115_v4).length)), new BigNumber((_dafny.Seq.Concat(_115_v4, _115_v4)).length)), _116_v5);
      let _rhs4 = _112_v1;
      let _lhs0 = globalState;
      let _lhs1 = globalState;
      let _lhs2 = _114_v3;
      let _lhs3 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_114_v3).length));
      _lhs0.f12 = _rhs0;
      _lhs1.f10 = _rhs1;
      _113_v2 = _rhs2;
      _lhs2[_lhs3] = _rhs3;
      _112_v1 = _rhs4;
      let _117_i0;
      _117_i0 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm0(_module.__default.fm5(_111_v0, globalState), globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_117_i0)) {
              break L0;
            }
            _117_i0 = (_117_i0).plus(_dafny.ONE);
            let _118_v6;
            let _init0 = ((_119_v1) => function (_120_i1) {
              return (_120_i1).minus(_119_v1);
            })(_112_v1);
            let _nw1 = Array((new BigNumber(20)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
              _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _118_v6 = _nw1;
            let _index2 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_118_v6).length));
            (_118_v6)[_index2] = (_112_v1).minus(_112_v1);
            let _index3 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_118_v6).length));
            (_118_v6)[_index3] = (_112_v1).multipliedBy((new BigNumber(-485)).plus(_112_v1));
            _115_v4 = (_114_v3)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_114_v3).length))];
            let _121_v7;
            let _nw2 = Array((new BigNumber(14)).toNumber());
            _121_v7 = _nw2;
            _121_v7 = _121_v7;
            let _122_v8;
            _122_v8 = _dafny.MultiSet.fromElements(new BigNumber(458));
            let _123_v9;
            _123_v9 = _dafny.Seq.of(_122_v8);
            let _124_v10;
            _124_v10 = _module.D23.create_DC58(_123_v9);
            (globalState).f12 = _dafny.areEqual(_dafny.Seq.Concat(_123_v9, _123_v9), (_124_v10).dtor_cf90);
          }
        }
      }
      let _125_v11;
      _125_v11 = _dafny.MultiSet.fromElements(false, _111_v0);
      let _126_v12;
      let _nw3 = new _module.C12();
      _nw3.__ctor(_112_v1, !(_111_v0), (((_125_v11).contains(!(_module.__default.fm0(new BigNumber(318), globalState)))) ? ((_125_v11).get(!(_module.__default.fm0(new BigNumber(318), globalState)))) : (_112_v1)), _114_v3);
      _126_v12 = _nw3;
      let _127_v13;
      _127_v13 = _module.D22.create_DC55((_126_v12).f26);
      let _nw4 = new _module.C6();
      _nw4.__ctor(_112_v1, (_127_v13).dtor_cf86);
      _126_v12 = _nw4;
      let _128_v14;
      let _nw5 = Array((new BigNumber(29)).toNumber()).fill(false);
      _128_v14 = _nw5;
      r0 = _128_v14;
      let _129_i2;
      _129_i2 = _dafny.ZERO;
      L1: {
        while (_111_v0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_129_i2)) {
              break L1;
            }
            _129_i2 = (_129_i2).plus(_dafny.ONE);
            _112_v1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), function (_130_i3) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            })).length));
            let _131_v15;
            _131_v15 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_112_v1), _module.__default.safeDivisionInt(_112_v1, new BigNumber(321)), _112_v1);
            _131_v15 = _131_v15;
            let _index4 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_128_v14).length));
            (_128_v14)[_index4] = _111_v0;
            let _index5 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_128_v14).length));
            (_128_v14)[_index5] = _111_v0;
            let _132_v16;
            let _nw6 = new _module.C12();
            _nw6.__ctor((_126_v12).f25, !(_111_v0), new BigNumber(432), (_126_v12).f26);
            _132_v16 = _nw6;
            _132_v16 = _132_v16;
          }
        }
      }
      let _133_v17;
      let _nw7 = new _module.C0();
      _nw7.__ctor();
      _133_v17 = _nw7;
      r0 = _128_v14;
      let _init1 = ((_134_v12) => function (_135_i4) {
        return (_135_i4).multipliedBy((_134_v12).f25);
      })(_126_v12);
      let _nw8 = Array((new BigNumber(13)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
        _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      r1 = _nw8;
      r2 = !_dafny.areEqual((_114_v3)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_114_v3).length))], _115_v4);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _136_v0;
      let _init2 = function (_137_i0) {
        return false;
      };
      let _nw9 = Array((new BigNumber(27)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw9.length); _i0_2++) {
        _nw9[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _136_v0 = _nw9;
      let _138_v1;
      _138_v1 = new BigNumber(454);
      let _139_v2;
      _139_v2 = _dafny.Seq.of(_138_v1);
      let _140_v3;
      _140_v3 = _dafny.Seq.UnicodeFromString("oxvkwml");
      let _141_v4;
      _141_v4 = _dafny.Seq.of(_140_v3);
      let _142_v5;
      _142_v5 = new _dafny.CodePoint('k'.codePointAt(0));
      let _143_v6;
      _143_v6 = _dafny.MultiSet.fromElements(_140_v3);
      let _144_v7;
      let _init3 = ((_145_v1) => function (_146_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(_145_v1,(_module.D0.create_DC2(true)).dtor_cf3);
      })(_138_v1);
      let _nw10 = Array((new BigNumber(23)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw10.length); _i0_3++) {
        _nw10[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _144_v7 = _nw10;
      let _147_v8;
      _147_v8 = false;
      let _148_v9;
      _148_v9 = _dafny.MultiSet.fromElements(_147_v8, _147_v8);
      let _149_v10;
      _149_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(249),_148_v9);
      let _150_v11;
      _150_v11 = _dafny.Set.fromElements(_147_v8, _147_v8);
      let _151_globalState;
      let _nw11 = new _module.GlobalState();
      _nw11.__ctor(new BigNumber(286), new BigNumber(656), true, new BigNumber(606), _136_v0, new BigNumber(150), _139_v2, new BigNumber(55), new BigNumber(667), _140_v3, _dafny.Seq.Concat((_141_v4)[_module.__default.safeIndex(_138_v1, new BigNumber((_141_v4).length))], _dafny.Seq.update(_140_v3, _module.__default.safeIndex(new BigNumber(478), new BigNumber((_140_v3).length)), _142_v5)), true, true, false, false, true, (_143_v6).Union(_143_v6), _144_v7, _149_v10, true, _150_v11, new BigNumber(517), new BigNumber(61), new _dafny.CodePoint('p'.codePointAt(0)));
      _151_globalState = _nw11;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_136_v0).length))) {
        let _152_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_152_i2)) && ((_152_i2).isLessThan(new BigNumber((_136_v0).length))))) {
          (_136_v0)[(_152_i2)] = _147_v8;
        }
      }
      let _153_v12;
      _153_v12 = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm0(_138_v1, _151_globalState));
      let _154_v13;
      _154_v13 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_153_v12).length));
      let _155_v14;
      _155_v14 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,(_154_v13).Merge(_dafny.Map.Empty.slice().updateUnsafe(_147_v8,new BigNumber((_140_v3).length))));
      let _156_v16;
      _156_v16 = _module.D0.create_DC3(_138_v1, _147_v8);
      let _157_v17;
      _157_v17 = _dafny.MultiSet.fromElements((_156_v16).dtor_cf4, _138_v1);
      _155_v14 = function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of (_157_v17).Elements) {
          let _158_v15 = _compr_22;
          if ((_157_v17).contains(_158_v15)) {
            _coll22.push([_module.__default.safeModuloInt(_158_v15, new BigNumber((_dafny.Seq.of(_147_v8, _147_v8, _147_v8, _147_v8)).length)),_154_v13]);
          }
        }
        return _coll22;
      }();
      let _159_v18;
      _159_v18 = _dafny.Seq.of(_147_v8);
      let _160_v19;
      _160_v19 = _dafny.Seq.of(_159_v18);
      let _161_i3;
      _161_i3 = _dafny.ZERO;
      L2: {
        while (!((_dafny.MultiSet.FromArray(_160_v19)).IsSubsetOf(_dafny.MultiSet.fromElements(_159_v18, _module.__default.fm1((_159_v18)[_module.__default.safeIndex(new BigNumber(-684), new BigNumber((_159_v18).length))], _151_globalState))))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_161_i3)) {
              break L2;
            }
            _161_i3 = (_161_i3).plus(_dafny.ONE);
            (_151_globalState).f23 = _142_v5;
            let _index6 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_136_v0).length));
            (_136_v0)[_index6] = _147_v8;
            let _index7 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_136_v0).length));
            (_136_v0)[_index7] = !_dafny.areEqual(_159_v18, _159_v18);
            let _162_v20;
            _162_v20 = _dafny.Set.fromElements(_138_v1);
            _138_v1 = ((_139_v2)[_module.__default.safeIndex(_138_v1, new BigNumber((_139_v2).length))]).plus((new BigNumber((_162_v20).length)).multipliedBy(_138_v1));
            _147_v8 = (_136_v0)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_136_v0).length))];
          }
        }
      }
      if ((new BigNumber((_dafny.Seq.Concat(_140_v3, _module.__default.fm2(_138_v1, _138_v1, _151_globalState))).length)).isEqualTo(new BigNumber((_139_v2).length))) {
        let _163_v21;
        _163_v21 = _module.D0.create_DC1(_147_v8, _147_v8);
        let _164_v22;
        _164_v22 = _dafny.Set.fromElements(_163_v21);
        (_151_globalState).f12 = ((_164_v22).Intersect(_164_v22)).IsSubsetOf(_164_v22);
        _138_v1 = _module.__default.safeDivisionInt(_138_v1, new BigNumber(347));
        let _index8 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_136_v0).length));
        (_136_v0)[_index8] = ((_147_v8) ? (_147_v8) : (_147_v8));
        let _165_v23;
        _165_v23 = _module.D1.create_DC4(_159_v18);
        let _index9 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_136_v0).length));
        (_136_v0)[_index9] = _dafny.Seq.contains(_140_v3, (_140_v3)[_module.__default.safeIndex(new BigNumber(((_165_v23).dtor_cf6).length), new BigNumber((_140_v3).length))]);
        _138_v1 = ((_147_v8) ? ((((_157_v17).contains(_138_v1)) ? ((_157_v17).get(_138_v1)) : (_138_v1))) : (_138_v1));
        let _166_v24;
        _166_v24 = _module.D1.create_DC6(_142_v5, (_136_v0)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((_136_v0).length))], _140_v3);
        let _167_v25;
        _167_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_138_v1, _151_globalState),_dafny.Seq.UnicodeFromString("qujrkn"));
        let _168_v26;
        let _nw12 = Array((new BigNumber(20)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), ((_169_v5) => function (_170_i4) {
          return _169_v5;
        })(_142_v5)), _140_v3);
        _nw12[(_dafny.ONE).toNumber()] = _140_v3;
        _nw12[(new BigNumber(2)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(3)).toNumber()] = (_166_v24).dtor_cf9;
        _nw12[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_140_v3, _140_v3);
        _nw12[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_140_v3, _dafny.Seq.UnicodeFromString("dakda"));
        _nw12[(new BigNumber(6)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(7)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(8)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(9)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(10)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(11)).toNumber()] = ((_147_v8) ? (_module.__default.fm2(_138_v1, _138_v1, _151_globalState)) : (_module.__default.fm2(_138_v1, new BigNumber(159), _151_globalState)));
        _nw12[(new BigNumber(12)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(13)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(14)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat((((_167_v25).contains((_136_v0)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((_136_v0).length))])) ? ((_167_v25).get((_136_v0)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((_136_v0).length))])) : (_140_v3)), _140_v3);
        _nw12[(new BigNumber(16)).toNumber()] = _140_v3;
        _nw12[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-947)), ((_171_v5) => function (_172_i5) {
          return _171_v5;
        })(_142_v5));
        _nw12[(new BigNumber(18)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-516)), ((_173_v5) => function (_174_i6) {
          return _173_v5;
        })(_142_v5));
        _nw12[(new BigNumber(19)).toNumber()] = _140_v3;
        _168_v26 = _nw12;
        let _index10 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_168_v26).length));
        (_168_v26)[_index10] = _140_v3;
        let _index11 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_168_v26).length));
        (_168_v26)[_index11] = _dafny.Seq.Concat(_140_v3, _dafny.Seq.UnicodeFromString("otnycwpnk"));
      } else {
        let _175_v27;
        let _nw13 = Array((new BigNumber(3)).toNumber()).fill([]);
        _175_v27 = _nw13;
        _175_v27 = ((_module.__default.fm0(new BigNumber((_159_v18).length), _151_globalState)) ? (_175_v27) : (_175_v27));
        let _176_v28;
        let _nw14 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _176_v28 = _nw14;
        let _177_v29;
        let _nw15 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _177_v29 = _nw15;
        let _178_v30;
        _178_v30 = _dafny.Map.Empty.slice().updateUnsafe(_176_v28,((false) ? (_177_v29) : (_177_v29)));
        _178_v30 = (_178_v30).update(_176_v28, _177_v29);
        let _179_v31;
        let _180_v32;
        let _181_v33;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = _module.__default.m0(_151_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _179_v31 = _out0;
        _180_v32 = _out1;
        _181_v33 = _out2;
        if ((_147_v8) === (false)) {
          _138_v1 = ((!(!(_138_v1).isEqualTo(_138_v1))) ? (new BigNumber((_139_v2).length)) : (_138_v1));
          let _182_v34;
          let _nw16 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _182_v34 = _nw16;
          let _183_v35;
          let _nw17 = new _module.C4();
          _nw17.__ctor(_182_v34, _138_v1, _182_v34);
          _183_v35 = _nw17;
          let _184_v36;
          _184_v36 = _dafny.Map.Empty.slice().updateUnsafe(_142_v5,_147_v8);
          _184_v36 = (_184_v36).update(_142_v5, (_181_v33) === (_147_v8));
          let _185_v37;
          _185_v37 = _module.D19.create_DC47(_136_v0, _147_v8);
          let _186_v38;
          _186_v38 = _module.D3.create_DC10(_181_v33, _136_v0, _181_v33, _138_v1, _138_v1);
          let _rhs5 = false;
          let _rhs6 = _module.__default.fm0((_186_v38).dtor_cf17, _151_globalState);
          let _rhs7 = _185_v37;
          let _lhs4 = _151_globalState;
          _181_v33 = _rhs5;
          _lhs4.f12 = _rhs6;
          _185_v37 = _rhs7;
          let _index12 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_136_v0).length));
          (_136_v0)[_index12] = _181_v33;
          let _187_v39;
          _187_v39 = _dafny.Map.Empty.slice().updateUnsafe(_142_v5,(_dafny.ZERO).minus((_139_v2)[_module.__default.safeIndex(_138_v1, new BigNumber((_139_v2).length))]));
          let _index13 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_180_v32).length));
          (_180_v32)[_index13] = (_dafny.ZERO).minus(new BigNumber((_187_v39).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_136_v0).length));
          let _index15 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_180_v32).length));
          let _rhs8 = !(_147_v8) || ((_138_v1).isEqualTo(_138_v1));
          let _rhs9 = _177_v29;
          let _rhs10 = _module.__default.safeModuloInt(new BigNumber(562), _138_v1);
          let _lhs5 = _136_v0;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_136_v0).length));
          let _lhs7 = _180_v32;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_180_v32).length));
          _lhs5[_lhs6] = _rhs8;
          _177_v29 = _rhs9;
          _lhs7[_lhs8] = _rhs10;
        } else {
          let _index16 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          (_136_v0)[_index16] = true;
          let _index17 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          let _rhs11 = !(_138_v1).isEqualTo(_module.__default.safeDivisionInt(_138_v1, new BigNumber((_140_v3).length)));
          let _rhs12 = _module.__default.fm31(_140_v3, _151_globalState);
          let _lhs9 = _136_v0;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          _lhs9[_lhs10] = _rhs11;
          _138_v1 = _rhs12;
          _180_v32 = _180_v32;
          _138_v1 = (_138_v1).minus((_138_v1).minus((_dafny.ZERO).minus(new BigNumber((_140_v3).length))));
          let _188_v40;
          _188_v40 = _dafny.MultiSet.fromElements(_142_v5);
          let _index18 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          let _rhs13 = false;
          let _rhs14 = _147_v8;
          let _rhs15 = (_dafny.ZERO).minus(new BigNumber((_159_v18).length));
          let _rhs16 = _module.__default.fm39(_dafny.Seq.IsProperPrefixOf(_140_v3, _140_v3), _151_globalState);
          let _rhs17 = _module.__default.safeModuloInt(_138_v1, (((_188_v40).contains(_142_v5)) ? ((_188_v40).get(_142_v5)) : (_138_v1)));
          let _lhs11 = _136_v0;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          let _lhs13 = _136_v0;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          let _lhs15 = _151_globalState;
          _lhs11[_lhs12] = _rhs13;
          _lhs13[_lhs14] = _rhs14;
          _138_v1 = _rhs15;
          _lhs15.f23 = _rhs16;
          _138_v1 = _rhs17;
          let _index20 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_136_v0).length));
          (_136_v0)[_index20] = _147_v8;
        }
        let _189_v41;
        _189_v41 = _module.D5.create_DC15(_157_v17);
        let _190_v42;
        _190_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(796),_189_v41);
        _190_v42 = (_190_v42).update((_139_v2)[_module.__default.safeIndex(new BigNumber((_150_v11).length), new BigNumber((_139_v2).length))], _189_v41);
      }
      let _191_v43;
      _191_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_139_v2).length),_147_v8);
      let _192_i7;
      _192_i7 = _dafny.ZERO;
      L3: {
        while ((((_191_v43).contains(_138_v1)) ? ((_191_v43).get(_138_v1)) : (_module.__default.fm0(new BigNumber(-719), _151_globalState)))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_192_i7)) {
              break L3;
            }
            _192_i7 = (_192_i7).plus(_dafny.ONE);
            let _193_v44;
            let _nw18 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
            _193_v44 = _nw18;
            let _194_v45;
            _194_v45 = _dafny.Map.Empty.slice().updateUnsafe(_140_v3,_138_v1);
            let _index21 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_193_v44).length));
            (_193_v44)[_index21] = ((false) ? (_194_v45) : (_194_v45));
            let _index22 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_193_v44).length));
            (_193_v44)[_index22] = (_194_v45).update(_dafny.Seq.Concat(_140_v3, _140_v3), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-743)), ((_195_v1) => function (_196_i8) {
              return _195_v1;
            })(_138_v1)), _139_v2)).length));
            _147_v8 = (_150_v11).IsDisjointFrom((_module.__default.fm20(_147_v8, _151_globalState)).Difference(_150_v11));
            let _197_v46;
            _197_v46 = _module.D19.create_DC47(_136_v0, (_159_v18)[_module.__default.safeIndex(_138_v1, new BigNumber((_159_v18).length))]);
            let _pat_let_tv0 = _136_v0;
            let _source8 = function (_pat_let0_0) {
              return function (_198_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_199_dt__update_hcf69_h0) {
                    return _module.D19.create_DC47(_199_dt__update_hcf69_h0, (_198_dt__update__tmp_h0).dtor_cf70);
                  }(_pat_let1_0);
                }(_pat_let_tv0);
              }(_pat_let0_0);
            }(_197_v46);
            if (_source8.is_DC47) {
              let _200___mcc_h0 = (_source8).cf69;
              let _201___mcc_h1 = (_source8).cf70;
              let _202_cf70 = _201___mcc_h1;
              let _203_cf69 = _200___mcc_h0;
              let _204_v47;
              _204_v47 = _module.D0.create_DC1(_202_cf70, _202_cf70);
              let _205_v48;
              _205_v48 = _dafny.Map.Empty.slice().updateUnsafe(_204_v47,(_140_v3)[_module.__default.safeIndex(_module.__default.fm31(_dafny.Seq.update(_module.__default.fm2(_138_v1, _138_v1, _151_globalState), _module.__default.safeIndex(new BigNumber((_159_v18).length), new BigNumber((_module.__default.fm2(_138_v1, _138_v1, _151_globalState)).length)), _142_v5), _151_globalState), new BigNumber((_140_v3).length))]);
              _205_v48 = (_205_v48).update(_204_v47, _142_v5);
              _138_v1 = _138_v1;
              let _206_v49;
              _206_v49 = _module.D15.create_DC40();
              _206_v49 = _206_v49;
              (_151_globalState).f10 = _140_v3;
            } else if (_source8.is_DC48) {
              let _207___mcc_h2 = (_source8).cf71;
              let _208_cf71 = _207___mcc_h2;
              let _209_v50;
              _209_v50 = _module.D0.create_DC0(!(_147_v8));
              let _210_v51;
              _210_v51 = _dafny.Map.Empty.slice().updateUnsafe(_209_v50,_147_v8);
              let _211_v52;
              let _nw19 = Array((new BigNumber(9)).toNumber());
              _nw19[(_dafny.ZERO).toNumber()] = (_module.__default.fm21(_210_v51, _138_v1, _147_v8, _140_v3, _151_globalState)).minus(_208_cf71);
              _nw19[(_dafny.ONE).toNumber()] = new BigNumber(488);
              _nw19[(new BigNumber(2)).toNumber()] = _208_cf71;
              _nw19[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("lgha"), _module.__default.safeIndex(_208_cf71, new BigNumber((_dafny.Seq.UnicodeFromString("lgha")).length)), new _dafny.CodePoint('c'.codePointAt(0)))).length);
              _nw19[(new BigNumber(4)).toNumber()] = _208_cf71;
              _nw19[(new BigNumber(5)).toNumber()] = new BigNumber(666);
              _nw19[(new BigNumber(6)).toNumber()] = new BigNumber(273);
              _nw19[(new BigNumber(7)).toNumber()] = _138_v1;
              _nw19[(new BigNumber(8)).toNumber()] = (_138_v1).minus(_208_cf71);
              _211_v52 = _nw19;
              _211_v52 = _211_v52;
              let _212_v53;
              _212_v53 = _191_v43;
              let _213_v54;
              let _nw20 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _213_v54 = _nw20;
              let _214_v55;
              _214_v55 = _module.D22.create_DC55(_213_v54);
              let _215_v56;
              let _nw21 = new _module.C10();
              _nw21.__ctor(_142_v5, new BigNumber(114), (_214_v55).dtor_cf86);
              _215_v56 = _nw21;
              let _rhs18 = (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_138_v1),_147_v8)).Merge(_191_v43);
              let _rhs19 = _215_v56;
              _212_v53 = _rhs18;
              _215_v56 = _rhs19;
              let _216_v57;
              let _217_v58;
              let _218_v59;
              let _out3;
              let _out4;
              let _out5;
              let _outcollector1 = _module.__default.m0(_151_globalState);
              _out3 = _outcollector1[0];
              _out4 = _outcollector1[1];
              _out5 = _outcollector1[2];
              _216_v57 = _out3;
              _217_v58 = _out4;
              _218_v59 = _out5;
              _208_cf71 = new BigNumber(((((_208_cf71).isEqualTo(_208_cf71)) ? (_159_v18) : (_159_v18))).length);
            } else {
              let _219___mcc_h3 = (_source8).cf68;
              let _220_cf68 = _219___mcc_h3;
              let _index23 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_136_v0).length));
              (_136_v0)[_index23] = !(_147_v8) || (_147_v8);
              let _index24 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_136_v0).length));
              (_136_v0)[_index24] = _147_v8;
              let _221_v60;
              let _nw22 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
              _221_v60 = _nw22;
              let _rhs20 = _221_v60;
              let _rhs21 = _dafny.Seq.UnicodeFromString("tkcekenj");
              let _rhs22 = _module.__default.safeDivisionInt((_138_v1).minus((_dafny.ZERO).minus(_module.__default.fm44((_136_v0)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_136_v0).length))], _147_v8, (_136_v0)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_136_v0).length))], _138_v1, _151_globalState))), _138_v1);
              let _lhs16 = _151_globalState;
              _221_v60 = _rhs20;
              _lhs16.f10 = _rhs21;
              _138_v1 = _rhs22;
              let _222_v61;
              let _nw23 = new _module.C5();
              _nw23.__ctor((_136_v0)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_136_v0).length))]);
              _222_v61 = _nw23;
              _222_v61 = ((false) ? (_222_v61) : (_222_v61));
              _138_v1 = new BigNumber(17);
            }
            _138_v1 = (_138_v1).minus(_138_v1);
          }
        }
      }
      if (true) {
        let _223_v62;
        let _init4 = ((_224_v1) => function (_225_i9) {
          return _module.__default.safeDivisionInt(_225_i9, _224_v1);
        })(_138_v1);
        let _nw24 = Array((new BigNumber(15)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw24.length); _i0_4++) {
          _nw24[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _223_v62 = _nw24;
        let _226_v63;
        let _nw25 = Array((new BigNumber(12)).toNumber());
        _nw25[(_dafny.ZERO).toNumber()] = _140_v3;
        _nw25[(_dafny.ONE).toNumber()] = _140_v3;
        _nw25[(new BigNumber(2)).toNumber()] = _140_v3;
        _nw25[(new BigNumber(3)).toNumber()] = _140_v3;
        _nw25[(new BigNumber(4)).toNumber()] = _module.__default.fm2(new BigNumber((_148_v9).cardinality()), _138_v1, _151_globalState);
        _nw25[(new BigNumber(5)).toNumber()] = _140_v3;
        _nw25[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("ik");
        _nw25[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("kgljbt");
        _nw25[(new BigNumber(8)).toNumber()] = _140_v3;
        _nw25[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("aomrfrrje");
        _nw25[(new BigNumber(10)).toNumber()] = _140_v3;
        _nw25[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("oroy");
        _226_v63 = _nw25;
        let _227_v64;
        let _nw26 = new _module.C12();
        _nw26.__ctor(_138_v1, _147_v8, _138_v1, _226_v63);
        _227_v64 = _nw26;
        let _228_v65;
        let _nw27 = new _module.C1();
        _nw27.__ctor(new BigNumber(-356), _226_v63);
        _228_v65 = _nw27;
        let _229_v66;
        _229_v66 = _dafny.Map.Empty.slice().updateUnsafe(_227_v64,_228_v65);
        let _index25 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length));
        (_223_v62)[_index25] = new BigNumber((_229_v66).length);
        let _index26 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length));
        (_223_v62)[_index26] = ((_227_v64).f27).multipliedBy((_227_v64).f27);
        _138_v1 = (((_148_v9).contains(_147_v8)) ? ((_148_v9).get(_147_v8)) : (_138_v1));
        let _index27 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length));
        (_223_v62)[_index27] = new BigNumber(308);
        let _230_v67;
        _230_v67 = _dafny.Set.fromElements((_223_v62)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length))]);
        if (((_230_v67).Union(_230_v67)).IsProperSubsetOf((_230_v67).Union(_230_v67))) {
          let _index28 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_136_v0).length));
          (_136_v0)[_index28] = _227_v64.f28;
          let _index29 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_136_v0).length));
          let _index30 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length));
          let _rhs23 = _dafny.Seq.update(_139_v2, _module.__default.safeIndex(_module.__default.safeModuloInt((_227_v64).f27, (_223_v62)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length))]), new BigNumber((_139_v2).length)), (_223_v62)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length))]);
          let _rhs24 = false;
          let _rhs25 = _227_v64.f28;
          let _rhs26 = _147_v8;
          let _rhs27 = new BigNumber(589);
          let _lhs17 = _227_v64;
          let _lhs18 = _136_v0;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_136_v0).length));
          let _lhs20 = _151_globalState;
          let _lhs21 = _223_v62;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length));
          _139_v2 = _rhs23;
          _lhs17.f28 = _rhs24;
          _lhs18[_lhs19] = _rhs25;
          _lhs20.f12 = _rhs26;
          _lhs21[_lhs22] = _rhs27;
          let _231_v68;
          let _nw28 = Array((new BigNumber(5)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = false;
          _nw28[(_dafny.ONE).toNumber()] = _227_v64.f28;
          _nw28[(new BigNumber(2)).toNumber()] = _227_v64.f28;
          _nw28[(new BigNumber(3)).toNumber()] = _147_v8;
          _nw28[(new BigNumber(4)).toNumber()] = _227_v64.f28;
          _231_v68 = _nw28;
          let _232_v69;
          _232_v69 = _module.D19.create_DC47(_231_v68, _227_v64.f28);
          let _233_v70;
          let _nw29 = Array((new BigNumber(11)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = _231_v68;
          _nw29[(_dafny.ONE).toNumber()] = _231_v68;
          _nw29[(new BigNumber(2)).toNumber()] = (_232_v69).dtor_cf69;
          _nw29[(new BigNumber(3)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(4)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(5)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(6)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(7)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(8)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(9)).toNumber()] = _231_v68;
          _nw29[(new BigNumber(10)).toNumber()] = _231_v68;
          _233_v70 = _nw29;
          let _234_v71;
          let _235_v72;
          let _out6;
          let _out7;
          let _outcollector2 = (_228_v65).m2(_233_v70, _147_v8, _227_v64.f28, _223_v62, _151_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _234_v71 = _out6;
          _235_v72 = _out7;
          let _236_v73;
          _236_v73 = _module.D1.create_DC5();
          let _237_v74;
          _237_v74 = _module.D8.create_DC24(_236_v73);
          let _238_v75;
          _238_v75 = _dafny.Map.Empty.slice().updateUnsafe((_136_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_136_v0).length))],_237_v74);
          _238_v75 = _238_v75;
          _234_v71 = (_139_v2)[_module.__default.safeIndex((new BigNumber((_157_v17).cardinality())).plus(_138_v1), new BigNumber((_139_v2).length))];
          _138_v1 = (_227_v64).f27;
        } else {
          let _239_v76;
          _239_v76 = _dafny.Seq.of((_223_v62)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length))], new BigNumber((_140_v3).length));
          let _240_v77;
          _240_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(701),(_223_v62)[_module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length))]);
          let _241_v78;
          let _242_v79;
          let _243_v80;
          let _out8;
          let _out9;
          let _out10;
          let _outcollector3 = (_227_v64).m5(_223_v62, _227_v64.f28, _dafny.Seq.update(_139_v2, _module.__default.safeIndex(new BigNumber((_240_v77).length), new BigNumber((_139_v2).length)), (_227_v64).f27), _147_v8, _151_globalState);
          _out8 = _outcollector3[0];
          _out9 = _outcollector3[1];
          _out10 = _outcollector3[2];
          _241_v78 = _out8;
          _242_v79 = _out9;
          _243_v80 = _out10;
          let _index31 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_136_v0).length));
          (_136_v0)[_index31] = _227_v64.f28;
          let _244_v81;
          _244_v81 = _dafny.Set.fromElements(_240_v77, (_240_v77).update(new BigNumber(109), _138_v1));
          let _245_v82;
          _245_v82 = _module.D21.create_DC53(_244_v81, true, _223_v62);
          let _index32 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_136_v0).length));
          (_136_v0)[_index32] = (_245_v82).dtor_cf83;
          let _246_v83;
          _246_v83 = _module.D1.create_DC4(_module.__default.fm1(_227_v64.f28, _151_globalState));
          (_227_v64).m4(_246_v83, _151_globalState);
          let _247_v85;
          _247_v85 = _dafny.Seq.of(function () {
            let _coll23 = new _dafny.Set();
            for (const _compr_23 of (_139_v2).Elements) {
              let _248_v84 = _compr_23;
              if (_dafny.Seq.contains(_139_v2, _248_v84)) {
                _coll23.add((_248_v84).minus(new BigNumber((_dafny.Seq.of(true)).length)));
              }
            }
            return _coll23;
          }(), _230_v67);
          let _249_v86;
          _249_v86 = _module.D18.create_DC44(new BigNumber((_240_v77).length), _242_v79, _140_v3, (_136_v0)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_136_v0).length))], _247_v85);
          let _index33 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_223_v62).length));
          (_223_v62)[_index33] = (_249_v86).dtor_cf63;
          let _250_v87;
          let _251_v88;
          let _252_v89;
          let _out11;
          let _out12;
          let _out13;
          let _outcollector4 = (_227_v64).m5(_223_v62, (_138_v1).isLessThan((_227_v64).f27), _139_v2, _147_v8, _151_globalState);
          _out11 = _outcollector4[0];
          _out12 = _outcollector4[1];
          _out13 = _outcollector4[2];
          _250_v87 = _out11;
          _251_v88 = _out12;
          _252_v89 = _out13;
        }
        let _253_v90;
        let _init5 = ((_254_v64) => function (_255_i10) {
          return _dafny.Seq.of(!(_254_v64.f28));
        })(_227_v64);
        let _nw30 = Array((new BigNumber(23)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw30.length); _i0_5++) {
          _nw30[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _253_v90 = _nw30;
        _253_v90 = _253_v90;
      } else {
        let _256_v91;
        let _257_v92;
        let _258_v93;
        let _out14;
        let _out15;
        let _out16;
        let _outcollector5 = _module.__default.m0(_151_globalState);
        _out14 = _outcollector5[0];
        _out15 = _outcollector5[1];
        _out16 = _outcollector5[2];
        _256_v91 = _out14;
        _257_v92 = _out15;
        _258_v93 = _out16;
        let _259_v94;
        _259_v94 = _dafny.Seq.of(_150_v11, _150_v11);
        if (((!(new BigNumber(((_259_v94)[_module.__default.safeIndex(_138_v1, new BigNumber((_259_v94).length))]).length)).isEqualTo(_138_v1)) ? (_258_v93) : (_258_v93))) {
          let _index34 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_256_v91).length));
          (_256_v91)[_index34] = _module.__default.fm0(_138_v1, _151_globalState);
          let _index35 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_256_v91).length));
          (_256_v91)[_index35] = ((_258_v93) ? (_147_v8) : (_147_v8));
          let _260_v95;
          let _nw31 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _260_v95 = _nw31;
          let _261_v96;
          let _nw32 = new _module.C6();
          _nw32.__ctor(_138_v1, _260_v95);
          _261_v96 = _nw32;
          let _rhs28 = _136_v0;
          let _rhs29 = (_139_v2)[_module.__default.safeIndex(_module.__default.fm31(_dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), function (_262_i11) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }), _151_globalState), new BigNumber((_139_v2).length))];
          _136_v0 = _rhs28;
          _138_v1 = _rhs29;
          let _263_v97;
          _263_v97 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,_142_v5);
          let _264_v98;
          _264_v98 = _module.D0.create_DC0(_module.__default.fm0(_138_v1, _151_globalState));
          let _265_v99;
          _265_v99 = _dafny.Map.Empty.slice().updateUnsafe(_264_v98,(_256_v91)[_module.__default.safeIndex(new BigNumber(79), new BigNumber((_256_v91).length))]);
          (_151_globalState).f23 = (((_263_v97).contains(_module.__default.fm21(_265_v99, _138_v1, _147_v8, _dafny.Seq.update(_140_v3, _module.__default.safeIndex(new BigNumber((_140_v3).length), new BigNumber((_140_v3).length)), _142_v5), _151_globalState))) ? ((_263_v97).get(_module.__default.fm21(_265_v99, _138_v1, _147_v8, _dafny.Seq.update(_140_v3, _module.__default.safeIndex(new BigNumber((_140_v3).length), new BigNumber((_140_v3).length)), _142_v5), _151_globalState))) : (_142_v5));
          (_151_globalState).f23 = _142_v5;
        } else {
          let _266_v100;
          let _nw33 = Array((new BigNumber(7)).toNumber());
          _266_v100 = _nw33;
          let _rhs30 = new BigNumber(-166);
          let _rhs31 = new BigNumber((_140_v3).length);
          let _rhs32 = _266_v100;
          _138_v1 = _rhs30;
          _138_v1 = _rhs31;
          _266_v100 = _rhs32;
          let _267_v101;
          let _nw34 = new _module.C13();
          _nw34.__ctor(!(_147_v8));
          _267_v101 = _nw34;
          let _268_v102;
          _268_v102 = _dafny.Seq.of(_267_v101);
          let _269_v103;
          _269_v103 = _dafny.Map.Empty.slice().updateUnsafe(_257_v92,(_268_v102)[_module.__default.safeIndex(_138_v1, new BigNumber((_268_v102).length))]);
          _269_v103 = (_269_v103).update(_257_v92, _267_v101);
          let _rhs33 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_140_v3, _140_v3), _dafny.Seq.update(_140_v3, _module.__default.safeIndex(_138_v1, new BigNumber((_140_v3).length)), _module.__default.fm33(new BigNumber(-47), _258_v93, _138_v1, _151_globalState))), _module.__default.safeIndex(_138_v1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_140_v3, _140_v3), _dafny.Seq.update(_140_v3, _module.__default.safeIndex(_138_v1, new BigNumber((_140_v3).length)), _module.__default.fm33(new BigNumber(-47), _258_v93, _138_v1, _151_globalState)))).length)), _142_v5);
          let _rhs34 = _138_v1;
          let _lhs23 = _151_globalState;
          _lhs23.f10 = _rhs33;
          _138_v1 = _rhs34;
          let _270_v104;
          let _nw35 = Array((new BigNumber(12)).toNumber());
          _270_v104 = _nw35;
          let _271_v105;
          let _nw36 = new _module.C11();
          _nw36.__ctor();
          _271_v105 = _nw36;
          let _index36 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_270_v104).length));
          (_270_v104)[_index36] = _271_v105;
          let _index37 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_270_v104).length));
          let _rhs35 = _271_v105;
          let _rhs36 = _136_v0;
          let _rhs37 = _138_v1;
          let _lhs24 = _270_v104;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_270_v104).length));
          _lhs24[_lhs25] = _rhs35;
          _136_v0 = _rhs36;
          _138_v1 = _rhs37;
          let _272_v106;
          let _nw37 = new _module.C5();
          _nw37.__ctor(_258_v93);
          _272_v106 = _nw37;
        }
        let _273_v107;
        let _nw38 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _273_v107 = _nw38;
        let _index38 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_273_v107).length));
        (_273_v107)[_index38] = _142_v5;
        let _index39 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_273_v107).length));
        (_273_v107)[_index39] = _module.__default.fm33(_138_v1, _258_v93, _138_v1, _151_globalState);
        (_151_globalState).f12 = _module.__default.fm0((_dafny.ZERO).minus(_138_v1), _151_globalState);
        if (_258_v93) {
          (_151_globalState).f23 = (_273_v107)[_module.__default.safeIndex(new BigNumber(885), new BigNumber((_273_v107).length))];
          let _274_v108;
          _274_v108 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,_159_v18);
          let _275_v109;
          let _nw39 = Array((new BigNumber(17)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _159_v18;
          _nw39[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((((_274_v108).contains(_138_v1)) ? ((_274_v108).get(_138_v1)) : (_159_v18)), _159_v18);
          _nw39[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_159_v18, _159_v18);
          _nw39[(new BigNumber(3)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(4)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_159_v18, _dafny.Seq.of(_147_v8));
          _nw39[(new BigNumber(6)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(7)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(8)).toNumber()] = _module.__default.fm1(_147_v8, _151_globalState);
          _nw39[(new BigNumber(9)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(10)).toNumber()] = _module.__default.fm1(_258_v93, _151_globalState);
          _nw39[(new BigNumber(11)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(12)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm1(_147_v8, _151_globalState), _159_v18);
          _nw39[(new BigNumber(14)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(15)).toNumber()] = _159_v18;
          _nw39[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_159_v18)[_module.__default.safeIndex(_138_v1, new BigNumber((_159_v18).length))], _147_v8), _module.__default.fm1(_258_v93, _151_globalState));
          _275_v109 = _nw39;
          let _index40 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_275_v109).length));
          (_275_v109)[_index40] = _159_v18;
          let _276_v110;
          _276_v110 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,(_dafny.ZERO).minus(_138_v1));
          let _277_v111;
          _277_v111 = _module.D1.create_DC6(new _dafny.CodePoint('n'.codePointAt(0)), _147_v8, _140_v3);
          let _278_v112;
          let _nw40 = Array((new BigNumber(14)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("sddesiht");
          _nw40[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_140_v3, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_147_v8)).length), new BigNumber((_140_v3).length)), _142_v5), _module.__default.safeIndex(_138_v1, new BigNumber((_dafny.Seq.update(_140_v3, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_147_v8)).length), new BigNumber((_140_v3).length)), _142_v5)).length)), (_273_v107)[_module.__default.safeIndex(new BigNumber(885), new BigNumber((_273_v107).length))]);
          _nw40[(new BigNumber(2)).toNumber()] = _140_v3;
          _nw40[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("mtvkkmo");
          _nw40[(new BigNumber(4)).toNumber()] = _module.__default.fm2(_138_v1, new BigNumber((_276_v110).length), _151_globalState);
          _nw40[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("rcb");
          _nw40[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(166)), function (_279_i12) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          });
          _nw40[(new BigNumber(7)).toNumber()] = _140_v3;
          _nw40[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("o");
          _nw40[(new BigNumber(9)).toNumber()] = _140_v3;
          _nw40[(new BigNumber(10)).toNumber()] = _140_v3;
          _nw40[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("ybch");
          _nw40[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_140_v3, _module.__default.safeIndex(_138_v1, new BigNumber((_140_v3).length)), (_273_v107)[_module.__default.safeIndex(new BigNumber(885), new BigNumber((_273_v107).length))]);
          _nw40[(new BigNumber(13)).toNumber()] = (_277_v111).dtor_cf9;
          _278_v112 = _nw40;
          let _280_v113;
          let _nw41 = Array((_dafny.ONE).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _278_v112;
          _280_v113 = _nw41;
          let _index41 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_280_v113).length));
          (_280_v113)[_index41] = _278_v112;
          let _index42 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_275_v109).length));
          let _index43 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_280_v113).length));
          let _rhs38 = _159_v18;
          let _rhs39 = _278_v112;
          let _lhs26 = _275_v109;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_275_v109).length));
          let _lhs28 = _280_v113;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_280_v113).length));
          _lhs26[_lhs27] = _rhs38;
          _lhs28[_lhs29] = _rhs39;
          let _rhs40 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs41 = _138_v1;
          _142_v5 = _rhs40;
          _138_v1 = _rhs41;
          (_151_globalState).f12 = _258_v93;
          (_151_globalState).f10 = _140_v3;
        } else {
          _138_v1 = _138_v1;
          _138_v1 = (_138_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("au")).length));
          let _281_v114;
          let _nw42 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _281_v114 = _nw42;
          let _282_v115;
          let _nw43 = new _module.C10();
          _nw43.__ctor((_273_v107)[_module.__default.safeIndex(new BigNumber(885), new BigNumber((_273_v107).length))], _138_v1, _281_v114);
          _282_v115 = _nw43;
          let _283_v116;
          let _nw44 = new _module.C1();
          _nw44.__ctor(_138_v1, _281_v114);
          _283_v116 = _nw44;
          let _284_v117;
          _284_v117 = _dafny.MultiSet.fromElements(_283_v116, _283_v116);
          let _285_v118;
          _285_v118 = _dafny.Set.fromElements(new BigNumber((_284_v117).cardinality()));
          let _286_v119;
          _286_v119 = _dafny.Map.Empty.slice().updateUnsafe((_285_v118).IsProperSubsetOf(_285_v118),_142_v5);
          _286_v119 = (_286_v119).update(_258_v93, _142_v5);
          (_151_globalState).f12 = _147_v8;
        }
      }
      let _287_v120;
      let _288_v121;
      let _289_v122;
      let _out17;
      let _out18;
      let _out19;
      let _outcollector6 = _module.__default.m0(_151_globalState);
      _out17 = _outcollector6[0];
      _out18 = _outcollector6[1];
      _out19 = _outcollector6[2];
      _287_v120 = _out17;
      _288_v121 = _out18;
      _289_v122 = _out19;
      _147_v8 = false;
      let _hi0 = _138_v1;
      for (let _290_i13 = _138_v1; _290_i13.isLessThan(_hi0); _290_i13 = _290_i13.plus(_dafny.ONE)) {
        _136_v0 = _136_v0;
        let _index44 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_136_v0).length));
        (_136_v0)[_index44] = _289_v122;
        let _291_v123;
        _291_v123 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,_138_v1);
        let _292_v124;
        _292_v124 = _module.D14.create_DC37(_291_v123);
        let _293_v125;
        _293_v125 = _module.D22.create_DC56(new _dafny.CodePoint('m'.codePointAt(0)), _147_v8);
        let _294_v126;
        _294_v126 = _dafny.Map.Empty.slice().updateUnsafe(_293_v125,_289_v122);
        let _295_v129;
        _295_v129 = _dafny.Seq.of(_293_v125);
        let _296_v130;
        _296_v130 = _dafny.Seq.of(_294_v126, function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of (_295_v129).Elements) {
              let _297_v128 = _compr_25;
              if (_dafny.Seq.contains(_295_v129, _297_v128)) {
                _coll25.push([_297_v128,_289_v122]);
              }
            }
            return _coll25;
          }()).Keys.Elements) {
            let _298_v127 = _compr_24;
            if ((function () {
              let _coll26 = new _dafny.Map();
              for (const _compr_26 of (_295_v129).Elements) {
                let _297_v128 = _compr_26;
                if (_dafny.Seq.contains(_295_v129, _297_v128)) {
                  _coll26.push([_297_v128,_289_v122]);
                }
              }
              return _coll26;
            }()).contains(_298_v127)) {
              _coll24.push([_298_v127,_289_v122]);
            }
          }
          return _coll24;
        }());
        let _299_v131;
        _299_v131 = _dafny.Map.Empty.slice().updateUnsafe(_292_v124,_296_v130);
        let _300_v133;
        _300_v133 = _dafny.MultiSet.fromElements(_294_v126, _294_v126, function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of (_module.__default.fm74(_151_globalState)).Elements) {
            let _301_v132 = _compr_27;
            if (_dafny.Seq.contains(_module.__default.fm74(_151_globalState), _301_v132)) {
              _coll27.push([_301_v132,_289_v122]);
            }
          }
          return _coll27;
        }());
        let _index45 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_136_v0).length));
        (_136_v0)[_index45] = (_300_v133).IsProperSubsetOf(_dafny.MultiSet.FromArray((((_299_v131).contains(_292_v124)) ? ((_299_v131).get(_292_v124)) : (_296_v130))));
        let _302_v134;
        let _nw45 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _302_v134 = _nw45;
        let _index46 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_302_v134).length));
        (_302_v134)[_index46] = _140_v3;
        let _index47 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_302_v134).length));
        (_302_v134)[_index47] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(213)), function (_303_i14) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), ((_147_v8) ? (_140_v3) : (_140_v3)));
        _288_v121 = _288_v121;
      }
      let _rhs42 = (_140_v3)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_140_v3).length))];
      let _rhs43 = (_141_v4)[_module.__default.safeIndex((_138_v1).multipliedBy(_module.__default.fm18(_151_globalState)), new BigNumber((_141_v4).length))];
      let _rhs44 = _289_v122;
      let _lhs30 = _151_globalState;
      let _lhs31 = _151_globalState;
      _lhs30.f23 = _rhs42;
      _lhs31.f10 = _rhs43;
      _289_v122 = _rhs44;
      _289_v122 = !(!(_147_v8));
      _191_v43 = (_191_v43).update(_138_v1, _147_v8);
      let _304_v135;
      _304_v135 = _module.D19.create_DC47(_136_v0, _147_v8);
      let _pat_let_tv1 = _287_v120;
      let _305_v136;
      _305_v136 = _dafny.MultiSet.fromElements(function (_pat_let2_0) {
        return function (_306_dt__update__tmp_h3) {
          return function (_pat_let3_0) {
            return function (_307_dt__update_hcf69_h1) {
              return _module.D19.create_DC47(_307_dt__update_hcf69_h1, (_306_dt__update__tmp_h3).dtor_cf70);
            }(_pat_let3_0);
          }(_pat_let_tv1);
        }(_pat_let2_0);
      }(_module.D19.create_DC47(_136_v0, _147_v8)), _304_v135, _304_v135, _304_v135);
      let _308_v137;
      _308_v137 = _dafny.MultiSet.fromElements(_142_v5);
      let _309_v138;
      _309_v138 = _dafny.Set.fromElements(new BigNumber((_308_v137).cardinality()));
      let _310_v139;
      _310_v139 = _dafny.Seq.update(_module.__default.fm19(!(_289_v122), _151_globalState), _module.__default.safeIndex(_138_v1, new BigNumber((_module.__default.fm19(!(_289_v122), _151_globalState)).length)), new BigNumber((_309_v138).length));
      let _pat_let_tv2 = _138_v1;
      let _rhs45 = _157_v17;
      let _rhs46 = function (_source9) {
        let _311___mcc_h4 = _source9;
        let _312_cf11 = _311___mcc_h4;
        return _pat_let_tv2;
      }(_310_v139);
      let _rhs47 = (((_289_v122) || (_147_v8)) ? ((_305_v136).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_304_v135, _304_v135)))) : (_305_v136));
      _157_v17 = _rhs45;
      _138_v1 = _rhs46;
      _305_v136 = _rhs47;
      _142_v5 = _142_v5;
      let _313_v140;
      let _nw46 = new _module.C11();
      _nw46.__ctor();
      _313_v140 = _nw46;
      _147_v8 = _147_v8;
      process.stdout.write(_dafny.toString((_136_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v0)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_138_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_139_v2, _dafny.Seq.of(new BigNumber(454)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_140_v3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_141_v4, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("oxvkwml")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_142_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v6).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("oxvkwml")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v7)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_v9).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_149_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(249),_dafny.MultiSet.fromElements(false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_150_v11).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f4)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_151_globalState.f6, _dafny.Seq.of(new BigNumber(454)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_151_globalState).f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_151_globalState.f10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState.f16).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("oxvkwml"), _dafny.Seq.UnicodeFromString("oxvkwml")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_globalState).f17)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(454),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f18).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(249),_dafny.MultiSet.fromElements(false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_globalState).f20).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_151_globalState.f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v16).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v16).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_v17).equals(_dafny.MultiSet.fromElements(new BigNumber(454), new BigNumber(454)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_159_v18, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_160_v19, _dafny.Seq.of(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_191_v43).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false).updateUnsafe(new BigNumber(2),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_192_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v120)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v121)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_v122));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_304_v135).dtor_cf69)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_304_v135).dtor_cf70));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_305_v136).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_308_v137).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_v138).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_310_v139), _dafny.Seq.of(_dafny.ONE, new BigNumber(166), _dafny.ONE, _dafny.ONE, new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC2(cf3) {
      let $dt = new D0(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC3(cf4, cf5) {
      let $dt = new D0(2);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC5() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC6(cf7, cf8, cf9) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC4(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC7(cf10) {
      let $dt = new D1(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5";
      } else if (this.$tag === 1) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + this.cf9.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC7" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC8(cf11) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf13, cf14, cf15, cf16, cf17) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13 && this.cf14 === other.cf14 && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, [], false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC13(cf19, cf20, cf21, cf22) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D4(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC14(cf23) {
      let $dt = new D4(3);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20 && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf25) {
      let $dt = new D5(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC17(cf26, cf27) {
      let $dt = new D5(1);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC15(cf24) {
      let $dt = new D5(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf29, cf30) {
      let $dt = new D6(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D6(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC20(cf31) {
      let $dt = new D6(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC19(_module.D0.Default(), _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf33, cf34) {
      let $dt = new D7(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC21(cf32) {
      let $dt = new D7(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf32 === other.cf32;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC22(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf36) {
      let $dt = new D8(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC23(cf35) {
      let $dt = new D8(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC25(cf37) {
      let $dt = new D8(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(_module.D1.Default());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf39, cf40) {
      let $dt = new D9(0);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf38) {
      let $dt = new D9(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC27(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf42, cf43, cf44) {
      let $dt = new D10(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC28(cf41) {
      let $dt = new D10(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42 && this.cf43 === other.cf43 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf45) {
      let $dt = new D11(0);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf47) {
      let $dt = new D12(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC33(cf48, cf49, cf50, cf51) {
      let $dt = new D12(1);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC31(cf46) {
      let $dt = new D12(2);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC34(cf52) {
      let $dt = new D12(3);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf48 === other.cf48 && this.cf49 === other.cf49 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf54, cf55, cf56) {
      let $dt = new D13(0);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC35(cf53) {
      let $dt = new D13(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54) && this.cf55 === other.cf55 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf53 === other.cf53;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC36(_dafny.Map.Empty, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf58) {
      let $dt = new D14(0);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC37(cf57) {
      let $dt = new D14(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC38(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC40() {
      let $dt = new D15(0);
      return $dt;
    }
    static create_DC39(cf59) {
      let $dt = new D15(1);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC40";
      } else if (this.$tag === 1) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC40();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf60) {
      let $dt = new D16(0);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf61) {
      let $dt = new D17(0);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC44(cf63, cf64, cf65, cf66, cf67) {
      let $dt = new D18(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC45() {
      let $dt = new D18(1);
      return $dt;
    }
    static create_DC43(cf62) {
      let $dt = new D18(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + this.cf65.toVerbatimString(true) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC45";
      } else if (this.$tag === 2) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63) && this.cf64 === other.cf64 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC44(_dafny.ZERO, false, _dafny.Seq.UnicodeFromString(""), false, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC47(cf69, cf70) {
      let $dt = new D19(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC48(cf71) {
      let $dt = new D19(1);
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC46(cf68) {
      let $dt = new D19(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC46() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC47" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf69 === other.cf69 && this.cf70 === other.cf70;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf68, other.cf68);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC47([], false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC50(cf73, cf74, cf75) {
      let $dt = new D20(0);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC49(cf72) {
      let $dt = new D20(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC50" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + this.cf75.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73) && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf72 === other.cf72;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC50(_module.D5.Default(), false, _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf77, cf78, cf79, cf80, cf81) {
      let $dt = new D21(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC53(cf82, cf83, cf84) {
      let $dt = new D21(1);
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC51(cf76) {
      let $dt = new D21(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC54(cf85) {
      let $dt = new D21(3);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get is_DC54() { return this.$tag === 3; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + this.cf79.toVerbatimString(true) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC53" + "(" + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC54" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79) && this.cf80 === other.cf80 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf82, other.cf82) && this.cf83 === other.cf83 && this.cf84 === other.cf84;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC52(_dafny.ZERO, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC56(cf87, cf88) {
      let $dt = new D22(0);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC57(cf89) {
      let $dt = new D22(1);
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC55(cf86) {
      let $dt = new D22(2);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC56" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC57" + "(" + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC55" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf87, other.cf87) && this.cf88 === other.cf88;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf89 === other.cf89;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf86 === other.cf86;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC56(new _dafny.CodePoint('D'.codePointAt(0)), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC59(cf91, cf92) {
      let $dt = new D23(0);
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC58(cf90) {
      let $dt = new D23(1);
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC60(cf93) {
      let $dt = new D23(2);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get is_DC60() { return this.$tag === 2; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC59" + "(" + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC58" + "(" + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC60" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf91 === other.cf91 && this.cf92 === other.cf92;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf90, other.cf90);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf93, other.cf93);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC59(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62(cf95) {
      let $dt = new D24(0);
      $dt.cf95 = cf95;
      return $dt;
    }
    static create_DC63(cf96, cf97, cf98, cf99) {
      let $dt = new D24(1);
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC64(cf100) {
      let $dt = new D24(2);
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC61(cf94) {
      let $dt = new D24(3);
      $dt.cf94 = cf94;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC64() { return this.$tag === 2; }
    get is_DC61() { return this.$tag === 3; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf94() { return this.cf94; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC62" + "(" + _dafny.toString(this.cf95) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC63" + "(" + _dafny.toString(this.cf96) + ", " + this.cf97.toVerbatimString(true) + ", " + _dafny.toString(this.cf98) + ", " + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC64" + "(" + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC61" + "(" + _dafny.toString(this.cf94) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf95 === other.cf95;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf96, other.cf96) && _dafny.areEqual(this.cf97, other.cf97) && _dafny.areEqual(this.cf98, other.cf98) && _dafny.areEqual(this.cf99, other.cf99);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf100, other.cf100);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf94, other.cf94);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC62(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC65(cf101) {
      let $dt = new D25(0);
      $dt.cf101 = cf101;
      return $dt;
    }
    get is_DC65() { return this.$tag === 0; }
    get dtor_cf101() { return this.cf101; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC65" + "(" + _dafny.toString(this.cf101) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf101, other.cf101);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f6 = _dafny.Seq.of();
      this.f10 = _dafny.Seq.UnicodeFromString("");
      this.f12 = false;
      this.f16 = _dafny.MultiSet.Empty;
      this.f23 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f4 = [];
      this._f5 = _dafny.ZERO;
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.Seq.UnicodeFromString("");
      this._f11 = false;
      this._f13 = false;
      this._f14 = false;
      this._f15 = false;
      this._f17 = [];
      this._f18 = _dafny.Map.Empty;
      this._f19 = false;
      this._f20 = _dafny.Set.Empty;
      this._f21 = _dafny.ZERO;
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this).f23 = f23;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm26(p0, p1, globalState) {
      let _this = this;
      let _source10 = _module.D4.create_DC14(_module.D4.create_DC12());
      if (_source10.is_DC12) {
        return new BigNumber((_dafny.Seq.of(!(!(false)))).length);
      } else if (_source10.is_DC13) {
        let _314___mcc_h0 = (_source10).cf19;
        let _315___mcc_h1 = (_source10).cf20;
        let _316___mcc_h2 = (_source10).cf21;
        let _317___mcc_h3 = (_source10).cf22;
        let _318_cf22 = _317___mcc_h3;
        let _319_cf21 = _316___mcc_h2;
        let _320_cf20 = _315___mcc_h1;
        let _321_cf19 = _314___mcc_h0;
        return new BigNumber((((_320_cf20) ? (_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(356)), (_module.D4.create_DC11(_dafny.Set.fromElements(_321_cf19))).dtor_cf18)) : (_dafny.Seq.of(_dafny.Set.fromElements(_321_cf19), function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of (function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of _dafny.IntegerRange(new BigNumber(345), new BigNumber(518))) {
              let _322_v0 = _compr_29;
              if (((new BigNumber(345)).isLessThanOrEqualTo(_322_v0)) && ((_322_v0).isLessThan(new BigNumber(518)))) {
                _coll29.add((_322_v0).plus(_321_cf19));
              }
            }
            return _coll29;
          }()).Elements) {
            let _323_v1 = _compr_28;
            if ((function () {
              let _coll30 = new _dafny.Set();
              for (const _compr_30 of _dafny.IntegerRange(new BigNumber(345), new BigNumber(518))) {
                let _324_v0 = _compr_30;
                if (((new BigNumber(345)).isLessThanOrEqualTo(_324_v0)) && ((_324_v0).isLessThan(new BigNumber(518)))) {
                  _coll30.add((_324_v0).plus(_321_cf19));
                }
              }
              return _coll30;
            }()).contains(_323_v1)) {
              _coll28.add(_module.__default.safeDivisionInt(_323_v1, new BigNumber(274)));
            }
          }
          return _coll28;
        }(), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("crfbyfmc")).length))), _dafny.Set.fromElements(_321_cf19))))).length);
      } else if (_source10.is_DC11) {
        let _325___mcc_h4 = (_source10).cf18;
        let _326_cf18 = _325___mcc_h4;
        return (new BigNumber(987)).multipliedBy(new BigNumber(75));
      } else {
        let _327___mcc_h5 = (_source10).cf23;
        let _328_cf23 = _327___mcc_h5;
        return new BigNumber(-90);
      }
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f25, f26) {
      let _this = this;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    m4(p0, globalState) {
      let _this = this;
      let _329_v0;
      _329_v0 = new BigNumber(-112);
      let _330_v1;
      _330_v1 = true;
      let _331_v2;
      _331_v2 = _dafny.Seq.of(_330_v1);
      let _332_v3;
      _332_v3 = new _dafny.CodePoint('p'.codePointAt(0));
      let _333_v4;
      _333_v4 = _dafny.Seq.of(_332_v3, _332_v3);
      _329_v0 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_331_v2, _module.__default.safeIndex((_this).f25, new BigNumber((_331_v2).length)), _330_v1)).length)), (_this).f25), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_332_v3), _333_v4)).length));
      _330_v1 = ((_330_v1) ? (!(_330_v1)) : (_330_v1));
      let _334_v5;
      _334_v5 = _dafny.Map.Empty.slice().updateUnsafe(_330_v1,new _dafny.CodePoint('x'.codePointAt(0)));
      _334_v5 = (_334_v5).update(_330_v1, _332_v3);
      let _335_v6;
      _335_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_329_v0, new BigNumber((_333_v4).length)),new BigNumber(495));
      _335_v6 = (_335_v6).update(((_dafny.ZERO).minus(_329_v0)).plus(_329_v0), _329_v0);
      let _336_v7;
      let _init6 = function (_337_i0) {
        return false;
      };
      let _nw47 = Array((new BigNumber(17)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw47.length); _i0_6++) {
        _nw47[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _336_v7 = _nw47;
      let _338_v8;
      _338_v8 = _module.D3.create_DC10(_330_v1, _336_v7, _330_v1, (_this).f25, _329_v0);
      let _339_v9;
      let _nw48 = Array((new BigNumber(20)).toNumber());
      _nw48[(_dafny.ZERO).toNumber()] = _336_v7;
      _nw48[(_dafny.ONE).toNumber()] = _336_v7;
      _nw48[(new BigNumber(2)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(3)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(4)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(5)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(6)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(7)).toNumber()] = ((_330_v1) ? (_336_v7) : (_336_v7));
      _nw48[(new BigNumber(8)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(9)).toNumber()] = (_338_v8).dtor_cf14;
      _nw48[(new BigNumber(10)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(11)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(12)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(13)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(14)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(15)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(16)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(17)).toNumber()] = _336_v7;
      _nw48[(new BigNumber(18)).toNumber()] = (_338_v8).dtor_cf14;
      _nw48[(new BigNumber(19)).toNumber()] = _336_v7;
      _339_v9 = _nw48;
      let _index48 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_339_v9).length));
      (_339_v9)[_index48] = _336_v7;
      let _340_v10;
      _340_v10 = _module.D1.create_DC5();
      let _pat_let_tv3 = _331_v2;
      let _pat_let_tv4 = _329_v0;
      let _pat_let_tv5 = _331_v2;
      let _pat_let_tv6 = _330_v1;
      let _pat_let_tv7 = _330_v1;
      let _index49 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_339_v9).length));
      let _rhs48 = _336_v7;
      let _rhs49 = function (_source11) {
        if (_source11.is_DC5) {
          return (_pat_let_tv3)[_module.__default.safeIndex(_pat_let_tv4, new BigNumber((_pat_let_tv5).length))];
        } else if (_source11.is_DC6) {
          let _341___mcc_h0 = (_source11).cf7;
          let _342___mcc_h1 = (_source11).cf8;
          let _343___mcc_h2 = (_source11).cf9;
          let _344_cf9 = _343___mcc_h2;
          let _345_cf8 = _342___mcc_h1;
          let _346_cf7 = _341___mcc_h0;
          return _345_cf8;
        } else if (_source11.is_DC4) {
          let _347___mcc_h3 = (_source11).cf6;
          let _348_cf6 = _347___mcc_h3;
          return _pat_let_tv6;
        } else {
          let _349___mcc_h4 = (_source11).cf10;
          let _350_cf10 = _349___mcc_h4;
          return _pat_let_tv7;
        }
      }(_340_v10);
      let _lhs32 = _339_v9;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_339_v9).length));
      let _lhs34 = globalState;
      _lhs32[_lhs33] = _rhs48;
      _lhs34.f12 = _rhs49;
      let _351_v11;
      let _nw49 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
      _351_v11 = _nw49;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_351_v11).length))) {
        let _352_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_352_i1)) && ((_352_i1).isLessThan(new BigNumber((_351_v11).length))))) {
          (_351_v11)[(_352_i1)] = _dafny.Seq.Concat(((_330_v1) ? (_dafny.Seq.of((_this).f25, new BigNumber(683))) : (_dafny.Seq.of(new BigNumber((_333_v4).length)))), _dafny.Seq.of((_this).f25, (_this).f25));
        }
      }
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _353_v0;
      _353_v0 = _dafny.MultiSet.fromElements(p2);
      let _354_v1;
      _354_v1 = _dafny.Seq.of(p1, p2, ((_this).f25).isLessThanOrEqualTo((_this).f25));
      _353_v0 = _dafny.MultiSet.FromArray(_354_v1);
      let _355_v2;
      _355_v2 = _dafny.Seq.UnicodeFromString("urjv");
      let _356_v3;
      _356_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_355_v2).length),_dafny.Seq.Concat(_355_v2, _355_v2));
      _356_v3 = _356_v3;
      r1 = ((_this).f25).plus((_this).f25);
      r0 = (_this).f25;
      r0 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f25, (_this).f25)), (_this).f25);
      _354_v1 = _354_v1;
      r0 = (_dafny.ZERO).minus(((_this).f25).multipliedBy((_this).f25));
      r1 = _module.__default.fm31(_355_v2, globalState);
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if (p1) {
        let _357_v0;
        _357_v0 = _module.D0.create_DC1(true, (!(p1)) && (p1));
        _357_v0 = _357_v0;
        let _358_v1;
        _358_v1 = _dafny.Seq.of((_this).f25, new BigNumber(995), (_this).f25, (_this).f25, (_this).f25);
        let _359_v2;
        _359_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(p1) || (p1),_module.__default.safeModuloInt((_358_v1)[_module.__default.safeIndex((_this).f25, new BigNumber((_358_v1).length))], (_this).f25));
        let _360_v3;
        _360_v3 = _dafny.MultiSet.fromElements(_module.__default.fm0((_this).f25, globalState), p1, p1);
        _359_v2 = (_359_v2).update(p1, new BigNumber((_360_v3).cardinality()));
        let _361_v4;
        _361_v4 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f25),(_358_v1)[_module.__default.safeIndex((_this).f25, new BigNumber((_358_v1).length))]);
        let _362_v5;
        _362_v5 = _module.D4.create_DC13(new BigNumber((_361_v4).length), !(p1), p1, p1);
        let _363_v6;
        _363_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_362_v5);
        _363_v6 = (_363_v6).update(p1, _362_v5);
        let _364_v7;
        let _nw50 = Array((_dafny.ONE).toNumber()).fill(false);
        _364_v7 = _nw50;
        let _365_v8;
        _365_v8 = _dafny.Map.Empty.slice().updateUnsafe(_364_v7,p1);
        if ((((_365_v8).contains(_364_v7)) ? ((_365_v8).get(_364_v7)) : (((_this).f25).isLessThan((_this).f25)))) {
          let _366_v9;
          _366_v9 = _dafny.Seq.UnicodeFromString("cvs");
          let _367_v10;
          _367_v10 = _dafny.Map.Empty.slice().updateUnsafe(p1,((_this).f25).isEqualTo(_module.__default.fm31(_366_v9, globalState)));
          let _368_v11;
          _368_v11 = _dafny.Seq.of(p1);
          let _369_v12;
          _369_v12 = new _dafny.CodePoint('s'.codePointAt(0));
          let _370_v13;
          _370_v13 = _dafny.Map.Empty.slice().updateUnsafe(_369_v12,_360_v3);
          _367_v10 = (_367_v10).update(p1, (_368_v11)[_module.__default.safeIndex(new BigNumber((_370_v13).length), new BigNumber((_368_v11).length))]);
          let _371_v14;
          _371_v14 = _module.D1.create_DC4(_dafny.Seq.update(_dafny.Seq.of(p1), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.of(p1)).length)), p1));
          _368_v11 = (_371_v14).dtor_cf6;
          (globalState).f10 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(443)), ((_372_v12) => function (_373_i0) {
            return _372_v12;
          })(_369_v12)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), ((_374_v12) => function (_375_i1) {
            return _374_v12;
          })(_369_v12)), _366_v9)), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(443)), ((_376_v12) => function (_377_i0) {
            return _376_v12;
          })(_369_v12)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), ((_378_v12) => function (_379_i1) {
            return _378_v12;
          })(_369_v12)), _366_v9))).length)), _369_v12);
          _361_v4 = (_361_v4).update(new BigNumber(-679), (new BigNumber((_368_v11).length)).multipliedBy((_this).f25));
          _359_v2 = (_359_v2).update(false, (_this).f25);
        } else {
          (globalState).f12 = !((p1) || (p1));
          let _380_v15;
          _380_v15 = _dafny.Seq.UnicodeFromString("pntpd");
          let _381_v16;
          _381_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), function (_382_i2) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          }), _380_v15));
          _381_v16 = (_381_v16).update(_module.__default.safeModuloInt((_this).f25, (_this).f25), (new BigNumber((_360_v3).cardinality())).isLessThanOrEqualTo((_this).f25));
          r0 = ((_this).f25).minus(((true) ? ((_this).f25) : ((_this).f25)));
          let _383_v17;
          _383_v17 = _dafny.MultiSet.fromElements((_this).f25, (_this).f25);
          let _384_v19;
          _384_v19 = _module.D5.create_DC15((_383_v17).update((_this).f25, _module.__default.abs((_this).f25)));
          let _385_v20;
          let _nw51 = Array((new BigNumber(29)).toNumber());
          _nw51[(_dafny.ZERO).toNumber()] = _383_v17;
          _nw51[(_dafny.ONE).toNumber()] = _module.__default.fm32((_this).f25, globalState);
          _nw51[(new BigNumber(2)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(3)).toNumber()] = ((_383_v17).update((_this).f25, _module.__default.abs(new BigNumber((function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(293), new BigNumber(694))) {
              let _386_v18 = _compr_31;
              if (((new BigNumber(293)).isLessThanOrEqualTo(_386_v18)) && ((_386_v18).isLessThan(new BigNumber(694)))) {
                _coll31.push([(_386_v18).minus((_this).f25),new _dafny.CodePoint('i'.codePointAt(0))]);
              }
            }
            return _coll31;
          }()).length)))).Difference(_dafny.MultiSet.fromElements((_this).f25));
          _nw51[(new BigNumber(4)).toNumber()] = (_383_v17).Union(_383_v17);
          _nw51[(new BigNumber(5)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(6)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(7)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements((_this).f25, _module.__default.fm31(_380_v15, globalState), (_this).f25, (_this).f25);
          _nw51[(new BigNumber(9)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(10)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(11)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(12)).toNumber()] = (_383_v17).Difference(_383_v17);
          _nw51[(new BigNumber(13)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.FromArray(_358_v1);
          _nw51[(new BigNumber(15)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(16)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(17)).toNumber()] = (_383_v17).Intersect(_dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_dafny.ZERO).minus(new BigNumber((_383_v17).cardinality())), (_this).f25, (_this).f25));
          _nw51[(new BigNumber(18)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(19)).toNumber()] = ((_module.__default.fm0((_this).f25, globalState)) ? ((_dafny.MultiSet.fromElements(new BigNumber((_380_v15).length))).update(_module.__default.fm31(_380_v15, globalState), _module.__default.abs((_this).f25))) : ((_383_v17).update((_this).f25, _module.__default.abs(new BigNumber(147)))));
          _nw51[(new BigNumber(20)).toNumber()] = (_383_v17).Intersect(_383_v17);
          _nw51[(new BigNumber(21)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(22)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(23)).toNumber()] = _module.__default.fm32(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(701)), function (_387_i3) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })).length), globalState);
          _nw51[(new BigNumber(24)).toNumber()] = (_383_v17).Intersect((_384_v19).dtor_cf24);
          _nw51[(new BigNumber(25)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(26)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(27)).toNumber()] = _383_v17;
          _nw51[(new BigNumber(28)).toNumber()] = _383_v17;
          _385_v20 = _nw51;
          let _388_v21;
          let _nw52 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _388_v21 = _nw52;
          let _index50 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_385_v20).length));
          (_385_v20)[_index50] = (_dafny.MultiSet.fromElements((_this).f25)).Intersect(_module.__default.fm32(new BigNumber((_dafny.Seq.of(_388_v21)).length), globalState));
          let _index51 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_385_v20).length));
          (_385_v20)[_index51] = _383_v17;
          let _389_v22;
          _389_v22 = _dafny.Seq.of(p1);
          let _390_v23;
          _390_v23 = _dafny.Set.fromElements((_this).f25, _module.__default.fm31(_dafny.Seq.UnicodeFromString("htgrnhtir"), globalState));
          let _391_v24;
          _391_v24 = _dafny.Map.Empty.slice().updateUnsafe((_389_v22)[_module.__default.safeIndex((_this).f25, new BigNumber((_389_v22).length))],(_dafny.Set.fromElements(new BigNumber(697), (_this).f25, (_this).f25, (_this).f25)).IsSubsetOf(_390_v23));
          let _392_v25;
          _392_v25 = _module.D1.create_DC6(_module.__default.fm33((_this).f25, !(p1), (_this).f25, globalState), p1, _380_v15);
          _391_v24 = (_391_v24).update(false, (_392_v25).dtor_cf8);
        }
        let _393_v26;
        _393_v26 = _dafny.Seq.UnicodeFromString("mqhshe");
        r0 = _module.__default.fm31(_dafny.Seq.Concat(_393_v26, _dafny.Seq.Create(_module.__default.abs(new BigNumber(898)), function (_394_i4) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })), globalState);
      } else {
        r0 = (_this).f25;
        let _395_v27;
        let _init7 = function (_396_i5) {
          return _module.__default.safeModuloInt(_396_i5, (_this).f25);
        };
        let _nw53 = Array((new BigNumber(26)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw53.length); _i0_7++) {
          _nw53[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _395_v27 = _nw53;
        let _index52 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_395_v27).length));
        (_395_v27)[_index52] = (_this).f25;
        let _397_v28;
        _397_v28 = _dafny.Seq.UnicodeFromString("hy");
        let _398_v29;
        _398_v29 = _dafny.Seq.of(new BigNumber((_397_v28).length));
        let _399_v30;
        _399_v30 = _dafny.Seq.of(_398_v29);
        let _400_v31;
        let _nw54 = Array((new BigNumber(7)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _399_v30;
        _nw54[(_dafny.ONE).toNumber()] = _399_v30;
        _nw54[(new BigNumber(2)).toNumber()] = _399_v30;
        _nw54[(new BigNumber(3)).toNumber()] = _399_v30;
        _nw54[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_399_v30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(392)), function (_401_i6) {
          return _dafny.Seq.of((_this).f25, (_dafny.ZERO).minus((_this).f25));
        }));
        _nw54[(new BigNumber(5)).toNumber()] = _399_v30;
        _nw54[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_399_v30, _module.__default.safeIndex((_this).f25, new BigNumber((_399_v30).length)), _398_v29);
        _400_v31 = _nw54;
        let _index53 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_400_v31).length));
        (_400_v31)[_index53] = _399_v30;
        let _402_v32;
        let _nw55 = Array((new BigNumber(29)).toNumber()).fill(false);
        _402_v32 = _nw55;
        let _index54 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_402_v32).length));
        (_402_v32)[_index54] = _dafny.Seq.contains(_397_v28, _module.__default.fm33((_this).f25, !(true), (_this).f25, globalState));
        let _403_v33;
        _403_v33 = _dafny.Map.Empty.slice().updateUnsafe(p1,_397_v28);
        let _404_v34;
        let _init8 = ((_405_p1) => function (_406_i7) {
          return _module.D5.create_DC16(_405_p1);
        })(p1);
        let _nw56 = Array((new BigNumber(23)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw56.length); _i0_8++) {
          _nw56[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _404_v34 = _nw56;
        let _407_v35;
        _407_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_404_v34);
        let _408_v36;
        _408_v36 = _dafny.MultiSet.fromElements((_this).f25, new BigNumber((_407_v35).length));
        let _409_v37;
        _409_v37 = _module.D5.create_DC15(_408_v36);
        let _410_v38;
        _410_v38 = _dafny.Seq.of(_409_v37, _409_v37, _409_v37);
        let _411_v39;
        _411_v39 = _dafny.Seq.of(_410_v38, _410_v38);
        let _index55 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_395_v27).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_400_v31).length));
        let _index57 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_402_v32).length));
        let _rhs50 = (new BigNumber((_dafny.Seq.of((_this).f25)).length)).multipliedBy((_this).f25);
        let _rhs51 = _module.__default.safeDivisionInt(((_this).f25).plus((_this).f25), (_this).f25);
        let _rhs52 = _module.__default.safeDivisionInt(new BigNumber((_403_v33).length), (_this).f25);
        let _rhs53 = _399_v30;
        let _rhs54 = _dafny.Seq.IsProperPrefixOf(_411_v39, _411_v39);
        let _lhs35 = _395_v27;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_395_v27).length));
        let _lhs37 = _400_v31;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_400_v31).length));
        let _lhs39 = _402_v32;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_402_v32).length));
        r0 = _rhs50;
        r0 = _rhs51;
        _lhs35[_lhs36] = _rhs52;
        _lhs37[_lhs38] = _rhs53;
        _lhs39[_lhs40] = _rhs54;
        let _index58 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_402_v32).length));
        let _rhs55 = p1;
        let _rhs56 = !((_395_v27)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_395_v27).length))]).isEqualTo((_this).f25);
        let _lhs41 = _402_v32;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_402_v32).length));
        let _lhs43 = globalState;
        _lhs41[_lhs42] = _rhs55;
        _lhs43.f12 = _rhs56;
        _408_v36 = _408_v36;
        (globalState).f12 = _module.__default.fm0(((_this).f25).multipliedBy((_395_v27)[_module.__default.safeIndex(new BigNumber(69), new BigNumber((_395_v27).length))]), globalState);
      }
      let _412_v40;
      _412_v40 = _dafny.Seq.UnicodeFromString("vp");
      let _413_v41;
      _413_v41 = _dafny.Set.fromElements(new BigNumber((_412_v40).length));
      let _414_v42;
      _414_v42 = _module.D4.create_DC11(_413_v41);
      let _415_v43;
      _415_v43 = _dafny.Seq.of(_414_v42, _module.__default.fm34(p1, p1, p1, (_this).f25, globalState));
      let _416_v44;
      _416_v44 = _dafny.Map.Empty.slice().updateUnsafe(_415_v43,false);
      _416_v44 = (_416_v44).update(_dafny.Seq.of(_414_v42, _414_v42, _414_v42), p1);
      let _417_v45;
      let _nw57 = Array((new BigNumber(23)).toNumber()).fill(false);
      _417_v45 = _nw57;
      _417_v45 = _417_v45;
      r0 = (_this).f25;
      let _hi1 = (_this).f25;
      for (let _418_i8 = (_this).f25; _418_i8.isLessThan(_hi1); _418_i8 = _418_i8.plus(_dafny.ONE)) {
        let _source12 = _414_v42;
        if (_source12.is_DC12) {
          let _419_v46;
          _419_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_413_v41).IsSubsetOf(_dafny.Set.fromElements((_this).f25)));
          let _420_v47;
          _420_v47 = new _dafny.CodePoint('j'.codePointAt(0));
          let _421_v48;
          _421_v48 = _dafny.MultiSet.fromElements(_420_v47, _420_v47, new _dafny.CodePoint('t'.codePointAt(0)));
          let _422_v49;
          _422_v49 = _dafny.Map.Empty.slice().updateUnsafe(_420_v47,_421_v48);
          let _rhs57 = _418_i8;
          let _rhs58 = _419_v46;
          let _rhs59 = p1;
          let _rhs60 = _412_v40;
          let _rhs61 = _422_v49;
          let _lhs44 = globalState;
          let _lhs45 = globalState;
          r0 = _rhs57;
          _419_v46 = _rhs58;
          _lhs44.f12 = _rhs59;
          _lhs45.f10 = _rhs60;
          _422_v49 = _rhs61;
          let _index59 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_417_v45).length));
          (_417_v45)[_index59] = p1;
          let _423_v50;
          _423_v50 = _dafny.Seq.of(p1, p1, !(p1) || (p1), p1);
          let _index60 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_417_v45).length));
          (_417_v45)[_index60] = (_423_v50)[_module.__default.safeIndex(_module.__default.fm31(_412_v40, globalState), new BigNumber((_423_v50).length))];
          r0 = _418_i8;
          (globalState).f12 = p1;
        } else if (_source12.is_DC13) {
          let _424___mcc_h0 = (_source12).cf19;
          let _425___mcc_h1 = (_source12).cf20;
          let _426___mcc_h2 = (_source12).cf21;
          let _427___mcc_h3 = (_source12).cf22;
          let _428_cf22 = _427___mcc_h3;
          let _429_cf21 = _426___mcc_h2;
          let _430_cf20 = _425___mcc_h1;
          let _431_cf19 = _424___mcc_h0;
          (globalState).f23 = new _dafny.CodePoint('k'.codePointAt(0));
          _431_cf19 = (_431_cf19).minus((new BigNumber(-388)).minus(_431_cf19));
          let _432_v51;
          let _nw58 = new _module.C0();
          _nw58.__ctor();
          _432_v51 = _nw58;
          let _433_v52;
          let _nw59 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _433_v52 = _nw59;
          let _434_v53;
          _434_v53 = _dafny.MultiSet.fromElements(_433_v52, _433_v52, ((_429_cf21) ? (_433_v52) : (_433_v52)), _433_v52);
          let _rhs62 = (_431_cf19).minus((new BigNumber(251)).multipliedBy(_431_cf19));
          let _rhs63 = _434_v53;
          r0 = _rhs62;
          _434_v53 = _rhs63;
        } else if (_source12.is_DC11) {
          let _435___mcc_h4 = (_source12).cf18;
          let _436_cf18 = _435___mcc_h4;
          let _437_v54;
          _437_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC17(new BigNumber((_module.__default.fm35(_418_i8, globalState)).length), p1),p1);
          let _438_v55;
          _438_v55 = _dafny.MultiSet.fromElements(_437_v54, (_437_v54).Merge(_437_v54));
          let _rhs64 = _438_v55;
          let _rhs65 = ((p1) ? ((_this).f25) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(551)), ((_439_i8) => function (_440_i9) {
            return (_dafny.ZERO).minus(_439_i8);
          })(_418_i8))).length)));
          let _rhs66 = _module.__default.fm33((_this).f25, !(p1), (_this).f25, globalState);
          let _lhs46 = globalState;
          _438_v55 = _rhs64;
          r0 = _rhs65;
          _lhs46.f23 = _rhs66;
          (globalState).f12 = _module.__default.fm0((_dafny.ZERO).minus(_418_i8), globalState);
          let _rhs67 = _418_i8;
          let _rhs68 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_this).f25, new BigNumber((_412_v40).length)), new BigNumber(226)));
          r0 = _rhs67;
          r0 = _rhs68;
          let _441_v56;
          _441_v56 = _dafny.MultiSet.fromElements(p1);
          let _442_v57;
          _442_v57 = _module.D0.create_DC2(true);
          let _443_v58;
          _443_v58 = _dafny.Set.fromElements(p1);
          let _rhs69 = (_441_v56).update(!((_442_v57).dtor_cf3), _module.__default.abs(_418_i8));
          let _rhs70 = !((_443_v58).IsProperSubsetOf((_443_v58).Difference(_443_v58)));
          let _lhs47 = globalState;
          _441_v56 = _rhs69;
          _lhs47.f12 = _rhs70;
        } else {
          let _444___mcc_h5 = (_source12).cf23;
          let _445_cf23 = _444___mcc_h5;
          let _446_v59;
          let _nw60 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _446_v59 = _nw60;
          let _index61 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_446_v59).length));
          (_446_v59)[_index61] = ((_dafny.ZERO).minus((_this).f25)).multipliedBy((_this).f25);
          let _447_v60;
          _447_v60 = _dafny.MultiSet.fromElements(((_this).f25).minus((_this).f25));
          let _index62 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_446_v59).length));
          let _rhs71 = new BigNumber((_447_v60).cardinality());
          let _rhs72 = p1;
          let _rhs73 = _412_v40;
          let _rhs74 = ((new BigNumber(((_dafny.MultiSet.fromElements(p1)).update(p1, _module.__default.abs(new BigNumber((_447_v60).cardinality())))).cardinality())).plus(_418_i8)).isEqualTo(_418_i8);
          let _lhs48 = _446_v59;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_446_v59).length));
          let _lhs50 = globalState;
          let _lhs51 = globalState;
          let _lhs52 = globalState;
          _lhs48[_lhs49] = _rhs71;
          _lhs50.f12 = _rhs72;
          _lhs51.f10 = _rhs73;
          _lhs52.f12 = _rhs74;
          (globalState).f12 = !((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25))).isEqualTo((_446_v59)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_446_v59).length))]);
          let _448_v61;
          let _nw61 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _448_v61 = _nw61;
          let _449_v62;
          _449_v62 = _dafny.Map.Empty.slice().updateUnsafe(_448_v61,p1);
          _449_v62 = (_449_v62).update(_448_v61, p1);
          let _index63 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_446_v59).length));
          (_446_v59)[_index63] = _418_i8;
        }
        let _450_v63;
        let _init9 = ((_451_v41) => function (_452_i10) {
          return _451_v41;
        })(_413_v41);
        let _nw62 = Array((new BigNumber(9)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw62.length); _i0_9++) {
          _nw62[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _450_v63 = _nw62;
        let _index64 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_450_v63).length));
        (_450_v63)[_index64] = _413_v41;
        let _index65 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_450_v63).length));
        (_450_v63)[_index65] = (_module.D4.create_DC11(_dafny.Set.fromElements(_418_i8))).dtor_cf18;
        let _453_v64;
        let _nw63 = Array((new BigNumber(21)).toNumber()).fill([]);
        _453_v64 = _nw63;
        let _454_v65;
        _454_v65 = _dafny.MultiSet.fromElements((_this).f25);
        let _455_v66;
        _455_v66 = _dafny.Seq.of(_418_i8, new BigNumber((_412_v40).length), _418_i8, new BigNumber((_454_v65).cardinality()));
        let _456_v67;
        _456_v67 = _455_v66;
        let _457_v68;
        _457_v68 = _dafny.Seq.of(p1);
        let _458_v69;
        let _nw64 = Array((new BigNumber(11)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), function (_459_i11) {
          return (_this).f25;
        });
        _nw64[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_460_i8) => function (_461_i12) {
          return _460_i8;
        })(_418_i8));
        _nw64[(new BigNumber(2)).toNumber()] = _455_v66;
        _nw64[(new BigNumber(3)).toNumber()] = _455_v66;
        _nw64[(new BigNumber(4)).toNumber()] = (_456_v67);
        _nw64[(new BigNumber(5)).toNumber()] = _dafny.Seq.of((_this).f25, new BigNumber((_457_v68).length));
        _nw64[(new BigNumber(6)).toNumber()] = _455_v66;
        _nw64[(new BigNumber(7)).toNumber()] = _455_v66;
        _nw64[(new BigNumber(8)).toNumber()] = _455_v66;
        _nw64[(new BigNumber(9)).toNumber()] = _455_v66;
        _nw64[(new BigNumber(10)).toNumber()] = _455_v66;
        _458_v69 = _nw64;
        let _index66 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_453_v64).length));
        (_453_v64)[_index66] = _458_v69;
        let _index67 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_453_v64).length));
        (_453_v64)[_index67] = _458_v69;
        if (false) {
          (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.Concat(_412_v40, _412_v40), _412_v40);
          (globalState).f12 = false;
          let _462_v70;
          let _init10 = ((_463_p1) => function (_464_i13) {
            return _module.__default.safeDivisionInt(_464_i13, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_463_p1,new _dafny.CodePoint('j'.codePointAt(0)))).length));
          })(p1);
          let _nw65 = Array((new BigNumber(18)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw65.length); _i0_10++) {
            _nw65[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _462_v70 = _nw65;
          _462_v70 = _462_v70;
          let _465_v71;
          _465_v71 = new _dafny.CodePoint('k'.codePointAt(0));
          r0 = new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_412_v40, _dafny.Seq.Concat(_412_v40, _dafny.Seq.update(_412_v40, _module.__default.safeIndex(_418_i8, new BigNumber((_412_v40).length)), _465_v71))), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f25), new BigNumber((_dafny.Seq.Concat(_412_v40, _dafny.Seq.Concat(_412_v40, _dafny.Seq.update(_412_v40, _module.__default.safeIndex(_418_i8, new BigNumber((_412_v40).length)), _465_v71)))).length)), _module.__default.fm33(new BigNumber((_457_v68).length), p1, (_this).f25, globalState)), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_412_v40, _dafny.Seq.Concat(_412_v40, _dafny.Seq.update(_412_v40, _module.__default.safeIndex(_418_i8, new BigNumber((_412_v40).length)), _465_v71))), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f25), new BigNumber((_dafny.Seq.Concat(_412_v40, _dafny.Seq.Concat(_412_v40, _dafny.Seq.update(_412_v40, _module.__default.safeIndex(_418_i8, new BigNumber((_412_v40).length)), _465_v71)))).length)), _module.__default.fm33(new BigNumber((_457_v68).length), p1, (_this).f25, globalState))).length)), ((false) ? (_465_v71) : (_465_v71)))).length);
          let _466_v72;
          let _nw66 = new _module.C0();
          _nw66.__ctor();
          _466_v72 = _nw66;
        } else {
          r0 = new BigNumber((_455_v66).length);
          (globalState).f12 = p1;
          let _467_v73;
          _467_v73 = _dafny.Map.Empty.slice().updateUnsafe(_418_i8,_412_v40);
          _467_v73 = (_467_v73).update((_this).f25, _dafny.Seq.Concat(_412_v40, _dafny.Seq.Create(_module.__default.abs(new BigNumber(688)), function (_468_i14) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })));
          let _469_v74;
          let _nw67 = Array((new BigNumber(22)).toNumber()).fill(_module.D5.Default());
          _469_v74 = _nw67;
          _469_v74 = _469_v74;
          let _rhs75 = _dafny.Seq.Concat(_dafny.Seq.Concat(_412_v40, _dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), function (_470_i15) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })), _dafny.Seq.UnicodeFromString("gltbwdt"));
          let _rhs76 = (_this).f25;
          let _rhs77 = _417_v45;
          let _lhs53 = globalState;
          _lhs53.f10 = _rhs75;
          r0 = _rhs76;
          _417_v45 = _rhs77;
        }
      }
      let _471_v75;
      let _nw68 = Array((new BigNumber(22)).toNumber()).fill([]);
      _471_v75 = _nw68;
      _471_v75 = _471_v75;
      r0 = new BigNumber(178);
      return r0;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm23(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-603)), function (_472_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length), new BigNumber(139)),new BigNumber(231))).length);
    };
    fm24(p0, globalState) {
      let _this = this;
      return new BigNumber(110);
    };
    m16(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.of();
      let r3 = _module.D1.Default();
      let _473_v0;
      _473_v0 = new BigNumber(840);
      let _hi2 = ((_dafny.ZERO).minus(_473_v0)).plus((_this).fm24(_473_v0, globalState));
      for (let _474_i0 = _473_v0; _474_i0.isLessThan(_hi2); _474_i0 = _474_i0.plus(_dafny.ONE)) {
        let _475_v1;
        let _init11 = function (_476_i1) {
          return _dafny.Seq.UnicodeFromString("cgmmbr");
        };
        let _nw69 = Array((_dafny.ONE).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw69.length); _i0_11++) {
          _nw69[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _475_v1 = _nw69;
        let _477_v2;
        _477_v2 = _dafny.Seq.UnicodeFromString("rmlqj");
        let _index68 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length));
        (_475_v1)[_index68] = _477_v2;
        let _478_v3;
        let _nw70 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _478_v3 = _nw70;
        let _479_v4;
        let _nw71 = Array((new BigNumber(12)).toNumber()).fill(_module.D3.Default());
        _479_v4 = _nw71;
        let _480_v5;
        _480_v5 = _dafny.Map.Empty.slice().updateUnsafe(_479_v4,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(416)), function (_481_i2) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length));
        let _index69 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_478_v3).length));
        (_478_v3)[_index69] = _480_v5;
        let _482_v6;
        _482_v6 = new _dafny.CodePoint('f'.codePointAt(0));
        let _483_v7;
        _483_v7 = false;
        let _484_v8;
        _484_v8 = _module.D1.create_DC6(_482_v6, _483_v7, _477_v2);
        let _index70 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length));
        let _index71 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_478_v3).length));
        let _rhs78 = (_484_v8).dtor_cf9;
        let _rhs79 = _480_v5;
        let _lhs54 = _475_v1;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length));
        let _lhs56 = _478_v3;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_478_v3).length));
        _lhs54[_lhs55] = _rhs78;
        _lhs56[_lhs57] = _rhs79;
        if (_483_v7) {
          (globalState).f12 = _module.__default.fm0(_474_i0, globalState);
          let _485_v9;
          let _nw72 = Array((new BigNumber(16)).toNumber()).fill(false);
          _485_v9 = _nw72;
          let _index72 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_485_v9).length));
          (_485_v9)[_index72] = !(_483_v7);
          let _index73 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_485_v9).length));
          (_485_v9)[_index73] = _module.__default.fm0(_module.__default.safeDivisionInt(_473_v0, _474_i0), globalState);
          let _486_v10;
          let _init12 = function (_487_i3) {
            return (_487_i3).multipliedBy(new BigNumber(981));
          };
          let _nw73 = Array((new BigNumber(16)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw73.length); _i0_12++) {
            _nw73[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _486_v10 = _nw73;
          _486_v10 = _486_v10;
          let _488_v11;
          _488_v11 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), function (_489_i4) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length)).multipliedBy(_473_v0),(_474_i0).isLessThan(new BigNumber(708)));
          let _490_v12;
          _490_v12 = _dafny.Seq.of(_483_v7);
          _488_v11 = (_488_v11).update((_dafny.ZERO).minus(_473_v0), (_490_v12)[_module.__default.safeIndex(_474_i0, new BigNumber((_490_v12).length))]);
          let _491_v13;
          _491_v13 = _module.D1.create_DC5();
          let _492_v14;
          let _nw74 = Array((new BigNumber(20)).toNumber());
          _nw74[(_dafny.ZERO).toNumber()] = _491_v13;
          _nw74[(_dafny.ONE).toNumber()] = _module.D1.create_DC5();
          _nw74[(new BigNumber(2)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(3)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(4)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(5)).toNumber()] = ((_483_v7) ? (_491_v13) : (_491_v13));
          _nw74[(new BigNumber(6)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(7)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(8)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(9)).toNumber()] = _module.__default.fm25(_483_v7, (_485_v9)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_485_v9).length))], _474_i0, (_485_v9)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_485_v9).length))], globalState);
          _nw74[(new BigNumber(10)).toNumber()] = _module.D1.create_DC5();
          _nw74[(new BigNumber(11)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(12)).toNumber()] = _module.D1.create_DC5();
          _nw74[(new BigNumber(13)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(14)).toNumber()] = ((!((_485_v9)[_module.__default.safeIndex(new BigNumber(419), new BigNumber((_485_v9).length))])) ? (_491_v13) : (_module.__default.fm25(_483_v7, _483_v7, _474_i0, _483_v7, globalState)));
          _nw74[(new BigNumber(15)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(16)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(17)).toNumber()] = _491_v13;
          _nw74[(new BigNumber(18)).toNumber()] = _module.__default.fm25(_483_v7, _483_v7, new BigNumber(153), _483_v7, globalState);
          _nw74[(new BigNumber(19)).toNumber()] = _491_v13;
          _492_v14 = _nw74;
          let _index74 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_492_v14).length));
          (_492_v14)[_index74] = _491_v13;
          let _index75 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_492_v14).length));
          (_492_v14)[_index75] = _491_v13;
        } else {
          let _rhs80 = _473_v0;
          let _rhs81 = _483_v7;
          let _rhs82 = (_473_v0).plus(_474_i0);
          let _lhs58 = globalState;
          r1 = _rhs80;
          _lhs58.f12 = _rhs81;
          r1 = _rhs82;
          (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.Concat((_475_v1)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length))], (_475_v1)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length))]), (_475_v1)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length))]);
          (globalState).f12 = _483_v7;
          let _493_v15;
          let _nw75 = new _module.C0();
          _nw75.__ctor();
          _493_v15 = _nw75;
          let _494_v16;
          _494_v16 = _dafny.Set.fromElements(_474_i0);
          let _495_v18;
          _495_v18 = _dafny.Set.fromElements(_494_v16, _494_v16, (_494_v16).Difference(_module.__default.fm27(_483_v7, function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of _dafny.IntegerRange(new BigNumber(470), new BigNumber(-771))) {
              let _496_v17 = _compr_32;
              if (((new BigNumber(470)).isLessThanOrEqualTo(_496_v17)) && ((_496_v17).isLessThan(new BigNumber(-771)))) {
                _coll32.push([_module.__default.safeDivisionInt(_496_v17, _473_v0),(_475_v1)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_475_v1).length))]]);
              }
            }
            return _coll32;
          }(), globalState)));
          _495_v18 = _495_v18;
        }
        let _497_v19;
        _497_v19 = _dafny.Map.Empty.slice().updateUnsafe(_483_v7,_483_v7);
        let _498_v20;
        _498_v20 = _module.D4.create_DC13(new BigNumber(((_497_v19).Merge(_497_v19)).length), !(_483_v7) || (_483_v7), _483_v7, _483_v7);
        _498_v20 = _498_v20;
        let _499_v21;
        let _init13 = ((_500_v7) => function (_501_i5) {
          return _dafny.Seq.Concat(_dafny.Seq.of(!(_500_v7)), _dafny.Seq.of(_500_v7));
        })(_483_v7);
        let _nw76 = Array((new BigNumber(11)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw76.length); _i0_13++) {
          _nw76[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _499_v21 = _nw76;
        let _502_v22;
        _502_v22 = _dafny.Seq.of(_483_v7, _483_v7, _483_v7, _483_v7);
        let _index76 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_499_v21).length));
        (_499_v21)[_index76] = _502_v22;
        let _index77 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_499_v21).length));
        (_499_v21)[_index77] = _dafny.Seq.of(true, _483_v7);
      }
      if (_module.__default.fm0(new BigNumber(-959), globalState)) {
        let _503_v23;
        _503_v23 = _dafny.Seq.UnicodeFromString("fjlplkxx");
        let _504_v24;
        _504_v24 = _dafny.Seq.of(_473_v0, new BigNumber((_503_v23).length), _473_v0);
        let _505_v25;
        _505_v25 = false;
        let _506_v26;
        _506_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_504_v24)[_module.__default.safeIndex(_473_v0, new BigNumber((_504_v24).length))]),_505_v25);
        let _507_v27;
        _507_v27 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_504_v24)[_module.__default.safeIndex(_473_v0, new BigNumber((_504_v24).length))],_505_v25), _506_v26);
        (globalState).f12 = !((_507_v27).IsSubsetOf(_507_v27));
        let _508_v28;
        let _nw77 = new _module.C0();
        _nw77.__ctor();
        _508_v28 = _nw77;
        _473_v0 = (_dafny.ZERO).minus((_this).fm23(_module.__default.fm28(_505_v25, _473_v0, _503_v23, _473_v0, globalState), _505_v25, !(_module.__default.fm0(_473_v0, globalState)) || (_505_v25), globalState));
        let _509_v29;
        let _init14 = ((_510_v0) => function (_511_i6) {
          return _module.__default.safeDivisionInt(_511_i6, _510_v0);
        })(_473_v0);
        let _nw78 = Array((new BigNumber(12)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw78.length); _i0_14++) {
          _nw78[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _509_v29 = _nw78;
        let _index78 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_509_v29).length));
        (_509_v29)[_index78] = _473_v0;
        let _index79 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_509_v29).length));
        (_509_v29)[_index79] = _473_v0;
        (globalState).f12 = _505_v25;
      } else {
        let _512_v30;
        _512_v30 = false;
        let _513_v31;
        _513_v31 = _dafny.Seq.of(_512_v30, true, _512_v30);
        if ((new BigNumber((_513_v31).length)).isEqualTo(_473_v0)) {
          let _514_v32;
          let _nw79 = Array((new BigNumber(8)).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = _513_v31;
          _nw79[(_dafny.ONE).toNumber()] = _513_v31;
          _nw79[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_512_v30);
          _nw79[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_513_v31, _module.__default.fm1(_512_v30, globalState));
          _nw79[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_512_v30);
          _nw79[(new BigNumber(5)).toNumber()] = _513_v31;
          _nw79[(new BigNumber(6)).toNumber()] = _513_v31;
          _nw79[(new BigNumber(7)).toNumber()] = _513_v31;
          _514_v32 = _nw79;
          let _index80 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_514_v32).length));
          (_514_v32)[_index80] = _513_v31;
          let _index81 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_514_v32).length));
          (_514_v32)[_index81] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_512_v30), _513_v31), _dafny.Seq.Concat(_dafny.Seq.update(_513_v31, _module.__default.safeIndex(_473_v0, new BigNumber((_513_v31).length)), _512_v30), _dafny.Seq.of(_512_v30)));
          let _515_v33;
          let _nw80 = new _module.C0();
          _nw80.__ctor();
          _515_v33 = _nw80;
          (globalState).f12 = !(!(_512_v30) || (_512_v30));
          let _516_v34;
          _516_v34 = _dafny.Seq.UnicodeFromString("fycjace");
          let _517_v35;
          _517_v35 = _dafny.Seq.of((_516_v34)[_module.__default.safeIndex(_473_v0, new BigNumber((_516_v34).length))]);
          (globalState).f6 = _module.__default.fm29(_dafny.Seq.Concat(_517_v35, _module.__default.fm2(_473_v0, _473_v0, globalState)), _512_v30, _473_v0, _517_v35, globalState);
          r0 = _473_v0;
        } else {
          let _518_v36;
          _518_v36 = _dafny.Seq.of(_473_v0);
          r1 = _module.__default.safeDivisionInt(new BigNumber((_518_v36).length), _473_v0);
          r0 = new BigNumber((function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(666), new BigNumber(838))) {
              let _519_v37 = _compr_33;
              if (((new BigNumber(666)).isLessThanOrEqualTo(_519_v37)) && ((_519_v37).isLessThan(new BigNumber(838)))) {
                _coll33.add(_module.__default.safeModuloInt(_519_v37, _473_v0));
              }
            }
            return _coll33;
          }()).length);
          let _520_v38;
          _520_v38 = _dafny.Map.Empty.slice().updateUnsafe((_512_v30) || (_512_v30),_512_v30);
          _520_v38 = _520_v38;
          let _521_v39;
          _521_v39 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_473_v0),_512_v30);
          let _522_v40;
          _522_v40 = _dafny.Map.Empty.slice().updateUnsafe(_512_v30,_473_v0);
          let _523_v41;
          _523_v41 = _dafny.Map.Empty.slice().updateUnsafe(!(_473_v0).isEqualTo((_this).fm23(_dafny.Map.Empty.slice().updateUnsafe(_512_v30,_473_v0), false, _512_v30, globalState)),_module.__default.fm30((_521_v39).update((_this).fm23(_522_v40, _512_v30, _512_v30, globalState), _module.__default.fm0(_473_v0, globalState)), globalState));
          let _524_v42;
          _524_v42 = new _dafny.CodePoint('x'.codePointAt(0));
          let _525_v43;
          _525_v43 = _dafny.Seq.UnicodeFromString("hymwetbrg");
          let _526_v44;
          _526_v44 = _module.D1.create_DC6(_524_v42, !(_512_v30), _525_v43);
          _523_v41 = (_523_v41).update(_512_v30, _526_v44);
          let _527_v45;
          let _nw81 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _527_v45 = _nw81;
          let _528_v46;
          let _529_v47;
          let _out20;
          let _out21;
          let _outcollector7 = (_this).m17((_473_v0).isLessThan(_473_v0), _527_v45, _512_v30, ((_512_v30) ? (_512_v30) : (_512_v30)), globalState);
          _out20 = _outcollector7[0];
          _out21 = _outcollector7[1];
          _528_v46 = _out20;
          _529_v47 = _out21;
        }
        let _530_v48;
        let _nw82 = new _module.C0();
        _nw82.__ctor();
        _530_v48 = _nw82;
        (globalState).f12 = !(_512_v30);
        let _531_v49;
        _531_v49 = _dafny.Seq.UnicodeFromString("singysn");
        r1 = (_473_v0).minus(new BigNumber((_531_v49).length));
        let _532_v50;
        _532_v50 = _dafny.Map.Empty.slice().updateUnsafe(_530_v48,_473_v0);
        let _533_v51;
        _533_v51 = _dafny.Set.fromElements(_473_v0);
        let _534_v52;
        _534_v52 = _dafny.Map.Empty.slice().updateUnsafe(_533_v51,_533_v51);
        r1 = (((_532_v50).contains(_530_v48)) ? ((_532_v50).get(_530_v48)) : (new BigNumber((_534_v52).length)));
      }
      let _hi3 = _473_v0;
      for (let _535_i7 = (new BigNumber(-324)).minus(_473_v0); _535_i7.isLessThan(_hi3); _535_i7 = _535_i7.plus(_dafny.ONE)) {
        let _536_v53;
        let _init15 = function (_537_i8) {
          return (_537_i8).multipliedBy(new BigNumber((_dafny.Set.fromElements(false)).length));
        };
        let _nw83 = Array((new BigNumber(19)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw83.length); _i0_15++) {
          _nw83[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _536_v53 = _nw83;
        _536_v53 = _536_v53;
        let _538_v54;
        _538_v54 = _dafny.Map.Empty.slice().updateUnsafe(_473_v0,_535_i7);
        let _539_v55;
        _539_v55 = _dafny.Seq.of((((_538_v54).contains(_535_i7)) ? ((_538_v54).get(_535_i7)) : (new BigNumber(655))), _473_v0, _535_i7);
        let _540_v56;
        _540_v56 = false;
        let _541_v57;
        _541_v57 = _dafny.Map.Empty.slice().updateUnsafe(_539_v55,_540_v56);
        let _542_v58;
        _542_v58 = _dafny.Seq.of(_540_v56);
        _541_v57 = (_541_v57).update(_539_v55, (_540_v56) === ((_542_v58)[_module.__default.safeIndex(_535_i7, new BigNumber((_542_v58).length))]));
        let _index82 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_536_v53).length));
        (_536_v53)[_index82] = new BigNumber(182);
        let _543_v59;
        _543_v59 = _dafny.Seq.UnicodeFromString("bunpdm");
        let _544_v60;
        _544_v60 = _dafny.MultiSet.fromElements(new BigNumber((_543_v59).length));
        let _index83 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_536_v53).length));
        let _rhs83 = !(_540_v56) || (_540_v56);
        let _rhs84 = _dafny.Seq.Concat(_543_v59, _543_v59);
        let _rhs85 = (((_544_v60).contains(((_540_v56) ? (_535_i7) : (_535_i7)))) ? ((_544_v60).get(((_540_v56) ? (_535_i7) : (_535_i7)))) : ((_dafny.ZERO).minus(_535_i7)));
        let _rhs86 = (_dafny.ZERO).minus(_473_v0);
        let _rhs87 = _541_v57;
        let _lhs59 = globalState;
        let _lhs60 = _536_v53;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_536_v53).length));
        _540_v56 = _rhs83;
        _lhs59.f10 = _rhs84;
        _lhs60[_lhs61] = _rhs85;
        r1 = _rhs86;
        _541_v57 = _rhs87;
        let _545_v61;
        _545_v61 = _module.D0.create_DC3(_535_i7, ((_540_v56) ? (true) : (_540_v56)));
        _545_v61 = _545_v61;
      }
      r1 = _473_v0;
      let _546_v62;
      let _nw84 = Array((new BigNumber(26)).toNumber()).fill(false);
      _546_v62 = _nw84;
      let _547_v63;
      _547_v63 = _dafny.Set.fromElements(_546_v62, _546_v62);
      let _548_i9;
      _548_i9 = _dafny.ZERO;
      L4: {
        while ((_dafny.Set.fromElements(_546_v62, _546_v62)).IsDisjointFrom((_547_v63).Union(_dafny.Set.fromElements(_546_v62)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_548_i9)) {
              break L4;
            }
            _548_i9 = (_548_i9).plus(_dafny.ONE);
            (globalState).f12 = (_473_v0).isLessThan(_module.__default.safeModuloInt(_473_v0, (_dafny.ZERO).minus(_473_v0)));
            let _549_v64;
            _549_v64 = true;
            if (_549_v64) {
              r1 = ((_dafny.ZERO).minus(_473_v0)).minus(_473_v0);
              let _550_v65;
              let _init16 = ((_551_v0) => function (_552_i10) {
                return (_552_i10).minus(_551_v0);
              })(_473_v0);
              let _nw85 = Array((new BigNumber(6)).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw85.length); _i0_16++) {
                _nw85[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _550_v65 = _nw85;
              let _553_v66;
              _553_v66 = _dafny.MultiSet.fromElements(_550_v65);
              let _554_v67;
              _554_v67 = _dafny.MultiSet.fromElements((_553_v66).IsProperSubsetOf(_dafny.MultiSet.fromElements(_550_v65)), !(_549_v64));
              r1 = new BigNumber((_554_v67).cardinality());
              let _555_v68;
              _555_v68 = _dafny.Set.fromElements(_549_v64, _549_v64);
              let _556_v69;
              _556_v69 = _dafny.Map.Empty.slice().updateUnsafe(_549_v64,_555_v68);
              let _557_v70;
              _557_v70 = _dafny.Map.Empty.slice().updateUnsafe(_556_v69,_473_v0);
              _557_v70 = (_557_v70).update((_556_v69).update(false, _555_v68), _473_v0);
              _473_v0 = (((_473_v0).isLessThan(_473_v0)) ? (_473_v0) : ((_dafny.ZERO).minus(_473_v0)));
              let _558_v71;
              _558_v71 = new _dafny.CodePoint('h'.codePointAt(0));
              (globalState).f23 = _558_v71;
            } else {
              _473_v0 = (_dafny.ZERO).minus(_473_v0);
              let _559_v72;
              _559_v72 = _module.D4.create_DC12();
              _559_v72 = ((_549_v64) ? (_559_v72) : (_559_v72));
              let _560_v73;
              _560_v73 = _dafny.Seq.UnicodeFromString("uunsmmj");
              let _561_v74;
              _561_v74 = _dafny.MultiSet.fromElements(_549_v64);
              let _562_v75;
              _562_v75 = new _dafny.CodePoint('v'.codePointAt(0));
              let _563_v76;
              _563_v76 = _module.D1.create_DC6(_562_v75, _549_v64, _dafny.Seq.UnicodeFromString("eyxjlokv"));
              let _564_v77;
              let _nw86 = Array((new BigNumber(7)).toNumber());
              _nw86[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("qu");
              _nw86[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(198)), function (_565_i11) {
                return new _dafny.CodePoint('y'.codePointAt(0));
              });
              _nw86[(new BigNumber(2)).toNumber()] = _560_v73;
              _nw86[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-519)), function (_566_i12) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              }), _module.__default.safeIndex(new BigNumber((_561_v74).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-519)), function (_567_i12) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              })).length)), _562_v75);
              _nw86[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(219)), ((_568_v75) => function (_569_i13) {
                return _568_v75;
              })(_562_v75));
              _nw86[(new BigNumber(5)).toNumber()] = (_563_v76).dtor_cf9;
              _nw86[(new BigNumber(6)).toNumber()] = _560_v73;
              _564_v77 = _nw86;
              let _570_v78;
              let _nw87 = new _module.C1();
              _nw87.__ctor(_473_v0, _564_v77);
              _570_v78 = _nw87;
              let _571_v79;
              _571_v79 = _dafny.Set.fromElements(_570_v78, _570_v78);
              (globalState).f12 = ((_571_v79).Intersect(_571_v79)).IsSubsetOf(((_549_v64) ? (_571_v79) : (_571_v79)));
              r1 = (_570_v78).f25;
              _546_v62 = _546_v62;
            }
            let _572_v80;
            _572_v80 = _dafny.Seq.UnicodeFromString("ybi");
            _473_v0 = new BigNumber((_572_v80).length);
            (globalState).f12 = _549_v64;
          }
        }
      }
      let _573_v81;
      _573_v81 = false;
      let _574_i14;
      _574_i14 = _dafny.ZERO;
      L5: {
        while (_573_v81) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_574_i14)) {
              break L5;
            }
            _574_i14 = (_574_i14).plus(_dafny.ONE);
            (globalState).f12 = _573_v81;
            _473_v0 = _473_v0;
            let _575_v82;
            _575_v82 = _dafny.Seq.UnicodeFromString("mvsjnjdbt");
            let _576_v83;
            _576_v83 = new _dafny.CodePoint('j'.codePointAt(0));
            let _577_v84;
            _577_v84 = _dafny.Map.Empty.slice().updateUnsafe((_473_v0).plus(new BigNumber((_575_v82).length)),!_dafny.Seq.contains(_575_v82, _576_v83));
            let _578_v85;
            _578_v85 = _dafny.MultiSet.fromElements(_573_v81, _573_v81, _573_v81);
            let _579_v86;
            _579_v86 = _dafny.Map.Empty.slice().updateUnsafe(_573_v81,new BigNumber(868));
            _577_v84 = (_577_v84).update((new BigNumber(86)).minus(new BigNumber(((_578_v85).update(_573_v81, _module.__default.abs((_this).fm23(_579_v86, _573_v81, _573_v81, globalState)))).cardinality())), _573_v81);
            let _580_v87;
            _580_v87 = _dafny.MultiSet.fromElements(_575_v82, _575_v82, _575_v82);
            let _581_v88;
            _581_v88 = _dafny.Seq.of(true, _573_v81, _573_v81, _573_v81);
            let _582_v89;
            _582_v89 = _dafny.Map.Empty.slice().updateUnsafe(_573_v81,_576_v83);
            let _583_v90;
            let _nw88 = Array((new BigNumber(7)).toNumber());
            _nw88[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_473_v0, _473_v0);
            _nw88[(_dafny.ONE).toNumber()] = _473_v0;
            _nw88[(new BigNumber(2)).toNumber()] = (((_580_v87).contains(_575_v82)) ? ((_580_v87).get(_575_v82)) : (_473_v0));
            _nw88[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_581_v88, _581_v88)).length);
            _nw88[(new BigNumber(4)).toNumber()] = _473_v0;
            _nw88[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_582_v89).length), _473_v0);
            _nw88[(new BigNumber(6)).toNumber()] = _473_v0;
            _583_v90 = _nw88;
            let _index84 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_583_v90).length));
            (_583_v90)[_index84] = new BigNumber((_575_v82).length);
            let _index85 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_583_v90).length));
            let _rhs88 = ((_dafny.ZERO).minus(_473_v0)).plus(new BigNumber(3));
            let _rhs89 = _576_v83;
            let _rhs90 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_dafny.ZERO).minus(_473_v0), _473_v0), _473_v0);
            let _rhs91 = _module.__default.safeModuloInt(_473_v0, _473_v0);
            let _lhs62 = _583_v90;
            let _lhs63 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_583_v90).length));
            r1 = _rhs88;
            _576_v83 = _rhs89;
            _lhs62[_lhs63] = _rhs90;
            _473_v0 = _rhs91;
          }
        }
      }
      r0 = _473_v0;
      r1 = _473_v0;
      let _584_v91;
      _584_v91 = _dafny.Seq.of(true);
      let _585_v92;
      _585_v92 = _dafny.MultiSet.fromElements((_584_v91)[_module.__default.safeIndex(new BigNumber((_584_v91).length), new BigNumber((_584_v91).length))]);
      let _586_v93;
      _586_v93 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(630),_573_v81);
      let _587_v94;
      _587_v94 = _dafny.Map.Empty.slice().updateUnsafe(_473_v0,_586_v93);
      r2 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_473_v0, (_dafny.ZERO).minus(_473_v0)), _module.__default.safeIndex((new BigNumber((_585_v92).cardinality())).multipliedBy(_473_v0), new BigNumber((_dafny.Seq.of(_473_v0, (_dafny.ZERO).minus(_473_v0))).length)), _module.__default.safeDivisionInt(_473_v0, new BigNumber((_587_v94).length))), _module.__default.safeIndex(_473_v0, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_473_v0, (_dafny.ZERO).minus(_473_v0)), _module.__default.safeIndex((new BigNumber((_585_v92).cardinality())).multipliedBy(_473_v0), new BigNumber((_dafny.Seq.of(_473_v0, (_dafny.ZERO).minus(_473_v0))).length)), _module.__default.safeDivisionInt(_473_v0, new BigNumber((_587_v94).length)))).length)), ((_573_v81) ? (_473_v0) : (_473_v0)));
      let _588_v95;
      _588_v95 = _module.D1.create_DC4(_dafny.Seq.of(_573_v81));
      r3 = ((_573_v81) ? (_588_v95) : (_588_v95));
      return [r0, r1, r2, r3];
    }
    m17(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _589_v0;
      let _nw89 = Array((new BigNumber(4)).toNumber()).fill(false);
      _589_v0 = _nw89;
      let _590_v1;
      _590_v1 = _dafny.Seq.of(p1);
      let _591_v2;
      _591_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_590_v1, _590_v1),_589_v0);
      _589_v0 = (((_591_v2).contains(_590_v1)) ? ((_591_v2).get(_590_v1)) : (_589_v0));
      let _592_v3;
      _592_v3 = new BigNumber(136);
      r1 = (_dafny.ZERO).minus(_592_v3);
      let _593_v4;
      _593_v4 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_592_v3),p2);
      let _594_i0;
      _594_i0 = _dafny.ZERO;
      L6: {
        while ((((_593_v4).contains(_592_v3)) ? ((_593_v4).get(_592_v3)) : (p0))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_594_i0)) {
              break L6;
            }
            _594_i0 = (_594_i0).plus(_dafny.ONE);
            let _595_v5;
            _595_v5 = _dafny.Set.fromElements(_589_v0, _589_v0, _589_v0);
            let _596_v6;
            _596_v6 = _module.D3.create_DC9(_595_v5);
            let _597_v7;
            _597_v7 = new _dafny.CodePoint('u'.codePointAt(0));
            let _598_v8;
            _598_v8 = _dafny.MultiSet.fromElements(_592_v3);
            let _pat_let_tv8 = _589_v0;
            let _pat_let_tv9 = _589_v0;
            let _pat_let_tv10 = _589_v0;
            let _pat_let_tv11 = _589_v0;
            _596_v6 = (((_598_v8).contains(new BigNumber((_module.__default.fm36(p2, _593_v4, _597_v7, globalState)).length))) ? (_596_v6) : (function (_pat_let4_0) {
              return function (_599_dt__update__tmp_h0) {
                return function (_pat_let5_0) {
                  return function (_600_dt__update_hcf12_h0) {
                    return _module.D3.create_DC9(_600_dt__update_hcf12_h0);
                  }(_pat_let5_0);
                }(_dafny.Set.fromElements(_pat_let_tv8, _pat_let_tv9, _pat_let_tv10, _pat_let_tv11));
              }(_pat_let4_0);
            }(_module.D3.create_DC9(_595_v5))));
            let _601_v9;
            let _nw90 = new _module.C0();
            _nw90.__ctor();
            _601_v9 = _nw90;
            let _602_v10;
            let _nw91 = Array((_dafny.ONE).toNumber());
            _nw91[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), ((_603_v7) => function (_604_i1) {
              return _603_v7;
            })(_597_v7));
            _602_v10 = _nw91;
            let _605_v11;
            let _nw92 = new _module.C1();
            _nw92.__ctor(_592_v3, _602_v10);
            _605_v11 = _nw92;
            _605_v11 = _605_v11;
            let _606_v12;
            _606_v12 = _dafny.Seq.UnicodeFromString("ppb");
            _592_v3 = _module.__default.fm31(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rudhwqxa"), _606_v12), globalState);
          }
        }
      }
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_589_v0).length))) {
        let _607_i2 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_607_i2)) && ((_607_i2).isLessThan(new BigNumber((_589_v0).length))))) {
          (_589_v0)[(_607_i2)] = p0;
        }
      }
      let _608_v13;
      _608_v13 = _dafny.Set.fromElements(_592_v3, _592_v3);
      let _609_v14;
      let _nw93 = Array((new BigNumber(6)).toNumber());
      _nw93[(_dafny.ZERO).toNumber()] = _592_v3;
      _nw93[(_dafny.ONE).toNumber()] = _592_v3;
      _nw93[(new BigNumber(2)).toNumber()] = _592_v3;
      _nw93[(new BigNumber(3)).toNumber()] = new BigNumber(796);
      _nw93[(new BigNumber(4)).toNumber()] = new BigNumber((_608_v13).length);
      _nw93[(new BigNumber(5)).toNumber()] = _592_v3;
      _609_v14 = _nw93;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_609_v14).length))) {
        let _610_i3 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_610_i3)) && ((_610_i3).isLessThan(new BigNumber((_609_v14).length))))) {
          (_609_v14)[(_610_i3)] = _module.__default.safeModuloInt(_610_i3, _592_v3);
        }
      }
      let _611_v15;
      let _nw94 = Array((new BigNumber(16)).toNumber()).fill(_module.D1.Default());
      _611_v15 = _nw94;
      let _612_v16;
      _612_v16 = _dafny.Seq.of(!(!(p3)));
      let _index86 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_611_v15).length));
      (_611_v15)[_index86] = _module.D1.create_DC4(_612_v16);
      let _613_v17;
      _613_v17 = _module.D1.create_DC4(_dafny.Seq.of(true));
      let _index87 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_611_v15).length));
      (_611_v15)[_index87] = _613_v17;
      r0 = p1;
      let _614_v18;
      _614_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(p2));
      let _615_v19;
      _615_v19 = _dafny.Seq.of(new BigNumber((_614_v18).length));
      r1 = _module.__default.safeDivisionInt(new BigNumber((((_module.__default.fm0(_592_v3, globalState)) ? (_dafny.Seq.update(_615_v19, _module.__default.safeIndex(_592_v3, new BigNumber((_615_v19).length)), _592_v3)) : (_dafny.Seq.of(new BigNumber(-40))))).length), new BigNumber(-753));
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
      this._f37 = _dafny.ZERO;
      this._f38 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f37, f38, f25, f26) {
      let _this = this;
      (_this)._f37 = f37;
      (_this)._f38 = f38;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm38(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_616_i0) {
        return (_this).f37;
      }), _dafny.Seq.of((_this).f25, (_this).f37));
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _617_v0;
      let _nw95 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
      _617_v0 = _nw95;
      let _618_v1;
      let _init17 = ((_619_p1) => function (_620_i0) {
        return !(_619_p1);
      })(p1);
      let _nw96 = Array((new BigNumber(16)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw96.length); _i0_17++) {
        _nw96[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _618_v1 = _nw96;
      let _index88 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_617_v0).length));
      (_617_v0)[_index88] = _dafny.MultiSet.fromElements(_618_v1, _618_v1, _618_v1);
      let _621_v2;
      _621_v2 = _dafny.MultiSet.fromElements(_618_v1);
      let _index89 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_617_v0).length));
      (_617_v0)[_index89] = (_621_v2).Intersect(_dafny.MultiSet.fromElements(_618_v1));
      let _622_v3;
      let _nw97 = new _module.C0();
      _nw97.__ctor();
      _622_v3 = _nw97;
      r0 = (_this).f37;
      if (p2) {
        let _index90 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length));
        (_618_v1)[_index90] = p2;
        let _623_v4;
        _623_v4 = _dafny.Seq.UnicodeFromString("gp");
        let _index91 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length));
        (_618_v1)[_index91] = (new BigNumber((_623_v4).length)).isLessThanOrEqualTo(new BigNumber(375));
        let _624_v5;
        let _nw98 = Array((new BigNumber(18)).toNumber()).fill(false);
        _624_v5 = _nw98;
        let _625_v6;
        _625_v6 = _dafny.Map.Empty.slice().updateUnsafe(_624_v5,(_this).f37);
        let _626_v7;
        _626_v7 = _dafny.Seq.of((_618_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length))]);
        let _627_v8;
        _627_v8 = _dafny.Set.fromElements((_this).f37, (_this).f25);
        let _628_v9;
        _628_v9 = _dafny.Seq.of(new BigNumber((_626_v7).length), new BigNumber((_627_v8).length));
        let _629_v10;
        _629_v10 = _module.D6.create_DC18((_dafny.Map.Empty.slice().updateUnsafe(_624_v5,(_this).f37)).Merge((_625_v6).update(_624_v5, new BigNumber((_628_v9).length))));
        let _630_v11;
        _630_v11 = _module.D0.create_DC1(p2, p2);
        let _631_v12;
        _631_v12 = _dafny.Set.fromElements((_630_v11).dtor_cf2);
        let _632_v13;
        _632_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_631_v12).length), (_this).f25),(_618_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length))]);
        let _633_v14;
        _633_v14 = _module.D1.create_DC5();
        let _634_v15;
        let _nw99 = Array((new BigNumber(24)).toNumber());
        _nw99[(_dafny.ZERO).toNumber()] = _633_v14;
        _nw99[(_dafny.ONE).toNumber()] = _633_v14;
        _nw99[(new BigNumber(2)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(3)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(4)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(5)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(6)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(7)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(8)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(9)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(10)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(11)).toNumber()] = _module.D1.create_DC5();
        _nw99[(new BigNumber(12)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(13)).toNumber()] = _module.D1.create_DC5();
        _nw99[(new BigNumber(14)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(15)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(16)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(17)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(18)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(19)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(20)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(21)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(22)).toNumber()] = _633_v14;
        _nw99[(new BigNumber(23)).toNumber()] = _633_v14;
        _634_v15 = _nw99;
        let _index92 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_634_v15).length));
        (_634_v15)[_index92] = _module.D1.create_DC5();
        let _635_v16;
        let _nw100 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _635_v16 = _nw100;
        let _636_v17;
        _636_v17 = _module.D7.create_DC21(p3);
        let _index93 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_634_v15).length));
        let _rhs92 = _629_v10;
        let _rhs93 = _632_v13;
        let _rhs94 = _633_v14;
        let _rhs95 = (p1) === ((_618_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length))]);
        let _rhs96 = (_636_v17).dtor_cf32;
        let _lhs64 = _634_v15;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_634_v15).length));
        let _lhs66 = globalState;
        _629_v10 = _rhs92;
        _632_v13 = _rhs93;
        _lhs64[_lhs65] = _rhs94;
        _lhs66.f12 = _rhs95;
        _635_v16 = _rhs96;
        let _637_v18;
        _637_v18 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f37).isLessThan((_this).f25),(_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f25, (_this).f25)));
        _637_v18 = (_637_v18).update((_this).f38, (_this).f25);
        let _638_v19;
        _638_v19 = new _dafny.CodePoint('p'.codePointAt(0));
        let _639_v20;
        _639_v20 = _module.D1.create_DC6(_638_v19, p1, _623_v4);
        let _640_v22;
        _640_v22 = _dafny.Map.Empty.slice().updateUnsafe(_639_v20,function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of (_627_v8).Elements) {
            let _641_v21 = _compr_34;
            if ((_627_v8).contains(_641_v21)) {
              _coll34.add((_641_v21).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true, false)).cardinality()))));
            }
          }
          return _coll34;
        }());
        _627_v8 = (((_640_v22).contains(_639_v20)) ? ((_640_v22).get(_639_v20)) : ((_627_v8).Intersect(_627_v8)));
        let _642_v23;
        _642_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), ((_643_v19) => function (_644_i1) {
          return _643_v19;
        })(_638_v19)));
        let _645_v24;
        let _nw101 = Array((new BigNumber(12)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_623_v4, _623_v4);
        _nw101[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_623_v4, _623_v4);
        _nw101[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("b");
        _nw101[(new BigNumber(3)).toNumber()] = _623_v4;
        _nw101[(new BigNumber(4)).toNumber()] = _623_v4;
        _nw101[(new BigNumber(5)).toNumber()] = _623_v4;
        _nw101[(new BigNumber(6)).toNumber()] = (((_642_v23).contains((_this).f37)) ? ((_642_v23).get((_this).f37)) : (_623_v4));
        _nw101[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_623_v4, _623_v4);
        _nw101[(new BigNumber(8)).toNumber()] = _623_v4;
        _nw101[(new BigNumber(9)).toNumber()] = _623_v4;
        _nw101[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_646_v4) => function (_647_i2) {
          return (_646_v4)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f25), new BigNumber((_646_v4).length))];
        })(_623_v4));
        _nw101[(new BigNumber(11)).toNumber()] = _623_v4;
        _645_v24 = _nw101;
        let _648_v25;
        _648_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber((_628_v9).length));
        let _index94 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length));
        let _rhs97 = _module.__default.safeDivisionInt((_this).f37, (_this).f25);
        let _rhs98 = _645_v24;
        let _rhs99 = (_618_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length))];
        let _rhs100 = _module.__default.safeDivisionInt(((_module.__default.fm0((_this).f37, globalState)) ? (new BigNumber((_631_v12).length)) : ((_622_v3).fm26(_dafny.ONE, (_618_v1)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length))], globalState))), (((_648_v25).contains((_this).f25)) ? ((_648_v25).get((_this).f25)) : ((_this).f37)));
        let _lhs67 = _618_v1;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_618_v1).length));
        r1 = _rhs97;
        _645_v24 = _rhs98;
        _lhs67[_lhs68] = _rhs99;
        r1 = _rhs100;
      } else {
        let _649_v26;
        _649_v26 = _module.D3.create_DC10((_this).f38, _618_v1, p1, (_this).f25, new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), _module.__default.fm39(p1, globalState))).length));
        let _650_v27;
        _650_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,(_649_v26).dtor_cf14);
        _650_v27 = _650_v27;
        let _651_v28;
        _651_v28 = _dafny.Seq.of(true);
        r0 = new BigNumber((_dafny.Seq.Concat(_651_v28, _dafny.Seq.update(_651_v28, _module.__default.safeIndex((_this).f37, new BigNumber((_651_v28).length)), p1))).length);
        (globalState).f10 = _dafny.Seq.UnicodeFromString("pjn");
        let _652_v29;
        _652_v29 = _module.D7.create_DC21(p3);
        let _653_v30;
        _653_v30 = _dafny.Map.Empty.slice().updateUnsafe(_618_v1,_652_v29);
        _653_v30 = _653_v30;
        r1 = (_this).f25;
      }
      let _654_v31;
      _654_v31 = _dafny.Map.Empty.slice().updateUnsafe(_618_v1,p1);
      let _655_v32;
      _655_v32 = _module.D8.create_DC23(_654_v31);
      if (((_655_v32).dtor_cf35).contains(((true) ? (_618_v1) : (_618_v1)))) {
        (globalState).f12 = (p2) || (p1);
        let _656_v33;
        _656_v33 = _dafny.Seq.of(p1);
        let _657_v34;
        _657_v34 = _module.D1.create_DC4(_656_v33);
        let _658_v35;
        _658_v35 = _module.D1.create_DC7(_657_v34);
        let _pat_let_tv12 = _657_v34;
        let _source13 = function (_pat_let6_0) {
          return function (_659_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_660_dt__update_hcf10_h0) {
                return _module.D1.create_DC7(_660_dt__update_hcf10_h0);
              }(_pat_let7_0);
            }(_pat_let_tv12);
          }(_pat_let6_0);
        }(_658_v35);
        if (_source13.is_DC5) {
          let _661_v36;
          _661_v36 = _dafny.Seq.UnicodeFromString("xlmwcmkr");
          let _index95 = _module.__default.safeIndex(new BigNumber(802), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index95] = _661_v36;
          let _index96 = _module.__default.safeIndex(new BigNumber(802), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index96] = _661_v36;
          let _662_v37;
          _662_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f25);
          let _663_v38;
          _663_v38 = _dafny.Set.fromElements(_662_v37, _662_v37);
          let _664_v39;
          _664_v39 = _module.D9.create_DC26(_662_v37);
          let _665_v40;
          _665_v40 = new _dafny.CodePoint('l'.codePointAt(0));
          let _666_v41;
          _666_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f37)).length),_665_v40);
          let _667_v42;
          _667_v42 = _dafny.Map.Empty.slice().updateUnsafe((_664_v39).dtor_cf38,_666_v41);
          let _668_v44;
          _668_v44 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm41(globalState),(_this).f37);
          let _669_v46;
          let _nw102 = Array((new BigNumber(21)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = _663_v38;
          _nw102[(_dafny.ONE).toNumber()] = _module.__default.fm40(globalState);
          _nw102[(new BigNumber(2)).toNumber()] = (_663_v38).Difference(_663_v38);
          _nw102[(new BigNumber(3)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(4)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(5)).toNumber()] = (((_this).f38) ? (_663_v38) : (_663_v38));
          _nw102[(new BigNumber(6)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_662_v37);
          _nw102[(new BigNumber(8)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(9)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(10)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((_662_v37).update(p1, (_this).f37), _module.__default.fm41(globalState));
          _nw102[(new BigNumber(12)).toNumber()] = function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of (_667_v42).Keys.Elements) {
              let _670_v43 = _compr_35;
              if ((_667_v42).contains(_670_v43)) {
                _coll35.add(_670_v43);
              }
            }
            return _coll35;
          }();
          _nw102[(new BigNumber(13)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(14)).toNumber()] = (_663_v38).Intersect(_663_v38);
          _nw102[(new BigNumber(15)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(16)).toNumber()] = _663_v38;
          _nw102[(new BigNumber(17)).toNumber()] = (_663_v38).Union(function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of (_668_v44).Keys.Elements) {
              let _671_v45 = _compr_36;
              if ((_668_v44).contains(_671_v45)) {
                _coll36.add(_671_v45);
              }
            }
            return _coll36;
          }());
          _nw102[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements(_module.__default.fm41(globalState), _662_v37);
          _nw102[(new BigNumber(19)).toNumber()] = (((_this).f38) ? (_module.__default.fm40(globalState)) : (_663_v38));
          _nw102[(new BigNumber(20)).toNumber()] = _663_v38;
          _669_v46 = _nw102;
          let _index97 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_669_v46).length));
          (_669_v46)[_index97] = _663_v38;
          let _672_v47;
          _672_v47 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_662_v37).length), globalState));
          let _673_v48;
          _673_v48 = _module.D4.create_DC13((_dafny.ZERO).minus((_this).f37), (_this).f38, (_672_v47).IsDisjointFrom(_672_v47), (_this).f38);
          let _674_v49;
          _674_v49 = _dafny.MultiSet.fromElements((_this).f25);
          let _index98 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_669_v46).length));
          let _rhs101 = (((!((_656_v33)[_module.__default.safeIndex((_this).f37, new BigNumber((_656_v33).length))])) && ((_this).f38)) ? ((new BigNumber((_674_v49).cardinality())).multipliedBy((_dafny.ZERO).minus((_this).f37))) : ((_this).f25));
          let _rhs102 = p1;
          let _rhs103 = _663_v38;
          let _rhs104 = !((_this).f38);
          let _rhs105 = _673_v48;
          let _lhs69 = globalState;
          let _lhs70 = _669_v46;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_669_v46).length));
          let _lhs72 = globalState;
          r1 = _rhs101;
          _lhs69.f12 = _rhs102;
          _lhs70[_lhs71] = _rhs103;
          _lhs72.f12 = _rhs104;
          _673_v48 = _rhs105;
          (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xehpc"), _dafny.Seq.Concat(_661_v36, ((_this).f26)[_module.__default.safeIndex(new BigNumber(802), new BigNumber(((_this).f26).length))]));
          let _index99 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((p3).length));
          (p3)[_index99] = (_this).f25;
          let _675_v50;
          _675_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(877),new BigNumber((_656_v33).length));
          let _index100 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((p3).length));
          (p3)[_index100] = (_dafny.ZERO).minus(((_622_v3).fm26(new BigNumber((_661_v36).length), !(true), globalState)).minus((_dafny.ZERO).minus((_622_v3).fm26((_dafny.ZERO).minus((_this).f37), !((_656_v33)[_module.__default.safeIndex((((_675_v50).contains(new BigNumber(107))) ? ((_675_v50).get(new BigNumber(107))) : ((_this).f25)), new BigNumber((_656_v33).length))]), globalState))));
        } else if (_source13.is_DC6) {
          let _676___mcc_h0 = (_source13).cf7;
          let _677___mcc_h1 = (_source13).cf8;
          let _678___mcc_h2 = (_source13).cf9;
          let _679_cf9 = _678___mcc_h2;
          let _680_cf8 = _677___mcc_h1;
          let _681_cf7 = _676___mcc_h0;
          let _682_v51;
          _682_v51 = _module.D5.create_DC16(!(p1));
          _682_v51 = _682_v51;
          r0 = ((_this).f25).plus((_this).f37);
          (globalState).f10 = _679_cf9;
          let _683_v52;
          _683_v52 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-429)), ((_684_cf7) => function (_685_i3) {
            return _684_cf7;
          })(_681_cf7))).length), (_this).f25);
          let _686_v53;
          _686_v53 = _dafny.Set.fromElements((_this).f38);
          let _687_v54;
          _687_v54 = _module.D3.create_DC10((_dafny.Set.fromElements((_this).f37, (_dafny.ZERO).minus(new BigNumber((_686_v53).length)))).IsSubsetOf(_683_v52), _618_v1, _module.__default.fm0(new BigNumber((_686_v53).length), globalState), (new BigNumber(795)).minus((_this).f25), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f37)));
          let _rhs106 = _687_v54;
          let _rhs107 = p1;
          _687_v54 = _rhs106;
          _680_cf8 = _rhs107;
        } else if (_source13.is_DC4) {
          let _688___mcc_h3 = (_source13).cf6;
          let _689_cf6 = _688___mcc_h3;
          (globalState).f12 = ((_this).f37).isLessThan((_this).f37);
          let _690_v55;
          let _nw103 = new _module.C0();
          _nw103.__ctor();
          _690_v55 = _nw103;
          (globalState).f12 = p2;
          r0 = (_dafny.ZERO).minus(((_this).f37).plus((_this).f37));
        } else {
          let _691___mcc_h4 = (_source13).cf10;
          let _692_cf10 = _691___mcc_h4;
          let _693_v56;
          _693_v56 = _dafny.Seq.UnicodeFromString("egqfhrepj");
          let _694_v58;
          _694_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,(_dafny.ZERO).minus((_this).f37));
          let _695_v59;
          _695_v59 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_693_v56).length),(_this).f37), (function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of _dafny.IntegerRange(new BigNumber(134), new BigNumber(936))) {
              let _696_v57 = _compr_37;
              if (((new BigNumber(134)).isLessThanOrEqualTo(_696_v57)) && ((_696_v57).isLessThan(new BigNumber(936)))) {
                _coll37.push([(_696_v57).plus((_this).f25),(_this).f25]);
              }
            }
            return _coll37;
          }()).Merge(_694_v58));
          let _697_v60;
          let _init18 = ((_698_p2) => function (_699_i4) {
            return _module.D0.create_DC3((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f37)).length)), _698_p2);
          })(p2);
          let _nw104 = Array((new BigNumber(12)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw104.length); _i0_18++) {
            _nw104[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _697_v60 = _nw104;
          let _700_v61;
          _700_v61 = _dafny.Set.fromElements(!(p1), (_this).f38, true, p1);
          let _index101 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_697_v60).length));
          (_697_v60)[_index101] = _module.__default.fm42(new BigNumber((_700_v61).length), globalState);
          let _index102 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          (_618_v1)[_index102] = false;
          let _701_v62;
          _701_v62 = _module.D0.create_DC3((_this).f25, p1);
          let _702_v63;
          _702_v63 = _module.D6.create_DC19(_701_v62, (_this).f25);
          let _703_v64;
          _703_v64 = new _dafny.CodePoint('m'.codePointAt(0));
          let _pat_let_tv13 = _701_v62;
          let _index103 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_697_v60).length));
          let _index104 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          let _rhs108 = _695_v59;
          let _rhs109 = _701_v62;
          let _rhs110 = !((_this).f38) || (_dafny.Seq.contains(_656_v33, !((_this).f38)));
          let _rhs111 = function (_pat_let8_0) {
            return function (_704_dt__update__tmp_h1) {
              return function (_pat_let9_0) {
                return function (_705_dt__update_hcf29_h0) {
                  return _module.D6.create_DC19(_705_dt__update_hcf29_h0, (_704_dt__update__tmp_h1).dtor_cf30);
                }(_pat_let9_0);
              }(_pat_let_tv13);
            }(_pat_let8_0);
          }(_702_v63);
          let _rhs112 = _703_v64;
          let _lhs73 = _697_v60;
          let _lhs74 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_697_v60).length));
          let _lhs75 = _618_v1;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          let _lhs77 = globalState;
          _695_v59 = _rhs108;
          _lhs73[_lhs74] = _rhs109;
          _lhs75[_lhs76] = _rhs110;
          _702_v63 = _rhs111;
          _lhs77.f23 = _rhs112;
          let _706_v65;
          let _init19 = ((_707_v33) => function (_708_i5) {
            return _707_v33;
          })(_656_v33);
          let _nw105 = Array((new BigNumber(18)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw105.length); _i0_19++) {
            _nw105[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _706_v65 = _nw105;
          let _709_v66;
          _709_v66 = _dafny.Seq.of(_656_v33, _656_v33);
          let _index105 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_706_v65).length));
          (_706_v65)[_index105] = _dafny.Seq.Concat(_dafny.Seq.of(p2), (_709_v66)[_module.__default.safeIndex((_this).f37, new BigNumber((_709_v66).length))]);
          let _710_v67;
          _710_v67 = _dafny.Map.Empty.slice().updateUnsafe(_703_v64,(_618_v1)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length))]);
          let _711_v69;
          let _nw106 = Array((new BigNumber(13)).toNumber());
          _nw106[(_dafny.ZERO).toNumber()] = (_710_v67).Merge(_710_v67);
          _nw106[(_dafny.ONE).toNumber()] = (_710_v67).Merge((_710_v67).update(new _dafny.CodePoint('u'.codePointAt(0)), p1));
          _nw106[(new BigNumber(2)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(3)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(4)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(5)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(6)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(7)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(8)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(9)).toNumber()] = function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_693_v56).Elements) {
              let _712_v68 = _compr_38;
              if (_dafny.Seq.contains(_693_v56, _712_v68)) {
                _coll38.push([_712_v68,false]);
              }
            }
            return _coll38;
          }();
          _nw106[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_703_v64,p2);
          _nw106[(new BigNumber(11)).toNumber()] = _710_v67;
          _nw106[(new BigNumber(12)).toNumber()] = _710_v67;
          _711_v69 = _nw106;
          let _713_v70;
          _713_v70 = _module.D10.create_DC28(_710_v67);
          let _714_v71;
          _714_v71 = _dafny.Seq.of(_710_v67, _710_v67, (_713_v70).dtor_cf41);
          let _index106 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_711_v69).length));
          (_711_v69)[_index106] = ((_714_v71)[_module.__default.safeIndex((_this).f37, new BigNumber((_714_v71).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe(_703_v64,(_this).f38));
          let _index107 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_706_v65).length));
          let _index108 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          let _index109 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          let _index110 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_711_v69).length));
          let _rhs113 = _656_v33;
          let _rhs114 = ((_this).f37).isEqualTo(((p1) ? ((_this).f37) : ((_this).f25)));
          let _rhs115 = p1;
          let _rhs116 = (_dafny.Map.Empty.slice().updateUnsafe(_703_v64,_module.__default.fm0((_this).f25, globalState))).Merge(_710_v67);
          let _lhs78 = _706_v65;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_706_v65).length));
          let _lhs80 = _618_v1;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          let _lhs82 = _618_v1;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          let _lhs84 = _711_v69;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_711_v69).length));
          _lhs78[_lhs79] = _rhs113;
          _lhs80[_lhs81] = _rhs114;
          _lhs82[_lhs83] = _rhs115;
          _lhs84[_lhs85] = _rhs116;
          let _715_v72;
          _715_v72 = _dafny.Seq.of((_this).f37);
          (globalState).f6 = _dafny.Seq.Concat(_715_v72, _715_v72);
          let _index111 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_618_v1).length));
          (_618_v1)[_index111] = (_656_v33)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f25), new BigNumber((_656_v33).length))];
        }
        let _716_v73;
        let _nw107 = Array((new BigNumber(10)).toNumber()).fill([]);
        _716_v73 = _nw107;
        let _717_v74;
        let _nw108 = Array((new BigNumber(5)).toNumber()).fill(_module.D7.Default());
        _717_v74 = _nw108;
        let _index112 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_716_v73).length));
        (_716_v73)[_index112] = _717_v74;
        let _index113 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_716_v73).length));
        (_716_v73)[_index113] = _717_v74;
        let _718_v75;
        _718_v75 = _dafny.Seq.UnicodeFromString("gkk");
        let _719_v76;
        _719_v76 = _dafny.Seq.of(new BigNumber((_718_v75).length), new BigNumber(-949), (_this).f25, (_this).f37);
        let _720_v77;
        _720_v77 = _719_v76;
        let _source14 = _720_v77;
        let _721___mcc_h5 = _source14;
        let _722_cf11 = _721___mcc_h5;
        (globalState).f12 = p2;
        let _723_v78;
        _723_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,true);
        let _724_v79;
        _724_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,p2);
        let _725_v80;
        _725_v80 = _dafny.Seq.of((_723_v78), (_724_v79).update(p2, true), _724_v79);
        _725_v80 = _dafny.Seq.Concat(_dafny.Seq.Concat(_725_v80, _725_v80), _dafny.Seq.of(_724_v79));
        (globalState).f23 = _module.__default.fm39((_this).f38, globalState);
        (globalState).f12 = p2;
        let _726_v82;
        _726_v82 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),true);
        (globalState).f12 = _module.__default.fm0(new BigNumber((function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of (_726_v82).Keys.Elements) {
            let _727_v81 = _compr_39;
            if ((_726_v82).contains(_727_v81)) {
              _coll39.push([_727_v81,p1]);
            }
          }
          return _coll39;
        }()).length), globalState);
      } else {
        let _728_v83;
        _728_v83 = _dafny.Seq.of(p3);
        (globalState).f12 = _dafny.Seq.IsPrefixOf(_728_v83, _728_v83);
        let _729_v84;
        _729_v84 = new _dafny.CodePoint('g'.codePointAt(0));
        let _730_v85;
        _730_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_729_v84);
        let _731_v86;
        _731_v86 = _dafny.Map.Empty.slice().updateUnsafe(_730_v85,p2);
        let _732_v87;
        _732_v87 = _dafny.Seq.UnicodeFromString("bh");
        let _733_v88;
        _733_v88 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f37, (_this).f37, (_this).f25)).length), (_this).f37);
        let _734_v89;
        _734_v89 = _module.D7.create_DC21(p3);
        let _735_v90;
        _735_v90 = _dafny.Map.Empty.slice().updateUnsafe(_734_v89,(_this).f37);
        let _736_v91;
        _736_v91 = _module.D3.create_DC10((_this).f38, _618_v1, _module.__default.fm0(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus((_this).f25))).length), globalState), new BigNumber(939), new BigNumber(-840));
        let _737_v92;
        _737_v92 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
        let _738_v93;
        _738_v93 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_733_v88)[_module.__default.safeIndex(new BigNumber((_732_v87).length), new BigNumber((_733_v88).length))]);
        let _739_v94;
        _739_v94 = _dafny.Seq.of(p2, false);
        let _740_v95;
        _740_v95 = _dafny.Set.fromElements(p1);
        let _741_v96;
        let _nw109 = Array((new BigNumber(27)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = (_this).f37;
        _nw109[(_dafny.ONE).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt((_this).f25, new BigNumber(-814));
        _nw109[(new BigNumber(3)).toNumber()] = new BigNumber(-329);
        _nw109[(new BigNumber(4)).toNumber()] = new BigNumber(((_731_v86).Merge(_731_v86)).length);
        _nw109[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_732_v87).length)).multipliedBy((_this).f25));
        _nw109[(new BigNumber(6)).toNumber()] = ((_this).f25).multipliedBy((_this).f37);
        _nw109[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_this).f25, new BigNumber(110));
        _nw109[(new BigNumber(8)).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(9)).toNumber()] = (_622_v3).fm26((_this).f37, p1, globalState);
        _nw109[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(11)).toNumber()] = (_733_v88)[_module.__default.safeIndex((((_735_v90).contains(_734_v89)) ? ((_735_v90).get(_734_v89)) : ((_this).f25)), new BigNumber((_733_v88).length))];
        _nw109[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus((_736_v91).dtor_cf17)), (_this).f25);
        _nw109[(new BigNumber(13)).toNumber()] = new BigNumber(662);
        _nw109[(new BigNumber(14)).toNumber()] = new BigNumber((_732_v87).length);
        _nw109[(new BigNumber(15)).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(16)).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(17)).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(18)).toNumber()] = (_this).f25;
        _nw109[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f25, (_this).f25));
        _nw109[(new BigNumber(20)).toNumber()] = new BigNumber(40);
        _nw109[(new BigNumber(21)).toNumber()] = new BigNumber((_737_v92).length);
        _nw109[(new BigNumber(22)).toNumber()] = (_622_v3).fm26((_this).f25, (_this).f38, globalState);
        _nw109[(new BigNumber(23)).toNumber()] = (_this).f37;
        _nw109[(new BigNumber(24)).toNumber()] = ((_622_v3).fm26((_this).f25, p1, globalState)).multipliedBy((_this).f25);
        _nw109[(new BigNumber(25)).toNumber()] = (((_738_v93).contains(p2)) ? ((_738_v93).get(p2)) : (new BigNumber((_739_v94).length)));
        _nw109[(new BigNumber(26)).toNumber()] = new BigNumber((_740_v95).length);
        _741_v96 = _nw109;
        let _nw110 = Array((new BigNumber(10)).toNumber());
        _nw110[(_dafny.ZERO).toNumber()] = (_this).f37;
        _nw110[(_dafny.ONE).toNumber()] = new BigNumber(-890);
        _nw110[(new BigNumber(2)).toNumber()] = (_this).f37;
        _nw110[(new BigNumber(3)).toNumber()] = (_this).f37;
        _nw110[(new BigNumber(4)).toNumber()] = (_622_v3).fm26((_this).f37, true, globalState);
        _nw110[(new BigNumber(5)).toNumber()] = (_this).f25;
        _nw110[(new BigNumber(6)).toNumber()] = (_this).f37;
        _nw110[(new BigNumber(7)).toNumber()] = (_this).f25;
        _nw110[(new BigNumber(8)).toNumber()] = (new BigNumber((_733_v88).length)).minus((_this).f25);
        _nw110[(new BigNumber(9)).toNumber()] = (_this).f37;
        _741_v96 = _nw110;
        let _742_v97;
        _742_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,(_this).f38);
        let _743_v98;
        _743_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_729_v84);
        let _rhs117 = (((_742_v97).contains((_this).f25)) ? ((_742_v97).get((_this).f25)) : ((new BigNumber((_732_v87).length)).isLessThan(new BigNumber((_743_v98).length))));
        let _rhs118 = (_this).f38;
        let _lhs86 = globalState;
        let _lhs87 = globalState;
        _lhs86.f12 = _rhs117;
        _lhs87.f12 = _rhs118;
        let _index114 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((p3).length));
        (p3)[_index114] = new BigNumber((_740_v95).length);
        let _index115 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((p3).length));
        let _rhs119 = new BigNumber((_dafny.Seq.Concat(_732_v87, _dafny.Seq.UnicodeFromString("rxowgeyge"))).length);
        let _rhs120 = p2;
        let _rhs121 = (_this).f25;
        let _lhs88 = p3;
        let _lhs89 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((p3).length));
        let _lhs90 = globalState;
        _lhs88[_lhs89] = _rhs119;
        _lhs90.f12 = _rhs120;
        r0 = _rhs121;
        (globalState).f23 = _729_v84;
      }
      let _744_v99;
      _744_v99 = new _dafny.CodePoint('l'.codePointAt(0));
      let _745_v100;
      _745_v100 = _dafny.Seq.UnicodeFromString("ymxtrxics");
      let _746_v101;
      _746_v101 = _dafny.Set.fromElements(p2);
      (globalState).f12 = !(_dafny.Seq.contains(_745_v100, _744_v99)) || ((_746_v101).IsProperSubsetOf(_746_v101));
      r0 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm1(false, globalState)).length));
      r1 = new BigNumber(-711);
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _747_v0;
      _747_v0 = _dafny.Seq.of((_this).f37);
      let _748_v1;
      _748_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
      let _749_v2;
      _749_v2 = _dafny.Seq.of(new BigNumber((_747_v0).length), (_this).f25, new BigNumber((_748_v1).length), (_this).f37);
      let _750_v3;
      _750_v3 = _dafny.Seq.UnicodeFromString("olbsqy");
      let _751_v4;
      _751_v4 = new _dafny.CodePoint('b'.codePointAt(0));
      let _752_v5;
      _752_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
      let _753_v6;
      _753_v6 = _dafny.Seq.of(p1);
      let _754_v7;
      let _nw111 = Array((new BigNumber(20)).toNumber());
      _nw111[(_dafny.ZERO).toNumber()] = new BigNumber((_750_v3).length);
      _nw111[(_dafny.ONE).toNumber()] = (_this).f37;
      _nw111[(new BigNumber(2)).toNumber()] = new BigNumber(801);
      _nw111[(new BigNumber(3)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(4)).toNumber()] = (_this).f37;
      _nw111[(new BigNumber(5)).toNumber()] = new BigNumber(355);
      _nw111[(new BigNumber(6)).toNumber()] = (_this).f37;
      _nw111[(new BigNumber(7)).toNumber()] = (_this).f37;
      _nw111[(new BigNumber(8)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(9)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(10)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(11)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(12)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(13)).toNumber()] = (_this).f37;
      _nw111[(new BigNumber(14)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(15)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_747_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_750_v3, _module.__default.safeIndex((_this).f25, new BigNumber((_750_v3).length)), _751_v4), _module.__default.safeIndex(new BigNumber((_752_v5).length), new BigNumber((_dafny.Seq.update(_750_v3, _module.__default.safeIndex((_this).f25, new BigNumber((_750_v3).length)), _751_v4)).length)), _751_v4)).length), new BigNumber((_747_v0).length))]);
      _nw111[(new BigNumber(17)).toNumber()] = new BigNumber((_753_v6).length);
      _nw111[(new BigNumber(18)).toNumber()] = (_this).f25;
      _nw111[(new BigNumber(19)).toNumber()] = (_this).f25;
      _754_v7 = _nw111;
      let _755_v8;
      _755_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,_754_v7);
      let _756_v9;
      _756_v9 = _dafny.Set.fromElements(!((_this).f38), p1, p1, (_this).f38, (_this).f38);
      let _757_v10;
      _757_v10 = _dafny.MultiSet.fromElements(_module.__default.fm2((_this).f25, (_this).f25, globalState), _dafny.Seq.UnicodeFromString("f"));
      let _758_v11;
      _758_v11 = _dafny.MultiSet.fromElements(false);
      let _759_v12;
      _759_v12 = _dafny.Seq.of(_758_v11, _758_v11, _758_v11, _dafny.MultiSet.FromArray(_753_v6), _758_v11);
      let _760_v13;
      _760_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_759_v12);
      let _761_v14;
      let _nw112 = Array((new BigNumber(20)).toNumber());
      _nw112[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).f37);
      _nw112[(_dafny.ONE).toNumber()] = (_this).f37;
      _nw112[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_747_v0, _dafny.Seq.of((_this).f37, (_this).f25))).length);
      _nw112[(new BigNumber(3)).toNumber()] = (_this).f37;
      _nw112[(new BigNumber(4)).toNumber()] = (((_this).f38) ? (new BigNumber(666)) : ((_this).f37));
      _nw112[(new BigNumber(5)).toNumber()] = ((p1) ? ((_this).f25) : ((_this).f25));
      _nw112[(new BigNumber(6)).toNumber()] = ((_this).f25).multipliedBy(new BigNumber((_755_v8).length));
      _nw112[(new BigNumber(7)).toNumber()] = new BigNumber((_756_v9).length);
      _nw112[(new BigNumber(8)).toNumber()] = new BigNumber((_748_v1).length);
      _nw112[(new BigNumber(9)).toNumber()] = (_this).f37;
      _nw112[(new BigNumber(10)).toNumber()] = (_this).f25;
      _nw112[(new BigNumber(11)).toNumber()] = (new BigNumber(685)).plus((_this).f37);
      _nw112[(new BigNumber(12)).toNumber()] = (_this).f25;
      _nw112[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(616), new BigNumber(-96));
      _nw112[(new BigNumber(14)).toNumber()] = ((_this).f37).minus((_dafny.ZERO).minus(new BigNumber((_757_v10).cardinality())));
      _nw112[(new BigNumber(15)).toNumber()] = ((_this).f37).minus((_this).f25);
      _nw112[(new BigNumber(16)).toNumber()] = (new BigNumber(979)).plus((_this).f25);
      _nw112[(new BigNumber(17)).toNumber()] = (_this).f37;
      _nw112[(new BigNumber(18)).toNumber()] = (_this).f25;
      _nw112[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.Concat((((_760_v13).contains((_this).f38)) ? ((_760_v13).get((_this).f38)) : (_759_v12)), _dafny.Seq.of(_758_v11, _758_v11))).length);
      _761_v14 = _nw112;
      let _index116 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
      (_754_v7)[_index116] = new BigNumber(764);
      let _index117 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
      (_754_v7)[_index117] = new BigNumber(919);
      let _index118 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
      let _rhs122 = new BigNumber((_dafny.Seq.Concat(_750_v3, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), ((_762_v4) => function (_763_i0) {
        return _762_v4;
      })(_751_v4)), _750_v3))).length);
      let _rhs123 = (_this).f37;
      let _lhs91 = _754_v7;
      let _lhs92 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
      r0 = _rhs122;
      _lhs91[_lhs92] = _rhs123;
      if (_module.__default.fm0(_module.__default.safeModuloInt((_this).f37, (_this).f25), globalState)) {
        let _source15 = _749_v2;
        let _764___mcc_h0 = _source15;
        let _765_cf11 = _764___mcc_h0;
        let _766_v15;
        _766_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,p1);
        let _767_v16;
        _767_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_766_v15);
        let _768_v17;
        _768_v17 = _dafny.Map.Empty.slice().updateUnsafe(_754_v7,_dafny.Map.Empty.slice().updateUnsafe((_this).f38,(_this).f38));
        let _769_v18;
        _769_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm44(p1, (_this).f38, p1, new BigNumber((_750_v3).length), globalState),_766_v15);
        let _770_v19;
        _770_v19 = _dafny.Seq.of(_766_v15);
        let _771_v20;
        let _nw113 = Array((new BigNumber(26)).toNumber());
        _nw113[(_dafny.ZERO).toNumber()] = _766_v15;
        _nw113[(_dafny.ONE).toNumber()] = (((_767_v16).contains(p1)) ? ((_767_v16).get(p1)) : (_766_v15));
        _nw113[(new BigNumber(2)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(3)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(4)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(5)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        _nw113[(new BigNumber(7)).toNumber()] = (((_768_v17).contains(_761_v14)) ? ((_768_v17).get(_761_v14)) : ((_dafny.Map.Empty.slice().updateUnsafe(p1,p1))));
        _nw113[(new BigNumber(8)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(9)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f38);
        _nw113[(new BigNumber(11)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(12)).toNumber()] = (_766_v15).update((_753_v6)[_module.__default.safeIndex(new BigNumber(-969), new BigNumber((_753_v6).length))], p1);
        _nw113[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f38);
        _nw113[(new BigNumber(14)).toNumber()] = (_766_v15).Merge(_766_v15);
        _nw113[(new BigNumber(15)).toNumber()] = (_766_v15).Merge(_766_v15);
        _nw113[(new BigNumber(16)).toNumber()] = (_module.__default.fm43((_this).f38, (_this).f38, globalState)).Merge((_766_v15).update((_this).f38, (_this).f38));
        _nw113[(new BigNumber(17)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(18)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(19)).toNumber()] = _766_v15;
        _nw113[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,p1);
        _nw113[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f38),p1);
        _nw113[(new BigNumber(22)).toNumber()] = ((((_769_v18).contains(_module.__default.fm44(p1, (_this).f38, p1, (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))], globalState))) ? ((_769_v18).get(_module.__default.fm44(p1, (_this).f38, p1, (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))], globalState))) : (_dafny.Map.Empty.slice().updateUnsafe((((_748_v1).contains((_this).f37)) ? ((_748_v1).get((_this).f37)) : (!(p1))),(_this).f38)))).Merge((_766_v15).update(p1, (_this).f38));
        _nw113[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f37, globalState),(_this).f38);
        _nw113[(new BigNumber(24)).toNumber()] = (_770_v19)[_module.__default.safeIndex((_this).f25, new BigNumber((_770_v19).length))];
        _nw113[(new BigNumber(25)).toNumber()] = _766_v15;
        _771_v20 = _nw113;
        let _index119 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_771_v20).length));
        (_771_v20)[_index119] = (_766_v15).Merge(_766_v15);
        let _772_v21;
        let _init20 = ((_773_p1) => function (_774_i1) {
          return _773_p1;
        })(p1);
        let _nw114 = Array((new BigNumber(10)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw114.length); _i0_20++) {
          _nw114[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _772_v21 = _nw114;
        let _775_v22;
        _775_v22 = _module.D3.create_DC10((_this).f38, _772_v21, p1, new BigNumber(43), (_this).f37);
        let _776_v23;
        _776_v23 = _dafny.Map.Empty.slice().updateUnsafe(_772_v21,_module.D3.create_DC10(p1, _772_v21, (_775_v22).dtor_cf13, new BigNumber((_module.__default.fm1((_this).f38, globalState)).length), _module.__default.fm44(!(false), p1, (_this).f38, (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))], globalState)));
        let _index120 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_771_v20).length));
        let _index121 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
        let _rhs124 = _766_v15;
        let _rhs125 = new BigNumber((_776_v23).length);
        let _lhs93 = _771_v20;
        let _lhs94 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_771_v20).length));
        let _lhs95 = _754_v7;
        let _lhs96 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
        _lhs93[_lhs94] = _rhs124;
        _lhs95[_lhs96] = _rhs125;
        let _777_v24;
        _777_v24 = _dafny.MultiSet.fromElements((_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]);
        (globalState).f12 = (_777_v24).contains(new BigNumber(-576));
        _748_v1 = _748_v1;
        let _778_v25;
        let _nw115 = new _module.C2();
        _nw115.__ctor();
        _778_v25 = _nw115;
        (globalState).f12 = _module.__default.fm0((_this).f25, globalState);
        let _779_v26;
        _779_v26 = _dafny.Set.fromElements(new BigNumber((_747_v0).length), new BigNumber((_750_v3).length));
        let _780_v27;
        _780_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_779_v26).Union(_779_v26)).length),(_this).f25);
        _780_v27 = (_780_v27).update(new BigNumber(541), ((_this).f37).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements((_this).f25)).cardinality())));
        _780_v27 = (_780_v27).update((_this).f25, (_this).f25);
        let _781_v28;
        _781_v28 = _dafny.Map.Empty.slice().updateUnsafe(_750_v3,_dafny.Seq.UnicodeFromString("dnufifda"));
        _781_v28 = (_781_v28).update(_dafny.Seq.update(_750_v3, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f37)), new BigNumber((_750_v3).length)), _751_v4), _750_v3);
      } else {
        let _782_v29;
        _782_v29 = _dafny.Seq.of(_752_v5, _752_v5, _752_v5);
        _782_v29 = _782_v29;
        let _783_v30;
        let _nw116 = Array((new BigNumber(9)).toNumber()).fill(_module.D5.Default());
        _783_v30 = _nw116;
        let _784_v31;
        _784_v31 = _module.D5.create_DC16((_this).f38);
        let _pat_let_tv14 = p1;
        let _index122 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_783_v30).length));
        (_783_v30)[_index122] = function (_pat_let10_0) {
          return function (_785_dt__update__tmp_h0) {
            return function (_pat_let11_0) {
              return function (_786_dt__update_hcf25_h0) {
                return _module.D5.create_DC16(_786_dt__update_hcf25_h0);
              }(_pat_let11_0);
            }(_pat_let_tv14);
          }(_pat_let10_0);
        }(_784_v31);
        let _index123 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_783_v30).length));
        (_783_v30)[_index123] = _module.__default.fm45(new BigNumber(637), _module.__default.fm44((_this).f38, p1, p1, (_this).f37, globalState), globalState);
        let _787_v32;
        _787_v32 = _module.D0.create_DC3((_this).f37, p1);
        let _pat_let_tv15 = _754_v7;
        let _pat_let_tv16 = _754_v7;
        let _788_v33;
        _788_v33 = _module.D6.create_DC19(function (_pat_let12_0) {
  return function (_789_dt__update__tmp_h1) {
    return function (_pat_let13_0) {
      return function (_790_dt__update_hcf4_h0) {
        return _module.D0.create_DC3(_790_dt__update_hcf4_h0, (_789_dt__update__tmp_h1).dtor_cf5);
      }(_pat_let13_0);
    }((_pat_let_tv16)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_pat_let_tv15).length))]);
  }(_pat_let12_0);
}(_787_v32), (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]);
        let _791_v34;
        let _nw117 = Array((new BigNumber(25)).toNumber()).fill(_module.D6.Default());
        _791_v34 = _nw117;
        let _pat_let_tv17 = _787_v32;
        let _792_v35;
        _792_v35 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let14_0) {
          return function (_793_dt__update__tmp_h2) {
            return function (_pat_let15_0) {
              return function (_794_dt__update_hcf29_h0) {
                return _module.D6.create_DC19(_794_dt__update_hcf29_h0, (_793_dt__update__tmp_h2).dtor_cf30);
              }(_pat_let15_0);
            }(_pat_let_tv17);
          }(_pat_let14_0);
        }(_788_v33),_791_v34);
        let _795_v36;
        _795_v36 = _module.D12.create_DC31(_792_v35);
        let _index124 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
        (_754_v7)[_index124] = new BigNumber((((_795_v36).dtor_cf46).update(_788_v33, _791_v34)).length);
        (globalState).f12 = _module.__default.fm0((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,p1)).length), _module.__default.fm44(p1, true, (_this).f38, (_this).f37, globalState))), globalState);
        _750_v3 = _dafny.Seq.Concat(_750_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(255)), function (_796_i2) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }));
      }
      let _797_v37;
      _797_v37 = _module.D9.create_DC27((_this).f37, (_dafny.ZERO).minus(new BigNumber(-605)));
      let _pat_let_tv18 = _754_v7;
      let _pat_let_tv19 = _754_v7;
      let _source16 = function (_pat_let16_0) {
        return function (_798_dt__update__tmp_h3) {
          return function (_pat_let17_0) {
            return function (_799_dt__update_hcf40_h0) {
              return _module.D9.create_DC27((_798_dt__update__tmp_h3).dtor_cf39, _799_dt__update_hcf40_h0);
            }(_pat_let17_0);
          }(((_dafny.ZERO).minus((_this).f37)).plus((_pat_let_tv19)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_pat_let_tv18).length))]));
        }(_pat_let16_0);
      }(_797_v37);
      if (_source16.is_DC27) {
        let _800___mcc_h1 = (_source16).cf39;
        let _801___mcc_h2 = (_source16).cf40;
        let _802_cf40 = _801___mcc_h2;
        let _803_cf39 = _800___mcc_h1;
        let _index125 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
        let _rhs126 = (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))];
        let _rhs127 = (_dafny.ZERO).minus(_803_cf39);
        let _rhs128 = (_this).f38;
        let _rhs129 = _module.__default.fm2((_802_cf40).multipliedBy((_dafny.ZERO).minus((_this).f25)), (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))], globalState);
        let _lhs97 = _754_v7;
        let _lhs98 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
        let _lhs99 = globalState;
        let _lhs100 = globalState;
        _802_cf40 = _rhs126;
        _lhs97[_lhs98] = _rhs127;
        _lhs99.f12 = _rhs128;
        _lhs100.f10 = _rhs129;
        let _804_v38;
        let _init21 = ((_805_v3) => function (_806_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_805_v3).length),(_this).f25);
        })(_750_v3);
        let _nw118 = Array((new BigNumber(29)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw118.length); _i0_21++) {
          _nw118[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _804_v38 = _nw118;
        let _807_v39;
        _807_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(464),(_this).f25);
        let _index126 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_804_v38).length));
        (_804_v38)[_index126] = _807_v39;
        let _808_v40;
        _808_v40 = _dafny.Map.Empty.slice().updateUnsafe(p1,_807_v39);
        let _index127 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_804_v38).length));
        (_804_v38)[_index127] = (_dafny.Map.Empty.slice().updateUnsafe(_803_cf39,new BigNumber((_808_v40).length))).Merge(_807_v39);
        let _809_v41;
        _809_v41 = _module.D1.create_DC6(_751_v4, (_this).f38, _750_v3);
        let _810_v42;
        _810_v42 = _module.__default.fm43((((_748_v1).contains((_this).f25)) ? ((_748_v1).get((_this).f25)) : ((_809_v41).dtor_cf8)), (_this).f38, globalState);
        let _811_v43;
        let _nw119 = Array((new BigNumber(7)).toNumber()).fill(false);
        _811_v43 = _nw119;
        let _812_v44;
        _812_v44 = _dafny.Seq.of(_811_v43, _811_v43, _811_v43);
        let _813_v45;
        _813_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(482),_811_v43);
        let _814_v46;
        _814_v46 = _dafny.Map.Empty.slice().updateUnsafe(_751_v4,_811_v43);
        let _815_v47;
        let _nw120 = Array((new BigNumber(29)).toNumber());
        _nw120[(_dafny.ZERO).toNumber()] = _811_v43;
        _nw120[(_dafny.ONE).toNumber()] = (_812_v44)[_module.__default.safeIndex(_803_cf39, new BigNumber((_812_v44).length))];
        _nw120[(new BigNumber(2)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(3)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(4)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(5)).toNumber()] = ((!(p1)) ? ((((_813_v45).contains((_this).f37)) ? ((_813_v45).get((_this).f37)) : (_811_v43))) : (_811_v43));
        _nw120[(new BigNumber(6)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(7)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(8)).toNumber()] = (((_814_v46).contains(_751_v4)) ? ((_814_v46).get(_751_v4)) : (_811_v43));
        _nw120[(new BigNumber(9)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(10)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(11)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(12)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(13)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(14)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(15)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(16)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(17)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(18)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(19)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(20)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(21)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(22)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(23)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(24)).toNumber()] = _811_v43;
        _nw120[(new BigNumber(25)).toNumber()] = ((p1) ? (_811_v43) : (_811_v43));
        _nw120[(new BigNumber(26)).toNumber()] = (_812_v44)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_812_v44).length))];
        _nw120[(new BigNumber(27)).toNumber()] = (((_813_v45).contains(_803_cf39)) ? ((_813_v45).get(_803_cf39)) : ((_812_v44)[_module.__default.safeIndex((_this).f37, new BigNumber((_812_v44).length))]));
        _nw120[(new BigNumber(28)).toNumber()] = _811_v43;
        _815_v47 = _nw120;
        let _index128 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_815_v47).length));
        (_815_v47)[_index128] = _811_v43;
        let _816_v48;
        let _nw121 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _816_v48 = _nw121;
        let _817_v49;
        _817_v49 = _dafny.Seq.of(_748_v1, _748_v1, _748_v1, _748_v1);
        let _index129 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_816_v48).length));
        (_816_v48)[_index129] = (_817_v49)[_module.__default.safeIndex((_this).f25, new BigNumber((_817_v49).length))];
        let _818_v50;
        _818_v50 = _dafny.MultiSet.fromElements((((_755_v8).contains((_this).f38)) ? ((_755_v8).get((_this).f38)) : (_761_v14)), _761_v14);
        let _index130 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_815_v47).length));
        let _index131 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_816_v48).length));
        let _rhs130 = _810_v42;
        let _rhs131 = (((_818_v50).contains(_761_v14)) ? ((_818_v50).get(_761_v14)) : ((_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]));
        let _rhs132 = _module.__default.safeDivisionInt(new BigNumber(904), (_this).f37);
        let _rhs133 = ((_dafny.Seq.IsPrefixOf(_753_v6, _dafny.Seq.of(p1))) ? (_811_v43) : (_811_v43));
        let _rhs134 = (_817_v49)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_802_cf40, (_this).f37), new BigNumber((_817_v49).length))];
        let _lhs101 = _815_v47;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_815_v47).length));
        let _lhs103 = _816_v48;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_816_v48).length));
        _810_v42 = _rhs130;
        _802_cf40 = _rhs131;
        _802_cf40 = _rhs132;
        _lhs101[_lhs102] = _rhs133;
        _lhs103[_lhs104] = _rhs134;
        _802_cf40 = (_this).f25;
      } else {
        let _819___mcc_h3 = (_source16).cf38;
        let _820_cf38 = _819___mcc_h3;
        (globalState).f12 = p1;
        if (p1) {
          let _821_v51;
          _821_v51 = _module.D10.create_DC29(p1, p1, (((_this).f38) ? ((_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]) : ((_this).f37)));
          _821_v51 = _module.__default.fm46(((false) ? ((_dafny.ZERO).minus((_this).f37)) : ((_this).f37)), (_this).f38, !((((_this).f38) ? ((_this).f38) : (false))), globalState);
          let _822_v52;
          let _nw122 = new _module.C2();
          _nw122.__ctor();
          _822_v52 = _nw122;
          let _823_v53;
          _823_v53 = _module.D0.create_DC3(new BigNumber(411), p1);
          r0 = (_823_v53).dtor_cf4;
          let _824_v54;
          let _init22 = function (_825_i4) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          };
          let _nw123 = Array((new BigNumber(11)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw123.length); _i0_22++) {
            _nw123[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _824_v54 = _nw123;
          let _index132 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_824_v54).length));
          (_824_v54)[_index132] = _module.__default.fm39(p1, globalState);
          let _826_v55;
          let _init23 = ((_827_v2) => function (_828_i5) {
            return _dafny.Seq.Concat(_827_v2, _827_v2);
          })(_749_v2);
          let _nw124 = Array((new BigNumber(12)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw124.length); _i0_23++) {
            _nw124[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _826_v55 = _nw124;
          let _index133 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_826_v55).length));
          (_826_v55)[_index133] = _dafny.Seq.of((_822_v52).fm24((_this).f25, globalState));
          let _index134 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_824_v54).length));
          let _index135 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_826_v55).length));
          let _rhs135 = _751_v4;
          let _rhs136 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), function (_829_i6) {
            return (_this).f37;
          }), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), function (_830_i6) {
            return (_this).f37;
          })).length)), ((_this).f25).plus((_this).f25));
          let _lhs105 = _824_v54;
          let _lhs106 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_824_v54).length));
          let _lhs107 = _826_v55;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_826_v55).length));
          _lhs105[_lhs106] = _rhs135;
          _lhs107[_lhs108] = _rhs136;
          (globalState).f12 = p1;
        } else {
          _753_v6 = _753_v6;
          let _index136 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length));
          (_754_v7)[_index136] = new BigNumber(-300);
          let _831_v56;
          _831_v56 = _dafny.MultiSet.fromElements((_this).f37);
          let _832_v57;
          _832_v57 = _dafny.Set.fromElements((_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]);
          let _833_v58;
          _833_v58 = _dafny.Seq.of((_831_v56).update((_this).f25, _module.__default.abs(new BigNumber((_752_v5).length))), _831_v56, _dafny.MultiSet.fromElements(new BigNumber((_832_v57).length)), _831_v56, _831_v56);
          let _834_v59;
          _834_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(613),_751_v4);
          let _835_v60;
          _835_v60 = _dafny.Set.fromElements(new BigNumber((_833_v58).length), new BigNumber((_834_v59).length), (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))], (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]);
          let _836_v61;
          let _init24 = ((_837_p1) => function (_838_i7) {
            return _837_p1;
          })(p1);
          let _nw125 = Array((new BigNumber(4)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw125.length); _i0_24++) {
            _nw125[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _836_v61 = _nw125;
          let _839_v62;
          _839_v62 = _dafny.Map.Empty.slice().updateUnsafe(_836_v61,_835_v60);
          let _840_v63;
          _840_v63 = _dafny.Seq.of(_836_v61, _836_v61);
          let _rhs137 = (((_839_v62).contains(((_module.__default.fm0((_this).f25, globalState)) ? ((_840_v63)[_module.__default.safeIndex((_this).f37, new BigNumber((_840_v63).length))]) : (_836_v61)))) ? ((_839_v62).get(((_module.__default.fm0((_this).f25, globalState)) ? ((_840_v63)[_module.__default.safeIndex((_this).f37, new BigNumber((_840_v63).length))]) : (_836_v61)))) : (_832_v57));
          let _rhs138 = (_dafny.Set.fromElements(new BigNumber((_747_v0).length), new BigNumber((_dafny.Set.fromElements((_this).f38)).length), (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))])).Intersect(_835_v60);
          _835_v60 = _rhs137;
          _832_v57 = _rhs138;
          let _841_v64;
          _841_v64 = _dafny.Seq.of(_761_v14, _754_v7, _761_v14, _754_v7, _754_v7);
          _841_v64 = _dafny.Seq.of(_754_v7, _761_v14, _761_v14, _761_v14);
          (globalState).f12 = (_this).f38;
        }
        let _842_v65;
        _842_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        r0 = _module.__default.fm44(((_842_v65).update(p1, false)).equals(_842_v65), p1, p1, (_this).f37, globalState);
        _748_v1 = (_748_v1).update(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f25), new BigNumber((_758_v11).cardinality())), (_this).f38);
      }
      let _843_v66;
      _843_v66 = _dafny.Seq.of(_750_v3, _750_v3, _dafny.Seq.UnicodeFromString("fjk"), _dafny.Seq.UnicodeFromString("igt"));
      let _844_v67;
      _844_v67 = _module.D4.create_DC13(new BigNumber((_843_v66).length), p1, _module.__default.fm0((_749_v2)[_module.__default.safeIndex(new BigNumber(-752), new BigNumber((_749_v2).length))], globalState), p1);
      let _845_v68;
      _845_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f25, (_this).f25),(_844_v67).dtor_cf19);
      _845_v68 = (_845_v68).update(new BigNumber((_module.__default.fm1(true, globalState)).length), (_754_v7)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_754_v7).length))]);
      (globalState).f12 = p1;
      r0 = (new BigNumber(-59)).plus((_this).f25);
      return r0;
    }
    m18(p0, p1, globalState) {
      let _this = this;
      let _846_i0;
      _846_i0 = _dafny.ZERO;
      L7: {
        while ((_this).f38) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_846_i0)) {
              break L7;
            }
            _846_i0 = (_846_i0).plus(_dafny.ONE);
            let _847_v0;
            _847_v0 = new _dafny.CodePoint('g'.codePointAt(0));
            (globalState).f23 = _847_v0;
            let _848_v1;
            _848_v1 = new BigNumber(880);
            _848_v1 = (new BigNumber(147)).multipliedBy(_848_v1);
            let _849_v2;
            _849_v2 = _module.D12.create_DC32(p0);
            _849_v2 = _849_v2;
            if ((_this).f38) {
              let _850_v3;
              _850_v3 = _dafny.MultiSet.fromElements(_847_v0);
              let _851_v4;
              _851_v4 = _dafny.Seq.UnicodeFromString("bcea");
              let _852_v5;
              _852_v5 = _dafny.Seq.of(_851_v4, _851_v4, _851_v4);
              let _853_v6;
              _853_v6 = _dafny.Set.fromElements((_this).f37, _848_v1, (_this).f37, new BigNumber(579));
              let _854_v7;
              _854_v7 = _dafny.Seq.of(new BigNumber((_851_v4).length));
              let _855_v8;
              _855_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_854_v7).length),_848_v1);
              let _856_v9;
              _856_v9 = _dafny.Map.Empty.slice().updateUnsafe(_851_v4,_module.__default.fm47(p1, _module.__default.fm44(p0, p0, true, _848_v1, globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f37,new BigNumber((_855_v8).length))).length), _847_v0, globalState));
              let _857_v10;
              _857_v10 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f25));
              let _858_v11;
              _858_v11 = _dafny.Seq.of((_this).f38);
              let _859_v12;
              _859_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_858_v11).length));
              let _rhs139 = ((!(p1)) ? (_dafny.MultiSet.fromElements(_847_v0, new _dafny.CodePoint('f'.codePointAt(0)))) : (_850_v3));
              let _rhs140 = _dafny.Seq.Concat(_852_v5, _852_v5);
              let _rhs141 = new BigNumber(829);
              let _rhs142 = (_853_v6).Intersect((((_856_v9).contains(_851_v4)) ? ((_856_v9).get(_851_v4)) : (_853_v6)));
              let _rhs143 = _module.__default.fm44((_857_v10).IsSubsetOf(_857_v10), _module.__default.fm0(_module.__default.fm44((_this).f38, p0, p0, (((_859_v12).contains(p0)) ? ((_859_v12).get(p0)) : (new BigNumber(993))), globalState), globalState), p0, (_this).f25, globalState);
              _850_v3 = _rhs139;
              _852_v5 = _rhs140;
              _848_v1 = _rhs141;
              _853_v6 = _rhs142;
              _848_v1 = _rhs143;
              (globalState).f6 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-146)), function (_860_i1) {
                return (_this).f37;
              });
              let _861_v13;
              _861_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)))));
              let _862_v14;
              _862_v14 = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,p0);
              let _863_v15;
              let _nw126 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
              _863_v15 = _nw126;
              let _864_v16;
              _864_v16 = _module.D0.create_DC3(new BigNumber((_dafny.Seq.of(_863_v15)).length), (_this).f38);
              let _865_v17;
              let _nw127 = Array((new BigNumber(23)).toNumber());
              _nw127[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_850_v3);
              _nw127[(_dafny.ONE).toNumber()] = (_861_v13).Merge((_861_v13).update((_this).f38, _850_v3));
              _nw127[(new BigNumber(2)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(3)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(4)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(5)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f38,_850_v3)).Merge(_861_v13);
              _nw127[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_850_v3);
              _nw127[(new BigNumber(8)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(9)).toNumber()] = (((_this).f38) ? (_861_v13) : (_861_v13));
              _nw127[(new BigNumber(10)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(11)).toNumber()] = _module.__default.fm48((_this).f38, _848_v1, _848_v1, new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_851_v4, _module.__default.safeIndex(new BigNumber((_862_v14).length), new BigNumber((_851_v4).length)), _847_v0), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.update(_851_v4, _module.__default.safeIndex(new BigNumber((_862_v14).length), new BigNumber((_851_v4).length)), _847_v0)).length)), _847_v0)).length), globalState);
              _nw127[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_850_v3);
              _nw127[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,(_850_v3).update(new _dafny.CodePoint('j'.codePointAt(0)), _module.__default.abs(_848_v1)));
              _nw127[(new BigNumber(14)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(15)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(16)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(17)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(18)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(19)).toNumber()] = (_861_v13).update((_864_v16).dtor_cf5, (_850_v3).update(_847_v0, _module.__default.abs((_this).f37)));
              _nw127[(new BigNumber(20)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(21)).toNumber()] = _861_v13;
              _nw127[(new BigNumber(22)).toNumber()] = _861_v13;
              _865_v17 = _nw127;
              let _index137 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_865_v17).length));
              (_865_v17)[_index137] = _861_v13;
              let _index138 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_865_v17).length));
              (_865_v17)[_index138] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_850_v3)).Merge(_module.__default.fm48(_module.__default.fm0((_this).f37, globalState), (_this).f37, (_this).f37, _848_v1, globalState));
              (globalState).f12 = (_module.__default.fm0(_848_v1, globalState)) || (_module.__default.fm0(new BigNumber((_858_v11).length), globalState));
              let _866_v18;
              let _nw128 = Array((new BigNumber(6)).toNumber()).fill([]);
              _866_v18 = _nw128;
              let _index139 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_866_v18).length));
              (_866_v18)[_index139] = _863_v15;
              let _index140 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_866_v18).length));
              let _rhs144 = _848_v1;
              let _rhs145 = _dafny.Seq.Concat(_dafny.Seq.of(p0), _858_v11);
              let _rhs146 = (_this).f38;
              let _rhs147 = _module.__default.safeModuloInt(_848_v1, (_this).f25);
              let _rhs148 = _863_v15;
              let _lhs109 = globalState;
              let _lhs110 = _866_v18;
              let _lhs111 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_866_v18).length));
              _848_v1 = _rhs144;
              _858_v11 = _rhs145;
              _lhs109.f12 = _rhs146;
              _848_v1 = _rhs147;
              _lhs110[_lhs111] = _rhs148;
            } else {
              let _867_v19;
              _867_v19 = _dafny.Seq.of((_this).f38, (_this).f38);
              _847_v0 = (((_867_v19)[_module.__default.safeIndex((_this).f25, new BigNumber((_867_v19).length))]) ? (_847_v0) : (new _dafny.CodePoint('l'.codePointAt(0))));
              let _868_v20;
              let _869_v21;
              let _870_v22;
              let _out22;
              let _out23;
              let _out24;
              let _outcollector8 = _module.__default.m0(globalState);
              _out22 = _outcollector8[0];
              _out23 = _outcollector8[1];
              _out24 = _outcollector8[2];
              _868_v20 = _out22;
              _869_v21 = _out23;
              _870_v22 = _out24;
              let _871_v23;
              let _872_v24;
              let _873_v25;
              let _out25;
              let _out26;
              let _out27;
              let _outcollector9 = _module.__default.m0(globalState);
              _out25 = _outcollector9[0];
              _out26 = _outcollector9[1];
              _out27 = _outcollector9[2];
              _871_v23 = _out25;
              _872_v24 = _out26;
              _873_v25 = _out27;
              let _874_v26;
              _874_v26 = _dafny.Seq.UnicodeFromString("ppywdo");
              let _rhs149 = _874_v26;
              let _rhs150 = (_this).f25;
              let _rhs151 = ((_873_v25) ? (_868_v20) : (_868_v20));
              let _rhs152 = (((_870_v22) ? ((_this).f38) : (p1))) === (p0);
              let _rhs153 = !(_873_v25);
              let _lhs112 = globalState;
              let _lhs113 = globalState;
              let _lhs114 = globalState;
              _lhs112.f10 = _rhs149;
              _848_v1 = _rhs150;
              _868_v20 = _rhs151;
              _lhs113.f12 = _rhs152;
              _lhs114.f12 = _rhs153;
              let _875_v27;
              _875_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_867_v19);
              _875_v27 = (_875_v27).update((_this).f37, _867_v19);
            }
          }
        }
      }
      let _876_v28;
      _876_v28 = _dafny.Seq.of((_this).f25);
      let _877_i2;
      _877_i2 = _dafny.ZERO;
      L8: {
        while (_module.__default.fm0(_module.__default.safeDivisionInt(_module.__default.fm44(p0, p0, p0, (_876_v28)[_module.__default.safeIndex((_this).f25, new BigNumber((_876_v28).length))], globalState), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_this).f25)).length))), globalState)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_877_i2)) {
              break L8;
            }
            _877_i2 = (_877_i2).plus(_dafny.ONE);
            let _878_v29;
            _878_v29 = new _dafny.CodePoint('r'.codePointAt(0));
            (globalState).f23 = _878_v29;
            let _879_v30;
            _879_v30 = _dafny.Seq.of(p0);
            let _880_v31;
            _880_v31 = _dafny.Seq.of(p0, (_879_v30)[_module.__default.safeIndex((_this).f37, new BigNumber((_879_v30).length))], (_this).f38, _module.__default.fm0((_this).f25, globalState));
            if (!_dafny.Seq.contains(_879_v30, (_this).f38)) {
              _878_v29 = _878_v29;
              let _881_v32;
              _881_v32 = new BigNumber(578);
              _881_v32 = ((_this).f37).plus((_881_v32).plus((_this).f25));
              (globalState).f12 = (_this).f38;
              let _882_v33;
              let _init25 = ((_883_v32) => function (_884_i3) {
                return _module.D9.create_DC27((_this).f37, _883_v32);
              })(_881_v32);
              let _nw129 = Array((new BigNumber(22)).toNumber());
              for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw129.length); _i0_25++) {
                _nw129[_i0_25] = _init25(new BigNumber(_i0_25));
              }
              _882_v33 = _nw129;
              let _885_v34;
              _885_v34 = _dafny.Seq.of(_882_v33, _882_v33, _882_v33);
              _885_v34 = _885_v34;
              let _886_v35;
              _886_v35 = _dafny.Seq.UnicodeFromString("u");
              let _887_v36;
              let _nw130 = Array((new BigNumber(3)).toNumber());
              _nw130[(_dafny.ZERO).toNumber()] = p1;
              _nw130[(_dafny.ONE).toNumber()] = (_this).f38;
              _nw130[(new BigNumber(2)).toNumber()] = p1;
              _887_v36 = _nw130;
              let _888_v37;
              _888_v37 = _dafny.Map.Empty.slice().updateUnsafe(_886_v35,_887_v36);
              _888_v37 = (_888_v37).update(_886_v35, _887_v36);
            } else {
              let _889_v38;
              let _nw131 = new _module.C1();
              _nw131.__ctor((_this).f37, (_this).f26);
              _889_v38 = _nw131;
              let _890_v39;
              let _init26 = function (_891_i4) {
                return _module.__default.safeDivisionInt(_891_i4, (_dafny.ZERO).minus((_this).f25));
              };
              let _nw132 = Array((new BigNumber(20)).toNumber());
              for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw132.length); _i0_26++) {
                _nw132[_i0_26] = _init26(new BigNumber(_i0_26));
              }
              _890_v39 = _nw132;
              let _index141 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_890_v39).length));
              (_890_v39)[_index141] = (_this).f37;
              let _index142 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_890_v39).length));
              (_890_v39)[_index142] = (_this).f37;
              let _892_v40;
              _892_v40 = _module.D5.create_DC15(_dafny.MultiSet.fromElements((_890_v39)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_890_v39).length))]));
              let _893_v41;
              _893_v41 = _module.D0.create_DC2((_this).f38);
              _892_v40 = _module.__default.fm49((_this).f37, p0, (((_893_v41).dtor_cf3) ? ((_this).f37) : ((_this).f25)), _878_v29, globalState);
              let _894_v42;
              _894_v42 = _module.D0.create_DC0(true);
              let _895_v43;
              let _nw133 = Array((new BigNumber(29)).toNumber());
              _nw133[(_dafny.ZERO).toNumber()] = (_this).f38;
              _nw133[(_dafny.ONE).toNumber()] = false;
              _nw133[(new BigNumber(2)).toNumber()] = p1;
              _nw133[(new BigNumber(3)).toNumber()] = p1;
              _nw133[(new BigNumber(4)).toNumber()] = p1;
              _nw133[(new BigNumber(5)).toNumber()] = true;
              _nw133[(new BigNumber(6)).toNumber()] = (_894_v42).dtor_cf0;
              _nw133[(new BigNumber(7)).toNumber()] = (_this).f38;
              _nw133[(new BigNumber(8)).toNumber()] = p1;
              _nw133[(new BigNumber(9)).toNumber()] = p1;
              _nw133[(new BigNumber(10)).toNumber()] = p1;
              _nw133[(new BigNumber(11)).toNumber()] = false;
              _nw133[(new BigNumber(12)).toNumber()] = p0;
              _nw133[(new BigNumber(13)).toNumber()] = p1;
              _nw133[(new BigNumber(14)).toNumber()] = p1;
              _nw133[(new BigNumber(15)).toNumber()] = true;
              _nw133[(new BigNumber(16)).toNumber()] = !((_this).f38);
              _nw133[(new BigNumber(17)).toNumber()] = p0;
              _nw133[(new BigNumber(18)).toNumber()] = p0;
              _nw133[(new BigNumber(19)).toNumber()] = p1;
              _nw133[(new BigNumber(20)).toNumber()] = p1;
              _nw133[(new BigNumber(21)).toNumber()] = p1;
              _nw133[(new BigNumber(22)).toNumber()] = false;
              _nw133[(new BigNumber(23)).toNumber()] = (_this).f38;
              _nw133[(new BigNumber(24)).toNumber()] = p0;
              _nw133[(new BigNumber(25)).toNumber()] = !(p0);
              _nw133[(new BigNumber(26)).toNumber()] = false;
              _nw133[(new BigNumber(27)).toNumber()] = !(_module.__default.fm0((_890_v39)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_890_v39).length))], globalState));
              _nw133[(new BigNumber(28)).toNumber()] = _module.__default.fm0((_this).f37, globalState);
              _895_v43 = _nw133;
              let _896_v44;
              _896_v44 = _dafny.Map.Empty.slice().updateUnsafe(_895_v43,(_this).f38);
              let _897_v45;
              _897_v45 = _dafny.Set.fromElements(_module.__default.fm0((_this).f25, globalState));
              let _898_v46;
              _898_v46 = _module.D12.create_DC32((_this).f38);
              let _899_v47;
              _899_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,false);
              let _900_v48;
              _900_v48 = _dafny.Map.Empty.slice().updateUnsafe(_898_v46,new BigNumber((_899_v47).length));
              let _901_v49;
              let _nw134 = Array((new BigNumber(28)).toNumber());
              _nw134[(_dafny.ZERO).toNumber()] = p1;
              _nw134[(_dafny.ONE).toNumber()] = p0;
              _nw134[(new BigNumber(2)).toNumber()] = (_this).f38;
              _nw134[(new BigNumber(3)).toNumber()] = ((_this).f38) === (p0);
              _nw134[(new BigNumber(4)).toNumber()] = p0;
              _nw134[(new BigNumber(5)).toNumber()] = !((p1) && ((_this).f38));
              _nw134[(new BigNumber(6)).toNumber()] = (_module.__default.fm0((_this).f25, globalState)) && ((((_896_v44).contains(_895_v43)) ? ((_896_v44).get(_895_v43)) : ((_this).f38)));
              _nw134[(new BigNumber(7)).toNumber()] = !(p1) || (p0);
              _nw134[(new BigNumber(8)).toNumber()] = (_this).f38;
              _nw134[(new BigNumber(9)).toNumber()] = (_897_v45).IsDisjointFrom(_897_v45);
              _nw134[(new BigNumber(10)).toNumber()] = p1;
              _nw134[(new BigNumber(11)).toNumber()] = ((_890_v39)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_890_v39).length))]).isLessThan(new BigNumber(130));
              _nw134[(new BigNumber(12)).toNumber()] = !((_879_v30)[_module.__default.safeIndex((_890_v39)[_module.__default.safeIndex(new BigNumber(960), new BigNumber((_890_v39).length))], new BigNumber((_879_v30).length))]) || (!(!(p0)));
              _nw134[(new BigNumber(13)).toNumber()] = p0;
              _nw134[(new BigNumber(14)).toNumber()] = !(_900_v48).contains(_898_v46);
              _nw134[(new BigNumber(15)).toNumber()] = true;
              _nw134[(new BigNumber(16)).toNumber()] = false;
              _nw134[(new BigNumber(17)).toNumber()] = p0;
              _nw134[(new BigNumber(18)).toNumber()] = !(!(_module.__default.fm0((_this).f37, globalState)));
              _nw134[(new BigNumber(19)).toNumber()] = (_this).f38;
              _nw134[(new BigNumber(20)).toNumber()] = p1;
              _nw134[(new BigNumber(21)).toNumber()] = p1;
              _nw134[(new BigNumber(22)).toNumber()] = _module.__default.fm0(new BigNumber(-100), globalState);
              _nw134[(new BigNumber(23)).toNumber()] = p1;
              _nw134[(new BigNumber(24)).toNumber()] = (_this).f38;
              _nw134[(new BigNumber(25)).toNumber()] = (_this).f38;
              _nw134[(new BigNumber(26)).toNumber()] = (_this).f38;
              _nw134[(new BigNumber(27)).toNumber()] = p1;
              _901_v49 = _nw134;
              let _index143 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_895_v43).length));
              (_895_v43)[_index143] = p0;
              let _index144 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_895_v43).length));
              (_895_v43)[_index144] = _module.__default.fm0((_this).f37, globalState);
              let _902_v50;
              _902_v50 = _dafny.Seq.UnicodeFromString("wow");
              let _903_v51;
              _903_v51 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),(_this).f25);
              let _904_v52;
              _904_v52 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_902_v50).length)).isLessThanOrEqualTo((_module.D0.create_DC3((_this).f37, p0)).dtor_cf4),((p0) ? (_903_v51) : (_903_v51)));
              _904_v52 = ((((_895_v43)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_895_v43).length))]) ? (_904_v52) : (_904_v52))).Merge(_904_v52);
            }
            let _905_v53;
            _905_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
            let _906_v54;
            _906_v54 = _dafny.Map.Empty.slice().updateUnsafe(_876_v28,_dafny.Map.Empty.slice().updateUnsafe(_905_v53,new BigNumber((_880_v31).length)));
            let _907_v55;
            _907_v55 = _dafny.Map.Empty.slice().updateUnsafe(_905_v53,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), function (_908_i5) {
              return (_this).f25;
            })).length));
            _906_v54 = (_906_v54).update(_876_v28, _907_v55);
            let _909_v56;
            _909_v56 = _module.D8.create_DC24(_module.D1.create_DC5());
            let _source17 = _909_v56;
            if (_source17.is_DC24) {
              let _910___mcc_h0 = (_source17).cf36;
              let _911_cf36 = _910___mcc_h0;
              let _912_v57;
              let _nw135 = Array((new BigNumber(7)).toNumber()).fill(false);
              _912_v57 = _nw135;
              let _index145 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length));
              (_912_v57)[_index145] = p1;
              let _index146 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length));
              (_912_v57)[_index146] = p0;
              let _913_v58;
              _913_v58 = _dafny.Seq.of(_876_v28);
              let _914_v59;
              _914_v59 = _module.D0.create_DC0((_this).f38);
              _905_v53 = (_905_v53).update(p1, ((_this).f37).isLessThan(new BigNumber((_dafny.Seq.update(_913_v58, _module.__default.safeIndex((_this).f25, new BigNumber((_913_v58).length)), (_this).fm38(_914_v59, globalState))).length)));
              let _915_v60;
              let _init27 = function (_916_i6) {
                return (_916_i6).multipliedBy((_this).f37);
              };
              let _nw136 = Array((new BigNumber(3)).toNumber());
              for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw136.length); _i0_27++) {
                _nw136[_i0_27] = _init27(new BigNumber(_i0_27));
              }
              _915_v60 = _nw136;
              let _917_v61;
              _917_v61 = _dafny.Map.Empty.slice().updateUnsafe(_915_v60,p1);
              let _918_v62;
              _918_v62 = _dafny.MultiSet.fromElements(false, (((_905_v53).contains((_912_v57)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length))])) ? ((_905_v53).get((_912_v57)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length))])) : ((_912_v57)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length))])));
              let _919_v63;
              _919_v63 = new BigNumber(497);
              let _920_v64;
              _920_v64 = _dafny.Seq.UnicodeFromString("hr");
              let _921_v65;
              _921_v65 = _dafny.Set.fromElements(_920_v64);
              let _922_v66;
              _922_v66 = _dafny.Seq.of(_dafny.Set.fromElements(_920_v64), _921_v65);
              let _rhs154 = ((((_this).f38) ? ((_922_v66)[_module.__default.safeIndex(_919_v63, new BigNumber((_922_v66).length))]) : (_921_v65))).IsDisjointFrom(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), ((_923_v64) => function (_924_i7) {
                return (_923_v64)[_module.__default.safeIndex((_this).f37, new BigNumber((_923_v64).length))];
              })(_920_v64)), _920_v64));
              let _rhs155 = _917_v61;
              let _rhs156 = (_dafny.MultiSet.fromElements(p1)).Union(_dafny.MultiSet.fromElements((_this).f38, true, (_912_v57)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length))]));
              let _rhs157 = (((_912_v57)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_912_v57).length))]) ? (_919_v63) : ((_this).f25));
              let _lhs115 = globalState;
              _lhs115.f12 = _rhs154;
              _917_v61 = _rhs155;
              _918_v62 = _rhs156;
              _919_v63 = _rhs157;
              _919_v63 = (_876_v28)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f25, _919_v63), new BigNumber((_876_v28).length))];
            } else if (_source17.is_DC23) {
              let _925___mcc_h1 = (_source17).cf35;
              let _926_cf35 = _925___mcc_h1;
              let _927_v67;
              let _nw137 = new _module.C1();
              _nw137.__ctor((_dafny.ZERO).minus((_this).f37), (_this).f26);
              _927_v67 = _nw137;
              let _928_v68;
              _928_v68 = new BigNumber(889);
              _928_v68 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f37, (_this).f25));
              (globalState).f23 = _878_v29;
              let _rhs158 = (new BigNumber(106)).multipliedBy(new BigNumber(398));
              let _rhs159 = _880_v31;
              let _rhs160 = _module.__default.safeModuloInt((_dafny.ZERO).minus(((p0) ? ((_dafny.ZERO).minus((_this).f37)) : (_928_v68))), (_this).f37);
              _928_v68 = _rhs158;
              _880_v31 = _rhs159;
              _928_v68 = _rhs160;
            } else {
              let _929___mcc_h2 = (_source17).cf37;
              let _930_cf37 = _929___mcc_h2;
              let _931_v69;
              _931_v69 = _dafny.Set.fromElements(p0, (_this).f38);
              let _932_v70;
              _932_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,(_this).f25);
              let _933_v71;
              _933_v71 = _dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC16((_this).f38)).dtor_cf25,new BigNumber(-328));
              let _934_v72;
              _934_v72 = _dafny.MultiSet.fromElements(p1, p0);
              let _935_v73;
              let _nw138 = Array((new BigNumber(22)).toNumber());
              _nw138[(_dafny.ZERO).toNumber()] = new BigNumber((_931_v69).length);
              _nw138[(_dafny.ONE).toNumber()] = new BigNumber(332);
              _nw138[(new BigNumber(2)).toNumber()] = new BigNumber((_932_v70).length);
              _nw138[(new BigNumber(3)).toNumber()] = (((_933_v71).contains((_this).f38)) ? ((_933_v71).get((_this).f38)) : ((_this).f25));
              _nw138[(new BigNumber(4)).toNumber()] = new BigNumber((_876_v28).length);
              _nw138[(new BigNumber(5)).toNumber()] = _module.__default.fm44(p0, (_this).f38, p1, (_this).f25, globalState);
              _nw138[(new BigNumber(6)).toNumber()] = (_this).f25;
              _nw138[(new BigNumber(7)).toNumber()] = (_this).f37;
              _nw138[(new BigNumber(8)).toNumber()] = (_this).f25;
              _nw138[(new BigNumber(9)).toNumber()] = (_this).f37;
              _nw138[(new BigNumber(10)).toNumber()] = (_this).f37;
              _nw138[(new BigNumber(11)).toNumber()] = (_this).f37;
              _nw138[(new BigNumber(12)).toNumber()] = (_this).f37;
              _nw138[(new BigNumber(13)).toNumber()] = (_this).f25;
              _nw138[(new BigNumber(14)).toNumber()] = _module.__default.fm44(!((_this).f38), p0, p0, (_this).f37, globalState);
              _nw138[(new BigNumber(15)).toNumber()] = new BigNumber(267);
              _nw138[(new BigNumber(16)).toNumber()] = new BigNumber((_934_v72).cardinality());
              _nw138[(new BigNumber(17)).toNumber()] = (_this).f25;
              _nw138[(new BigNumber(18)).toNumber()] = (_this).f25;
              _nw138[(new BigNumber(19)).toNumber()] = (_this).f25;
              _nw138[(new BigNumber(20)).toNumber()] = new BigNumber(562);
              _nw138[(new BigNumber(21)).toNumber()] = (_this).f25;
              _935_v73 = _nw138;
              let _936_v74;
              _936_v74 = _dafny.Map.Empty.slice().updateUnsafe(_935_v73,(_this).f37);
              _936_v74 = (_936_v74).update(_935_v73, (_this).f37);
              let _937_v75;
              _937_v75 = new BigNumber(547);
              _937_v75 = (new BigNumber((function () {
                let _coll40 = new _dafny.Map();
                for (const _compr_40 of (_876_v28).Elements) {
                  let _938_v76 = _compr_40;
                  if (_dafny.Seq.contains(_876_v28, _938_v76)) {
                    _coll40.push([_module.__default.safeDivisionInt(_938_v76, _937_v75),_934_v72]);
                  }
                }
                return _coll40;
              }()).length)).multipliedBy(_937_v75);
              _933_v71 = (_933_v71).update(false, (new BigNumber(291)).plus(new BigNumber(832)));
              let _index147 = _module.__default.safeIndex(new BigNumber(332), new BigNumber(((_this).f26).length));
              ((_this).f26)[_index147] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(562)), ((_939_v29) => function (_940_i8) {
                return _939_v29;
              })(_878_v29));
              let _941_v77;
              _941_v77 = _dafny.Seq.UnicodeFromString("rkvfdt");
              let _index148 = _module.__default.safeIndex(new BigNumber(332), new BigNumber(((_this).f26).length));
              ((_this).f26)[_index148] = _941_v77;
            }
          }
        }
      }
      let _942_v78;
      let _init28 = function (_943_i9) {
        return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("mmh"), _dafny.Seq.UnicodeFromString("nlcfnt"));
      };
      let _nw139 = Array((new BigNumber(26)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw139.length); _i0_28++) {
        _nw139[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _942_v78 = _nw139;
      let _index149 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
      (_942_v78)[_index149] = (_this).f38;
      let _index150 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
      (_942_v78)[_index150] = (p0) === (_module.__default.fm0(new BigNumber((_876_v28).length), globalState));
      let _944_i10;
      _944_i10 = _dafny.ZERO;
      L9: {
        while ((_this).f38) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_944_i10)) {
              break L9;
            }
            _944_i10 = (_944_i10).plus(_dafny.ONE);
            let _945_v79;
            let _init29 = ((_946_p1) => function (_947_i11) {
              return _dafny.Seq.of(_946_p1, _946_p1);
            })(p1);
            let _nw140 = Array((new BigNumber(15)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw140.length); _i0_29++) {
              _nw140[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _945_v79 = _nw140;
            let _948_v80;
            let _init30 = function (_949_i12) {
              return (_949_i12).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements((_this).f25, (_this).f37))).length));
            };
            let _nw141 = Array((new BigNumber(26)).toNumber());
            for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw141.length); _i0_30++) {
              _nw141[_i0_30] = _init30(new BigNumber(_i0_30));
            }
            _948_v80 = _nw141;
            let _index151 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_948_v80).length));
            (_948_v80)[_index151] = new BigNumber(697);
            let _950_v81;
            _950_v81 = new BigNumber(484);
            let _index152 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_948_v80).length));
            let _rhs161 = _945_v79;
            let _rhs162 = _module.__default.safeModuloInt(new BigNumber(519), _950_v81);
            let _rhs163 = true;
            let _rhs164 = (_this).f37;
            let _lhs116 = _948_v80;
            let _lhs117 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_948_v80).length));
            let _lhs118 = globalState;
            _945_v79 = _rhs161;
            _lhs116[_lhs117] = _rhs162;
            _lhs118.f12 = _rhs163;
            _950_v81 = _rhs164;
            let _951_v82;
            _951_v82 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
            let _952_v83;
            _952_v83 = _dafny.MultiSet.fromElements(_876_v28, _876_v28);
            let _953_v84;
            _953_v84 = _dafny.Set.fromElements(new BigNumber(540), (_this).f37, new BigNumber((_952_v83).cardinality()));
            _951_v82 = (_951_v82).update((_953_v84).IsSubsetOf(_953_v84), (_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]);
            let _index153 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
            (_942_v78)[_index153] = false;
            _953_v84 = _953_v84;
          }
        }
      }
      let _954_v85;
      _954_v85 = _dafny.Seq.UnicodeFromString("iei");
      let _955_v86;
      _955_v86 = new _dafny.CodePoint('h'.codePointAt(0));
      let _956_v87;
      _956_v87 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_954_v85, _module.__default.safeIndex((_this).f37, new BigNumber((_954_v85).length)), _955_v86)).length));
      let _hi4 = new BigNumber((_956_v87).cardinality());
      for (let _957_i13 = (_this).f37; _957_i13.isLessThan(_hi4); _957_i13 = _957_i13.plus(_dafny.ONE)) {
        let _958_v88;
        _958_v88 = _dafny.Set.fromElements((_876_v28)[_module.__default.safeIndex((_this).f25, new BigNumber((_876_v28).length))]);
        let _959_v89;
        _959_v89 = _dafny.Map.Empty.slice().updateUnsafe(_957_i13,p0);
        let _960_v90;
        _960_v90 = _module.D5.create_DC17(new BigNumber((_958_v88).length), (((_959_v89).contains(_957_i13)) ? ((_959_v89).get(_957_i13)) : (p0)));
        let _961_v91;
        _961_v91 = _dafny.Seq.of(p1, (_this).f38);
        let _962_v92;
        _962_v92 = _dafny.MultiSet.fromElements((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]);
        let _963_v93;
        _963_v93 = _dafny.Set.fromElements(_module.__default.fm50(_957_i13, _954_v85, (_960_v90).dtor_cf27, new BigNumber((_961_v91).length), globalState), _962_v92, _962_v92, (_962_v92).update((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))], _module.__default.abs((_dafny.ZERO).minus((_this).f37))), _962_v92);
        let _964_v94;
        _964_v94 = new BigNumber(821);
        let _965_v95;
        _965_v95 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,_957_i13);
        let _index154 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        let _rhs165 = (_956_v87).IsProperSubsetOf((_956_v87).Union(_956_v87));
        let _rhs166 = _module.__default.fm51((_958_v88).Difference(_module.__default.fm47((_this).f38, new BigNumber(170), new BigNumber((_965_v95).length), _955_v86, globalState)), new BigNumber(-486), _955_v86, _964_v94, globalState);
        let _rhs167 = (_876_v28)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_876_v28)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f37), new BigNumber((_876_v28).length))], (_this).f37), new BigNumber((_876_v28).length))];
        let _lhs119 = _942_v78;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        _lhs119[_lhs120] = _rhs165;
        _963_v93 = _rhs166;
        _964_v94 = _rhs167;
        _961_v91 = _dafny.Seq.Concat(_module.__default.fm1((_this).f38, globalState), _dafny.Seq.of((_this).f38));
        let _966_v96;
        _966_v96 = _dafny.Set.fromElements(p1, p0);
        let _index155 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        let _rhs168 = false;
        let _rhs169 = _964_v94;
        let _rhs170 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_966_v96).length), new BigNumber(-109)), new BigNumber(-219));
        let _rhs171 = (_961_v91)[_module.__default.safeIndex((_module.__default.fm44((_this).f38, (_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))], true, (_this).f37, globalState)).multipliedBy(_957_i13), new BigNumber((_961_v91).length))];
        let _rhs172 = _964_v94;
        let _lhs121 = _942_v78;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        let _lhs123 = globalState;
        _lhs121[_lhs122] = _rhs168;
        _964_v94 = _rhs169;
        _964_v94 = _rhs170;
        _lhs123.f12 = _rhs171;
        _964_v94 = _rhs172;
        let _index156 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        (_942_v78)[_index156] = p0;
      }
      let _967_v97;
      _967_v97 = _module.D12.create_DC32(false);
      let _968_v98;
      _968_v98 = _module.D12.create_DC34(_967_v97);
      let _source18 = _968_v98;
      if (_source18.is_DC32) {
        let _969___mcc_h3 = (_source18).cf47;
        let _970_cf47 = _969___mcc_h3;
        _970_cf47 = (_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))];
        let _971_v99;
        let _nw142 = new _module.C1();
        _nw142.__ctor(new BigNumber(802), (_this).f26);
        _971_v99 = _nw142;
        let _972_v100;
        _972_v100 = _module.D0.create_DC0(_970_cf47);
        let _973_v101;
        _973_v101 = _dafny.Set.fromElements(new BigNumber((_954_v85).length), new BigNumber(((_this).fm38(_972_v100, globalState)).length));
        let _974_v102;
        _974_v102 = _dafny.Map.Empty.slice().updateUnsafe(_973_v101,(_this).f25);
        let _975_v104;
        _975_v104 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f37),(_this).f25);
        _974_v102 = (_974_v102).update((((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]) ? (_973_v101) : (function () {
          let _coll41 = new _dafny.Set();
          for (const _compr_41 of (_dafny.MultiSet.FromArray(_876_v28)).Elements) {
            let _976_v103 = _compr_41;
            if ((_dafny.MultiSet.FromArray(_876_v28)).contains(_976_v103)) {
              _coll41.add((_976_v103).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length)));
            }
          }
          return _coll41;
        }())), (((_975_v104).contains((_this).f37)) ? ((_975_v104).get((_this).f37)) : ((_this).f37)));
        let _977_v105;
        _977_v105 = _dafny.Seq.of(_973_v101);
        let _978_v106;
        _978_v106 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,_955_v86);
        let _979_v108;
        _979_v108 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,_970_cf47);
        let _980_v111;
        let _nw143 = Array((new BigNumber(23)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = _module.__default.fm47(!(!(p1)), (_this).f37, new BigNumber((_973_v101).length), _955_v86, globalState);
        _nw143[(_dafny.ONE).toNumber()] = _973_v101;
        _nw143[(new BigNumber(2)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(3)).toNumber()] = (_973_v101).Difference(_973_v101);
        _nw143[(new BigNumber(4)).toNumber()] = (_977_v105)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_977_v105).length))];
        _nw143[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements((_this).f37, new BigNumber((_956_v87).cardinality()), (_this).f25);
        _nw143[(new BigNumber(6)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(7)).toNumber()] = (_973_v101).Intersect(_973_v101);
        _nw143[(new BigNumber(8)).toNumber()] = (_973_v101).Intersect(_973_v101);
        _nw143[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements((_this).f25);
        _nw143[(new BigNumber(10)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(11)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(12)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(13)).toNumber()] = (_module.__default.fm47((_this).f38, (_this).f25, (_this).f37, _955_v86, globalState)).Union(_973_v101);
        _nw143[(new BigNumber(14)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(15)).toNumber()] = _973_v101;
        _nw143[(new BigNumber(16)).toNumber()] = (_973_v101).Difference(function () {
          let _coll42 = new _dafny.Set();
          for (const _compr_42 of (_978_v106).Keys.Elements) {
            let _981_v107 = _compr_42;
            if ((_978_v106).contains(_981_v107)) {
              _coll42.add(_module.__default.safeModuloInt(_981_v107, new BigNumber(-123)));
            }
          }
          return _coll42;
        }());
        _nw143[(new BigNumber(17)).toNumber()] = (_973_v101).Union(_973_v101);
        _nw143[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements((_dafny.ZERO).minus(_module.__default.fm44(true, (_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))], (((_979_v108).contains((_this).f25)) ? ((_979_v108).get((_this).f25)) : (true)), _module.__default.fm44(_970_cf47, p1, (_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))], (_this).f37, globalState), globalState)));
        _nw143[(new BigNumber(19)).toNumber()] = (_973_v101).Intersect(_973_v101);
        _nw143[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements((_this).f37, (_this).f37);
        _nw143[(new BigNumber(21)).toNumber()] = (((_this).f38) ? (function () {
          let _coll43 = new _dafny.Set();
          for (const _compr_43 of (_975_v104).Keys.Elements) {
            let _982_v109 = _compr_43;
            if ((_975_v104).contains(_982_v109)) {
              _coll43.add((_982_v109).minus(new BigNumber(172)));
            }
          }
          return _coll43;
        }()) : ((_module.D4.create_DC11(function () {
  let _coll44 = new _dafny.Set();
  for (const _compr_44 of (_973_v101).Elements) {
    let _983_v110 = _compr_44;
    if ((_973_v101).contains(_983_v110)) {
      _coll44.add((_983_v110).minus(new BigNumber((_dafny.Seq.UnicodeFromString("f")).length)));
    }
  }
  return _coll44;
}())).dtor_cf18));
        _nw143[(new BigNumber(22)).toNumber()] = _973_v101;
        _980_v111 = _nw143;
        let _index157 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_980_v111).length));
        (_980_v111)[_index157] = _973_v101;
        let _984_v112;
        _984_v112 = _dafny.Seq.of((_this).f38);
        let _985_v113;
        _985_v113 = new BigNumber(533);
        let _986_v114;
        let _nw144 = Array((new BigNumber(2)).toNumber());
        _nw144[(_dafny.ZERO).toNumber()] = _985_v113;
        _nw144[(_dafny.ONE).toNumber()] = ((_this).f25).multipliedBy((_this).f37);
        _986_v114 = _nw144;
        let _index158 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_986_v114).length));
        (_986_v114)[_index158] = (new BigNumber((_979_v108).length)).minus((_this).f25);
        let _index159 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_980_v111).length));
        let _index160 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        let _index161 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_986_v114).length));
        let _rhs173 = _973_v101;
        let _rhs174 = _dafny.Seq.Concat(_984_v112, _984_v112);
        let _rhs175 = new BigNumber(947);
        let _rhs176 = (_module.__default.safeDivisionInt((_this).f37, _985_v113)).isEqualTo((new BigNumber((_954_v85).length)).plus((_this).f25));
        let _rhs177 = (_this).f25;
        let _lhs124 = _980_v111;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_980_v111).length));
        let _lhs126 = _942_v78;
        let _lhs127 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
        let _lhs128 = _986_v114;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_986_v114).length));
        _lhs124[_lhs125] = _rhs173;
        _984_v112 = _rhs174;
        _985_v113 = _rhs175;
        _lhs126[_lhs127] = _rhs176;
        _lhs128[_lhs129] = _rhs177;
      } else if (_source18.is_DC33) {
        let _987___mcc_h4 = (_source18).cf48;
        let _988___mcc_h5 = (_source18).cf49;
        let _989___mcc_h6 = (_source18).cf50;
        let _990___mcc_h7 = (_source18).cf51;
        let _991_cf51 = _990___mcc_h7;
        let _992_cf50 = _989___mcc_h6;
        let _993_cf49 = _988___mcc_h5;
        let _994_cf48 = _987___mcc_h4;
        let _995_v116;
        _995_v116 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hglnqof"), _dafny.Seq.update(_954_v85, _module.__default.safeIndex(new BigNumber((_954_v85).length), new BigNumber((_954_v85).length)), _955_v86));
        let _rhs178 = !(false);
        let _rhs179 = (_this).f37;
        let _rhs180 = (_dafny.ZERO).minus((_this).f37);
        let _rhs181 = _991_cf51;
        let _rhs182 = new BigNumber((function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of (_995_v116).Elements) {
            let _996_v115 = _compr_45;
            if ((_995_v116).contains(_996_v115)) {
              _coll45.push([_996_v115,false]);
            }
          }
          return _coll45;
        }()).length);
        _992_cf50 = _rhs178;
        _991_cf51 = _rhs179;
        _991_cf51 = _rhs180;
        _991_cf51 = _rhs181;
        _991_cf51 = _rhs182;
        let _997_v117;
        _997_v117 = _dafny.MultiSet.fromElements((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]);
        if (!((((_992_cf50) && (p1)) ? (_module.__default.fm0((_this).f37, globalState)) : ((_997_v117).IsDisjointFrom(_997_v117))))) {
          let _998_v118;
          _998_v118 = _dafny.Map.Empty.slice().updateUnsafe(_991_cf51,_954_v85);
          let _999_v119;
          _999_v119 = _dafny.Map.Empty.slice().updateUnsafe(_954_v85,_998_v118);
          _999_v119 = (_999_v119).update(_954_v85, _998_v118);
          (globalState).f12 = (_this).f38;
          (globalState).f12 = (((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]) ? ((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]) : (p1));
          let _1000_v120;
          _1000_v120 = _dafny.Map.Empty.slice().updateUnsafe((_this).f37,p0);
          let _1001_v121;
          _1001_v121 = _dafny.Seq.of(_1000_v120);
          let _1002_v122;
          _1002_v122 = _dafny.Set.fromElements(new BigNumber((_1001_v121).length), (_this).f25, (_this).f37);
          let _1003_v123;
          _1003_v123 = _dafny.MultiSet.fromElements(_1002_v122, _1002_v122);
          _991_cf51 = (((_1003_v123).contains(_1002_v122)) ? ((_1003_v123).get(_1002_v122)) : (_991_cf51));
          let _1004_v124;
          let _init31 = ((_1005_v28) => function (_1006_i14) {
            return (_1006_i14).multipliedBy(new BigNumber((_1005_v28).length));
          })(_876_v28);
          let _nw145 = Array((new BigNumber(8)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw145.length); _i0_31++) {
            _nw145[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1004_v124 = _nw145;
          _1004_v124 = _1004_v124;
        } else {
          let _1007_v125;
          _1007_v125 = _dafny.Set.fromElements(new BigNumber(738));
          let _1008_v126;
          _1008_v126 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f38);
          let _1009_v127;
          _1009_v127 = _module.D9.create_DC27(new BigNumber((_1007_v125).length), new BigNumber((_1008_v126).length));
          _991_cf51 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_1009_v127).dtor_cf40), _991_cf51);
          let _1010_v128;
          let _nw146 = new _module.C1();
          _nw146.__ctor((_this).f25, (_this).f26);
          _1010_v128 = _nw146;
          let _1011_v129;
          _1011_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1010_v128,_module.__default.safeModuloInt(new BigNumber(982), (_this).f37));
          _1011_v129 = _1011_v129;
          _876_v28 = _876_v28;
          (globalState).f12 = false;
          let _1012_v130;
          let _nw147 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Set.Empty);
          _1012_v130 = _nw147;
          let _index162 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_1012_v130).length));
          (_1012_v130)[_index162] = _1007_v125;
          let _1013_v131;
          _1013_v131 = _module.D4.create_DC11(_module.__default.fm47(_module.__default.fm0(_991_cf51, globalState), (_this).f37, (_this).f25, new _dafny.CodePoint('v'.codePointAt(0)), globalState));
          let _index163 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_1012_v130).length));
          let _rhs183 = (_1013_v131).dtor_cf18;
          let _rhs184 = _954_v85;
          let _rhs185 = (_this).f37;
          let _rhs186 = (_dafny.ZERO).minus(((!(_991_cf51).isEqualTo(new BigNumber(747))) ? ((_this).f37) : (((_dafny.ZERO).minus(_991_cf51)).minus((_this).f37))));
          let _lhs130 = _1012_v130;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_1012_v130).length));
          _lhs130[_lhs131] = _rhs183;
          _954_v85 = _rhs184;
          _991_cf51 = _rhs185;
          _991_cf51 = _rhs186;
        }
        if (_992_cf50) {
          _991_cf51 = (_this).f37;
          let _1014_v132;
          let _nw148 = new _module.C2();
          _nw148.__ctor();
          _1014_v132 = _nw148;
          let _1015_v133;
          let _nw149 = Array((new BigNumber(26)).toNumber()).fill([]);
          _1015_v133 = _nw149;
          let _1016_v134;
          let _init32 = ((_1017_v85) => function (_1018_i15) {
            return (_1017_v85)[_module.__default.safeIndex((_this).f37, new BigNumber((_1017_v85).length))];
          })(_954_v85);
          let _nw150 = Array((new BigNumber(26)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw150.length); _i0_32++) {
            _nw150[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1016_v134 = _nw150;
          let _1019_v135;
          _1019_v135 = _module.D13.create_DC35(_1016_v134);
          let _index164 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_1015_v133).length));
          (_1015_v133)[_index164] = (_1019_v135).dtor_cf53;
          let _index165 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_1015_v133).length));
          let _init33 = ((_1020_v86) => function (_1021_i16) {
            return _1020_v86;
          })(_955_v86);
          let _nw151 = Array((new BigNumber(7)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw151.length); _i0_33++) {
            _nw151[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          (_1015_v133)[_index165] = _nw151;
          (globalState).f23 = _module.__default.fm39(p1, globalState);
          (globalState).f12 = !(((_this).f37).isLessThan(_991_cf51)) || ((_this).f38);
        } else {
          let _1022_v136;
          _1022_v136 = _dafny.Map.Empty.slice().updateUnsafe(!(p0) || (false),(_this).f37);
          let _1023_v137;
          _1023_v137 = _dafny.Map.Empty.slice().updateUnsafe(_991_cf51,(_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]);
          let _rhs187 = _954_v85;
          let _rhs188 = (((_1022_v136).contains(!(_992_cf50) || ((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]))) ? ((_1022_v136).get(!(_992_cf50) || ((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]))) : (_module.__default.fm44(!((((_1023_v137).contains(_991_cf51)) ? ((_1023_v137).get(_991_cf51)) : ((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]))), p0, p0, (_dafny.ZERO).minus((_this).f25), globalState)));
          let _lhs132 = globalState;
          _lhs132.f10 = _rhs187;
          _991_cf51 = _rhs188;
          (globalState).f12 = (_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))];
          let _1024_v138;
          let _nw152 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1024_v138 = _nw152;
          let _index166 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1024_v138).length));
          (_1024_v138)[_index166] = new BigNumber((function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of (_module.__default.fm52(_992_cf50, (_this).f37, _991_cf51, globalState)).Keys.Elements) {
              let _1025_v139 = _compr_46;
              if ((_module.__default.fm52(_992_cf50, (_this).f37, _991_cf51, globalState)).contains(_1025_v139)) {
                _coll46.push([_1025_v139,new BigNumber(3)]);
              }
            }
            return _coll46;
          }()).length);
          let _index167 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1024_v138).length));
          (_1024_v138)[_index167] = ((_this).f25).minus(new BigNumber((_module.__default.fm47(true, new BigNumber((_954_v85).length), new BigNumber((_1023_v137).length), _955_v86, globalState)).length));
          let _index168 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1024_v138).length));
          (_1024_v138)[_index168] = _module.__default.safeModuloInt((_this).f37, (_this).f25);
          let _1026_v140;
          _1026_v140 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("y"));
          let _index169 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
          let _rhs189 = _dafny.Seq.update(_dafny.Seq.Concat(_954_v85, _954_v85), _module.__default.safeIndex(_module.__default.fm44(_module.__default.fm0((_1024_v138)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_1024_v138).length))], globalState), _module.__default.fm0((_this).f25, globalState), !(false), (_1024_v138)[_module.__default.safeIndex(new BigNumber(114), new BigNumber((_1024_v138).length))], globalState), new BigNumber((_dafny.Seq.Concat(_954_v85, _954_v85)).length)), _955_v86);
          let _rhs190 = !(((_1026_v140).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), ((_1027_v86) => function (_1028_i17) {
            return _1027_v86;
          })(_955_v86)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-740)), function (_1029_i18) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          })))).IsSubsetOf(_1026_v140));
          let _lhs133 = globalState;
          let _lhs134 = _942_v78;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length));
          _lhs133.f10 = _rhs189;
          _lhs134[_lhs135] = _rhs190;
        }
        _991_cf51 = (_991_cf51).minus(_991_cf51);
      } else if (_source18.is_DC31) {
        let _1030___mcc_h8 = (_source18).cf46;
        let _1031_cf46 = _1030___mcc_h8;
        let _1032_v141;
        _1032_v141 = _dafny.Set.fromElements((_942_v78)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_942_v78).length))]);
        let _1033_v142;
        _1033_v142 = _dafny.Set.fromElements((_this).f25, (_this).f25);
        let _1034_v143;
        _1034_v143 = _module.D10.create_DC29(p1, (_this).f38, new BigNumber((_1033_v142).length));
        let _1035_v144;
        _1035_v144 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f38),(_1032_v141).Difference(_dafny.Set.fromElements((_1034_v143).dtor_cf43)));
        _1035_v144 = (_1035_v144).update(p1, _1032_v141);
        let _1036_v145;
        let _nw153 = new _module.C1();
        _nw153.__ctor(((_this).f25).minus((_this).f37), (_this).f26);
        _1036_v145 = _nw153;
        let _1037_v146;
        let _nw154 = new _module.C0();
        _nw154.__ctor();
        _1037_v146 = _nw154;
        let _1038_v147;
        _1038_v147 = new BigNumber(385);
        _1038_v147 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_954_v85, _954_v85)).length));
      } else {
        let _1039___mcc_h9 = (_source18).cf52;
        let _1040_cf52 = _1039___mcc_h9;
        let _1041_v148;
        let _nw155 = new _module.C2();
        _nw155.__ctor();
        _1041_v148 = _nw155;
        (globalState).f23 = _955_v86;
        let _1042_v149;
        let _nw156 = new _module.C1();
        _nw156.__ctor((_this).f37, (_this).f26);
        _1042_v149 = _nw156;
        let _1043_v150;
        _1043_v150 = new BigNumber(519);
        _1043_v150 = new BigNumber((_dafny.Seq.Concat(_954_v85, _954_v85)).length);
      }
      return;
    }
    get f37() {
      let _this = this;
      return _this._f37;
    };
    get f38() {
      let _this = this;
      return _this._f38;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
      this._f36 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f36, f25, f26) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      r0 = (_this).f25;
      if (!(p1)) {
        let _index170 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length));
        (p3)[_index170] = ((_this).f25).minus((_this).f25);
        let _index171 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length));
        (p3)[_index171] = (_this).f25;
        let _1044_v0;
        _1044_v0 = _dafny.Seq.UnicodeFromString("w");
        let _1045_v1;
        _1045_v1 = _dafny.Set.fromElements(p1, p1);
        let _1046_v2;
        _1046_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1045_v1);
        let _1047_v3;
        _1047_v3 = _dafny.Map.Empty.slice().updateUnsafe(((!(p2)) ? (_1045_v1) : ((((_1046_v2).contains((_this).f25)) ? ((_1046_v2).get((_this).f25)) : (_dafny.Set.fromElements(p1, _module.__default.fm0(new BigNumber(834), globalState), true, p2))))),new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2)).update(new BigNumber((_1045_v1).length), !(p1))).length));
        let _index172 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length));
        let _rhs191 = (_this).f25;
        let _rhs192 = _dafny.Seq.Concat(_1044_v0, _dafny.Seq.UnicodeFromString("ctwwm"));
        let _rhs193 = (((_1047_v3).contains((_1045_v1).Union(_1045_v1))) ? ((_1047_v3).get((_1045_v1).Union(_1045_v1))) : ((_this).f25));
        let _lhs136 = globalState;
        let _lhs137 = p3;
        let _lhs138 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length));
        r1 = _rhs191;
        _lhs136.f10 = _rhs192;
        _lhs137[_lhs138] = _rhs193;
        (globalState).f12 = !(p1);
        let _1048_v4;
        let _init34 = function (_1049_i0) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        };
        let _nw157 = Array((new BigNumber(13)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw157.length); _i0_34++) {
          _nw157[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1048_v4 = _nw157;
        let _1050_v5;
        _1050_v5 = new _dafny.CodePoint('a'.codePointAt(0));
        let _index173 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1048_v4).length));
        (_1048_v4)[_index173] = _1050_v5;
        let _index174 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1048_v4).length));
        (_1048_v4)[_index174] = _1050_v5;
        let _1051_v6;
        _1051_v6 = _module.D0.create_DC0(p2);
        let _1052_v7;
        _1052_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1051_v6,p1);
        let _1053_v8;
        _1053_v8 = _dafny.Seq.of((_this).f25);
        let _1054_v9;
        _1054_v9 = _dafny.Map.Empty.slice().updateUnsafe(((p3)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length))]).plus(_module.__default.fm21(_1052_v7, new BigNumber(-60), false, _1044_v0, globalState)),(_1053_v8)[_module.__default.safeIndex((p3)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length))], new BigNumber((_1053_v8).length))]);
        _1054_v9 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((p3)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length))]),(p3)[_module.__default.safeIndex(new BigNumber(990), new BigNumber((p3).length))]);
      } else {
        let _1055_v10;
        _1055_v10 = _dafny.Seq.UnicodeFromString("bliyyoh");
        (globalState).f10 = _dafny.Seq.Concat(_module.__default.fm2((_this).f25, (_this).f25, globalState), _dafny.Seq.Concat(_1055_v10, _1055_v10));
        let _1056_v11;
        _1056_v11 = _dafny.Seq.of(_module.__default.fm0((_dafny.ZERO).minus((_this).f25), globalState), !(p2) || (p2));
        (globalState).f12 = (_1056_v11)[_module.__default.safeIndex((_this).f25, new BigNumber((_1056_v11).length))];
        if (false) {
          r0 = _module.__default.safeModuloInt((_this).f25, (_this).f25);
          let _1057_v12;
          _1057_v12 = _dafny.Set.fromElements(p2);
          let _1058_v13;
          let _nw158 = Array((new BigNumber(3)).toNumber());
          _nw158[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1057_v12).length)), (_this).f25);
          _nw158[(_dafny.ONE).toNumber()] = (_this).f25;
          _nw158[(new BigNumber(2)).toNumber()] = (_this).f25;
          _1058_v13 = _nw158;
          let _1059_v14;
          _1059_v14 = _module.D1.create_DC4(_dafny.Seq.of(p2));
          let _1060_v15;
          _1060_v15 = _module.D1.create_DC7(_1059_v14);
          let _1061_v16;
          _1061_v16 = _dafny.Seq.of(_1060_v15);
          let _rhs194 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_1062_p2, _1063_v10) => function (_1064_i1) {
            return _module.D1.create_DC7(_module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), _1062_p2, _1063_v10));
          })(p2, _1055_v10)), _1061_v16)).length)), (_this).f25);
          let _rhs195 = p2;
          let _rhs196 = _1058_v13;
          let _lhs139 = globalState;
          r1 = _rhs194;
          _lhs139.f12 = _rhs195;
          _1058_v13 = _rhs196;
          let _1065_v17;
          _1065_v17 = new _dafny.CodePoint('h'.codePointAt(0));
          (globalState).f23 = _1065_v17;
          (globalState).f12 = false;
          (globalState).f12 = p1;
        } else {
          let _index175 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length));
          (p3)[_index175] = (_this).f25;
          let _index176 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length));
          (p3)[_index176] = (_dafny.ZERO).minus((_this).f25);
          let _index177 = _module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index177] = _1055_v10;
          let _index178 = _module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index178] = _dafny.Seq.Concat(_1055_v10, _1055_v10);
          (globalState).f23 = (_1055_v10)[_module.__default.safeIndex(_module.__default.safeDivisionInt((p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))], (_this).f25), new BigNumber((_1055_v10).length))];
          let _1066_v18;
          _1066_v18 = _dafny.Map.Empty.slice().updateUnsafe((p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))],_1056_v11);
          let _1067_v19;
          _1067_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1055_v10,new BigNumber((_dafny.Seq.Concat((((_1066_v18).contains((p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))])) ? ((_1066_v18).get((p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))])) : (_dafny.Seq.update(_1056_v11, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kytjdmn")).length), new BigNumber((_1056_v11).length)), p2))), _1056_v11)).length));
          _1067_v19 = (_1067_v19).update(_dafny.Seq.update(_dafny.Seq.Concat(_1055_v10, ((_this).f26)[_module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length))]), _module.__default.safeIndex((p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))], new BigNumber((_dafny.Seq.Concat(_1055_v10, ((_this).f26)[_module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length))])).length)), new _dafny.CodePoint('f'.codePointAt(0))), _module.__default.safeModuloInt((p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))], (p3)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((p3).length))]));
          let _1068_v20;
          let _nw159 = Array((new BigNumber(29)).toNumber()).fill(false);
          _1068_v20 = _nw159;
          let _1069_v21;
          _1069_v21 = _dafny.Seq.of(_1068_v20);
          let _1070_v22;
          _1070_v22 = new _dafny.CodePoint('v'.codePointAt(0));
          let _index179 = _module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length));
          let _rhs197 = (_module.D1.create_DC6(_1070_v22, true, _1055_v10)).dtor_cf9;
          let _rhs198 = (_1056_v11)[_module.__default.safeIndex(new BigNumber((((_this).f26)[_module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length))]).length), new BigNumber((_1056_v11).length))];
          let _rhs199 = _1070_v22;
          let _rhs200 = _1069_v21;
          let _lhs140 = (_this).f26;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(297), new BigNumber(((_this).f26).length));
          let _lhs142 = globalState;
          let _lhs143 = globalState;
          _lhs140[_lhs141] = _rhs197;
          _lhs142.f12 = _rhs198;
          _lhs143.f23 = _rhs199;
          _1069_v21 = _rhs200;
        }
        if (p1) {
          let _1071_v23;
          let _nw160 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _1071_v23 = _nw160;
          let _1072_v24;
          _1072_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(630));
          let _1073_v25;
          _1073_v25 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1072_v24);
          let _1074_v26;
          _1074_v26 = _dafny.Seq.of(_1073_v25);
          let _index180 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1071_v23).length));
          (_1071_v23)[_index180] = (_1074_v26)[_module.__default.safeIndex((_this).f25, new BigNumber((_1074_v26).length))];
          let _index181 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_1071_v23).length));
          (_1071_v23)[_index181] = _1073_v25;
          let _1075_v27;
          let _nw161 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1075_v27 = _nw161;
          let _index182 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1075_v27).length));
          (_1075_v27)[_index182] = !(p1);
          let _1076_v28;
          _1076_v28 = _dafny.MultiSet.fromElements((_this).f25);
          let _1077_v29;
          _1077_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1076_v28);
          let _1078_v30;
          _1078_v30 = _dafny.Set.fromElements(p2, p1);
          let _1079_v31;
          _1079_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber(((((_1077_v29).contains(new BigNumber((_1078_v30).length))) ? ((_1077_v29).get(new BigNumber((_1078_v30).length))) : (_1076_v28))).cardinality()));
          let _1080_v32;
          _1080_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(new BigNumber((_1079_v31).length)).plus((_this).f25));
          let _1081_v33;
          _1081_v33 = _dafny.Seq.of((_this).f25, ((_this).f25).plus((_this).f25));
          let _1082_v34;
          _1082_v34 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _index183 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1075_v27).length));
          let _rhs201 = !(p2);
          let _rhs202 = _1080_v32;
          let _rhs203 = _1081_v33;
          let _rhs204 = ((new BigNumber((_1079_v31).length)).minus(new BigNumber((_1055_v10).length))).multipliedBy((new BigNumber((_1072_v24).length)).minus(new BigNumber((_1082_v34).length)));
          let _lhs144 = _1075_v27;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1075_v27).length));
          let _lhs146 = globalState;
          _lhs144[_lhs145] = _rhs201;
          _1080_v32 = _rhs202;
          _lhs146.f6 = _rhs203;
          r1 = _rhs204;
          let _1083_v35;
          _1083_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1055_v10);
          _1083_v35 = (_1083_v35).Merge(_1083_v35);
          (globalState).f12 = p2;
          let _1084_v36;
          _1084_v36 = _module.D0.create_DC3((_this).f25, false);
          (globalState).f12 = (_1084_v36).dtor_cf5;
        } else {
          let _1085_v37;
          _1085_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,(p2) && (p2));
          (globalState).f12 = (((_1085_v37).contains(!((_this).f25).isEqualTo((_this).f25))) ? ((_1085_v37).get(!((_this).f25).isEqualTo((_this).f25))) : (p2));
          let _1086_v38;
          _1086_v38 = _dafny.Set.fromElements(p3);
          let _1087_v39;
          let _nw162 = Array((new BigNumber(4)).toNumber());
          _nw162[(_dafny.ZERO).toNumber()] = ((false) ? (_dafny.Set.fromElements(p3)) : (_dafny.Set.fromElements(p3, p3, p3)));
          _nw162[(_dafny.ONE).toNumber()] = (_1086_v38).Union(_1086_v38);
          _nw162[(new BigNumber(2)).toNumber()] = _1086_v38;
          _nw162[(new BigNumber(3)).toNumber()] = _1086_v38;
          _1087_v39 = _nw162;
          let _index184 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_1087_v39).length));
          (_1087_v39)[_index184] = (_dafny.Set.fromElements(p3)).Intersect(_1086_v38);
          let _index185 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_1087_v39).length));
          (_1087_v39)[_index185] = _1086_v38;
          let _1088_v40;
          let _nw163 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _1088_v40 = _nw163;
          let _index186 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1088_v40).length));
          (_1088_v40)[_index186] = (_dafny.Map.Empty.slice().updateUnsafe(p2,false)).Merge(_1085_v37);
          let _index187 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((p3).length));
          (p3)[_index187] = ((_this).f25).plus((_dafny.ZERO).minus((_this).f25));
          let _1089_v41;
          _1089_v41 = _dafny.Seq.of((_this).f25, (_this).f25);
          let _index188 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((p3).length));
          (p3)[_index188] = new BigNumber((_1089_v41).length);
          let _1090_v42;
          _1090_v42 = _module.D0.create_DC0(p1);
          let _1091_v43;
          _1091_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1090_v42,p1);
          let _index189 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1088_v40).length));
          let _index190 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((p3).length));
          let _index191 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((p3).length));
          let _rhs205 = _1085_v37;
          let _rhs206 = (((_1056_v11)[_module.__default.safeIndex(_module.__default.fm21(_1091_v43, new BigNumber((_dafny.Seq.UnicodeFromString("mlkroi")).length), p2, _dafny.Seq.UnicodeFromString("aufslk"), globalState), new BigNumber((_1056_v11).length))]) ? ((_this).f25) : ((_this).f25));
          let _rhs207 = _dafny.Seq.Concat(_1055_v10, _1055_v10);
          let _rhs208 = new BigNumber((_1056_v11).length);
          let _lhs147 = _1088_v40;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1088_v40).length));
          let _lhs149 = p3;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((p3).length));
          let _lhs151 = globalState;
          let _lhs152 = p3;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((p3).length));
          _lhs147[_lhs148] = _rhs205;
          _lhs149[_lhs150] = _rhs206;
          _lhs151.f10 = _rhs207;
          _lhs152[_lhs153] = _rhs208;
          let _index192 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((p3).length));
          (p3)[_index192] = (_module.__default.safeDivisionInt((p3)[_module.__default.safeIndex(new BigNumber(105), new BigNumber((p3).length))], (p3)[_module.__default.safeIndex(new BigNumber(105), new BigNumber((p3).length))])).multipliedBy((_this).f25);
          let _1092_v44;
          let _nw164 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _1092_v44 = _nw164;
          let _1093_v45;
          _1093_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,false);
          let _index193 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_1092_v44).length));
          (_1092_v44)[_index193] = _1093_v45;
          let _index194 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_1092_v44).length));
          (_1092_v44)[_index194] = _1093_v45;
        }
        if (!(!(((_this).f25).isLessThan((_this).f25)))) {
          (globalState).f12 = (_module.D0.create_DC2(p2)).dtor_cf3;
          let _1094_v46;
          let _nw165 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1094_v46 = _nw165;
          let _1095_v47;
          _1095_v47 = _dafny.MultiSet.fromElements(p2, p2);
          let _index195 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1094_v46).length));
          (_1094_v46)[_index195] = (_dafny.MultiSet.fromElements(p2)).Union(_1095_v47);
          let _index196 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1094_v46).length));
          let _rhs209 = (((_1095_v47).update(p1, _module.__default.abs(new BigNumber((_1055_v10).length)))).Union(_1095_v47)).update(_module.__default.fm0((_this).f25, globalState), _module.__default.abs((_this).f25));
          let _rhs210 = !(p2);
          let _lhs154 = _1094_v46;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1094_v46).length));
          let _lhs156 = globalState;
          _lhs154[_lhs155] = _rhs209;
          _lhs156.f12 = _rhs210;
          let _1096_v48;
          _1096_v48 = _module.D0.create_DC0(p2);
          let _1097_v49;
          _1097_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1096_v48,!(p1));
          let _1098_v50;
          _1098_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1055_v10).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(360)), function (_1099_i2) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }));
          let _1100_v51;
          _1100_v51 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1101_v52;
          _1101_v52 = _module.D1.create_DC6(_1100_v51, false, _1055_v10);
          let _1102_v53;
          _1102_v53 = _dafny.Seq.of((_1101_v52).dtor_cf9, _1055_v10);
          r0 = _module.__default.fm21(_1097_v49, (_this).f25, false, (((_1098_v50).contains(new BigNumber(((_1102_v53)[_module.__default.safeIndex((_this).f25, new BigNumber((_1102_v53).length))]).length))) ? ((_1098_v50).get(new BigNumber(((_1102_v53)[_module.__default.safeIndex((_this).f25, new BigNumber((_1102_v53).length))]).length))) : (_1055_v10)), globalState);
          (globalState).f12 = !((_dafny.ZERO).minus(((_this).f25).multipliedBy((_this).f25))).isEqualTo(new BigNumber(-63));
          let _1103_v54;
          _1103_v54 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f36);
          let _1104_v55;
          _1104_v55 = _dafny.Set.fromElements((_this).f25, (_this).f25);
          _1103_v54 = (_1103_v54).update((_1104_v55).IsSubsetOf(_1104_v55), (_this).f26);
        } else {
          let _1105_v56;
          let _nw166 = Array((new BigNumber(3)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1105_v56 = _nw166;
          let _1106_v57;
          let _nw167 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1106_v57 = _nw167;
          let _1107_v58;
          _1107_v58 = _dafny.MultiSet.fromElements(_1106_v57, _1106_v57);
          let _index197 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1105_v56).length));
          (_1105_v56)[_index197] = (_1107_v58).update(_1106_v57, _module.__default.abs(new BigNumber((_1056_v11).length)));
          let _1108_v59;
          _1108_v59 = _dafny.MultiSet.fromElements(p2);
          let _index198 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1105_v56).length));
          (_1105_v56)[_index198] = (((p1) ? (_1107_v58) : (_1107_v58))).update(_1106_v57, _module.__default.abs(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1108_v59).cardinality()))).length)));
          let _1109_v60;
          let _init35 = ((_1110_p2, _1111_p1) => function (_1112_i3) {
            return _module.D4.create_DC13((_this).f25, _1110_p2, _1111_p1, true);
          })(p2, p1);
          let _nw168 = Array((new BigNumber(18)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw168.length); _i0_35++) {
            _nw168[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1109_v60 = _nw168;
          let _1113_v61;
          _1113_v61 = _module.D4.create_DC13((_this).f25, p1, p2, p2);
          let _index199 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1109_v60).length));
          (_1109_v60)[_index199] = _1113_v61;
          let _1114_v62;
          _1114_v62 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality()), (_this).f25);
          let _1115_v63;
          _1115_v63 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_1056_v11)).cardinality()), new BigNumber((_1114_v62).cardinality()), (_this).f25, (_this).f25);
          let _index200 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1109_v60).length));
          (_1109_v60)[_index200] = _module.__default.fm22(_1115_v63, (_this).f25, globalState);
          let _1116_v64;
          let _nw169 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _1116_v64 = _nw169;
          let _1117_v65;
          _1117_v65 = _dafny.Set.fromElements(p2);
          let _index201 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1116_v64).length));
          (_1116_v64)[_index201] = _1117_v65;
          let _index202 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1116_v64).length));
          (_1116_v64)[_index202] = (_1117_v65).Intersect((_1117_v65).Difference(_1117_v65));
          let _1118_v66;
          _1118_v66 = _dafny.Seq.of((_this).f25);
          let _1119_v67;
          _1119_v67 = _dafny.Map.Empty.slice().updateUnsafe(((p2) ? (_1118_v66) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-301)), function (_1120_i4) {
            return (_this).f25;
          }))),p2);
          _1119_v67 = ((_1119_v67).Merge(_1119_v67)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1118_v66,false)).Merge(_1119_v67));
          let _1121_v68;
          _1121_v68 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f25);
          let _1122_v69;
          let _1123_v70;
          let _1124_v71;
          let _out28;
          let _out29;
          let _out30;
          let _outcollector10 = (_this).m15(!(p1) || (!(p2)), _1121_v68, new BigNumber((_1055_v10).length), globalState);
          _out28 = _outcollector10[0];
          _out29 = _outcollector10[1];
          _out30 = _outcollector10[2];
          _1122_v69 = _out28;
          _1123_v70 = _out29;
          _1124_v71 = _out30;
        }
      }
      let _1125_v72;
      _1125_v72 = _module.D0.create_DC2(false);
      let _1126_v73;
      _1126_v73 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1125_v72);
      let _1127_v74;
      _1127_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      let _hi5 = (((_1127_v74).contains((_dafny.ZERO).minus((_this).f25))) ? ((_1127_v74).get((_dafny.ZERO).minus((_this).f25))) : ((_this).f25));
      for (let _1128_i5 = new BigNumber((_1126_v73).length); _1128_i5.isLessThan(_hi5); _1128_i5 = _1128_i5.plus(_dafny.ONE)) {
        let _1129_v75;
        _1129_v75 = _dafny.Seq.of(p2);
        let _1130_v76;
        _1130_v76 = _module.D1.create_DC4(_1129_v75);
        let _1131_v77;
        _1131_v77 = _dafny.Set.fromElements(_1130_v76, _module.D1.create_DC4(_dafny.Seq.of(p1)), _1130_v76);
        _1131_v77 = ((_1131_v77).Difference(_1131_v77)).Union((_1131_v77).Difference(_1131_v77));
        let _1132_v78;
        _1132_v78 = _module.D4.create_DC12();
        let _rhs211 = (_dafny.ZERO).minus((_this).f25);
        let _rhs212 = _1132_v78;
        r0 = _rhs211;
        _1132_v78 = _rhs212;
        if (true) {
          let _1133_v79;
          _1133_v79 = _dafny.Seq.UnicodeFromString("opd");
          r0 = (_1128_i5).plus(new BigNumber((_dafny.Seq.Concat(_1133_v79, _1133_v79)).length));
          let _index203 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length));
          (p3)[_index203] = (_this).f25;
          let _index204 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length));
          (p3)[_index204] = ((_this).f25).minus(new BigNumber(816));
          let _1134_v80;
          _1134_v80 = _dafny.Set.fromElements((_this).f25);
          let _1135_v81;
          _1135_v81 = _dafny.Map.Empty.slice().updateUnsafe(p1,!(p1));
          let _1136_v82;
          _1136_v82 = _dafny.Seq.of((p3)[_module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length))], new BigNumber((_1135_v81).length), new BigNumber((_1134_v80).length), new BigNumber((_1129_v75).length), _1128_i5);
          let _index205 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length));
          let _rhs213 = !(p1);
          let _rhs214 = _module.__default.safeModuloInt(new BigNumber((_1133_v79).length), ((p3)[_module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length))]).plus((_dafny.ZERO).minus(new BigNumber((_1134_v80).length))));
          let _rhs215 = ((p1) ? (new BigNumber(-232)) : ((_1136_v82)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_1136_v82).length))]));
          let _lhs157 = globalState;
          let _lhs158 = p3;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length));
          _lhs157.f12 = _rhs213;
          r0 = _rhs214;
          _lhs158[_lhs159] = _rhs215;
          let _1137_v83;
          _1137_v83 = _module.D0.create_DC0(p1);
          let _1138_v84;
          _1138_v84 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1139_v85;
          _1139_v85 = _dafny.MultiSet.fromElements(_module.__default.fm0(new BigNumber(166), globalState));
          let _1140_v86;
          _1140_v86 = _dafny.Map.Empty.slice().updateUnsafe((p3)[_module.__default.safeIndex(new BigNumber(770), new BigNumber((p3).length))],p2);
          r0 = _module.__default.safeDivisionInt((_module.__default.fm21((_dafny.Map.Empty.slice().updateUnsafe(_1137_v83,_module.__default.fm0((_this).f25, globalState))).update(_1137_v83, p1), _1128_i5, false, _dafny.Seq.update((_module.D1.create_DC6(new _dafny.CodePoint('c'.codePointAt(0)), p2, _1133_v79)).dtor_cf9, _module.__default.safeIndex((_this).f25, new BigNumber(((_module.D1.create_DC6(new _dafny.CodePoint('c'.codePointAt(0)), p2, _1133_v79)).dtor_cf9).length)), _1138_v84), globalState)).multipliedBy(new BigNumber(-305)), (((_1139_v85).contains((((_1140_v86).contains((_this).f25)) ? ((_1140_v86).get((_this).f25)) : (p1)))) ? ((_1139_v85).get((((_1140_v86).contains((_this).f25)) ? ((_1140_v86).get((_this).f25)) : (p1)))) : (_1128_i5)));
          (globalState).f12 = (new BigNumber(934)).isLessThanOrEqualTo(_1128_i5);
        } else {
          let _1141_v87;
          _1141_v87 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1142_v88;
          _1142_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1141_v87,(_this).f25);
          let _1143_v89;
          _1143_v89 = _dafny.Set.fromElements((_this).f25);
          (globalState).f12 = (_1128_i5).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_1142_v88).length), new BigNumber((_1143_v89).length)));
          let _1144_v90;
          _1144_v90 = _dafny.Seq.UnicodeFromString("xtm");
          let _1145_v91;
          _1145_v91 = _dafny.MultiSet.fromElements(false, p1, _dafny.Seq.IsPrefixOf(_1144_v90, _1144_v90));
          let _1146_v92;
          _1146_v92 = _dafny.Set.fromElements(p1, p1);
          let _1147_v93;
          _1147_v93 = _dafny.Set.fromElements(_1141_v87);
          r1 = (((_1145_v91).contains((_1146_v92).IsSubsetOf(_1146_v92))) ? ((_1145_v91).get((_1146_v92).IsSubsetOf(_1146_v92))) : (new BigNumber((_1147_v93).length)));
          (globalState).f12 = ((!(p1) || (!(_module.__default.fm0((_this).f25, globalState)))) ? (p2) : (p2));
          let _1148_v94;
          _1148_v94 = _module.D0.create_DC0(p1);
          let _pat_let_tv20 = p1;
          let _1149_v95;
          _1149_v95 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let18_0) {
            return function (_1150_dt__update__tmp_h0) {
              return function (_pat_let19_0) {
                return function (_1151_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_1151_dt__update_hcf0_h0);
                }(_pat_let19_0);
              }(_pat_let_tv20);
            }(_pat_let18_0);
          }(_1148_v94),p1);
          (globalState).f12 = !((_this).f25).isEqualTo(_module.__default.safeDivisionInt((_this).f25, _module.__default.fm21(_1149_v95, new BigNumber(555), p1, _dafny.Seq.UnicodeFromString("dbb"), globalState)));
          _1129_v75 = _1129_v75;
        }
        let _1152_v96;
        _1152_v96 = _module.D0.create_DC0(p2);
        let _1153_v97;
        _1153_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1152_v96,p1);
        let _1154_v98;
        _1154_v98 = _dafny.Seq.UnicodeFromString("sjlaotpx");
        let _1155_v99;
        let _init36 = ((_1156_p1) => function (_1157_i6) {
          return _1156_p1;
        })(p1);
        let _nw170 = Array((new BigNumber(20)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw170.length); _i0_36++) {
          _nw170[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1155_v99 = _nw170;
        let _1158_v100;
        _1158_v100 = _dafny.Seq.of(_1155_v99, _1155_v99, _1155_v99);
        let _1159_v101;
        _1159_v101 = _module.D4.create_DC13((_module.__default.fm21(_1153_v97, new BigNumber((_1154_v98).length), p1, _1154_v98, globalState)).multipliedBy(new BigNumber((_1129_v75).length)), false, _module.__default.fm0((((_1127_v74).contains(new BigNumber((_1158_v100).length))) ? ((_1127_v74).get(new BigNumber((_1158_v100).length))) : (new BigNumber(546))), globalState), !(!(p2)));
        _1159_v101 = _1159_v101;
      }
      (globalState).f12 = (p2) || ((p1) && (p1));
      let _1160_v102;
      let _nw171 = Array((new BigNumber(8)).toNumber()).fill(false);
      _1160_v102 = _nw171;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1160_v102).length))) {
        let _1161_i7 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1161_i7)) && ((_1161_i7).isLessThan(new BigNumber((_1160_v102).length))))) {
          (_1160_v102)[(_1161_i7)] = ((_dafny.Set.fromElements((_this).f25, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f25)).length)))).Difference(function () {
            let _coll47 = new _dafny.Set();
            for (const _compr_47 of _dafny.IntegerRange(new BigNumber(36), new BigNumber(152))) {
              let _1162_v103 = _compr_47;
              if (((new BigNumber(36)).isLessThanOrEqualTo(_1162_v103)) && ((_1162_v103).isLessThan(new BigNumber(152)))) {
                _coll47.add(_module.__default.safeModuloInt(_1162_v103, (_dafny.ZERO).minus((_this).f25)));
              }
            }
            return _coll47;
          }())).IsSubsetOf((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), function (_1163_i8) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          })).length))).Difference(_dafny.Set.fromElements((_module.D4.create_DC13(new BigNumber((_1127_v74).length), p1, p1, !(p2))).dtor_cf19, (_this).f25)));
        }
      }
      let _1164_v104;
      _1164_v104 = _module.D0.create_DC0(!(true));
      let _1165_v105;
      _1165_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1164_v104,_module.__default.fm0((_this).f25, globalState));
      let _1166_v106;
      _1166_v106 = _dafny.MultiSet.fromElements((_this).f25);
      let _1167_v107;
      _1167_v107 = _dafny.Set.fromElements(_1166_v106);
      let _1168_v108;
      _1168_v108 = _dafny.Seq.of(_module.__default.fm21(_1165_v105, new BigNumber((_1167_v107).length), p2, _dafny.Seq.UnicodeFromString("kgfs"), globalState));
      let _1169_v109;
      _1169_v109 = _1168_v108;
      let _pat_let_tv21 = p2;
      (globalState).f12 = function (_source19) {
        let _1170___mcc_h0 = _source19;
        let _1171_cf11 = _1170___mcc_h0;
        return _pat_let_tv21;
      }(_1169_v109);
      r0 = (_this).f25;
      let _1172_v110;
      _1172_v110 = _module.D1.create_DC4(_dafny.Seq.of(!(p1)));
      let _1173_v111;
      _1173_v111 = _dafny.MultiSet.fromElements(_1172_v110);
      r1 = _module.__default.safeDivisionInt((((_1173_v111).contains(_1172_v110)) ? ((_1173_v111).get(_1172_v110)) : ((_this).f25)), (_this).f25);
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1174_v0;
      _1174_v0 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_this).f25).plus((_this).f25)),p1);
      if ((((_1174_v0).contains((_this).f25)) ? ((_1174_v0).get((_this).f25)) : (true))) {
        let _1175_v1;
        _1175_v1 = _dafny.Seq.UnicodeFromString("y");
        if (!(_dafny.Seq.IsPrefixOf(_1175_v1, _1175_v1))) {
          r0 = (_this).f25;
          let _1176_v2;
          let _nw172 = new _module.C2();
          _nw172.__ctor();
          _1176_v2 = _nw172;
          let _1177_v3;
          _1177_v3 = _dafny.Set.fromElements((_this).f25, (_this).f25, (_this).f25);
          (globalState).f12 = (_1177_v3).IsDisjointFrom((_1177_v3).Intersect(_1177_v3));
          let _1178_v4;
          let _init37 = ((_1179_p1) => function (_1180_i0) {
            return _1179_p1;
          })(p1);
          let _nw173 = Array((new BigNumber(4)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw173.length); _i0_37++) {
            _nw173[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1178_v4 = _nw173;
          let _1181_v5;
          _1181_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1178_v4,(_this).f25);
          let _1182_v6;
          _1182_v6 = _module.D6.create_DC18(_1181_v5);
          let _rhs216 = _dafny.Seq.UnicodeFromString("lk");
          let _rhs217 = !(((_1182_v6).dtor_cf28).Merge(_1181_v5)).equals(_dafny.Map.Empty.slice().updateUnsafe(_1178_v4,(_this).f25));
          let _lhs160 = globalState;
          _1175_v1 = _rhs216;
          _lhs160.f12 = _rhs217;
          (globalState).f12 = _module.__default.fm0((_this).f25, globalState);
        } else {
          let _1183_v7;
          let _nw174 = new _module.C1();
          _nw174.__ctor((_this).f25, (_this).f26);
          _1183_v7 = _nw174;
          let _1184_v8;
          _1184_v8 = _dafny.Seq.of(p1, p1);
          let _1185_v9;
          _1185_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1184_v8)[_module.__default.safeIndex((_this).f25, new BigNumber((_1184_v8).length))],!(true));
          let _1186_v10;
          _1186_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p1,p1));
          let _1187_v11;
          _1187_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1185_v9,new BigNumber((_1186_v10).length));
          r0 = (((_1187_v11).contains(_1185_v9)) ? ((_1187_v11).get(_1185_v9)) : ((_this).f25));
          r0 = (_this).f25;
          let _1188_v12;
          _1188_v12 = _module.D0.create_DC3(new BigNumber((_1185_v9).length), !(p1));
          let _1189_v13;
          _1189_v13 = _module.D6.create_DC19(_1188_v12, (_this).f25);
          let _1190_v14;
          let _nw175 = new _module.C1();
          _nw175.__ctor(((_dafny.ZERO).minus((_this).f25)).plus((_1189_v13).dtor_cf30), (_this).f36);
          _1190_v14 = _nw175;
          r0 = ((_this).f25).plus((_this).f25);
        }
        let _1191_v15;
        _1191_v15 = _dafny.Set.fromElements(p1, p1, p1, p1);
        let _1192_v16;
        _1192_v16 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f25).isLessThan((_this).f25),_1191_v15);
        _1192_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1191_v15);
        let _1193_v17;
        let _init38 = ((_1194_p1) => function (_1195_i1) {
          return !(_1194_p1);
        })(p1);
        let _nw176 = Array((new BigNumber(14)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw176.length); _i0_38++) {
          _nw176[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1193_v17 = _nw176;
        _1193_v17 = _1193_v17;
        let _rhs218 = ((_dafny.ZERO).minus((_this).f25)).isLessThanOrEqualTo((_this).f25);
        let _rhs219 = _1193_v17;
        let _lhs161 = globalState;
        _lhs161.f12 = _rhs218;
        _1193_v17 = _rhs219;
        let _index206 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1193_v17).length));
        (_1193_v17)[_index206] = true;
        let _index207 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_1193_v17).length));
        (_1193_v17)[_index207] = true;
      } else {
        let _1196_v18;
        _1196_v18 = _dafny.Set.fromElements(p1, p1);
        let _1197_v19;
        _1197_v19 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1198_v20;
        _1198_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1197_v19);
        let _1199_v21;
        _1199_v21 = _dafny.MultiSet.fromElements(_1198_v20, _1198_v20, _1198_v20, _1198_v20, _1198_v20);
        let _1200_v22;
        let _nw177 = Array((new BigNumber(16)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = !(_module.__default.fm0((_this).f25, globalState));
        _nw177[(_dafny.ONE).toNumber()] = !(p1);
        _nw177[(new BigNumber(2)).toNumber()] = p1;
        _nw177[(new BigNumber(3)).toNumber()] = _module.__default.fm0((_this).f25, globalState);
        _nw177[(new BigNumber(4)).toNumber()] = p1;
        _nw177[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_module.__default.fm0((_this).f25, globalState))).IsProperSubsetOf(_1196_v18);
        _nw177[(new BigNumber(6)).toNumber()] = (p1) && ((((_1174_v0).contains((_this).f25)) ? ((_1174_v0).get((_this).f25)) : (p1)));
        _nw177[(new BigNumber(7)).toNumber()] = (_1199_v21).equals(_1199_v21);
        _nw177[(new BigNumber(8)).toNumber()] = p1;
        _nw177[(new BigNumber(9)).toNumber()] = p1;
        _nw177[(new BigNumber(10)).toNumber()] = p1;
        _nw177[(new BigNumber(11)).toNumber()] = p1;
        _nw177[(new BigNumber(12)).toNumber()] = (((_1174_v0).contains((_this).f25)) ? ((_1174_v0).get((_this).f25)) : (!(true)));
        _nw177[(new BigNumber(13)).toNumber()] = !(false);
        _nw177[(new BigNumber(14)).toNumber()] = false;
        _nw177[(new BigNumber(15)).toNumber()] = p1;
        _1200_v22 = _nw177;
        let _1201_v23;
        _1201_v23 = _dafny.Seq.of(_1200_v22);
        let _index208 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1200_v22).length));
        (_1200_v22)[_index208] = !(_dafny.Seq.IsPrefixOf(_1201_v23, _1201_v23));
        let _1202_v24;
        _1202_v24 = _dafny.Seq.UnicodeFromString("yqpkwm");
        let _1203_v25;
        _1203_v25 = _dafny.MultiSet.fromElements((_this).f25, (_dafny.ZERO).minus((_this).f25));
        let _1204_v26;
        _1204_v26 = _dafny.MultiSet.fromElements(_1203_v25, _dafny.MultiSet.fromElements((_this).f25));
        let _index209 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1200_v22).length));
        let _rhs220 = _dafny.Seq.IsPrefixOf(_1202_v24, _1202_v24);
        let _rhs221 = new BigNumber((_module.__default.fm37(_1204_v26, globalState)).length);
        let _lhs162 = _1200_v22;
        let _lhs163 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1200_v22).length));
        _lhs162[_lhs163] = _rhs220;
        r0 = _rhs221;
        let _1205_v27;
        _1205_v27 = _dafny.Seq.of(_1202_v24);
        _1202_v24 = _module.__default.fm2(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1205_v27, _module.__default.safeIndex((_this).f25, new BigNumber((_1205_v27).length)), _1202_v24), _1205_v27)).length), (_this).f25, globalState);
        let _index210 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1200_v22).length));
        let _rhs222 = _1202_v24;
        let _rhs223 = p1;
        let _lhs164 = _1200_v22;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1200_v22).length));
        _1202_v24 = _rhs222;
        _lhs164[_lhs165] = _rhs223;
        r0 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f25), (_dafny.ZERO).minus(new BigNumber(-231))), (new BigNumber((_1174_v0).length)).plus((_this).f25));
        let _1206_v28;
        let _nw178 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1206_v28 = _nw178;
        let _index211 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1206_v28).length));
        (_1206_v28)[_index211] = (_this).f25;
        let _index212 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1206_v28).length));
        (_1206_v28)[_index212] = (new BigNumber(855)).minus((_this).f25);
      }
      r0 = (_this).f25;
      let _1207_v29;
      let _init39 = ((_1208_p1) => function (_1209_i2) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_1208_p1,_dafny.Seq.UnicodeFromString("embe"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_1208_p1),_dafny.Seq.UnicodeFromString("tvjmfhrm")));
      })(p1);
      let _nw179 = Array((new BigNumber(4)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw179.length); _i0_39++) {
        _nw179[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1207_v29 = _nw179;
      let _1210_v30;
      _1210_v30 = _dafny.Seq.UnicodeFromString("awybtal");
      let _1211_v31;
      _1211_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1210_v30);
      let _index213 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1207_v29).length));
      (_1207_v29)[_index213] = _1211_v31;
      let _index214 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1207_v29).length));
      (_1207_v29)[_index214] = ((_1211_v31).Merge(_1211_v31)).Merge((_1211_v31).Merge(_1211_v31));
      let _hi6 = ((_this).f25).multipliedBy((_this).f25);
      for (let _1212_i3 = (_this).f25; _1212_i3.isLessThan(_hi6); _1212_i3 = _1212_i3.plus(_dafny.ONE)) {
        let _1213_v32;
        _1213_v32 = _dafny.MultiSet.fromElements(new BigNumber((_1210_v30).length));
        let _1214_v33;
        _1214_v33 = _dafny.Seq.of(p1);
        let _1215_v34;
        _1215_v34 = _dafny.Map.Empty.slice().updateUnsafe((((_1174_v0).contains(_1212_i3)) ? ((_1174_v0).get(_1212_i3)) : (_module.__default.fm0(new BigNumber((_1214_v33).length), globalState))),p1);
        let _1216_v35;
        _1216_v35 = _module.D0.create_DC3((((_1213_v32).contains(_1212_i3)) ? ((_1213_v32).get(_1212_i3)) : (new BigNumber((_1215_v34).length))), p1);
        _1216_v35 = _1216_v35;
        let _1217_v36;
        let _init40 = ((_1218_p1, _1219_i3) => function (_1220_i4) {
          return ((_1218_p1) ? (_dafny.Map.Empty.slice().updateUnsafe(_1218_p1,(_module.D5.create_DC17(_1219_i3, _1218_p1)).dtor_cf26)) : (_dafny.Map.Empty.slice().updateUnsafe(_1218_p1,(_this).f25)));
        })(p1, _1212_i3);
        let _nw180 = Array((new BigNumber(16)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw180.length); _i0_40++) {
          _nw180[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1217_v36 = _nw180;
        let _nw181 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _1217_v36 = _nw181;
        (globalState).f12 = ((p1) ? (p1) : (!(p1)));
        let _1221_v37;
        _1221_v37 = _module.D0.create_DC0(p1);
        let _1222_v38;
        _1222_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1221_v37,p1);
        let _1223_v39;
        _1223_v39 = _dafny.Seq.of(_module.__default.fm21(_1222_v38, _1212_i3, _module.__default.fm0(new BigNumber((_1210_v30).length), globalState), _dafny.Seq.UnicodeFromString("noqaw"), globalState), (_this).f25, _1212_i3);
        let _1224_v40;
        _1224_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1223_v39,_1212_i3);
        r0 = (((_1224_v40).contains(_1223_v39)) ? ((_1224_v40).get(_1223_v39)) : (((_this).f25).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_1212_i3)).length)))));
      }
      let _1225_v41;
      let _nw182 = new _module.C3();
      _nw182.__ctor((_this).f25, p1, (_this).f25, (_this).f36);
      _1225_v41 = _nw182;
      _1225_v41 = _1225_v41;
      let _1226_v42;
      _1226_v42 = new _dafny.CodePoint('y'.codePointAt(0));
      let _1227_v43;
      _1227_v43 = _module.D1.create_DC6(_1226_v42, p1, _1210_v30);
      let _1228_i5;
      _1228_i5 = _dafny.ZERO;
      L10: {
        while ((_module.D1.create_DC6(_1226_v42, p1, (_1227_v43).dtor_cf9)).dtor_cf8) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1228_i5)) {
              break L10;
            }
            _1228_i5 = (_1228_i5).plus(_dafny.ONE);
            let _1229_v44;
            let _nw183 = new _module.C3();
            _nw183.__ctor((_this).f25, p1, (_dafny.ZERO).minus((_this).f25), (_1225_v41).f26);
            _1229_v44 = _nw183;
            let _1230_v45;
            let _init41 = ((_1231_p1) => function (_1232_i6) {
              return _dafny.MultiSet.fromElements(_1231_p1);
            })(p1);
            let _nw184 = Array((new BigNumber(3)).toNumber());
            for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw184.length); _i0_41++) {
              _nw184[_i0_41] = _init41(new BigNumber(_i0_41));
            }
            _1230_v45 = _nw184;
            _1230_v45 = _1230_v45;
            let _1233_v46;
            _1233_v46 = _dafny.MultiSet.fromElements((_1229_v44).f38);
            let _1234_v47;
            _1234_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v46,p1);
            let _1235_v49;
            _1235_v49 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm0((_1225_v41).f25, globalState)),(_1234_v47).Merge(function () {
              let _coll48 = new _dafny.Map();
              for (const _compr_48 of (_1234_v47).Keys.Elements) {
                let _1236_v48 = _compr_48;
                if ((_1234_v47).contains(_1236_v48)) {
                  _coll48.push([_1236_v48,p1]);
                }
              }
              return _coll48;
            }()));
            _1235_v49 = (_1235_v49).update(((_1229_v44).f37).isLessThan(new BigNumber(754)), _dafny.Map.Empty.slice().updateUnsafe(_1233_v46,(_1229_v44).f38));
            let _1237_v50;
            let _nw185 = Array((new BigNumber(3)).toNumber()).fill(false);
            _1237_v50 = _nw185;
            let _index215 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_1237_v50).length));
            (_1237_v50)[_index215] = (_1229_v44).f38;
            let _index216 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_1237_v50).length));
            (_1237_v50)[_index216] = !((_1229_v44).f38);
          }
        }
      }
      let _1238_v51;
      let _nw186 = Array((new BigNumber(28)).toNumber()).fill(false);
      _1238_v51 = _nw186;
      let _1239_v52;
      _1239_v52 = _dafny.MultiSet.fromElements(p1, _module.__default.fm0(new BigNumber((_dafny.MultiSet.fromElements(_1238_v51)).cardinality()), globalState));
      r0 = (((_1239_v52).contains(p1)) ? ((_1239_v52).get(p1)) : (new BigNumber(38)));
      return r0;
    }
    m15(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let _1240_v0;
      _1240_v0 = new BigNumber(227);
      let _1241_v1;
      _1241_v1 = _module.D0.create_DC0(p0);
      let _1242_v2;
      _1242_v2 = _dafny.Seq.UnicodeFromString("q");
      let _1243_v3;
      let _nw187 = Array((new BigNumber(25)).toNumber());
      _nw187[(_dafny.ZERO).toNumber()] = _module.D0.create_DC0(p0);
      _nw187[(_dafny.ONE).toNumber()] = _module.D0.create_DC0(p0);
      _nw187[(new BigNumber(2)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(3)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(4)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(5)).toNumber()] = _module.D0.create_DC0(!(p0));
      _nw187[(new BigNumber(6)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(7)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(8)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(9)).toNumber()] = _module.__default.fm53(p0, p2, p0, globalState);
      _nw187[(new BigNumber(10)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(11)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(12)).toNumber()] = ((p0) ? (_1241_v1) : (_1241_v1));
      _nw187[(new BigNumber(13)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(14)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(15)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(16)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(17)).toNumber()] = _module.D0.create_DC0(p0);
      _nw187[(new BigNumber(18)).toNumber()] = _module.D0.create_DC0(true);
      _nw187[(new BigNumber(19)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(20)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(21)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(22)).toNumber()] = _1241_v1;
      _nw187[(new BigNumber(23)).toNumber()] = ((!((_module.__default.fm45(new BigNumber((_1242_v2).length), _1240_v0, globalState)).dtor_cf25)) ? (_module.D0.create_DC0(p0)) : (_1241_v1));
      _nw187[(new BigNumber(24)).toNumber()] = _module.D0.create_DC0(p0);
      _1243_v3 = _nw187;
      let _pat_let_tv22 = p0;
      let _pat_let_tv23 = p0;
      let _index217 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1243_v3).length));
      (_1243_v3)[_index217] = function (_pat_let20_0) {
        return function (_1246_dt__update__tmp_h1) {
          return function (_pat_let23_0) {
            return function (_1247_dt__update_hcf0_h1) {
              return _module.D0.create_DC0(_1247_dt__update_hcf0_h1);
            }(_pat_let23_0);
          }(_pat_let_tv23);
        }(_pat_let20_0);
      }(function (_pat_let21_0) {
        return function (_1244_dt__update__tmp_h0) {
          return function (_pat_let22_0) {
            return function (_1245_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_1245_dt__update_hcf0_h0);
            }(_pat_let22_0);
          }(_pat_let_tv22);
        }(_pat_let21_0);
      }(_module.D0.create_DC0(p0)));
      let _1248_v4;
      _1248_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v2,!(p0));
      let _1249_v5;
      _1249_v5 = _dafny.Seq.of(_1242_v2);
      let _index218 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1243_v3).length));
      let _rhs224 = p0;
      let _rhs225 = p0;
      let _rhs226 = _1240_v0;
      let _rhs227 = _1241_v1;
      let _rhs228 = (((_1248_v4).contains((_1249_v5)[_module.__default.safeIndex(p2, new BigNumber((_1249_v5).length))])) ? ((_1248_v4).get((_1249_v5)[_module.__default.safeIndex(p2, new BigNumber((_1249_v5).length))])) : (((_this).f25).isEqualTo((_dafny.ZERO).minus(p2))));
      let _lhs166 = globalState;
      let _lhs167 = _1243_v3;
      let _lhs168 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1243_v3).length));
      _lhs166.f12 = _rhs224;
      r1 = _rhs225;
      _1240_v0 = _rhs226;
      _lhs167[_lhs168] = _rhs227;
      r1 = _rhs228;
      let _1250_v6;
      _1250_v6 = new _dafny.CodePoint('d'.codePointAt(0));
      _1242_v2 = _dafny.Seq.update(_1242_v2, _module.__default.safeIndex(new BigNumber(-86), new BigNumber((_1242_v2).length)), _1250_v6);
      let _1251_v7;
      _1251_v7 = _module.D1.create_DC6(_1250_v6, false, _1242_v2);
      let _1252_v8;
      let _init42 = ((_1253_p0) => function (_1254_i0) {
        return _1253_p0;
      })(p0);
      let _nw188 = Array((new BigNumber(24)).toNumber());
      for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw188.length); _i0_42++) {
        _nw188[_i0_42] = _init42(new BigNumber(_i0_42));
      }
      _1252_v8 = _nw188;
      let _1255_v9;
      _1255_v9 = _module.D12.create_DC33(_1252_v8, _1252_v8, p0, _1240_v0);
      let _1256_v10;
      _1256_v10 = _dafny.MultiSet.fromElements(_1255_v9);
      let _1257_v11;
      let _nw189 = Array((new BigNumber(17)).toNumber());
      _nw189[(_dafny.ZERO).toNumber()] = false;
      _nw189[(_dafny.ONE).toNumber()] = p0;
      _nw189[(new BigNumber(2)).toNumber()] = p0;
      _nw189[(new BigNumber(3)).toNumber()] = p0;
      _nw189[(new BigNumber(4)).toNumber()] = p0;
      _nw189[(new BigNumber(5)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf((_1251_v7).dtor_cf9, _1242_v2));
      _nw189[(new BigNumber(6)).toNumber()] = p0;
      _nw189[(new BigNumber(7)).toNumber()] = p0;
      _nw189[(new BigNumber(8)).toNumber()] = !(false);
      _nw189[(new BigNumber(9)).toNumber()] = (_1256_v10).IsProperSubsetOf(_1256_v10);
      _nw189[(new BigNumber(10)).toNumber()] = p0;
      _nw189[(new BigNumber(11)).toNumber()] = (false) || (!(p0));
      _nw189[(new BigNumber(12)).toNumber()] = !(p0);
      _nw189[(new BigNumber(13)).toNumber()] = p0;
      _nw189[(new BigNumber(14)).toNumber()] = (p0) || (!(p0));
      _nw189[(new BigNumber(15)).toNumber()] = p0;
      _nw189[(new BigNumber(16)).toNumber()] = p0;
      _1257_v11 = _nw189;
      let _1258_v12;
      _1258_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1257_v11);
      _1257_v11 = (((_1258_v12).contains(p2)) ? ((_1258_v12).get(p2)) : (_1257_v11));
      let _1259_v13;
      _1259_v13 = _dafny.Seq.of(true);
      let _1260_v14;
      _1260_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1259_v13).length),p0);
      let _1261_v15;
      _1261_v15 = _dafny.Seq.of((((_1260_v14).contains((((p1).contains(p0)) ? ((p1).get(p0)) : (_module.__default.fm31(_1242_v2, globalState))))) ? ((_1260_v14).get((((p1).contains(p0)) ? ((p1).get(p0)) : (_module.__default.fm31(_1242_v2, globalState))))) : (p0)));
      let _1262_v16;
      _1262_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(998),new BigNumber((_1259_v13).length));
      let _1263_v19;
      _1263_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1262_v16);
      let _1264_v21;
      _1264_v21 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_1242_v2)).cardinality()), _1240_v0, (_this).f25);
      let _1265_v22;
      _1265_v22 = _module.D14.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(_1240_v0,_1240_v0));
      let _1266_v24;
      _1266_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1240_v0);
      let _1267_v25;
      let _nw190 = Array((new BigNumber(18)).toNumber());
      _nw190[(_dafny.ZERO).toNumber()] = _1262_v16;
      _nw190[(_dafny.ONE).toNumber()] = _1262_v16;
      _nw190[(new BigNumber(2)).toNumber()] = function () {
        let _coll49 = new _dafny.Map();
        for (const _compr_49 of (function () {
          let _coll50 = new _dafny.Set();
          for (const _compr_50 of _dafny.IntegerRange(new BigNumber(-449), new BigNumber(406))) {
            let _1268_v18 = _compr_50;
            if (((new BigNumber(-449)).isLessThanOrEqualTo(_1268_v18)) && ((_1268_v18).isLessThan(new BigNumber(406)))) {
              _coll50.add((_1268_v18).multipliedBy(_1240_v0));
            }
          }
          return _coll50;
        }()).Elements) {
          let _1269_v17 = _compr_49;
          if ((function () {
            let _coll51 = new _dafny.Set();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(-449), new BigNumber(406))) {
              let _1270_v18 = _compr_51;
              if (((new BigNumber(-449)).isLessThanOrEqualTo(_1270_v18)) && ((_1270_v18).isLessThan(new BigNumber(406)))) {
                _coll51.add((_1270_v18).multipliedBy(_1240_v0));
              }
            }
            return _coll51;
          }()).contains(_1269_v17)) {
            _coll49.push([(_1269_v17).plus(p2),p2]);
          }
        }
        return _coll49;
      }();
      _nw190[(new BigNumber(3)).toNumber()] = _1262_v16;
      _nw190[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1242_v2).length),p2);
      _nw190[(new BigNumber(5)).toNumber()] = _1262_v16;
      _nw190[(new BigNumber(6)).toNumber()] = _1262_v16;
      _nw190[(new BigNumber(7)).toNumber()] = (((((_1263_v19).contains((_this).f25)) ? ((_1263_v19).get((_this).f25)) : (_1262_v16))).update((_this).f25, (_this).f25)).Merge(_1262_v16);
      _nw190[(new BigNumber(8)).toNumber()] = function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of (_1264_v21).Elements) {
          let _1271_v20 = _compr_52;
          if (_dafny.Seq.contains(_1264_v21, _1271_v20)) {
            _coll52.push([_module.__default.safeDivisionInt(_1271_v20, new BigNumber(-531)),new BigNumber(-150)]);
          }
        }
        return _coll52;
      }();
      _nw190[(new BigNumber(9)).toNumber()] = (_1262_v16).update(p2, _1240_v0);
      _nw190[(new BigNumber(10)).toNumber()] = _1262_v16;
      _nw190[(new BigNumber(11)).toNumber()] = _1262_v16;
      _nw190[(new BigNumber(12)).toNumber()] = ((_1265_v22).dtor_cf57).Merge(function () {
        let _coll53 = new _dafny.Map();
        for (const _compr_53 of _dafny.IntegerRange(new BigNumber(-965), new BigNumber(817))) {
          let _1272_v23 = _compr_53;
          if (((new BigNumber(-965)).isLessThanOrEqualTo(_1272_v23)) && ((_1272_v23).isLessThan(new BigNumber(817)))) {
            _coll53.push([(_1272_v23).minus(p2),_1240_v0]);
          }
        }
        return _coll53;
      }());
      _nw190[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1240_v0,_module.__default.fm31(_1242_v2, globalState));
      _nw190[(new BigNumber(14)).toNumber()] = (_1262_v16).Merge(_1262_v16);
      _nw190[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1266_v24).length),(_this).f25);
      _nw190[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      _nw190[(new BigNumber(17)).toNumber()] = _1262_v16;
      _1267_v25 = _nw190;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1267_v25).length))) {
        let _1273_i1 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1273_i1)) && ((_1273_i1).isLessThan(new BigNumber((_1267_v25).length))))) {
          (_1267_v25)[(_1273_i1)] = _1262_v16;
        }
      }
      let _index219 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length));
      (_1252_v8)[_index219] = p0;
      let _index220 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length));
      (_1252_v8)[_index220] = p0;
      let _1274_v26;
      _1274_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1252_v8,(p2).isLessThanOrEqualTo((_this).f25));
      if ((((_1274_v26).contains(_1252_v8)) ? ((_1274_v26).get(_1252_v8)) : (p0))) {
        if ((!(!(p0)) || ((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))])) || (p0)) {
          let _1275_v27;
          let _nw191 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1275_v27 = _nw191;
          let _1276_v28;
          _1276_v28 = _module.D13.create_DC35(_1275_v27);
          _1276_v28 = ((p0) ? (_module.D13.create_DC35(_1275_v27)) : (_module.D13.create_DC35(_1275_v27)));
          let _index221 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length));
          (_1252_v8)[_index221] = (((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]) ? (((false) ? (p0) : (!((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))])))) : (p0));
          (globalState).f23 = _1250_v6;
          let _1277_v29;
          _1277_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1262_v16,new BigNumber(((p1).update((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))], _1240_v0)).length));
          _1277_v29 = (_1277_v29).update(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber((function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_1262_v16).Keys.Elements) {
              let _1278_v30 = _compr_54;
              if ((_1262_v16).contains(_1278_v30)) {
                _coll54.add(_module.__default.safeModuloInt(_1278_v30, new BigNumber((_dafny.Set.fromElements(true, !(true), false, true, true)).length)));
              }
            }
            return _coll54;
          }()).length)), (new BigNumber((p1).length)).minus(new BigNumber((_1259_v13).length)));
          (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-11)), ((_1279_v6) => function (_1280_i2) {
            return _1279_v6;
          })(_1250_v6)), _1242_v2);
        } else {
          let _1281_v31;
          let _nw192 = new _module.C3();
          _nw192.__ctor((p2).multipliedBy(_1240_v0), !(!(p0)), new BigNumber(953), (_this).f26);
          _1281_v31 = _nw192;
          let _1282_v32;
          _1282_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1281_v31).f38,p0);
          let _1283_v33;
          _1283_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1281_v31).f38,(((_1282_v32).contains((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))])) ? ((_1282_v32).get((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))])) : ((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))])));
          let _1284_v34;
          _1284_v34 = _module.D13.create_DC36((_1283_v33).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))],!(!((_1281_v31).f38)))), p0, ((_this).f25).minus((_dafny.ZERO).minus((_1281_v31).f37)));
          _1284_v34 = _1284_v34;
          _1240_v0 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(_1240_v0, (_1281_v31).f37)).multipliedBy((_1281_v31).f37));
          (globalState).f12 = (((p0) ? (p2) : ((_dafny.ZERO).minus((_1281_v31).f37)))).isLessThan(((_dafny.ZERO).minus((_this).f25)).minus((_1281_v31).f37));
          let _1285_v35;
          let _nw193 = new _module.C2();
          _nw193.__ctor();
          _1285_v35 = _nw193;
        }
        let _1286_v36;
        let _init43 = ((_1287_v0) => function (_1288_i3) {
          return (_1288_i3).multipliedBy(_1287_v0);
        })(_1240_v0);
        let _nw194 = Array((new BigNumber(16)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw194.length); _i0_43++) {
          _nw194[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1286_v36 = _nw194;
        let _1289_v37;
        _1289_v37 = _module.D7.create_DC21(_1286_v36);
        _1286_v36 = (_1289_v37).dtor_cf32;
        let _1290_v38;
        _1290_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))],(_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]);
        let _1291_v39;
        _1291_v39 = _dafny.MultiSet.fromElements(true);
        let _1292_v40;
        _1292_v40 = _module.D15.create_DC39(_1291_v39);
        _1290_v38 = (_1290_v38).update(p0, !((_dafny.MultiSet.fromElements(!(true))).IsProperSubsetOf((_1292_v40).dtor_cf59)));
        let _1293_v41;
        let _nw195 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
        _1293_v41 = _nw195;
        let _1294_v42;
        _1294_v42 = _module.D9.create_DC26(p1);
        let _1295_v43;
        _1295_v43 = _dafny.Seq.of(_1294_v42, _1294_v42, _module.D9.create_DC26(p1), _1294_v42, _1294_v42);
        let _1296_v44;
        _1296_v44 = _dafny.Seq.of(_dafny.Seq.of(_1294_v42), _1295_v43);
        let _index222 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1293_v41).length));
        (_1293_v41)[_index222] = _1296_v44;
        let _index223 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_1293_v41).length));
        (_1293_v41)[_index223] = (((_1261_v15)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("fwphtpp"), _module.__default.safeIndex(new BigNumber((_1242_v2).length), new BigNumber((_dafny.Seq.UnicodeFromString("fwphtpp")).length)), _1250_v6)).length), new BigNumber((_1261_v15).length))]) ? (_dafny.Seq.of(_dafny.Seq.of(_1294_v42, _1294_v42))) : (_1296_v44));
        let _1297_v45;
        _1297_v45 = _module.D12.create_DC32(p0);
        let _1298_v46;
        _1298_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm33(p2, (_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))], p2, globalState),(_1297_v45).dtor_cf47);
        let _1299_v47;
        _1299_v47 = _module.D10.create_DC28(_1298_v46);
        let _index224 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length));
        let _rhs229 = _1299_v47;
        let _rhs230 = (_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))];
        let _rhs231 = p0;
        let _rhs232 = (_1290_v38).equals((_1290_v38).Merge(_1290_v38));
        let _rhs233 = _1248_v4;
        let _lhs169 = _1252_v8;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length));
        _1299_v47 = _rhs229;
        r1 = _rhs230;
        _lhs169[_lhs170] = _rhs231;
        r1 = _rhs232;
        _1248_v4 = _rhs233;
      } else {
        let _1300_v48;
        let _init44 = function (_1301_i4) {
          return (_1301_i4).plus(new BigNumber(502));
        };
        let _nw196 = Array((new BigNumber(14)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw196.length); _i0_44++) {
          _nw196[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1300_v48 = _nw196;
        let _index225 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length));
        (_1300_v48)[_index225] = p2;
        let _index226 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length));
        (_1300_v48)[_index226] = new BigNumber(339);
        let _1302_v49;
        let _nw197 = new _module.C2();
        _nw197.__ctor();
        _1302_v49 = _nw197;
        let _1303_v50;
        _1303_v50 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1259_v13).length),_1240_v0)).length));
        let _1304_v51;
        _1304_v51 = _dafny.MultiSet.fromElements(p0, p0);
        let _1305_v52;
        let _nw198 = new _module.C3();
        _nw198.__ctor((_this).f25, (_1303_v50).IsDisjointFrom(_1303_v50), (((_1304_v51).contains(true)) ? ((_1304_v51).get(true)) : (new BigNumber(228))), (_this).f26);
        _1305_v52 = _nw198;
        _1305_v52 = _1305_v52;
        _1300_v48 = _1300_v48;
        if ((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]) {
          let _index227 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length));
          (_1252_v8)[_index227] = !((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]);
          let _index228 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length));
          (_1300_v48)[_index228] = new BigNumber((_dafny.Seq.Concat(_1261_v15, _1261_v15)).length);
          _1240_v0 = (_module.D12.create_DC33(_1257_v11, _1252_v8, (_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))], (_1305_v52).f37)).dtor_cf51;
          let _1306_v53;
          let _nw199 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _1306_v53 = _nw199;
          let _index229 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1306_v53).length));
          (_1306_v53)[_index229] = _1261_v15;
          let _index230 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1306_v53).length));
          let _index231 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length));
          let _rhs234 = (_1300_v48)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length))];
          let _rhs235 = p2;
          let _rhs236 = _dafny.Seq.Concat(_1261_v15, _dafny.Seq.Concat(_module.__default.fm1((((_1260_v14).contains((_1305_v52).f37)) ? ((_1260_v14).get((_1305_v52).f37)) : ((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))])), globalState), _1261_v15));
          let _rhs237 = (_1300_v48)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length))];
          let _rhs238 = _1240_v0;
          let _lhs171 = _1306_v53;
          let _lhs172 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1306_v53).length));
          let _lhs173 = _1300_v48;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length));
          _1240_v0 = _rhs234;
          _1240_v0 = _rhs235;
          _lhs171[_lhs172] = _rhs236;
          _1240_v0 = _rhs237;
          _lhs173[_lhs174] = _rhs238;
          let _1307_v54;
          _1307_v54 = _dafny.Seq.of(_1252_v8, _1257_v11);
          let _1308_v55;
          _1308_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1307_v54, _1307_v54),!(!(p0)));
          _1308_v55 = (_1308_v55).update(_1307_v54, _module.__default.fm0((_1264_v21)[_module.__default.safeIndex((_1305_v52).f37, new BigNumber((_1264_v21).length))], globalState));
        } else {
          (globalState).f12 = (_1305_v52).f38;
          let _1309_v56;
          _1309_v56 = _module.D0.create_DC2(p0);
          let _pat_let_tv24 = _1261_v15;
          let _pat_let_tv25 = _1242_v2;
          let _pat_let_tv26 = _1240_v0;
          let _pat_let_tv27 = _1242_v2;
          let _pat_let_tv28 = _1250_v6;
          let _pat_let_tv29 = _1261_v15;
          let _1310_v57;
          let _nw200 = Array((new BigNumber(8)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _1309_v56;
          _nw200[(_dafny.ONE).toNumber()] = _1309_v56;
          _nw200[(new BigNumber(2)).toNumber()] = _1309_v56;
          _nw200[(new BigNumber(3)).toNumber()] = (((_1305_v52).f38) ? (_1309_v56) : (_module.__default.fm54(_dafny.MultiSet.FromArray(_1261_v15), globalState)));
          _nw200[(new BigNumber(4)).toNumber()] = _1309_v56;
          _nw200[(new BigNumber(5)).toNumber()] = _module.D0.create_DC2((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]);
          _nw200[(new BigNumber(6)).toNumber()] = _1309_v56;
          _nw200[(new BigNumber(7)).toNumber()] = function (_pat_let24_0) {
            return function (_1311_dt__update__tmp_h2) {
              return function (_pat_let25_0) {
                return function (_1312_dt__update_hcf3_h0) {
                  return _module.D0.create_DC2(_1312_dt__update_hcf3_h0);
                }(_pat_let25_0);
              }((_pat_let_tv24)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_pat_let_tv25, _module.__default.safeIndex(_pat_let_tv26, new BigNumber((_pat_let_tv27).length)), _pat_let_tv28)).length), new BigNumber((_pat_let_tv29).length))]);
            }(_pat_let24_0);
          }(_1309_v56);
          _1310_v57 = _nw200;
          let _index232 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1310_v57).length));
          (_1310_v57)[_index232] = _module.D0.create_DC2(p0);
          let _1313_v58;
          let _init45 = ((_1314_v50) => function (_1315_i5) {
            return _1314_v50;
          })(_1303_v50);
          let _nw201 = Array((new BigNumber(5)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw201.length); _i0_45++) {
            _nw201[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1313_v58 = _nw201;
          let _index233 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1310_v57).length));
          let _rhs239 = _1303_v50;
          let _rhs240 = ((((_1300_v48)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length))]).isLessThanOrEqualTo((_1305_v52).f37)) ? (_1309_v56) : (_1309_v56));
          let _rhs241 = _1313_v58;
          let _lhs175 = _1310_v57;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1310_v57).length));
          _1303_v50 = _rhs239;
          _lhs175[_lhs176] = _rhs240;
          _1313_v58 = _rhs241;
          let _1316_v59;
          _1316_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1305_v52).f38,(_1305_v52).f38);
          (globalState).f12 = !((((_1300_v48)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1300_v48).length))]).multipliedBy(new BigNumber((_1316_v59).length))).isLessThanOrEqualTo(_module.__default.fm31(_1242_v2, globalState)));
          (globalState).f12 = p0;
          _1259_v13 = _dafny.Seq.of(true, !((_1305_v52).f38));
        }
      }
      r0 = _module.D0.create_DC2(p0);
      r1 = (_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))];
      let _1317_v60;
      _1317_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, (_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]),((true) ? ((_1252_v8)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_1252_v8).length))]) : (p0)));
      r2 = _1317_v60;
      return [r0, r1, r2];
    }
    get f36() {
      let _this = this;
      return _this._f36;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f35 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f35) {
      let _this = this;
      (_this)._f35 = f35;
      return;
    }
    m13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let r2 = false;
      let r3 = [];
      r2 = (p1) === (!((_this).f35));
      let _1318_v0;
      _1318_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      _1318_v0 = (_1318_v0).update(p2, _module.__default.fm18(globalState));
      r0 = p0;
      if (p2) {
        (globalState).f12 = (new BigNumber(-283)).isLessThan(_module.__default.safeDivisionInt(p0, p0));
        r0 = p0;
        (globalState).f6 = _dafny.Seq.Concat(_module.__default.fm19(p1, globalState), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), ((_1319_p0) => function (_1320_i0) {
          return _1319_p0;
        })(p0)), _dafny.Seq.of(p0)));
        let _1321_v1;
        let _nw202 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1321_v1 = _nw202;
        let _index234 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_1321_v1).length));
        (_1321_v1)[_index234] = p0;
        let _index235 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_1321_v1).length));
        (_1321_v1)[_index235] = _module.__default.safeDivisionInt(new BigNumber(((_module.__default.fm20(_module.__default.fm0(p0, globalState), globalState)).Union(_dafny.Set.fromElements((_this).f35, p2, p2, p2))).length), _module.__default.safeDivisionInt(p0, p0));
        _1321_v1 = _1321_v1;
      } else {
        (globalState).f12 = (_this).f35;
        if (((p2) ? (p1) : (!((_this).f35)))) {
          let _1322_v2;
          _1322_v2 = _dafny.Seq.of(p0, p0);
          let _1323_v3;
          let _nw203 = Array((_dafny.ONE).toNumber());
          _nw203[(_dafny.ZERO).toNumber()] = _1322_v2;
          _1323_v3 = _nw203;
          let _index236 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1323_v3).length));
          (_1323_v3)[_index236] = _dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex(_module.__default.fm18(globalState), new BigNumber((_dafny.Seq.of(p0)).length)), p0);
          let _1324_v4;
          _1324_v4 = _dafny.Seq.UnicodeFromString("d");
          let _1325_v5;
          _1325_v5 = _dafny.Set.fromElements(new BigNumber(778), new BigNumber((_1324_v4).length));
          let _1326_v6;
          _1326_v6 = _module.D4.create_DC11(_1325_v5);
          let _index237 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_1323_v3).length));
          (_1323_v3)[_index237] = _dafny.Seq.of(p0, new BigNumber(((_1326_v6).dtor_cf18).length), _module.__default.fm18(globalState));
          r0 = p0;
          r0 = p0;
          r0 = p0;
          let _1327_v7;
          let _init46 = ((_1328_p1) => function (_1329_i1) {
            return (((_this).f35) ? (_module.D0.create_DC1(_1328_p1, !(_1328_p1))) : (_module.D0.create_DC1((_this).f35, _1328_p1)));
          })(p1);
          let _nw204 = Array((new BigNumber(7)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw204.length); _i0_46++) {
            _nw204[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1327_v7 = _nw204;
          let _1330_v8;
          _1330_v8 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1331_v9;
          _1331_v9 = _module.D0.create_DC1((_module.D1.create_DC6(_1330_v8, (_this).f35, _1324_v4)).dtor_cf8, p1);
          let _index238 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1327_v7).length));
          (_1327_v7)[_index238] = _1331_v9;
          let _index239 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1327_v7).length));
          (_1327_v7)[_index239] = _module.D0.create_DC1((p0).isLessThanOrEqualTo(p0), p2);
        } else {
          let _1332_v10;
          _1332_v10 = _module.D4.create_DC13(p0, (_this).f35, _module.__default.fm0(p0, globalState), p2);
          (globalState).f12 = _module.__default.fm0(_module.__default.safeDivisionInt(_module.__default.fm18(globalState), (_1332_v10).dtor_cf19), globalState);
          let _1333_v11;
          let _nw205 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1333_v11 = _nw205;
          r3 = _1333_v11;
          let _1334_v12;
          _1334_v12 = _module.D0.create_DC0((p2) && ((_this).f35));
          let _1335_v13;
          _1335_v13 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1336_v14;
          _1336_v14 = _dafny.Seq.of(p0, p0, _module.__default.fm18(globalState), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_1335_v13)).length)), p0);
          let _rhs242 = (_1336_v14)[_module.__default.safeIndex(new BigNumber(565), new BigNumber((_1336_v14).length))];
          let _rhs243 = _1334_v12;
          let _rhs244 = p0;
          let _rhs245 = (_module.__default.fm0(p0, globalState)) || (p1);
          r0 = _rhs242;
          _1334_v12 = _rhs243;
          r0 = _rhs244;
          r2 = _rhs245;
          (globalState).f12 = p1;
          let _1337_v15;
          _1337_v15 = _dafny.Seq.of(_1333_v11);
          let _1338_v16;
          let _nw206 = Array((new BigNumber(21)).toNumber());
          _nw206[(_dafny.ZERO).toNumber()] = _1333_v11;
          _nw206[(_dafny.ONE).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(2)).toNumber()] = (_1337_v15)[_module.__default.safeIndex(p0, new BigNumber((_1337_v15).length))];
          _nw206[(new BigNumber(3)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(4)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(5)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(6)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(7)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(8)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(9)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(10)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(11)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(12)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(13)).toNumber()] = ((_module.__default.fm0(p0, globalState)) ? (_1333_v11) : (_1333_v11));
          _nw206[(new BigNumber(14)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(15)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(16)).toNumber()] = (((_this).f35) ? (_1333_v11) : (_1333_v11));
          _nw206[(new BigNumber(17)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(18)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(19)).toNumber()] = _1333_v11;
          _nw206[(new BigNumber(20)).toNumber()] = _1333_v11;
          _1338_v16 = _nw206;
          let _index240 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_1338_v16).length));
          (_1338_v16)[_index240] = _1333_v11;
          let _index241 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_1338_v16).length));
          (_1338_v16)[_index241] = _1333_v11;
        }
        let _1339_v17;
        let _nw207 = Array((new BigNumber(14)).toNumber()).fill(_module.D3.Default());
        _1339_v17 = _nw207;
        let _1340_v18;
        let _nw208 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1340_v18 = _nw208;
        let _1341_v19;
        _1341_v19 = _module.D3.create_DC9(_dafny.Set.fromElements(_1340_v18));
        let _index242 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1339_v17).length));
        (_1339_v17)[_index242] = _1341_v19;
        let _1342_v20;
        _1342_v20 = new _dafny.CodePoint('n'.codePointAt(0));
        let _index243 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1339_v17).length));
        let _rhs246 = _1341_v19;
        let _rhs247 = (p0).plus(p0);
        let _rhs248 = _1342_v20;
        let _lhs177 = _1339_v17;
        let _lhs178 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1339_v17).length));
        let _lhs179 = globalState;
        _lhs177[_lhs178] = _rhs246;
        r0 = _rhs247;
        _lhs179.f23 = _rhs248;
        (globalState).f23 = _1342_v20;
        (globalState).f12 = (p0).isLessThan(new BigNumber(778));
      }
      let _1343_v21;
      _1343_v21 = _dafny.MultiSet.fromElements(p0, p0);
      let _hi7 = p0;
      for (let _1344_i2 = new BigNumber((((_1343_v21).update(p0, _module.__default.abs(p0))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(p2, p2, true)).length), p0, p0, p0))).cardinality()); _1344_i2.isLessThan(_hi7); _1344_i2 = _1344_i2.plus(_dafny.ONE)) {
        (globalState).f12 = !(_module.__default.fm0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), ((_1345_i2) => function (_1346_i3) {
          return _1345_i2;
        })(_1344_i2))).length), globalState));
        r2 = p1;
        let _1347_v22;
        _1347_v22 = _dafny.Seq.of(p2, p1);
        let _1348_v23;
        _1348_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1347_v22,p2);
        let _1349_v24;
        let _nw209 = Array((new BigNumber(7)).toNumber());
        _nw209[(_dafny.ZERO).toNumber()] = _1348_v23;
        _nw209[(_dafny.ONE).toNumber()] = _1348_v23;
        _nw209[(new BigNumber(2)).toNumber()] = _1348_v23;
        _nw209[(new BigNumber(3)).toNumber()] = (_1348_v23).Merge(_1348_v23);
        _nw209[(new BigNumber(4)).toNumber()] = _1348_v23;
        _nw209[(new BigNumber(5)).toNumber()] = (_1348_v23).Merge(_1348_v23);
        _nw209[(new BigNumber(6)).toNumber()] = _1348_v23;
        _1349_v24 = _nw209;
        let _index244 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_1349_v24).length));
        (_1349_v24)[_index244] = _1348_v23;
        let _index245 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_1349_v24).length));
        (_1349_v24)[_index245] = (_1348_v23).Merge(_1348_v23);
        let _1350_v25;
        _1350_v25 = _dafny.Set.fromElements(true);
        if (((_this).f35) === ((_dafny.Set.fromElements(false)).IsProperSubsetOf(_1350_v25))) {
          (globalState).f12 = p1;
          let _1351_v26;
          _1351_v26 = new _dafny.CodePoint('x'.codePointAt(0));
          (globalState).f23 = _1351_v26;
          let _1352_v27;
          _1352_v27 = _dafny.Seq.UnicodeFromString("r");
          let _1353_v28;
          _1353_v28 = _module.D5.create_DC17(p0, p2);
          let _1354_v29;
          _1354_v29 = _dafny.Seq.of(_1344_i2, (_1353_v28).dtor_cf26, _1344_i2);
          let _1355_v30;
          _1355_v30 = _dafny.Seq.of(new BigNumber((_1354_v29).length));
          let _1356_v31;
          _1356_v31 = _dafny.Seq.of(_1344_i2, new BigNumber((_1352_v27).length), new BigNumber((_1354_v29).length), new BigNumber((_1354_v29).length));
          let _1357_v32;
          let _init47 = ((_1358_v27) => function (_1359_i4) {
            return _1358_v27;
          })(_1352_v27);
          let _nw210 = Array((new BigNumber(16)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw210.length); _i0_47++) {
            _nw210[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1357_v32 = _nw210;
          let _1360_v33;
          let _nw211 = new _module.C3();
          _nw211.__ctor((_1355_v30)[_module.__default.safeIndex(_1344_i2, new BigNumber((_1355_v30).length))], (_1350_v25).IsDisjointFrom(_1350_v25), _module.__default.safeModuloInt(_1344_i2, _1344_i2), _1357_v32);
          _1360_v33 = _nw211;
          let _1361_v34;
          let _nw212 = new _module.C1();
          _nw212.__ctor(_1344_i2, _1357_v32);
          _1361_v34 = _nw212;
          let _1362_v35;
          let _nw213 = Array((_dafny.ONE).toNumber());
          _nw213[(_dafny.ZERO).toNumber()] = new BigNumber((_1347_v22).length);
          _1362_v35 = _nw213;
          let _1363_v36;
          _1363_v36 = _module.D7.create_DC21(_1362_v35);
          r3 = (_1363_v36).dtor_cf32;
        } else {
          let _1364_v37;
          let _nw214 = Array((new BigNumber(29)).toNumber()).fill(false);
          _1364_v37 = _nw214;
          let _1365_v38;
          _1365_v38 = _module.D12.create_DC33(_1364_v37, _1364_v37, p2, p0);
          let _1366_v39;
          _1366_v39 = _dafny.Map.Empty.slice().updateUnsafe((_1365_v38).dtor_cf51,_1364_v37);
          _1366_v39 = (_1366_v39).update(_1344_i2, _1364_v37);
          let _1367_v40;
          let _init48 = ((_1368_i2) => function (_1369_i5) {
            return (_1369_i5).multipliedBy(_1368_i2);
          })(_1344_i2);
          let _nw215 = Array((new BigNumber(28)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw215.length); _i0_48++) {
            _nw215[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1367_v40 = _nw215;
          let _index246 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1367_v40).length));
          (_1367_v40)[_index246] = p0;
          let _index247 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1367_v40).length));
          (_1367_v40)[_index247] = _1344_i2;
          let _1370_v41;
          _1370_v41 = _module.D3.create_DC10(((_this).f35) && (p1), _1364_v37, (_this).f35, (_1344_i2).multipliedBy(new BigNumber((_1318_v0).length)), _module.__default.fm18(globalState));
          let _1371_v42;
          let _nw216 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1371_v42 = _nw216;
          let _1372_v43;
          _1372_v43 = new _dafny.CodePoint('f'.codePointAt(0));
          let _index248 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1371_v42).length));
          (_1371_v42)[_index248] = _1372_v43;
          let _index249 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1371_v42).length));
          let _rhs249 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), function (_1373_i6) {
            return (((_this).f35) ? (new _dafny.CodePoint('f'.codePointAt(0))) : (new _dafny.CodePoint('q'.codePointAt(0))));
          });
          let _rhs250 = _1370_v41;
          let _rhs251 = _1372_v43;
          let _lhs180 = globalState;
          let _lhs181 = _1371_v42;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1371_v42).length));
          _lhs180.f10 = _rhs249;
          _1370_v41 = _rhs250;
          _lhs181[_lhs182] = _rhs251;
          (globalState).f10 = _dafny.Seq.UnicodeFromString("bdpspo");
          let _index250 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1367_v40).length));
          (_1367_v40)[_index250] = _1344_i2;
        }
      }
      let _1374_v44;
      _1374_v44 = _dafny.Seq.UnicodeFromString("kvtp");
      let _1375_v45;
      _1375_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _rhs252 = _1374_v44;
      let _rhs253 = ((p0).isEqualTo(new BigNumber((_1375_v45).length))) && (p1);
      let _rhs254 = _1374_v44;
      let _lhs183 = globalState;
      let _lhs184 = globalState;
      let _lhs185 = globalState;
      _lhs183.f10 = _rhs252;
      _lhs184.f12 = _rhs253;
      _lhs185.f10 = _rhs254;
      r0 = p0;
      let _1376_v46;
      _1376_v46 = _dafny.Set.fromElements(p0);
      r1 = ((_1376_v46).Difference(function () {
        let _coll55 = new _dafny.Set();
        for (const _compr_55 of _dafny.IntegerRange(new BigNumber(896), new BigNumber(496))) {
          let _1377_v47 = _compr_55;
          if (((new BigNumber(896)).isLessThanOrEqualTo(_1377_v47)) && ((_1377_v47).isLessThan(new BigNumber(496)))) {
            _coll55.add(_module.__default.safeModuloInt(_1377_v47, p0));
          }
        }
        return _coll55;
      }())).Intersect(_1376_v46);
      r2 = _module.__default.fm0(p0, globalState);
      let _1378_v48;
      let _nw217 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1378_v48 = _nw217;
      let _1379_v49;
      _1379_v49 = _dafny.Map.Empty.slice().updateUnsafe((_1343_v21).Difference(_1343_v21),_1378_v48);
      r3 = (((_1379_v49).contains(_1343_v21)) ? ((_1379_v49).get(_1343_v21)) : (_1378_v48));
      return [r0, r1, r2, r3];
    }
    m14(p0, p1, p2, globalState) {
      let _this = this;
      if ((new BigNumber(-963)).isLessThanOrEqualTo(p0)) {
        let _1380_v0;
        let _nw218 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1380_v0 = _nw218;
        let _1381_v1;
        let _nw219 = new _module.C1();
        _nw219.__ctor(new BigNumber((_dafny.Seq.of(false)).length), _1380_v0);
        _1381_v1 = _nw219;
        let _1382_v2;
        let _nw220 = Array((new BigNumber(24)).toNumber()).fill(false);
        _1382_v2 = _nw220;
        let _index251 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_1382_v2).length));
        (_1382_v2)[_index251] = (_this).f35;
        let _index252 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_1382_v2).length));
        (_1382_v2)[_index252] = p2;
        let _1383_v3;
        _1383_v3 = new BigNumber(-49);
        let _1384_v4;
        let _init49 = ((_1385_p0) => function (_1386_i0) {
          return (_1386_i0).minus(_1385_p0);
        })(p0);
        let _nw221 = Array((new BigNumber(3)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw221.length); _i0_49++) {
          _nw221[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1384_v4 = _nw221;
        let _1387_v5;
        _1387_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1384_v4,(_1382_v2)[_module.__default.safeIndex(new BigNumber(783), new BigNumber((_1382_v2).length))]);
        _1383_v3 = new BigNumber((_1387_v5).length);
        let _1388_v6;
        _1388_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,(_this).f35);
        let _1389_v7;
        _1389_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1383_v3,_1388_v6);
        _1389_v7 = (_1389_v7).update(new BigNumber(-393), (_1388_v6).Merge(_1388_v6));
        let _1390_v8;
        _1390_v8 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1391_v9;
        _1391_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,_1390_v8);
        let _1392_v10;
        let _nw222 = new _module.C4();
        _nw222.__ctor(_1380_v0, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("fhtp"), _module.__default.safeIndex(_1383_v3, new BigNumber((_dafny.Seq.UnicodeFromString("fhtp")).length)), (((_1391_v9).contains((_this).f35)) ? ((_1391_v9).get((_this).f35)) : (_1390_v8)))).length), _1380_v0);
        _1392_v10 = _nw222;
      } else {
        let _1393_v11;
        _1393_v11 = new BigNumber(108);
        let _1394_v12;
        _1394_v12 = _module.D0.create_DC0(p2);
        let _1395_v13;
        _1395_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1394_v12,true);
        let _1396_v14;
        _1396_v14 = _dafny.Seq.UnicodeFromString("wsxb");
        let _1397_v15;
        _1397_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f35,_1396_v14);
        let _1398_v16;
        _1398_v16 = new _dafny.CodePoint('t'.codePointAt(0));
        _1393_v11 = _module.__default.safeDivisionInt(_module.__default.fm21(_1395_v13, _module.__default.fm31(_1396_v14, globalState), p2, _dafny.Seq.update((((_1397_v15).contains((_this).f35)) ? ((_1397_v15).get((_this).f35)) : (_1396_v14)), _module.__default.safeIndex(_1393_v11, new BigNumber(((((_1397_v15).contains((_this).f35)) ? ((_1397_v15).get((_this).f35)) : (_1396_v14))).length)), _1398_v16), globalState), _1393_v11);
        let _1399_v17;
        let _nw223 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _1399_v17 = _nw223;
        let _1400_v18;
        _1400_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1399_v17,p2);
        let _1401_v19;
        _1401_v19 = _dafny.Seq.of(p2, (_this).f35);
        _1400_v18 = (_1400_v18).update(_1399_v17, _dafny.areEqual(_1401_v19, _dafny.Seq.of(p2)));
        let _1402_v20;
        _1402_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        _1402_v20 = _1402_v20;
        let _1403_v21;
        _1403_v21 = _module.D4.create_DC12();
        let _1404_v22;
        _1404_v22 = _module.D4.create_DC14(_1403_v21);
        let _1405_v23;
        _1405_v23 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(p2, !(p2)), _dafny.MultiSet.FromArray(_1401_v19));
        let _1406_v24;
        _1406_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1404_v22,_1405_v23);
        _1406_v24 = (_1406_v24).update(_1404_v22, (_1405_v23).Difference(_1405_v23));
        let _1407_v25;
        let _nw224 = Array((new BigNumber(24)).toNumber());
        _nw224[(_dafny.ZERO).toNumber()] = _1396_v14;
        _nw224[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("xyqt");
        _nw224[(new BigNumber(2)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(3)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(4)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(5)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(6)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), ((_1408_v16) => function (_1409_i1) {
          return _1408_v16;
        })(_1398_v16));
        _nw224[(new BigNumber(8)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(9)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("om");
        _nw224[(new BigNumber(11)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(12)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(13)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(14)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(15)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("cupth");
        _nw224[(new BigNumber(17)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(18)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(19)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(20)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(21)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(22)).toNumber()] = _1396_v14;
        _nw224[(new BigNumber(23)).toNumber()] = _1396_v14;
        _1407_v25 = _nw224;
        let _1410_v26;
        let _nw225 = new _module.C1();
        _nw225.__ctor((_dafny.ZERO).minus(_1393_v11), _1407_v25);
        _1410_v26 = _nw225;
      }
      let _1411_v27;
      _1411_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      _1411_v27 = (_1411_v27).update(p0, (p0).isLessThan(p0));
      let _1412_v28;
      let _init50 = ((_1413_v27) => function (_1414_i3) {
        return _1413_v27;
      })(_1411_v27);
      let _nw226 = Array((_dafny.ONE).toNumber());
      for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw226.length); _i0_50++) {
        _nw226[_i0_50] = _init50(new BigNumber(_i0_50));
      }
      _1412_v28 = _nw226;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1412_v28).length))) {
        let _1415_i2 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1415_i2)) && ((_1415_i2).isLessThan(new BigNumber((_1412_v28).length))))) {
          (_1412_v28)[(_1415_i2)] = (_1411_v27).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(-527))).length),(_this).f35));
        }
      }
      let _1416_v29;
      let _nw227 = Array((new BigNumber(13)).toNumber()).fill(false);
      _1416_v29 = _nw227;
      let _1417_v30;
      _1417_v30 = _dafny.Set.fromElements(true);
      let _index253 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1416_v29).length));
      (_1416_v29)[_index253] = (_1417_v30).equals(_1417_v30);
      let _index254 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1416_v29).length));
      (_1416_v29)[_index254] = (_this).f35;
      let _1418_v31;
      let _init51 = ((_1419_p1) => function (_1420_i5) {
        return _1419_p1;
      })(p1);
      let _nw228 = Array((new BigNumber(11)).toNumber());
      for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw228.length); _i0_51++) {
        _nw228[_i0_51] = _init51(new BigNumber(_i0_51));
      }
      _1418_v31 = _nw228;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1418_v31).length))) {
        let _1421_i4 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1421_i4)) && ((_1421_i4).isLessThan(new BigNumber((_1418_v31).length))))) {
          (_1418_v31)[(_1421_i4)] = p1;
        }
      }
      _1411_v27 = _1411_v27;
      return;
    }
    get f35() {
      let _this = this;
      return _this._f35;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f25, f26) {
      let _this = this;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    m4(p0, globalState) {
      let _this = this;
      let _1422_v0;
      _1422_v0 = _dafny.Seq.of((_this).f25, (_this).f25, (_this).f25, (_this).f25);
      let _1423_v1;
      let _nw229 = new _module.C4();
      _nw229.__ctor((_this).f26, new BigNumber((_1422_v0).length), (_this).f26);
      _1423_v1 = _nw229;
      let _1424_v2;
      _1424_v2 = false;
      if (_1424_v2) {
        let _1425_v3;
        let _nw230 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _1425_v3 = _nw230;
        let _1426_v4;
        _1426_v4 = _module.D7.create_DC21(_1425_v3);
        let _source20 = _1426_v4;
        if (_source20.is_DC22) {
          let _1427___mcc_h0 = (_source20).cf33;
          let _1428___mcc_h1 = (_source20).cf34;
          let _1429_cf34 = _1428___mcc_h1;
          let _1430_cf33 = _1427___mcc_h0;
          let _1431_v5;
          let _nw231 = new _module.C2();
          _nw231.__ctor();
          _1431_v5 = _nw231;
          let _1432_v6;
          _1432_v6 = _dafny.Seq.UnicodeFromString("vgkrb");
          let _1433_v7;
          _1433_v7 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1434_v8;
          _1434_v8 = _module.D1.create_DC6(_1433_v7, _1424_v2, _1432_v6);
          (globalState).f10 = _dafny.Seq.Concat(_1432_v6, (_1434_v8).dtor_cf9);
          let _1435_v9;
          _1435_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1432_v6,!(!(_1424_v2) || (_1424_v2)));
          _1435_v9 = (_1435_v9).update(_1432_v6, (new BigNumber(-573)).isLessThanOrEqualTo(new BigNumber(535)));
          (globalState).f12 = (((_this).f25).multipliedBy((_this).f25)).isLessThan((_this).f25);
        } else {
          let _1436___mcc_h2 = (_source20).cf32;
          let _1437_cf32 = _1436___mcc_h2;
          let _1438_v10;
          _1438_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v2,_1437_cf32);
          let _rhs255 = (((_1438_v10).contains(_1424_v2)) ? ((_1438_v10).get(_1424_v2)) : (((_1424_v2) ? (_1425_v3) : (_1425_v3))));
          let _rhs256 = _1424_v2;
          let _lhs186 = globalState;
          _1437_cf32 = _rhs255;
          _lhs186.f12 = _rhs256;
          let _index255 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1437_cf32).length));
          (_1437_cf32)[_index255] = (_this).f25;
          let _1439_v11;
          _1439_v11 = _dafny.Seq.UnicodeFromString("opuevflja");
          let _1440_v12;
          _1440_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v2,(_this).f25);
          let _index256 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1437_cf32).length));
          (_1437_cf32)[_index256] = ((_1424_v2) ? (new BigNumber((_1439_v11).length)) : (_module.__default.safeDivisionInt((_this).f25, new BigNumber((_1440_v12).length))));
          let _1441_v13;
          _1441_v13 = new _dafny.CodePoint('s'.codePointAt(0));
          let _index257 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1437_cf32).length));
          (_1437_cf32)[_index257] = _module.__default.fm31(_dafny.Seq.Concat(_1439_v11, _dafny.Seq.update(_1439_v11, _module.__default.safeIndex((_this).f25, new BigNumber((_1439_v11).length)), _1441_v13)), globalState);
          let _1442_v14;
          _1442_v14 = _dafny.MultiSet.fromElements(true, _1424_v2);
          let _1443_v15;
          _1443_v15 = _dafny.Seq.of(_1440_v12);
          let _1444_v16;
          _1444_v16 = _module.D9.create_DC26(_1440_v12);
          let _1445_v17;
          let _nw232 = Array((new BigNumber(19)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = _1440_v12;
          _nw232[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1424_v2,(_dafny.ZERO).minus(new BigNumber(-439)))).update(!(false), (_dafny.ZERO).minus((((_1442_v14).contains(_1424_v2)) ? ((_1442_v14).get(_1424_v2)) : ((_this).f25))));
          _nw232[(new BigNumber(2)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(3)).toNumber()] = (_1440_v12).update(_1424_v2, (_1437_cf32)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1437_cf32).length))]);
          _nw232[(new BigNumber(4)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(5)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(6)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(7)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(8)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(9)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(10)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(11)).toNumber()] = (_1443_v15)[_module.__default.safeIndex(new BigNumber(603), new BigNumber((_1443_v15).length))];
          _nw232[(new BigNumber(12)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(13)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(14)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(15)).toNumber()] = _1440_v12;
          _nw232[(new BigNumber(16)).toNumber()] = (_1444_v16).dtor_cf38;
          _nw232[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f25, globalState),(_this).f25);
          _nw232[(new BigNumber(18)).toNumber()] = _1440_v12;
          _1445_v17 = _nw232;
          let _1446_v18;
          _1446_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1445_v17);
          _1446_v18 = (_1446_v18).update((_dafny.ZERO).minus(((_1437_cf32)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1437_cf32).length))]).plus((_1437_cf32)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1437_cf32).length))])), _1445_v17);
        }
        let _1447_v19;
        _1447_v19 = _dafny.Seq.UnicodeFromString("vhdow");
        let _1448_v20;
        _1448_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v19,true);
        _1448_v20 = function () {
          let _coll56 = new _dafny.Map();
          for (const _compr_56 of (_1448_v20).Keys.Elements) {
            let _1449_v21 = _compr_56;
            if ((_1448_v20).contains(_1449_v21)) {
              _coll56.push([_1449_v21,_1424_v2]);
            }
          }
          return _coll56;
        }();
        let _source21 = _module.D9.create_DC27((_this).f25, (_this).f25);
        if (_source21.is_DC27) {
          let _1450___mcc_h3 = (_source21).cf39;
          let _1451___mcc_h4 = (_source21).cf40;
          let _1452_cf40 = _1451___mcc_h4;
          let _1453_cf39 = _1450___mcc_h3;
          let _1454_v22;
          _1454_v22 = _dafny.MultiSet.fromElements(_1452_cf40);
          let _1455_v23;
          _1455_v23 = _dafny.MultiSet.fromElements(_1454_v22, _dafny.MultiSet.FromArray(_1422_v0), (_dafny.MultiSet.fromElements(new BigNumber(155), (_this).f25, new BigNumber(974))).Union(_1454_v22));
          _1455_v23 = (_1455_v23).Difference((_1455_v23).Difference(_1455_v23));
          _1424_v2 = !(true);
          let _1456_v24;
          _1456_v24 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1457_v25;
          _1457_v25 = _module.D1.create_DC6(_1456_v24, !(_1424_v2), _1447_v19);
          let _1458_v26;
          _1458_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v2,(_dafny.MultiSet.fromElements(new BigNumber(((_1457_v25).dtor_cf9).length))).Intersect(_1454_v22));
          let _rhs257 = _1456_v24;
          let _rhs258 = _1424_v2;
          let _rhs259 = _1458_v26;
          let _lhs187 = globalState;
          let _lhs188 = globalState;
          _lhs187.f23 = _rhs257;
          _lhs188.f12 = _rhs258;
          _1458_v26 = _rhs259;
          (globalState).f12 = _1424_v2;
        } else {
          let _1459___mcc_h5 = (_source21).cf38;
          let _1460_cf38 = _1459___mcc_h5;
          let _1461_v27;
          let _nw233 = new _module.C1();
          _nw233.__ctor(new BigNumber(-131), (_1423_v1).f36);
          _1461_v27 = _nw233;
          let _1462_v28;
          let _nw234 = Array((new BigNumber(4)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = _1461_v27;
          _nw234[(_dafny.ONE).toNumber()] = _1461_v27;
          _nw234[(new BigNumber(2)).toNumber()] = _1461_v27;
          _nw234[(new BigNumber(3)).toNumber()] = _1461_v27;
          _1462_v28 = _nw234;
          let _index258 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1462_v28).length));
          (_1462_v28)[_index258] = _1461_v27;
          let _index259 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1462_v28).length));
          (_1462_v28)[_index259] = ((_1424_v2) ? (_1461_v27) : (_1461_v27));
          let _1463_v29;
          let _nw235 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _1463_v29 = _nw235;
          let _1464_v30;
          let _nw236 = Array((new BigNumber(11)).toNumber());
          _nw236[(_dafny.ZERO).toNumber()] = _1424_v2;
          _nw236[(_dafny.ONE).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(2)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(3)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(4)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(5)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(6)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(7)).toNumber()] = true;
          _nw236[(new BigNumber(8)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(9)).toNumber()] = _1424_v2;
          _nw236[(new BigNumber(10)).toNumber()] = _1424_v2;
          _1464_v30 = _nw236;
          let _1465_v31;
          _1465_v31 = _dafny.Seq.of(_1464_v30, _1464_v30);
          let _index260 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1463_v29).length));
          (_1463_v29)[_index260] = _1465_v31;
          let _index261 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1463_v29).length));
          (_1463_v29)[_index261] = _1465_v31;
          let _index262 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1464_v30).length));
          (_1464_v30)[_index262] = _1424_v2;
          let _index263 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1464_v30).length));
          (_1464_v30)[_index263] = _1424_v2;
          let _1466_v32;
          _1466_v32 = new _dafny.CodePoint('k'.codePointAt(0));
          (globalState).f10 = ((_dafny.Seq.contains(_1422_v0, (_this).f25)) ? (_module.__default.fm2((_this).f25, (_this).f25, globalState)) : (_dafny.Seq.update(_1447_v19, _module.__default.safeIndex((_this).f25, new BigNumber((_1447_v19).length)), _1466_v32)));
        }
        let _1467_v33;
        let _nw237 = new _module.C0();
        _nw237.__ctor();
        _1467_v33 = _nw237;
        let _1468_v34;
        _1468_v34 = new BigNumber(731);
        _1468_v34 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1447_v19).length)), (_this).f25);
      } else {
        let _1469_v35;
        _1469_v35 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1470_v36;
        _1470_v36 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),_1469_v35);
        let _1471_v37;
        _1471_v37 = _dafny.Set.fromElements(_1470_v36);
        let _1472_v38;
        _1472_v38 = _dafny.MultiSet.fromElements(false);
        let _1473_v39;
        _1473_v39 = _dafny.Seq.UnicodeFromString("wrtlea");
        let _rhs260 = !(_dafny.MultiSet.fromElements(_1424_v2, _1424_v2, _1424_v2)).equals((_1472_v38).Difference(_1472_v38));
        let _rhs261 = (_1473_v39)[_module.__default.safeIndex(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f25), (_this).f25), new BigNumber((_1473_v39).length))];
        let _rhs262 = _1471_v37;
        let _lhs189 = globalState;
        let _lhs190 = globalState;
        _lhs189.f12 = _rhs260;
        _lhs190.f23 = _rhs261;
        _1471_v37 = _rhs262;
        let _1474_v40;
        let _init52 = ((_1475_v38) => function (_1476_i0) {
          return _module.__default.safeDivisionInt(_1476_i0, new BigNumber((_1475_v38).cardinality()));
        })(_1472_v38);
        let _nw238 = Array((new BigNumber(5)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw238.length); _i0_52++) {
          _nw238[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1474_v40 = _nw238;
        let _1477_v41;
        _1477_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_1424_v2),_1424_v2);
        let _index264 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length));
        (_1474_v40)[_index264] = _module.__default.fm21(_1477_v41, (_this).f25, _1424_v2, _dafny.Seq.update(_1473_v39, _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1473_v39).length)), _1469_v35), globalState);
        let _index265 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length));
        (_1474_v40)[_index265] = (_this).f25;
        let _1478_v42;
        _1478_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v2,!(_1424_v2));
        let _1479_v43;
        _1479_v43 = _module.D13.create_DC36(_1478_v42, _1424_v2, (_this).f25);
        let _1480_v44;
        _1480_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1479_v43,new BigNumber(827));
        let _1481_v45;
        _1481_v45 = _module.D1.create_DC6(_1469_v35, _1424_v2, _1473_v39);
        let _1482_v46;
        _1482_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v2,new BigNumber((_dafny.Seq.of(new BigNumber(((_1481_v45).dtor_cf9).length), (_this).f25)).length));
        _1424_v2 = ((((_1472_v38).contains(_1424_v2)) ? ((_1472_v38).get(_1424_v2)) : ((_1474_v40)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length))]))).isLessThanOrEqualTo(((_1424_v2) ? ((_dafny.ZERO).minus(_module.__default.fm18(globalState))) : ((((_1480_v44).contains(_1479_v43)) ? ((_1480_v44).get(_1479_v43)) : (new BigNumber((_1482_v46).length))))));
        let _index266 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length));
        (_1474_v40)[_index266] = new BigNumber(272);
        let _1483_v47;
        _1483_v47 = _dafny.MultiSet.fromElements(new BigNumber((_1473_v39).length));
        if (((_1483_v47).Intersect(_dafny.MultiSet.fromElements((_this).f25))).IsProperSubsetOf(((_1424_v2) ? (_1483_v47) : (_dafny.MultiSet.fromElements(new BigNumber(-675)))))) {
          let _1484_v48;
          let _nw239 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1484_v48 = _nw239;
          let _1485_v49;
          _1485_v49 = _dafny.Set.fromElements(_1484_v48, _1484_v48);
          let _1486_v50;
          _1486_v50 = _dafny.MultiSet.fromElements(_1485_v49);
          let _1487_v51;
          _1487_v51 = _dafny.MultiSet.fromElements(_1469_v35);
          let _1488_v52;
          _1488_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1487_v51,_1472_v38);
          let _index267 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length));
          (_1474_v40)[_index267] = (((_1486_v50).contains(_1485_v49)) ? ((_1486_v50).get(_1485_v49)) : (new BigNumber(((((_1488_v52).contains(_module.__default.fm55(_1424_v2, new BigNumber((_1487_v51).cardinality()), new BigNumber((_1473_v39).length), globalState))) ? ((_1488_v52).get(_module.__default.fm55(_1424_v2, new BigNumber((_1487_v51).cardinality()), new BigNumber((_1473_v39).length), globalState))) : (_1472_v38))).cardinality())));
          (globalState).f23 = _1469_v35;
          (globalState).f12 = _module.__default.fm0(_module.__default.safeModuloInt(new BigNumber((_1422_v0).length), new BigNumber(571)), globalState);
          let _1489_v53;
          _1489_v53 = _dafny.Seq.of(_1424_v2, _1424_v2, _1424_v2, _1424_v2, _1424_v2);
          let _1490_v54;
          _1490_v54 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1474_v40)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length))]),(_1474_v40)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length))]);
          let _1491_v55;
          _1491_v55 = _dafny.Map.Empty.slice().updateUnsafe((_1474_v40)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length))],_dafny.Seq.of(_1424_v2, false));
          let _1492_v56;
          _1492_v56 = _module.D10.create_DC29(_1424_v2, _1424_v2, (_this).f25);
          let _pat_let_tv30 = _1489_v53;
          let _1493_v57;
          let _nw240 = Array((new BigNumber(24)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = _1489_v53;
          _nw240[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_1489_v53, _module.__default.safeIndex(new BigNumber((_1490_v54).length), new BigNumber((_1489_v53).length)), _1424_v2);
          _nw240[(new BigNumber(2)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(3)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(4)).toNumber()] = (function (_pat_let26_0) {
            return function (_1494_dt__update__tmp_h0) {
              return function (_pat_let27_0) {
                return function (_1495_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4(_1495_dt__update_hcf6_h0);
                }(_pat_let27_0);
              }(_pat_let_tv30);
            }(_pat_let26_0);
          }(p0)).dtor_cf6;
          _nw240[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm1(_1424_v2, globalState), _dafny.Seq.of(true, _1424_v2, _1424_v2));
          _nw240[(new BigNumber(6)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(7)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(8)).toNumber()] = (_module.D1.create_DC4(_dafny.Seq.of(_1424_v2))).dtor_cf6;
          _nw240[(new BigNumber(9)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(10)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(11)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(12)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1489_v53, _1489_v53);
          _nw240[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _1489_v53);
          _nw240[(new BigNumber(15)).toNumber()] = (((_1491_v55).contains((_1492_v56).dtor_cf44)) ? ((_1491_v55).get((_1492_v56).dtor_cf44)) : (_1489_v53));
          _nw240[(new BigNumber(16)).toNumber()] = _module.__default.fm1(_1424_v2, globalState);
          _nw240[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_1424_v2, _1424_v2, (_1489_v53)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_1489_v53).length))]);
          _nw240[(new BigNumber(18)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(19)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(20)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_1489_v53, _1489_v53);
          _nw240[(new BigNumber(22)).toNumber()] = _1489_v53;
          _nw240[(new BigNumber(23)).toNumber()] = _1489_v53;
          _1493_v57 = _nw240;
          let _index268 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1493_v57).length));
          (_1493_v57)[_index268] = ((_module.__default.fm0((_this).f25, globalState)) ? (_1489_v53) : (_1489_v53));
          let _1496_v58;
          _1496_v58 = _module.D0.create_DC1(true, _1424_v2);
          let _index269 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1493_v57).length));
          let _rhs263 = _1424_v2;
          let _rhs264 = _module.__default.fm0((_this).f25, globalState);
          let _rhs265 = _dafny.Seq.Concat(_1489_v53, _1489_v53);
          let _rhs266 = _module.__default.fm56(_1424_v2, globalState);
          let _lhs191 = globalState;
          let _lhs192 = _1493_v57;
          let _lhs193 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1493_v57).length));
          _lhs191.f12 = _rhs263;
          _1424_v2 = _rhs264;
          _lhs192[_lhs193] = _rhs265;
          _1496_v58 = _rhs266;
          let _index270 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_1484_v48).length));
          (_1484_v48)[_index270] = _1424_v2;
          let _index271 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_1484_v48).length));
          (_1484_v48)[_index271] = _1424_v2;
        } else {
          (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mrtsnphg"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rcfgfvj"), _1473_v39));
          (globalState).f10 = _dafny.Seq.Concat(_1473_v39, _dafny.Seq.Create(_module.__default.abs(new BigNumber(438)), ((_1497_v35) => function (_1498_i1) {
            return _1497_v35;
          })(_1469_v35)));
          let _index272 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length));
          (_1474_v40)[_index272] = (_this).f25;
          _1474_v40 = _1474_v40;
          let _1499_v59;
          _1499_v59 = _dafny.Seq.of(_1473_v39);
          let _index273 = _module.__default.safeIndex(new BigNumber(680), new BigNumber(((_1423_v1).f36).length));
          ((_1423_v1).f36)[_index273] = _dafny.Seq.Concat((_1499_v59)[_module.__default.safeIndex((_1474_v40)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_1474_v40).length))], new BigNumber((_1499_v59).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(-695)), ((_1500_v35) => function (_1501_i2) {
            return _1500_v35;
          })(_1469_v35)));
          let _index274 = _module.__default.safeIndex(new BigNumber(680), new BigNumber(((_1423_v1).f36).length));
          ((_1423_v1).f36)[_index274] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("mipgqij"), _module.__default.safeIndex(((_this).f25).plus(new BigNumber((_1473_v39).length)), new BigNumber((_dafny.Seq.UnicodeFromString("mipgqij")).length)), _1469_v35);
        }
      }
      _1424_v2 = _1424_v2;
      let _1502_v60;
      _1502_v60 = _dafny.Seq.of(_1424_v2, _1424_v2);
      let _1503_v61;
      _1503_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,!((_1502_v60)[_module.__default.safeIndex((_this).f25, new BigNumber((_1502_v60).length))]));
      _1503_v61 = _1503_v61;
      let _1504_v62;
      let _init53 = ((_1505_v2) => function (_1506_i3) {
        return ((_1505_v2) ? (true) : (_1505_v2));
      })(_1424_v2);
      let _nw241 = Array((new BigNumber(2)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw241.length); _i0_53++) {
        _nw241[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _1504_v62 = _nw241;
      let _index275 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1504_v62).length));
      (_1504_v62)[_index275] = !(_1424_v2);
      let _index276 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1504_v62).length));
      (_1504_v62)[_index276] = _1424_v2;
      let _1507_v63;
      _1507_v63 = new _dafny.CodePoint('y'.codePointAt(0));
      if (!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("fiqda"), _1507_v63)) {
        let _1508_v64;
        _1508_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_module.__default.fm18(globalState));
        _1508_v64 = (_1508_v64).update((_this).f25, (_this).f25);
        let _1509_v65;
        _1509_v65 = _module.D9.create_DC27(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), ((_1510_v63) => function (_1511_i4) {
  return _1510_v63;
})(_1507_v63))).length), (_this).f25);
        let _1512_v66;
        _1512_v66 = _dafny.MultiSet.fromElements(_1509_v65, _1509_v65, _1509_v65);
        let _1513_v67;
        _1513_v67 = _dafny.Set.fromElements(_1512_v66);
        _1513_v67 = (_1513_v67).Intersect(_1513_v67);
        let _1514_v68;
        let _nw242 = new _module.C2();
        _nw242.__ctor();
        _1514_v68 = _nw242;
        let _1515_v69;
        let _nw243 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1515_v69 = _nw243;
        let _1516_v70;
        _1516_v70 = _dafny.Seq.of(_1515_v69);
        _1515_v69 = (_1516_v70)[_module.__default.safeIndex(((_this).f25).minus((_this).f25), new BigNumber((_1516_v70).length))];
        let _1517_v71;
        _1517_v71 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25),!(_1424_v2));
        let _1518_v72;
        _1518_v72 = _dafny.Seq.UnicodeFromString("a");
        _1517_v71 = (_1517_v71).update((_1508_v64).Merge(_module.__default.fm57((_1504_v62)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_1504_v62).length))], _module.__default.fm0((_this).f25, globalState), (_this).f25, globalState)), _dafny.Seq.contains(_1518_v72, _1507_v63));
      } else {
        let _1519_v73;
        let _nw244 = new _module.C2();
        _nw244.__ctor();
        _1519_v73 = _nw244;
        let _index277 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1504_v62).length));
        let _rhs267 = !((_1504_v62)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_1504_v62).length))]);
        let _rhs268 = _1504_v62;
        let _lhs194 = _1504_v62;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1504_v62).length));
        _lhs194[_lhs195] = _rhs267;
        _1504_v62 = _rhs268;
        let _1520_v74;
        _1520_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1424_v2,_1424_v2);
        let _1521_v75;
        _1521_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1520_v74);
        _1521_v75 = (_1521_v75).update((_1519_v73).fm24((_this).f25, globalState), _1520_v74);
        let _1522_v76;
        _1522_v76 = new BigNumber(413);
        _1522_v76 = (_this).f25;
        let _1523_v77;
        let _init54 = ((_1524_v76) => function (_1525_i5) {
          return _module.__default.safeModuloInt(_1525_i5, _1524_v76);
        })(_1522_v76);
        let _nw245 = Array((new BigNumber(8)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw245.length); _i0_54++) {
          _nw245[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1523_v77 = _nw245;
        let _1526_v78;
        _1526_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1523_v77,new BigNumber(452));
        let _index278 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1523_v77).length));
        (_1523_v77)[_index278] = ((_this).f25).multipliedBy((((_1526_v78).contains(_1523_v77)) ? ((_1526_v78).get(_1523_v77)) : (new BigNumber(296))));
        let _1527_v79;
        _1527_v79 = _dafny.Set.fromElements(false, _1424_v2, _1424_v2, _1424_v2, _1424_v2);
        let _1528_v80;
        _1528_v80 = _dafny.Seq.of(_1527_v79, _1527_v79);
        let _index279 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1523_v77).length));
        (_1523_v77)[_index279] = (new BigNumber(((_1527_v79).Intersect((_1528_v80)[_module.__default.safeIndex(_1522_v76, new BigNumber((_1528_v80).length))])).length)).minus(_1522_v76);
      }
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      if (p1) {
        r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25));
        let _1529_v0;
        _1529_v0 = _dafny.Seq.UnicodeFromString("olyddmha");
        let _1530_v1;
        let _nw246 = new _module.C1();
        _nw246.__ctor(_module.__default.fm18(globalState), (_this).f26);
        _1530_v1 = _nw246;
        let _1531_v2;
        _1531_v2 = _dafny.MultiSet.fromElements(_1530_v1, _1530_v1);
        let _1532_v3;
        _1532_v3 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1529_v0).length)).plus((_this).f25),_1531_v2);
        _1532_v3 = (_1532_v3).update((_this).f25, _1531_v2);
        r0 = (_this).f25;
        (globalState).f12 = _dafny.Seq.IsProperPrefixOf(_1529_v0, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(444)), function (_1533_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        }), _1529_v0));
        let _index280 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((p3).length));
        (p3)[_index280] = (_this).f25;
        let _index281 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((p3).length));
        (p3)[_index281] = (_this).f25;
      } else {
        let _1534_v4;
        _1534_v4 = _dafny.MultiSet.fromElements(false, p1);
        let _1535_v5;
        _1535_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _1536_v6;
        _1536_v6 = _dafny.Set.fromElements((_1534_v4).IsSubsetOf(_1534_v4), (((_1535_v5).contains(p1)) ? ((_1535_v5).get(p1)) : (p2)));
        _1536_v6 = ((_1536_v6).Intersect(_1536_v6)).Intersect(_1536_v6);
        let _1537_v7;
        _1537_v7 = _dafny.Seq.of(p1);
        (globalState).f12 = (_dafny.Set.fromElements(p2, (_1537_v7)[_module.__default.safeIndex((_this).f25, new BigNumber((_1537_v7).length))], p1, p2)).IsSubsetOf((_dafny.Set.fromElements(p1, p2)).Difference(_1536_v6));
        let _1538_v8;
        let _nw247 = Array((new BigNumber(16)).toNumber()).fill(_module.D10.Default());
        _1538_v8 = _nw247;
        _1538_v8 = _1538_v8;
        let _1539_v9;
        _1539_v9 = _module.D0.create_DC1(p2, p1);
        let _1540_v10;
        let _nw248 = Array((new BigNumber(21)).toNumber());
        _nw248[(_dafny.ZERO).toNumber()] = p1;
        _nw248[(_dafny.ONE).toNumber()] = p2;
        _nw248[(new BigNumber(2)).toNumber()] = p1;
        _nw248[(new BigNumber(3)).toNumber()] = p1;
        _nw248[(new BigNumber(4)).toNumber()] = p2;
        _nw248[(new BigNumber(5)).toNumber()] = p1;
        _nw248[(new BigNumber(6)).toNumber()] = p1;
        _nw248[(new BigNumber(7)).toNumber()] = false;
        _nw248[(new BigNumber(8)).toNumber()] = p2;
        _nw248[(new BigNumber(9)).toNumber()] = p1;
        _nw248[(new BigNumber(10)).toNumber()] = false;
        _nw248[(new BigNumber(11)).toNumber()] = p1;
        _nw248[(new BigNumber(12)).toNumber()] = (_1539_v9).dtor_cf1;
        _nw248[(new BigNumber(13)).toNumber()] = p2;
        _nw248[(new BigNumber(14)).toNumber()] = p2;
        _nw248[(new BigNumber(15)).toNumber()] = p2;
        _nw248[(new BigNumber(16)).toNumber()] = p1;
        _nw248[(new BigNumber(17)).toNumber()] = !(p1);
        _nw248[(new BigNumber(18)).toNumber()] = p2;
        _nw248[(new BigNumber(19)).toNumber()] = p1;
        _nw248[(new BigNumber(20)).toNumber()] = p1;
        _1540_v10 = _nw248;
        let _1541_v11;
        _1541_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1540_v10);
        let _1542_v12;
        _1542_v12 = _module.D4.create_DC13((_this).f25, !(p1), p1, true);
        _1541_v11 = (_1541_v11).update(!(p2) || ((_1542_v12).dtor_cf21), _1540_v10);
        let _1543_v15;
        let _init55 = function (_1544_i1) {
          return (function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_this).f25)).cardinality()))).Elements) {
              let _1545_v13 = _compr_57;
              if ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_this).f25)).cardinality()))).contains(_1545_v13)) {
                _coll57.add(_module.__default.safeModuloInt(_1545_v13, new BigNumber((function () {
                  let _coll58 = new _dafny.Map();
                  for (const _compr_58 of _dafny.IntegerRange(new BigNumber(384), new BigNumber(189))) {
                    let _1546_v14 = _compr_58;
                    if (((new BigNumber(384)).isLessThanOrEqualTo(_1546_v14)) && ((_1546_v14).isLessThan(new BigNumber(189)))) {
                      _coll58.push([_module.__default.safeModuloInt(_1546_v14, new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).length)),false]);
                    }
                  }
                  return _coll58;
                }()).length)));
              }
            }
            return _coll57;
          }()).Difference(_dafny.Set.fromElements((_this).f25));
        };
        let _nw249 = Array((new BigNumber(19)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw249.length); _i0_55++) {
          _nw249[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _1543_v15 = _nw249;
        let _1547_v16;
        _1547_v16 = _dafny.Set.fromElements((_this).f25, (_this).f25, (_this).f25);
        let _index282 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1543_v15).length));
        (_1543_v15)[_index282] = _1547_v16;
        let _1548_v17;
        _1548_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1547_v16);
        let _1549_v18;
        _1549_v18 = _dafny.Seq.UnicodeFromString("wbpat");
        let _index283 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1543_v15).length));
        (_1543_v15)[_index283] = (((_1548_v17).contains(_module.__default.fm31(_1549_v18, globalState))) ? ((_1548_v17).get(_module.__default.fm31(_1549_v18, globalState))) : (_1547_v16));
      }
      if (p1) {
        let _1550_v19;
        _1550_v19 = _dafny.Seq.of(p2, p2, p1, p2, true);
        let _1551_v20;
        _1551_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
        let _1552_v21;
        _1552_v21 = _dafny.Seq.of(_1550_v19, _1550_v19, ((p2) ? (_1550_v19) : (_dafny.Seq.of((((_1551_v20).contains(true)) ? ((_1551_v20).get(true)) : (p2)), p2, p1, p2, (((_1551_v20).contains(p1)) ? ((_1551_v20).get(p1)) : (p1))))));
        _1552_v21 = _dafny.Seq.update(_1552_v21, _module.__default.safeIndex((_this).f25, new BigNumber((_1552_v21).length)), _1550_v19);
        let _1553_v22;
        let _nw250 = new _module.C4();
        _nw250.__ctor((_this).f26, (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f25, (_this).f25))), (_this).f26);
        _1553_v22 = _nw250;
        _1553_v22 = _1553_v22;
        let _1554_v23;
        let _nw251 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1554_v23 = _nw251;
        let _1555_v24;
        _1555_v24 = _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_1554_v23,p2));
        let _1556_v25;
        _1556_v25 = _module.D8.create_DC25(_1555_v24);
        let _1557_v26;
        _1557_v26 = _module.D8.create_DC25(_module.D8.create_DC25(_module.D8.create_DC25(_1556_v25)));
        let _1558_v27;
        _1558_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("esapojtap"),_1557_v26);
        _1558_v27 = (_1558_v27).update(_dafny.Seq.UnicodeFromString("cxdasflhy"), _1557_v26);
        let _1559_v28;
        _1559_v28 = _dafny.Seq.UnicodeFromString("f");
        (globalState).f10 = _1559_v28;
        r0 = _module.__default.fm21(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(p2),true), (_this).f25, _dafny.Seq.contains(_1550_v19, p2), _1559_v28, globalState);
      } else {
        (globalState).f12 = p2;
        (globalState).f12 = false;
        if (p2) {
          let _1560_v29;
          _1560_v29 = _dafny.Map.Empty.slice().updateUnsafe(((p2) ? ((_this).f25) : ((_this).f25)),(_this).f25);
          let _1561_v30;
          let _init56 = ((_1562_p1) => function (_1563_i2) {
            return _1562_p1;
          })(p1);
          let _nw252 = Array((new BigNumber(2)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw252.length); _i0_56++) {
            _nw252[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _1561_v30 = _nw252;
          let _1564_v31;
          _1564_v31 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1561_v30);
          let _1565_v32;
          _1565_v32 = _dafny.Set.fromElements((((_1564_v31).contains(false)) ? ((_1564_v31).get(false)) : (_1561_v30)));
          let _1566_v33;
          _1566_v33 = _dafny.Seq.of(_1565_v32);
          let _1567_v34;
          _1567_v34 = _module.D3.create_DC9((_1566_v33)[_module.__default.safeIndex((_this).f25, new BigNumber((_1566_v33).length))]);
          let _1568_v35;
          _1568_v35 = _dafny.Seq.of(_module.D3.create_DC9(_1565_v32), _1567_v34, _1567_v34);
          _1560_v29 = (_1560_v29).update(_module.__default.safeModuloInt((_this).f25, (_this).f25), ((_this).f25).multipliedBy(new BigNumber((_1568_v35).length)));
          let _1569_v36;
          _1569_v36 = _dafny.Seq.UnicodeFromString("tfim");
          let _index284 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((p3).length));
          (p3)[_index284] = new BigNumber((_1569_v36).length);
          let _index285 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((p3).length));
          (p3)[_index285] = new BigNumber(664);
          let _1570_v37;
          _1570_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1569_v36,(_this).f25);
          let _1571_v38;
          _1571_v38 = _module.D5.create_DC16(p1);
          let _1572_v39;
          _1572_v39 = _module.D6.create_DC19(_module.D0.create_DC3(new BigNumber((_1570_v37).length), (_1571_v38).dtor_cf25), new BigNumber(-841));
          r0 = (_1572_v39).dtor_cf30;
          let _index286 = _module.__default.safeIndex(new BigNumber(288), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index286] = _1569_v36;
          let _1573_v40;
          _1573_v40 = _dafny.Seq.of(_1569_v36, _1569_v36);
          let _1574_v41;
          _1574_v41 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((p3)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((p3).length))]),p1);
          let _index287 = _module.__default.safeIndex(new BigNumber(288), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index287] = (_1573_v40)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_1574_v41).length), (p3)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((p3).length))]), new BigNumber((_1573_v40).length))];
          (globalState).f12 = p1;
        } else {
          let _1575_v42;
          let _nw253 = new _module.C2();
          _nw253.__ctor();
          _1575_v42 = _nw253;
          _1575_v42 = _1575_v42;
          let _1576_v43;
          let _nw254 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1576_v43 = _nw254;
          _1576_v43 = _1576_v43;
          let _index288 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length));
          (_1576_v43)[_index288] = p2;
          let _index289 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length));
          let _rhs269 = p1;
          let _rhs270 = ((_this).f25).minus((_this).f25);
          let _rhs271 = (_dafny.ZERO).minus((_this).f25);
          let _rhs272 = p2;
          let _lhs196 = globalState;
          let _lhs197 = _1576_v43;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length));
          _lhs196.f12 = _rhs269;
          r1 = _rhs270;
          r1 = _rhs271;
          _lhs197[_lhs198] = _rhs272;
          let _1577_v44;
          _1577_v44 = _dafny.Seq.of(new BigNumber(928));
          let _1578_v45;
          _1578_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-80)), function (_1579_i3) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }),p2);
          let _1580_v46;
          _1580_v46 = _dafny.Seq.UnicodeFromString("f");
          let _1581_v47;
          _1581_v47 = _dafny.Set.fromElements(p1, _module.__default.fm0((_this).f25, globalState), (((_1578_v45).contains(_1580_v46)) ? ((_1578_v45).get(_1580_v46)) : ((_1576_v43)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length))])), !(p2), (_1576_v43)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length))]);
          let _1582_v48;
          _1582_v48 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements((_1576_v43)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length))], _module.__default.fm0(_module.__default.fm44(_module.__default.fm0((_1577_v44)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_1577_v44).length))], globalState), (_1576_v43)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_1576_v43).length))], p1, (_this).f25, globalState), globalState))).Union(_1581_v47),_1577_v44);
          _1582_v48 = _1582_v48;
          let _1583_v49;
          let _init57 = function (_1584_i4) {
            return (_1584_i4).multipliedBy((_this).f25);
          };
          let _nw255 = Array((new BigNumber(10)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw255.length); _i0_57++) {
            _nw255[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _1583_v49 = _nw255;
          _1583_v49 = _1583_v49;
        }
        let _1585_v50;
        let _nw256 = Array((new BigNumber(12)).toNumber());
        _nw256[(_dafny.ZERO).toNumber()] = _module.__default.fm0((_this).f25, globalState);
        _nw256[(_dafny.ONE).toNumber()] = p2;
        _nw256[(new BigNumber(2)).toNumber()] = !(p2);
        _nw256[(new BigNumber(3)).toNumber()] = !(p2);
        _nw256[(new BigNumber(4)).toNumber()] = p1;
        _nw256[(new BigNumber(5)).toNumber()] = p2;
        _nw256[(new BigNumber(6)).toNumber()] = p2;
        _nw256[(new BigNumber(7)).toNumber()] = p2;
        _nw256[(new BigNumber(8)).toNumber()] = p1;
        _nw256[(new BigNumber(9)).toNumber()] = p2;
        _nw256[(new BigNumber(10)).toNumber()] = false;
        _nw256[(new BigNumber(11)).toNumber()] = p2;
        _1585_v50 = _nw256;
        let _1586_v51;
        _1586_v51 = _dafny.Seq.of(_1585_v50);
        let _1587_v52;
        _1587_v52 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1588_v53;
        _1588_v53 = _dafny.Seq.of(_dafny.Seq.of(_1587_v52, _1587_v52, _1587_v52));
        let _1589_v54;
        _1589_v54 = _module.D0.create_DC0(p1);
        let _1590_v55;
        _1590_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1589_v54,p1);
        let _1591_v56;
        _1591_v56 = _dafny.Seq.UnicodeFromString("vpr");
        let _1592_v58;
        _1592_v58 = _dafny.Seq.of(_1589_v54);
        let _rhs273 = _module.__default.fm21(_1590_v55, _module.__default.safeModuloInt((_this).f25, new BigNumber((_1588_v53).length)), (p1) || (p2), _1591_v56, globalState);
        let _rhs274 = _dafny.Seq.update(_1586_v51, _module.__default.safeIndex((_this).f25, new BigNumber((_1586_v51).length)), _1585_v50);
        let _rhs275 = _dafny.Seq.Concat(_dafny.Seq.of(_1591_v56, _dafny.Seq.of(_1587_v52, _module.__default.fm39(p1, globalState), _1587_v52, _1587_v52, _1587_v52)), _module.__default.fm58(_1591_v56, _module.__default.fm21(function () {
          let _coll59 = new _dafny.Map();
          for (const _compr_59 of (_1592_v58).Elements) {
            let _1593_v57 = _compr_59;
            if (_dafny.Seq.contains(_1592_v58, _1593_v57)) {
              _coll59.push([_1593_v57,p2]);
            }
          }
          return _coll59;
        }(), (_this).f25, p1, _dafny.Seq.update(_1591_v56, _module.__default.safeIndex((_this).f25, new BigNumber((_1591_v56).length)), _1587_v52), globalState), p1, (_this).f25, globalState));
        let _rhs276 = (_this).f25;
        r1 = _rhs273;
        _1586_v51 = _rhs274;
        _1588_v53 = _rhs275;
        r1 = _rhs276;
        let _1594_v59;
        _1594_v59 = _module.D12.create_DC32(p1);
        let _1595_v60;
        let _nw257 = new _module.C3();
        _nw257.__ctor((_this).f25, (_1594_v59).dtor_cf47, (_this).f25, (_this).f26);
        _1595_v60 = _nw257;
      }
      let _index290 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length));
      (p3)[_index290] = _module.__default.safeModuloInt((_this).f25, (_this).f25);
      let _1596_v61;
      _1596_v61 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(_this).f25);
      let _1597_v62;
      _1597_v62 = _module.D9.create_DC26(_1596_v61);
      let _index291 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length));
      (p3)[_index291] = (_dafny.ZERO).minus(new BigNumber((function (_source22) {
        if (_source22.is_DC27) {
          let _1598___mcc_h0 = (_source22).cf39;
          let _1599___mcc_h1 = (_source22).cf40;
          let _1600_cf40 = _1599___mcc_h1;
          let _1601_cf39 = _1598___mcc_h0;
          return (_dafny.MultiSet.fromElements((_this).f25, _1600_cf40)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1600_cf40,_1601_cf39)).length))));
        } else {
          let _1602___mcc_h2 = (_source22).cf38;
          let _1603_cf38 = _1602___mcc_h2;
          return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kjvuwsmtm")).length), new BigNumber(183), (_dafny.ZERO).minus((_this).f25), (_dafny.ZERO).minus((_this).f25), (_this).f25);
        }
      }(_1597_v62)).cardinality()));
      let _1604_v63;
      _1604_v63 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jfbukxl"),new BigNumber(223));
      let _1605_v64;
      _1605_v64 = _dafny.Seq.UnicodeFromString("t");
      let _1606_v65;
      _1606_v65 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f25),(_this).f25);
      let _1607_v66;
      _1607_v66 = _dafny.Set.fromElements((((_1606_v65).contains(new BigNumber(467))) ? ((_1606_v65).get(new BigNumber(467))) : ((_this).f25)), new BigNumber((_1605_v64).length));
      let _1608_v67;
      _1608_v67 = _module.D4.create_DC11(_1607_v66);
      _1604_v63 = (_1604_v63).Merge(_module.__default.fm59(_1605_v64, _1608_v67, globalState));
      let _1609_v68;
      _1609_v68 = _dafny.Seq.of(_dafny.Seq.Concat(_1605_v64, _1605_v64), _1605_v64, _dafny.Seq.Create(_module.__default.abs(new BigNumber(373)), function (_1610_i5) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }));
      _1605_v64 = (_1609_v68)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm44(p2, p1, true, (_this).f25, globalState)), new BigNumber((_1609_v68).length))];
      let _1611_v69;
      _1611_v69 = _dafny.MultiSet.fromElements((p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))]);
      _1611_v69 = _dafny.MultiSet.fromElements((p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))], (p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))], (p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))]);
      r0 = (p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))];
      r1 = (_module.__default.safeModuloInt((p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))], (p3)[_module.__default.safeIndex(new BigNumber(853), new BigNumber((p3).length))])).multipliedBy(_module.__default.fm44(false, !(p2), p1, new BigNumber(7), globalState));
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      (globalState).f12 = !((_this).f25).isEqualTo(new BigNumber(702));
      r0 = (_dafny.ZERO).minus((new BigNumber(701)).minus((_this).f25));
      let _1612_v0;
      let _nw258 = new _module.C0();
      _nw258.__ctor();
      _1612_v0 = _nw258;
      let _1613_i0;
      _1613_i0 = _dafny.ZERO;
      L11: {
        while (_module.__default.fm0((_this).f25, globalState)) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1613_i0)) {
              break L11;
            }
            _1613_i0 = (_1613_i0).plus(_dafny.ONE);
            r0 = new BigNumber((_dafny.Seq.UnicodeFromString("dwkghx")).length);
            (globalState).f12 = p1;
            let _1614_v1;
            let _nw259 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
            _1614_v1 = _nw259;
            let _index292 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1614_v1).length));
            (_1614_v1)[_index292] = _module.__default.safeDivisionInt((_this).f25, (_dafny.ZERO).minus((_this).f25));
            let _1615_v2;
            _1615_v2 = _dafny.Seq.of((_this).f25, (_this).f25);
            let _1616_v3;
            _1616_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f25),p1);
            let _index293 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1614_v1).length));
            (_1614_v1)[_index293] = (new BigNumber(513)).plus(((p1) ? ((_1615_v2)[_module.__default.safeIndex(new BigNumber((_1616_v3).length), new BigNumber((_1615_v2).length))]) : ((_this).f25)));
            (globalState).f12 = p1;
          }
        }
      }
      let _1617_v4;
      _1617_v4 = _dafny.Set.fromElements((_this).f25, (_this).f25);
      r0 = new BigNumber(((_1617_v4).Union(_1617_v4)).length);
      let _1618_v5;
      let _nw260 = new _module.C0();
      _nw260.__ctor();
      _1618_v5 = _nw260;
      r0 = _module.__default.safeDivisionInt(new BigNumber(-364), (_this).f25);
      return r0;
    }
    m11(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D3.Default();
      let r1 = false;
      let _1619_v0;
      _1619_v0 = new _dafny.CodePoint('q'.codePointAt(0));
      let _1620_v1;
      _1620_v1 = _dafny.Seq.UnicodeFromString("nxtquho");
      let _1621_v2;
      _1621_v2 = _module.D1.create_DC6(_1619_v0, p1, _1620_v1);
      let _source23 = _1621_v2;
      if (_source23.is_DC5) {
        r1 = p1;
        (globalState).f12 = p1;
        let _1622_v3;
        _1622_v3 = _dafny.Seq.of(p0, p0, p0);
        _1622_v3 = _1622_v3;
        if (p1) {
          let _1623_v4;
          _1623_v4 = new BigNumber(834);
          _1623_v4 = new BigNumber(-982);
          (globalState).f12 = !(new BigNumber(291)).isEqualTo(_1623_v4);
          _1623_v4 = _1623_v4;
          _1623_v4 = _1623_v4;
          let _1624_v5;
          _1624_v5 = _dafny.MultiSet.fromElements(p1);
          _1624_v5 = _dafny.MultiSet.fromElements(_dafny.Seq.IsProperPrefixOf(_1620_v1, _1620_v1));
        } else {
          let _1625_v6;
          _1625_v6 = _module.D0.create_DC3((_this).f25, p1);
          let _1626_v7;
          _1626_v7 = _module.D6.create_DC19(function (_pat_let28_0) {
  return function (_1627_dt__update__tmp_h0) {
    return function (_pat_let29_0) {
      return function (_1628_dt__update_hcf4_h0) {
        return _module.D0.create_DC3(_1628_dt__update_hcf4_h0, (_1627_dt__update__tmp_h0).dtor_cf5);
      }(_pat_let29_0);
    }((_this).f25);
  }(_pat_let28_0);
}(_1625_v6), (_module.__default.fm44(p1, p1, p1, (_this).f25, globalState)).multipliedBy(new BigNumber(915)));
          _1626_v7 = _module.__default.fm60(globalState);
          let _1629_v8;
          let _init58 = ((_1630_p1) => function (_1631_i0) {
            return (_1630_p1) || (_1630_p1);
          })(p1);
          let _nw261 = Array((new BigNumber(18)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw261.length); _i0_58++) {
            _nw261[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _1629_v8 = _nw261;
          let _index294 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length));
          (_1629_v8)[_index294] = p1;
          let _1632_v9;
          let _nw262 = new _module.C4();
          _nw262.__ctor((_this).f26, (_this).f25, (_this).f26);
          _1632_v9 = _nw262;
          let _index295 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_1629_v8).length));
          (_1629_v8)[_index295] = false;
          let _1633_v10;
          _1633_v10 = _dafny.Set.fromElements(!(p1), p1);
          let _1634_v11;
          _1634_v11 = _dafny.MultiSet.fromElements(_1619_v0, _1619_v0, _1619_v0);
          let _index296 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length));
          let _index297 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_1629_v8).length));
          let _rhs277 = (new BigNumber(((_1633_v10).Difference(_1633_v10)).length)).isLessThan((_1632_v9).f25);
          let _rhs278 = (((new BigNumber((_1634_v11).cardinality())).isLessThan((_1632_v9).f25)) ? (_1632_v9) : (_1632_v9));
          let _rhs279 = p1;
          let _lhs199 = _1629_v8;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length));
          let _lhs201 = _1629_v8;
          let _lhs202 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_1629_v8).length));
          _lhs199[_lhs200] = _rhs277;
          _1632_v9 = _rhs278;
          _lhs201[_lhs202] = _rhs279;
          let _index298 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length));
          let _rhs280 = (_1629_v8)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length))];
          let _rhs281 = p1;
          let _rhs282 = (_1629_v8)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length))];
          let _rhs283 = _module.__default.fm0(new BigNumber(-567), globalState);
          let _rhs284 = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("muieyuiqy"), _1619_v0);
          let _lhs203 = globalState;
          let _lhs204 = globalState;
          let _lhs205 = _1629_v8;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1629_v8).length));
          let _lhs207 = globalState;
          _lhs203.f12 = _rhs280;
          _lhs204.f12 = _rhs281;
          _lhs205[_lhs206] = _rhs282;
          _lhs207.f12 = _rhs283;
          r1 = _rhs284;
          (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1620_v1, _1620_v1), _dafny.Seq.Concat(_1620_v1, _1620_v1));
          let _1635_v12;
          _1635_v12 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          _1635_v12 = (_1635_v12).update(p1, p0);
        }
      } else if (_source23.is_DC6) {
        let _1636___mcc_h0 = (_source23).cf7;
        let _1637___mcc_h1 = (_source23).cf8;
        let _1638___mcc_h2 = (_source23).cf9;
        let _1639_cf9 = _1638___mcc_h2;
        let _1640_cf8 = _1637___mcc_h1;
        let _1641_cf7 = _1636___mcc_h0;
        let _1642_v13;
        _1642_v13 = new BigNumber(873);
        _1642_v13 = _module.__default.safeDivisionInt(_module.__default.fm18(globalState), (_this).f25);
        _1642_v13 = (_this).f25;
        (globalState).f12 = _1640_cf8;
        let _1643_v14;
        _1643_v14 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1640_cf8,p1)).length));
        _1643_v14 = _1643_v14;
      } else if (_source23.is_DC4) {
        let _1644___mcc_h3 = (_source23).cf6;
        let _1645_cf6 = _1644___mcc_h3;
        if (p1) {
          let _1646_v15;
          _1646_v15 = new BigNumber(903);
          let _1647_v16;
          let _nw263 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1647_v16 = _nw263;
          let _1648_v17;
          _1648_v17 = _dafny.MultiSet.fromElements(_1646_v15, _1646_v15, (_this).f25);
          let _index299 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length));
          (_1647_v16)[_index299] = (_1645_cf6)[_module.__default.safeIndex((((_1648_v17).contains((_this).f25)) ? ((_1648_v17).get((_this).f25)) : (_1646_v15)), new BigNumber((_1645_cf6).length))];
          let _index300 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
          (p0)[_index300] = _1646_v15;
          let _1649_v18;
          _1649_v18 = _module.D10.create_DC29(p1, p1, (_this).f25);
          let _index301 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length));
          let _index302 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
          let _rhs285 = (_dafny.ZERO).minus((_module.D14.create_DC38(_1646_v15)).dtor_cf58);
          let _rhs286 = p1;
          let _rhs287 = (_dafny.ZERO).minus(((_this).f25).plus((_1649_v18).dtor_cf44));
          let _rhs288 = (_this).f25;
          let _rhs289 = _module.__default.fm0(_module.__default.safeDivisionInt(_1646_v15, (_this).f25), globalState);
          let _lhs208 = _1647_v16;
          let _lhs209 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length));
          let _lhs210 = p0;
          let _lhs211 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
          let _lhs212 = globalState;
          _1646_v15 = _rhs285;
          _lhs208[_lhs209] = _rhs286;
          _lhs210[_lhs211] = _rhs287;
          _1646_v15 = _rhs288;
          _lhs212.f12 = _rhs289;
          let _1650_v19;
          _1650_v19 = _dafny.MultiSet.fromElements((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))], p1, p1, p1, (_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))]);
          let _1651_v20;
          _1651_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))],(((_1650_v19).contains(p1)) ? ((_1650_v19).get(p1)) : (new BigNumber(279))));
          _1651_v20 = (_1651_v20).update(p1, new BigNumber(-990));
          let _1652_v22;
          _1652_v22 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Create(_module.__default.abs(new BigNumber(212)), function (_1653_i1) {
            return function () {
              let _coll60 = new _dafny.Set();
              for (const _compr_60 of _dafny.IntegerRange(new BigNumber(621), new BigNumber(454))) {
                let _1654_v21 = _compr_60;
                if (((new BigNumber(621)).isLessThanOrEqualTo(_1654_v21)) && ((_1654_v21).isLessThan(new BigNumber(454)))) {
                  _coll60.add((_1654_v21).minus(new BigNumber(969)));
                }
              }
              return _coll60;
            }();
          }));
          let _1655_v23;
          _1655_v23 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fyq")).length));
          let _1656_v24;
          _1656_v24 = _dafny.Seq.of(_1655_v23, _1655_v23, _1655_v23, _1655_v23);
          let _1657_v26;
          _1657_v26 = _dafny.Set.fromElements(_module.__default.fm47((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))], (p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))], _1646_v15, _1619_v0, globalState));
          let _1658_v28;
          _1658_v28 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))],_1646_v15);
          let _1659_v30;
          _1659_v30 = _dafny.Seq.of(_module.__default.fm44(true, p1, (_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))], (_this).f25, globalState));
          let _1660_v34;
          let _nw264 = Array((new BigNumber(29)).toNumber());
          _nw264[(_dafny.ZERO).toNumber()] = function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of ((((_1652_v22).contains((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))])) ? ((_1652_v22).get((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))])) : (_1656_v24))).Elements) {
              let _1661_v25 = _compr_61;
              if (_dafny.Seq.contains((((_1652_v22).contains((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))])) ? ((_1652_v22).get((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))])) : (_1656_v24)), _1661_v25)) {
                _coll61.add(_1661_v25);
              }
            }
            return _coll61;
          }();
          _nw264[(_dafny.ONE).toNumber()] = (((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))]) ? (_1657_v26) : (_dafny.Set.fromElements((_1656_v24)[_module.__default.safeIndex(_1646_v15, new BigNumber((_1656_v24).length))])));
          _nw264[(new BigNumber(2)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(3)).toNumber()] = (_1657_v26).Difference(_1657_v26);
          _nw264[(new BigNumber(4)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(5)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(6)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(7)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(8)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(9)).toNumber()] = (_1657_v26).Difference(_1657_v26);
          _nw264[(new BigNumber(10)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(11)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(_1655_v23, function () {
            let _coll62 = new _dafny.Set();
            for (const _compr_62 of (_1655_v23).Elements) {
              let _1662_v27 = _compr_62;
              if ((_1655_v23).contains(_1662_v27)) {
                _coll62.add(_module.__default.safeDivisionInt(_1662_v27, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ohtqnsan")).length),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(383), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length),new BigNumber((_dafny.Seq.UnicodeFromString("dvft")).length))).length), new BigNumber(-285))).cardinality()))).length)));
              }
            }
            return _coll62;
          }(), function () {
            let _coll63 = new _dafny.Set();
            for (const _compr_63 of (_1658_v28).Keys.Elements) {
              let _1663_v29 = _compr_63;
              if ((_1658_v28).contains(_1663_v29)) {
                _coll63.add((_1663_v29).minus(new BigNumber(873)));
              }
            }
            return _coll63;
          }(), _dafny.Set.fromElements(new BigNumber((_1659_v30).length)));
          _nw264[(new BigNumber(13)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(14)).toNumber()] = (_dafny.Set.fromElements(_1655_v23, _dafny.Set.fromElements(new BigNumber((function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of (_1655_v23).Elements) {
              let _1664_v31 = _compr_64;
              if ((_1655_v23).contains(_1664_v31)) {
                _coll64.push([(_1664_v31).multipliedBy(new BigNumber((_1650_v19).cardinality())),(_1645_cf6)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))],(_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))])).length), new BigNumber((_1645_cf6).length))]]);
              }
            }
            return _coll64;
          }()).length), _1646_v15))).Union(_1657_v26);
          _nw264[(new BigNumber(15)).toNumber()] = (_dafny.Set.fromElements(_dafny.Set.fromElements((p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))]), _1655_v23)).Difference(_1657_v26);
          _nw264[(new BigNumber(16)).toNumber()] = (_dafny.Set.fromElements(_1655_v23, _1655_v23)).Intersect(_dafny.Set.fromElements(function () {
            let _coll65 = new _dafny.Set();
            for (const _compr_65 of _dafny.IntegerRange(new BigNumber(-431), new BigNumber(-120))) {
              let _1665_v32 = _compr_65;
              if (((new BigNumber(-431)).isLessThanOrEqualTo(_1665_v32)) && ((_1665_v32).isLessThan(new BigNumber(-120)))) {
                _coll65.add((_1665_v32).plus(new BigNumber(987)));
              }
            }
            return _coll65;
          }(), _1655_v23));
          _nw264[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(_1655_v23, _1655_v23);
          _nw264[(new BigNumber(18)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(19)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(20)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(21)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(22)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(23)).toNumber()] = _module.__default.fm61(_1646_v15, (_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))], p1, globalState);
          _nw264[(new BigNumber(24)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(25)).toNumber()] = _module.__default.fm61((_this).f25, !((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))]), (_1645_cf6)[_module.__default.safeIndex(_1646_v15, new BigNumber((_1645_cf6).length))], globalState);
          _nw264[(new BigNumber(26)).toNumber()] = _1657_v26;
          _nw264[(new BigNumber(27)).toNumber()] = function () {
            let _coll66 = new _dafny.Set();
            for (const _compr_66 of (_1656_v24).Elements) {
              let _1666_v33 = _compr_66;
              if (_dafny.Seq.contains(_1656_v24, _1666_v33)) {
                _coll66.add(_1666_v33);
              }
            }
            return _coll66;
          }();
          _nw264[(new BigNumber(28)).toNumber()] = _module.__default.fm61((p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))], p1, p1, globalState);
          _1660_v34 = _nw264;
          let _index303 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_1660_v34).length));
          (_1660_v34)[_index303] = _module.__default.fm61((_this).f25, p1, (_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))], globalState);
          let _1667_v35;
          _1667_v35 = _module.D9.create_DC27((_this).f25, (p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))]);
          let _index304 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_1660_v34).length));
          let _index305 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
          let _rhs290 = ((_module.__default.fm0(_1646_v15, globalState)) ? ((_1657_v26).Difference(_1657_v26)) : (_1657_v26));
          let _rhs291 = (false) && ((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))]);
          let _rhs292 = new BigNumber(((((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))]) ? (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gf"), _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))], new BigNumber((_dafny.Seq.UnicodeFromString("gf")).length)), _1619_v0)) : ((_module.D1.create_DC6(_1619_v0, (_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))], _dafny.Seq.update(_1620_v1, _module.__default.safeIndex((_this).f25, new BigNumber((_1620_v1).length)), _1619_v0))).dtor_cf9))).length);
          let _rhs293 = new BigNumber((_1659_v30).length);
          let _rhs294 = _module.D9.create_DC27(((_this).f25).minus((p0)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length))]), new BigNumber(-817));
          let _lhs213 = _1660_v34;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_1660_v34).length));
          let _lhs215 = globalState;
          let _lhs216 = p0;
          let _lhs217 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
          _lhs213[_lhs214] = _rhs290;
          _lhs215.f12 = _rhs291;
          _1646_v15 = _rhs292;
          _lhs216[_lhs217] = _rhs293;
          _1667_v35 = _rhs294;
          let _1668_v36;
          let _nw265 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1668_v36 = _nw265;
          let _1669_v37;
          _1669_v37 = _dafny.Map.Empty.slice().updateUnsafe((_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))],_1668_v36);
          _1669_v37 = (_1669_v37).update(p1, _1668_v36);
          let _1670_v38;
          _1670_v38 = _dafny.Set.fromElements(p1);
          let _1671_v39;
          _1671_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1670_v38,_dafny.Map.Empty.slice().updateUnsafe(p1,(_1647_v16)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1647_v16).length))]));
          let _index306 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((p0).length));
          (p0)[_index306] = new BigNumber(((_1671_v39).Merge(_1671_v39)).length);
        } else {
          (globalState).f12 = !(p1) || (!(p1));
          let _1672_v40;
          _1672_v40 = new BigNumber(747);
          _1672_v40 = (_this).f25;
          let _1673_v41;
          _1673_v41 = _module.D14.create_DC38(_1672_v40);
          let _1674_v42;
          _1674_v42 = _dafny.Set.fromElements(_1673_v41);
          let _1675_v43;
          _1675_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1674_v42,_1620_v1);
          _1675_v43 = _1675_v43;
          _1672_v40 = new BigNumber((_module.__default.fm20(!(p1) || (p1), globalState)).length);
          let _1676_v44;
          _1676_v44 = _dafny.MultiSet.fromElements(p1, false, _dafny.Seq.IsProperPrefixOf(_1620_v1, _1620_v1));
          _1676_v44 = _module.__default.fm50(new BigNumber(168), ((p1) ? (_1620_v1) : (_1620_v1)), p1, new BigNumber(18), globalState);
        }
        if (p1) {
          let _1677_v45;
          _1677_v45 = _dafny.Set.fromElements(p1, p1);
          _1677_v45 = ((_1677_v45).Intersect(_1677_v45)).Union(_1677_v45);
          let _1678_v46;
          let _nw266 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1678_v46 = _nw266;
          let _1679_v47;
          _1679_v47 = _dafny.Seq.of(_1678_v46, _1678_v46, _1678_v46);
          _1678_v46 = (_1679_v47)[_module.__default.safeIndex((_this).f25, new BigNumber((_1679_v47).length))];
          let _1680_v48;
          _1680_v48 = new BigNumber(309);
          let _1681_v49;
          _1681_v49 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f25, globalState),_module.__default.fm44(p1, p1, p1, new BigNumber(661), globalState));
          let _1682_v50;
          _1682_v50 = _dafny.Set.fromElements(_1680_v48, (((_1681_v49).contains(p1)) ? ((_1681_v49).get(p1)) : (_1680_v48)));
          let _1683_v51;
          _1683_v51 = _dafny.Seq.of(_1680_v48);
          _1680_v48 = new BigNumber(((_1682_v50).Difference(function () {
            let _coll67 = new _dafny.Set();
            for (const _compr_67 of (_1683_v51).Elements) {
              let _1684_v52 = _compr_67;
              if (_dafny.Seq.contains(_1683_v51, _1684_v52)) {
                _coll67.add((_1684_v52).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))).cardinality())));
              }
            }
            return _coll67;
          }())).length);
          let _1685_v53;
          _1685_v53 = _module.D1.create_DC4(_1645_cf6);
          let _1686_v54;
          _1686_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1685_v53);
          let _pat_let_tv31 = _1645_cf6;
          _1686_v54 = (_1686_v54).update(p0, function (_pat_let30_0) {
            return function (_1687_dt__update__tmp_h1) {
              return function (_pat_let31_0) {
                return function (_1688_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4(_1688_dt__update_hcf6_h0);
                }(_pat_let31_0);
              }(_pat_let_tv31);
            }(_pat_let30_0);
          }(_module.D1.create_DC4(_1645_cf6)));
          let _1689_v55;
          _1689_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1620_v1).length),p1);
          let _1690_v57;
          _1690_v57 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1691_v58;
          _1691_v58 = _dafny.Seq.of(_1689_v55);
          let _1692_v59;
          _1692_v59 = _1689_v55;
          let _1693_v60;
          let _nw267 = Array((new BigNumber(26)).toNumber());
          _nw267[(_dafny.ZERO).toNumber()] = (_1689_v55).Merge(_1689_v55);
          _nw267[(_dafny.ONE).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(2)).toNumber()] = (_1689_v55).Merge(_1689_v55);
          _nw267[(new BigNumber(3)).toNumber()] = (function () {
            let _coll68 = new _dafny.Map();
            for (const _compr_68 of _dafny.IntegerRange(new BigNumber(629), new BigNumber(-124))) {
              let _1694_v56 = _compr_68;
              if (((new BigNumber(629)).isLessThanOrEqualTo(_1694_v56)) && ((_1694_v56).isLessThan(new BigNumber(-124)))) {
                _coll68.push([_module.__default.safeModuloInt(_1694_v56, new BigNumber((_1683_v51).length)),p1]);
              }
            }
            return _coll68;
          }()).Merge(_1689_v55);
          _nw267[(new BigNumber(4)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(5)).toNumber()] = (_1689_v55).Merge(_1689_v55);
          _nw267[(new BigNumber(6)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm44(p1, !(p1), p1, _1680_v48, globalState),(((_1690_v57).contains(p1)) ? ((_1690_v57).get(p1)) : (p1)));
          _nw267[(new BigNumber(8)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(9)).toNumber()] = (_1689_v55).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1680_v48,p1));
          _nw267[(new BigNumber(10)).toNumber()] = _module.__default.fm36(p1, _1689_v55, _1619_v0, globalState);
          _nw267[(new BigNumber(11)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
          _nw267[(new BigNumber(13)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(14)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(15)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(16)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(17)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC3((_this).f25, _module.__default.fm0(new BigNumber((_1682_v50).length), globalState))).dtor_cf4,p1);
          _nw267[(new BigNumber(19)).toNumber()] = (_1689_v55).update(_1680_v48, p1);
          _nw267[(new BigNumber(20)).toNumber()] = (_1689_v55).Merge(_1689_v55);
          _nw267[(new BigNumber(21)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(22)).toNumber()] = (_1689_v55).Merge(_1689_v55);
          _nw267[(new BigNumber(23)).toNumber()] = (_1691_v58)[_module.__default.safeIndex(_1680_v48, new BigNumber((_1691_v58).length))];
          _nw267[(new BigNumber(24)).toNumber()] = _1689_v55;
          _nw267[(new BigNumber(25)).toNumber()] = ((_1692_v59)).Merge(_1689_v55);
          _1693_v60 = _nw267;
          let _init59 = ((_1695_v55) => function (_1696_i2) {
            return _1695_v55;
          })(_1689_v55);
          let _nw268 = Array((new BigNumber(14)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw268.length); _i0_59++) {
            _nw268[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _1693_v60 = _nw268;
        } else {
          let _1697_v61;
          _1697_v61 = new BigNumber(750);
          _1697_v61 = ((_dafny.ZERO).minus((_this).f25)).plus(new BigNumber((_1645_cf6).length));
          let _1698_v62;
          _1698_v62 = _dafny.Seq.of((_this).f26, (_this).f26, (_this).f26);
          let _1699_v63;
          _1699_v63 = _dafny.MultiSet.fromElements(_1619_v0);
          let _1700_v64;
          _1700_v64 = _dafny.Seq.of(new BigNumber((_1699_v63).cardinality()));
          let _1701_v65;
          _1701_v65 = _dafny.Set.fromElements(_1700_v64);
          let _1702_v66;
          let _nw269 = new _module.C4();
          _nw269.__ctor((_1698_v62)[_module.__default.safeIndex((_this).f25, new BigNumber((_1698_v62).length))], new BigNumber(((_1701_v65).Difference(_dafny.Set.fromElements(_1700_v64))).length), (_this).f26);
          _1702_v66 = _nw269;
          let _1703_v67;
          _1703_v67 = _dafny.MultiSet.fromElements(p1);
          _1697_v61 = _module.__default.safeDivisionInt(new BigNumber(102), _module.__default.safeDivisionInt(new BigNumber((_1703_v67).cardinality()), _1697_v61));
          let _1704_v68;
          let _init60 = ((_1705_v67) => function (_1706_i3) {
            return (_1705_v67).Union(_1705_v67);
          })(_1703_v67);
          let _nw270 = Array((new BigNumber(17)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw270.length); _i0_60++) {
            _nw270[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _1704_v68 = _nw270;
          let _index307 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_1704_v68).length));
          (_1704_v68)[_index307] = _module.__default.fm50(new BigNumber(-87), _1620_v1, p1, (_this).f25, globalState);
          let _index308 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_1704_v68).length));
          (_1704_v68)[_index308] = (_dafny.MultiSet.FromArray(_1645_cf6)).Union((_1703_v67).Difference(_1703_v67));
          let _1707_v69;
          _1707_v69 = _dafny.MultiSet.fromElements(_1697_v61, new BigNumber(((_1704_v68)[_module.__default.safeIndex(new BigNumber(403), new BigNumber((_1704_v68).length))]).cardinality()), _1697_v61);
          _1707_v69 = (_dafny.MultiSet.fromElements((_this).f25, (_this).f25)).Union(_1707_v69);
        }
        let _1708_v70;
        let _nw271 = Array((new BigNumber(9)).toNumber());
        _nw271[(_dafny.ZERO).toNumber()] = !(p1);
        _nw271[(_dafny.ONE).toNumber()] = p1;
        _nw271[(new BigNumber(2)).toNumber()] = p1;
        _nw271[(new BigNumber(3)).toNumber()] = p1;
        _nw271[(new BigNumber(4)).toNumber()] = _module.__default.fm0((_this).f25, globalState);
        _nw271[(new BigNumber(5)).toNumber()] = false;
        _nw271[(new BigNumber(6)).toNumber()] = p1;
        _nw271[(new BigNumber(7)).toNumber()] = p1;
        _nw271[(new BigNumber(8)).toNumber()] = p1;
        _1708_v70 = _nw271;
        let _1709_v71;
        _1709_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1708_v70,(_this).f25);
        _1709_v71 = (_1709_v71).update(_1708_v70, (_this).f25);
        (globalState).f12 = p1;
      } else {
        let _1710___mcc_h4 = (_source23).cf10;
        let _1711_cf10 = _1710___mcc_h4;
        let _1712_v72;
        _1712_v72 = _module.D7.create_DC21(p0);
        let _1713_v73;
        _1713_v73 = new BigNumber(299);
        let _1714_v74;
        _1714_v74 = _dafny.Seq.of(_1620_v1);
        let _1715_v75;
        _1715_v75 = _module.D9.create_DC27((_this).f25, _1713_v73);
        let _1716_v76;
        _1716_v76 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(p1)).length), new BigNumber((_1620_v1).length));
        let _rhs295 = _1712_v72;
        let _rhs296 = p1;
        let _rhs297 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_1714_v74, _module.__default.fm58(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), _1619_v0, _1619_v0), (_this).f25, p1, _module.__default.fm44(p1, _module.__default.fm0((_this).f25, globalState), p1, new BigNumber((_dafny.Seq.of((_this).f25, new BigNumber((_1620_v1).length), (_this).f25, _1713_v73)).length), globalState), globalState)), _module.__default.safeIndex(new BigNumber((_1620_v1).length), new BigNumber((_dafny.Seq.Concat(_1714_v74, _module.__default.fm58(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), _1619_v0, _1619_v0), (_this).f25, p1, _module.__default.fm44(p1, _module.__default.fm0((_this).f25, globalState), p1, new BigNumber((_dafny.Seq.of((_this).f25, new BigNumber((_1620_v1).length), (_this).f25, _1713_v73)).length), globalState), globalState))).length)), _1620_v1), _dafny.Seq.Concat(_dafny.Seq.of(_1620_v1, _1620_v1, _1620_v1), _1714_v74));
        let _rhs298 = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements((_this).f25, (_this).f25, (_this).f25, (_1715_v75).dtor_cf39, _1713_v73)).length), (_this).f25);
        let _rhs299 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1620_v1, _1620_v1), _dafny.Seq.Concat(_dafny.Seq.of(_1619_v0, _1619_v0, _1619_v0, new _dafny.CodePoint('j'.codePointAt(0)), _1619_v0), (_1714_v74)[_module.__default.safeIndex(new BigNumber((_1716_v76).length), new BigNumber((_1714_v74).length))]));
        let _lhs218 = globalState;
        let _lhs219 = globalState;
        _1712_v72 = _rhs295;
        _lhs218.f12 = _rhs296;
        r1 = _rhs297;
        _1713_v73 = _rhs298;
        _lhs219.f10 = _rhs299;
        let _1717_v77;
        _1717_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1619_v0,(_this).f25);
        let _1718_v78;
        _1718_v78 = _dafny.Seq.of(p1, p1, false);
        let _1719_v79;
        _1719_v79 = _dafny.Map.Empty.slice().updateUnsafe((_1718_v78)[_module.__default.safeIndex(new BigNumber(56), new BigNumber((_1718_v78).length))],(_this).f25);
        _1717_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1619_v0,_module.__default.safeModuloInt(new BigNumber((_1719_v79).length), (_this).f25));
        let _1720_v80;
        _1720_v80 = _module.D5.create_DC17((_this).f25, (_1718_v78)[_module.__default.safeIndex(_1713_v73, new BigNumber((_1718_v78).length))]);
        let _1721_v81;
        _1721_v81 = _dafny.Set.fromElements(p1);
        let _pat_let_tv32 = _1721_v81;
        let _pat_let_tv33 = _1721_v81;
        let _source24 = function (_pat_let32_0) {
          return function (_1724_dt__update__tmp_h3) {
            return function (_pat_let35_0) {
              return function (_1725_dt__update_hcf27_h0) {
                return _module.D5.create_DC17((_1724_dt__update__tmp_h3).dtor_cf26, _1725_dt__update_hcf27_h0);
              }(_pat_let35_0);
            }((_pat_let_tv32).IsDisjointFrom(_pat_let_tv33));
          }(_pat_let32_0);
        }(function (_pat_let33_0) {
          return function (_1722_dt__update__tmp_h2) {
            return function (_pat_let34_0) {
              return function (_1723_dt__update_hcf26_h0) {
                return _module.D5.create_DC17(_1723_dt__update_hcf26_h0, (_1722_dt__update__tmp_h2).dtor_cf27);
              }(_pat_let34_0);
            }((_this).f25);
          }(_pat_let33_0);
        }(_1720_v80));
        if (_source24.is_DC16) {
          let _1726___mcc_h5 = (_source24).cf25;
          let _1727_cf25 = _1726___mcc_h5;
          let _1728_v83;
          _1728_v83 = _dafny.Map.Empty.slice().updateUnsafe(!(_1727_cf25),_module.D10.create_DC28(function () {
  let _coll69 = new _dafny.Map();
  for (const _compr_69 of (_1620_v1).Elements) {
    let _1729_v82 = _compr_69;
    if (_dafny.Seq.contains(_1620_v1, _1729_v82)) {
      _coll69.push([_1729_v82,p1]);
    }
  }
  return _coll69;
}()));
          let _1730_v84;
          let _nw272 = new _module.C0();
          _nw272.__ctor();
          _1730_v84 = _nw272;
          let _1731_v85;
          _1731_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1619_v0,_1727_cf25);
          let _1732_v86;
          _1732_v86 = _module.D10.create_DC28(_1731_v85);
          let _1733_v87;
          _1733_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1712_v72,_1730_v84)).length),_dafny.Map.Empty.slice().updateUnsafe(p1,_1732_v86));
          let _1734_v88;
          let _nw273 = Array((new BigNumber(14)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = _1728_v83;
          _nw273[(_dafny.ONE).toNumber()] = _1728_v83;
          _nw273[(new BigNumber(2)).toNumber()] = _1728_v83;
          _nw273[(new BigNumber(3)).toNumber()] = (((_1733_v87).contains(_1713_v73)) ? ((_1733_v87).get(_1713_v73)) : (_1728_v83));
          _nw273[(new BigNumber(4)).toNumber()] = _1728_v83;
          _nw273[(new BigNumber(5)).toNumber()] = _1728_v83;
          _nw273[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1727_cf25,_1732_v86);
          _nw273[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_1732_v86);
          _nw273[(new BigNumber(8)).toNumber()] = _1728_v83;
          _nw273[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_1732_v86);
          _nw273[(new BigNumber(10)).toNumber()] = (_1728_v83).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1727_cf25,_1732_v86));
          _nw273[(new BigNumber(11)).toNumber()] = (_1728_v83).Merge(_1728_v83);
          _nw273[(new BigNumber(12)).toNumber()] = _module.__default.fm62(globalState);
          _nw273[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_1732_v86);
          _1734_v88 = _nw273;
          let _index309 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1734_v88).length));
          (_1734_v88)[_index309] = _1728_v83;
          let _index310 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1734_v88).length));
          (_1734_v88)[_index310] = ((_1728_v83).Merge(_1728_v83)).Merge(((_1728_v83).update(p1, _1732_v86)).Merge(_1728_v83));
          _1713_v73 = (_1730_v84).fm26((((_1719_v79).contains(p1)) ? ((_1719_v79).get(p1)) : (_1713_v73)), ((false) ? (_1727_cf25) : (_1727_cf25)), globalState);
          let _1735_v89;
          _1735_v89 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f25).minus(new BigNumber(574)),_1727_cf25);
          _1735_v89 = (_1735_v89).update(((_this).f25).multipliedBy((_this).f25), _1727_cf25);
          (globalState).f6 = _1716_v76;
        } else if (_source24.is_DC17) {
          let _1736___mcc_h6 = (_source24).cf26;
          let _1737___mcc_h7 = (_source24).cf27;
          let _1738_cf27 = _1737___mcc_h7;
          let _1739_cf26 = _1736___mcc_h6;
          let _1740_v90;
          let _nw274 = Array((new BigNumber(20)).toNumber()).fill(false);
          _1740_v90 = _nw274;
          let _init61 = ((_1741_v78) => function (_1742_i4) {
            return !_dafny.areEqual(_1741_v78, _1741_v78);
          })(_1718_v78);
          let _nw275 = Array((_dafny.ONE).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw275.length); _i0_61++) {
            _nw275[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _1740_v90 = _nw275;
          let _1743_v91;
          _1743_v91 = _module.D0.create_DC3((_this).f25, _1738_cf27);
          _1719_v79 = (_1719_v79).update(((_1743_v91).dtor_cf4).isLessThan(new BigNumber(622)), _1739_cf26);
          let _1744_v92;
          _1744_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p1),p1);
          _1744_v92 = (_1744_v92).update(_dafny.Seq.of(p1, p1, _1738_cf27), ((_this).f25).isLessThan((_this).f25));
          let _1745_v93;
          let _init62 = ((_1746_v0) => function (_1747_i5) {
            return _1746_v0;
          })(_1619_v0);
          let _nw276 = Array((new BigNumber(26)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw276.length); _i0_62++) {
            _nw276[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _1745_v93 = _nw276;
          let _index311 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1745_v93).length));
          (_1745_v93)[_index311] = _module.__default.fm39(_1738_cf27, globalState);
          let _index312 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1745_v93).length));
          (_1745_v93)[_index312] = _1619_v0;
          let _index313 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1745_v93).length));
          let _index314 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1745_v93).length));
          let _rhs300 = ((p1) ? (_1740_v90) : (_1740_v90));
          let _rhs301 = _1619_v0;
          let _rhs302 = _1619_v0;
          let _lhs220 = _1745_v93;
          let _lhs221 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1745_v93).length));
          let _lhs222 = _1745_v93;
          let _lhs223 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1745_v93).length));
          _1740_v90 = _rhs300;
          _lhs220[_lhs221] = _rhs301;
          _lhs222[_lhs223] = _rhs302;
        } else {
          let _1748___mcc_h8 = (_source24).cf24;
          let _1749_cf24 = _1748___mcc_h8;
          (globalState).f12 = p1;
          let _1750_v94;
          _1750_v94 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(p1, globalState),false);
          let _1751_v95;
          _1751_v95 = _module.D0.create_DC2(p1);
          let _1752_v96;
          _1752_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1750_v94,_1751_v95);
          _1752_v96 = (_1752_v96).update(_1750_v94, _1751_v95);
          let _1753_v97;
          let _init63 = function (_1754_i6) {
            return false;
          };
          let _nw277 = Array((new BigNumber(10)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw277.length); _i0_63++) {
            _nw277[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _1753_v97 = _nw277;
          let _index315 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1753_v97).length));
          (_1753_v97)[_index315] = p1;
          let _index316 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1753_v97).length));
          (_1753_v97)[_index316] = _dafny.Seq.IsPrefixOf(_1620_v1, _dafny.Seq.Concat(_1620_v1, _dafny.Seq.UnicodeFromString("xum")));
          let _1755_v98;
          _1755_v98 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), ((_1756_v0) => function (_1757_i7) {
            return _1756_v0;
          })(_1619_v0)),_1753_v97);
          let _1758_v99;
          _1758_v99 = _dafny.Map.Empty.slice().updateUnsafe((((_1755_v98).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), ((_1761_v0) => function (_1762_i8) {
            return _1761_v0;
          })(_1619_v0)))) ? ((_1755_v98).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), ((_1759_v0) => function (_1760_i8) {
            return _1759_v0;
          })(_1619_v0)))) : (_1753_v97)),(_1753_v97)[_module.__default.safeIndex(new BigNumber(741), new BigNumber((_1753_v97).length))]);
          let _1763_v100;
          _1763_v100 = _module.D8.create_DC23(_1758_v99);
          let _1764_v101;
          _1764_v101 = _module.D8.create_DC25(_1763_v100);
          let _1765_v102;
          _1765_v102 = _dafny.Set.fromElements(_1764_v101);
          _1713_v73 = _module.__default.safeModuloInt(new BigNumber(569), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1620_v1, _module.__default.safeIndex(new BigNumber((_1765_v102).length), new BigNumber((_1620_v1).length)), _1619_v0), _1620_v1)).length));
        }
        let _rhs303 = p1;
        let _rhs304 = _1619_v0;
        let _rhs305 = _1713_v73;
        let _rhs306 = _module.__default.fm39(p1, globalState);
        let _rhs307 = _1620_v1;
        let _lhs224 = globalState;
        let _lhs225 = globalState;
        r1 = _rhs303;
        _lhs224.f23 = _rhs304;
        _1713_v73 = _rhs305;
        _lhs225.f23 = _rhs306;
        _1620_v1 = _rhs307;
      }
      let _1766_v103;
      let _nw278 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1766_v103 = _nw278;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1766_v103).length))) {
        let _1767_i9 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1767_i9)) && ((_1767_i9).isLessThan(new BigNumber((_1766_v103).length))))) {
          (_1766_v103)[(_1767_i9)] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(404)), ((_1768_v0) => function (_1769_i10) {
            return _1768_v0;
          })(_1619_v0));
        }
      }
      let _1770_v104;
      let _nw279 = Array((new BigNumber(11)).toNumber()).fill([]);
      _1770_v104 = _nw279;
      let _1771_v105;
      _1771_v105 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
      let _1772_v106;
      _1772_v106 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1771_v105);
      let _1773_v107;
      _1773_v107 = _dafny.Map.Empty.slice().updateUnsafe((_1772_v106).contains((_this).f25),_1770_v104);
      let _1774_v108;
      _1774_v108 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      _1770_v104 = (((_1773_v107).contains((((_1774_v108).contains(p1)) ? ((_1774_v108).get(p1)) : (p1)))) ? ((_1773_v107).get((((_1774_v108).contains(p1)) ? ((_1774_v108).get(p1)) : (p1)))) : (_1770_v104));
      let _1775_v109;
      _1775_v109 = new BigNumber(573);
      _1775_v109 = (_1775_v109).plus(_1775_v109);
      let _hi8 = _module.__default.safeModuloInt(_1775_v109, (_this).f25);
      for (let _1776_i11 = (_this).f25; _1776_i11.isLessThan(_hi8); _1776_i11 = _1776_i11.plus(_dafny.ONE)) {
        let _1777_v110;
        _1777_v110 = _dafny.MultiSet.fromElements(p1);
        _1775_v109 = ((((_this).f25).isLessThanOrEqualTo(_1776_i11)) ? ((((_1777_v110).contains(p1)) ? ((_1777_v110).get(p1)) : ((_this).f25))) : (_1776_i11));
        let _index317 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((p0).length));
        (p0)[_index317] = ((p1) ? (_1776_i11) : (_1776_i11));
        let _1778_v111;
        _1778_v111 = _dafny.Seq.of(_1777_v110, _1777_v110, _1777_v110);
        let _1779_v112;
        let _nw280 = new _module.C4();
        _nw280.__ctor((_this).f26, _1775_v109, _1766_v103);
        _1779_v112 = _nw280;
        let _1780_v113;
        _1780_v113 = _dafny.Map.Empty.slice().updateUnsafe(_1779_v112,p1);
        let _1781_v114;
        _1781_v114 = _dafny.Seq.of(_1780_v113);
        let _1782_v115;
        _1782_v115 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dtllwx")).length),(_1779_v112).f25);
        let _index318 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((p0).length));
        let _rhs308 = _dafny.Seq.IsPrefixOf(_1781_v114, _1781_v114);
        let _rhs309 = _module.__default.safeDivisionInt(_1775_v109, _module.__default.safeDivisionInt((_1779_v112).f25, (_dafny.ZERO).minus(new BigNumber((_1782_v115).length))));
        let _rhs310 = (_dafny.ZERO).minus((_this).f25);
        let _rhs311 = ((_1779_v112).f25).minus(_1776_i11);
        let _rhs312 = _1778_v111;
        let _lhs226 = globalState;
        let _lhs227 = p0;
        let _lhs228 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((p0).length));
        _lhs226.f12 = _rhs308;
        _1775_v109 = _rhs309;
        _lhs227[_lhs228] = _rhs310;
        _1775_v109 = _rhs311;
        _1778_v111 = _rhs312;
        let _1783_v116;
        let _nw281 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1783_v116 = _nw281;
        let _index319 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1783_v116).length));
        (_1783_v116)[_index319] = _module.__default.fm0((_dafny.ZERO).minus((_1779_v112).f25), globalState);
        let _index320 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1783_v116).length));
        (_1783_v116)[_index320] = true;
        let _1784_v117;
        _1784_v117 = _dafny.Map.Empty.slice().updateUnsafe((_1783_v116)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_1783_v116).length))],_1620_v1);
        let _1785_v118;
        _1785_v118 = _dafny.Seq.of((((_1777_v110).contains(p1)) ? ((_1777_v110).get(p1)) : ((_1779_v112).f25)));
        let _1786_v119;
        _1786_v119 = _dafny.MultiSet.fromElements(new BigNumber((_1785_v118).length), (_this).f25);
        let _1787_v120;
        _1787_v120 = _dafny.Seq.of(_module.__default.fm18(globalState), _module.__default.fm31((((_1784_v117).contains(p1)) ? ((_1784_v117).get(p1)) : (_1620_v1)), globalState), (((_1786_v119).contains(_1776_i11)) ? ((_1786_v119).get(_1776_i11)) : (_1776_i11)));
        let _rhs313 = ((_dafny.ZERO).minus(new BigNumber((_1777_v110).cardinality()))).multipliedBy(_1776_i11);
        let _rhs314 = (_1785_v118)[_module.__default.safeIndex(((_this).f25).minus(new BigNumber((_1771_v105).length)), new BigNumber((_1785_v118).length))];
        let _rhs315 = _1620_v1;
        let _rhs316 = !((((_1777_v110).IsSubsetOf(_dafny.MultiSet.fromElements(p1, p1))) ? (p1) : ((_1783_v116)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_1783_v116).length))])));
        let _rhs317 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_1620_v1, _1620_v1)).length));
        let _lhs229 = globalState;
        let _lhs230 = globalState;
        _1775_v109 = _rhs313;
        _1775_v109 = _rhs314;
        _lhs229.f10 = _rhs315;
        r1 = _rhs316;
        _lhs230.f6 = _rhs317;
      }
      let _1788_i12;
      _1788_i12 = _dafny.ZERO;
      L12: {
        while (p1) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1788_i12)) {
              break L12;
            }
            _1788_i12 = (_1788_i12).plus(_dafny.ONE);
            _1774_v108 = (_1774_v108).update(p1, (_1775_v109).isLessThanOrEqualTo(new BigNumber((_module.__default.fm63(p1, (_this).f25, p1, _module.__default.fm0((_this).f25, globalState), globalState)).cardinality())));
            if (p1) {
              let _index321 = _module.__default.safeIndex(new BigNumber(186), new BigNumber(((_this).f26).length));
              ((_this).f26)[_index321] = _dafny.Seq.UnicodeFromString("lubqn");
              let _index322 = _module.__default.safeIndex(new BigNumber(186), new BigNumber(((_this).f26).length));
              ((_this).f26)[_index322] = _1620_v1;
              let _1789_v121;
              _1789_v121 = _dafny.MultiSet.fromElements(p1, p1, p1, p1);
              let _1790_v122;
              _1790_v122 = _dafny.MultiSet.fromElements(_1789_v121);
              _1790_v122 = _1790_v122;
              _1775_v109 = _1775_v109;
              let _1791_v123;
              _1791_v123 = _dafny.Seq.of((_module.D13.create_DC36(_1774_v108, p1, (_this).f25)).dtor_cf56, (_this).f25);
              let _1792_v124;
              _1792_v124 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1791_v123, _module.__default.safeIndex((_this).f25, new BigNumber((_1791_v123).length)), _1775_v109),(_dafny.ZERO).minus((_this).f25));
              let _1793_v125;
              _1793_v125 = _dafny.Seq.of(_1791_v123);
              _1792_v124 = (_1792_v124).update(_dafny.Seq.Concat(_dafny.Seq.of(_1775_v109), (_1793_v125)[_module.__default.safeIndex((_this).f25, new BigNumber((_1793_v125).length))]), ((((_1771_v105).contains(p1)) ? ((_1771_v105).get(p1)) : (new BigNumber(-199)))).plus(_1775_v109));
              let _1794_v126;
              _1794_v126 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), ((_1795_v0) => function (_1796_i13) {
                return _1795_v0;
              })(_1619_v0))).length));
              let _1797_v127;
              _1797_v127 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1794_v126).cardinality()),p1);
              let _1798_v128;
              _1798_v128 = _module.D13.create_DC36(_1774_v108, p1, new BigNumber((_1797_v127).length));
              let _1799_v129;
              let _nw282 = Array((new BigNumber(24)).toNumber());
              _nw282[(_dafny.ZERO).toNumber()] = true;
              _nw282[(_dafny.ONE).toNumber()] = p1;
              _nw282[(new BigNumber(2)).toNumber()] = p1;
              _nw282[(new BigNumber(3)).toNumber()] = p1;
              _nw282[(new BigNumber(4)).toNumber()] = true;
              _nw282[(new BigNumber(5)).toNumber()] = p1;
              _nw282[(new BigNumber(6)).toNumber()] = p1;
              _nw282[(new BigNumber(7)).toNumber()] = p1;
              _nw282[(new BigNumber(8)).toNumber()] = p1;
              _nw282[(new BigNumber(9)).toNumber()] = true;
              _nw282[(new BigNumber(10)).toNumber()] = p1;
              _nw282[(new BigNumber(11)).toNumber()] = p1;
              _nw282[(new BigNumber(12)).toNumber()] = p1;
              _nw282[(new BigNumber(13)).toNumber()] = false;
              _nw282[(new BigNumber(14)).toNumber()] = !(p1);
              _nw282[(new BigNumber(15)).toNumber()] = p1;
              _nw282[(new BigNumber(16)).toNumber()] = p1;
              _nw282[(new BigNumber(17)).toNumber()] = p1;
              _nw282[(new BigNumber(18)).toNumber()] = p1;
              _nw282[(new BigNumber(19)).toNumber()] = p1;
              _nw282[(new BigNumber(20)).toNumber()] = p1;
              _nw282[(new BigNumber(21)).toNumber()] = (_1798_v128).dtor_cf55;
              _nw282[(new BigNumber(22)).toNumber()] = p1;
              _nw282[(new BigNumber(23)).toNumber()] = p1;
              _1799_v129 = _nw282;
              let _1800_v130;
              _1800_v130 = _dafny.MultiSet.fromElements(_1799_v129);
              (globalState).f12 = (_1800_v130).IsSubsetOf(_1800_v130);
            } else {
              let _1801_v131;
              _1801_v131 = _module.D0.create_DC1(false, p1);
              let _pat_let_tv34 = globalState;
              let _1802_v132;
              let _nw283 = new _module.C3();
              _nw283.__ctor(new BigNumber(398), (((function (_pat_let36_0) {
                return function (_1803_dt__update__tmp_h4) {
                  return function (_pat_let37_0) {
                    return function (_1804_dt__update_hcf2_h0) {
                      return _module.D0.create_DC1((_1803_dt__update__tmp_h4).dtor_cf1, _1804_dt__update_hcf2_h0);
                    }(_pat_let37_0);
                  }(_module.__default.fm0((_this).f25, _pat_let_tv34));
                }(_pat_let36_0);
              }(_1801_v131)).dtor_cf2) ? (p1) : (p1)), ((p1) ? ((_this).f25) : ((_this).f25)), (_this).f26);
              _1802_v132 = _nw283;
              _1775_v109 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_1775_v109, globalState),p1)).length)).multipliedBy(new BigNumber(932));
              let _1805_v133;
              let _nw284 = Array((new BigNumber(13)).toNumber()).fill(false);
              _1805_v133 = _nw284;
              let _1806_v134;
              _1806_v134 = _dafny.Set.fromElements(!((_1802_v132).f38));
              let _1807_v135;
              _1807_v135 = _dafny.Seq.of(_1806_v134);
              let _1808_v136;
              _1808_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1805_v133,(_1807_v135)[_module.__default.safeIndex((_this).f25, new BigNumber((_1807_v135).length))]);
              let _1809_v137;
              _1809_v137 = _dafny.Seq.of(_1775_v109);
              _1808_v136 = ((_dafny.Seq.IsPrefixOf(_1809_v137, _1809_v137)) ? (_dafny.Map.Empty.slice().updateUnsafe(_1805_v133,_1806_v134)) : (_1808_v136));
              r1 = false;
              (globalState).f12 = false;
            }
            _1775_v109 = (_dafny.ZERO).minus((_this).f25);
            let _1810_v138;
            _1810_v138 = _module.D5.create_DC16(p1);
            let _1811_v139;
            _1811_v139 = _dafny.Set.fromElements(_1810_v138);
            let _1812_v140;
            _1812_v140 = _dafny.Map.Empty.slice().updateUnsafe((_1811_v139).Difference(_1811_v139),(_this).f25);
            let _1813_v141;
            _1813_v141 = _dafny.Seq.of(p1);
            _1812_v140 = (_1812_v140).update((_1811_v139).Union(_1811_v139), new BigNumber((_1813_v141).length));
          }
        }
      }
      let _1814_v142;
      let _init64 = ((_1815_p1) => function (_1816_i14) {
        return _1815_p1;
      })(p1);
      let _nw285 = Array((new BigNumber(6)).toNumber());
      for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw285.length); _i0_64++) {
        _nw285[_i0_64] = _init64(new BigNumber(_i0_64));
      }
      _1814_v142 = _nw285;
      let _1817_v143;
      _1817_v143 = _module.D3.create_DC10(p1, _1814_v142, p1, new BigNumber(-294), (_this).f25);
      let _1818_v144;
      _1818_v144 = _dafny.Seq.of(p1, p1, p1);
      let _pat_let_tv35 = _1818_v144;
      let _pat_let_tv36 = _1818_v144;
      let _pat_let_tv37 = _1814_v142;
      r0 = function (_pat_let38_0) {
        return function (_1819_dt__update__tmp_h5) {
          return function (_pat_let39_0) {
            return function (_1820_dt__update_hcf15_h0) {
              return function (_pat_let40_0) {
                return function (_1821_dt__update_hcf14_h0) {
                  return _module.D3.create_DC10((_1819_dt__update__tmp_h5).dtor_cf13, _1821_dt__update_hcf14_h0, _1820_dt__update_hcf15_h0, (_1819_dt__update__tmp_h5).dtor_cf16, (_1819_dt__update__tmp_h5).dtor_cf17);
                }(_pat_let40_0);
              }(_pat_let_tv37);
            }(_pat_let39_0);
          }((_pat_let_tv35)[_module.__default.safeIndex((_this).f25, new BigNumber((_pat_let_tv36).length))]);
        }(_pat_let38_0);
      }(_1817_v143);
      r1 = p1;
      return [r0, r1];
    }
    m12(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.MultiSet.Empty;
      let r2 = false;
      let _1822_v0;
      let _init65 = function (_1823_i0) {
        return ((_this).f25).isEqualTo((_this).f25);
      };
      let _nw286 = Array((new BigNumber(26)).toNumber());
      for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw286.length); _i0_65++) {
        _nw286[_i0_65] = _init65(new BigNumber(_i0_65));
      }
      _1822_v0 = _nw286;
      let _1824_v1;
      _1824_v1 = true;
      let _index323 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length));
      (_1822_v0)[_index323] = _1824_v1;
      let _index324 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length));
      (_1822_v0)[_index324] = !(!(true) || (_1824_v1));
      let _1825_i1;
      _1825_i1 = _dafny.ZERO;
      L13: {
        while (((_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))]) === ((_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))])) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1825_i1)) {
              break L13;
            }
            _1825_i1 = (_1825_i1).plus(_dafny.ONE);
            let _1826_v2;
            let _init66 = function (_1827_i2) {
              return _module.__default.safeDivisionInt(_1827_i2, (_this).f25);
            };
            let _nw287 = Array((new BigNumber(2)).toNumber());
            for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw287.length); _i0_66++) {
              _nw287[_i0_66] = _init66(new BigNumber(_i0_66));
            }
            _1826_v2 = _nw287;
            let _index325 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            (_1826_v2)[_index325] = (_this).f25;
            let _index326 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            (_1826_v2)[_index326] = (new BigNumber(567)).multipliedBy((_this).f25);
            let _1828_v3;
            _1828_v3 = _dafny.Map.Empty.slice().updateUnsafe((_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))],(_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))]);
            let _1829_v4;
            _1829_v4 = _dafny.Seq.of(!((((_1828_v3).contains(_1824_v1)) ? ((_1828_v3).get(_1824_v1)) : (!(_1824_v1)))));
            let _index327 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            let _index328 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            let _index329 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length));
            let _index330 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            let _rhs318 = (_1826_v2)[_module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length))];
            let _rhs319 = (_this).f25;
            let _rhs320 = _1822_v0;
            let _rhs321 = ((_1829_v4)[_module.__default.safeIndex((_this).f25, new BigNumber((_1829_v4).length))]) && (((_1826_v2)[_module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length))]).isLessThan((_1826_v2)[_module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length))]));
            let _rhs322 = (_dafny.ZERO).minus((((_dafny.ZERO).minus(new BigNumber(-32))).multipliedBy((_1826_v2)[_module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length))])).multipliedBy((_this).f25));
            let _lhs231 = _1826_v2;
            let _lhs232 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            let _lhs233 = _1826_v2;
            let _lhs234 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            let _lhs235 = _1822_v0;
            let _lhs236 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length));
            let _lhs237 = _1826_v2;
            let _lhs238 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1826_v2).length));
            _lhs231[_lhs232] = _rhs318;
            _lhs233[_lhs234] = _rhs319;
            _1822_v0 = _rhs320;
            _lhs235[_lhs236] = _rhs321;
            _lhs237[_lhs238] = _rhs322;
            let _1830_v5;
            _1830_v5 = new _dafny.CodePoint('o'.codePointAt(0));
            let _1831_v6;
            _1831_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(_1824_v1),_1830_v5);
            (globalState).f23 = (((_1831_v6).contains(_1824_v1)) ? ((_1831_v6).get(_1824_v1)) : (_1830_v5));
            let _1832_v7;
            _1832_v7 = _dafny.Seq.UnicodeFromString("qbnx");
            let _1833_v8;
            _1833_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1832_v7,(_this).f25);
            let _rhs323 = (_1833_v8).Merge(_1833_v8);
            let _rhs324 = _dafny.Seq.IsPrefixOf(_1832_v7, _1832_v7);
            let _rhs325 = (_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))];
            let _lhs239 = globalState;
            let _lhs240 = globalState;
            _1833_v8 = _rhs323;
            _lhs239.f12 = _rhs324;
            _lhs240.f12 = _rhs325;
          }
        }
      }
      let _1834_i3;
      _1834_i3 = _dafny.ZERO;
      L14: {
        while ((_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))]) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1834_i3)) {
              break L14;
            }
            _1834_i3 = (_1834_i3).plus(_dafny.ONE);
            let _1835_v9;
            _1835_v9 = new BigNumber(782);
            _1835_v9 = _1835_v9;
            let _1836_v10;
            let _init67 = ((_1837_v9) => function (_1838_i4) {
              return _module.__default.safeDivisionInt(_1838_i4, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_1837_v9), _1837_v9)).length));
            })(_1835_v9);
            let _nw288 = Array((new BigNumber(28)).toNumber());
            for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw288.length); _i0_67++) {
              _nw288[_i0_67] = _init67(new BigNumber(_i0_67));
            }
            _1836_v10 = _nw288;
            let _index331 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1836_v10).length));
            (_1836_v10)[_index331] = (_this).f25;
            let _1839_v11;
            _1839_v11 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus((_this).f25)).multipliedBy(_module.__default.fm31(_dafny.Seq.UnicodeFromString("tdedfs"), globalState)),(_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))]);
            let _index332 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1836_v10).length));
            (_1836_v10)[_index332] = (_dafny.ZERO).minus(new BigNumber((_1839_v11).length));
            let _1840_v12;
            let _nw289 = new _module.C0();
            _nw289.__ctor();
            _1840_v12 = _nw289;
            let _1841_v13;
            let _nw290 = new _module.C1();
            _nw290.__ctor(((_this).f25).plus(new BigNumber((_1839_v11).length)), (_this).f26);
            _1841_v13 = _nw290;
          }
        }
      }
      let _1842_v14;
      _1842_v14 = _dafny.Seq.of((_this).f25);
      let _1843_v15;
      _1843_v15 = _dafny.MultiSet.fromElements(_1824_v1, _module.__default.fm0(_module.__default.fm31(_dafny.Seq.Create(_module.__default.abs(new BigNumber(117)), function (_1844_i6) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }), globalState), globalState), (_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))], false, (_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))]);
      let _1845_v16;
      let _nw291 = Array((new BigNumber(8)).toNumber());
      _nw291[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw291[(_dafny.ONE).toNumber()] = (new BigNumber((_1842_v14).length)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("edjrpfsui")).length));
      _nw291[(new BigNumber(2)).toNumber()] = (_this).f25;
      _nw291[(new BigNumber(3)).toNumber()] = (new BigNumber((_1842_v14).length)).multipliedBy(new BigNumber((_dafny.Seq.update(_1842_v14, _module.__default.safeIndex((_this).f25, new BigNumber((_1842_v14).length)), (_this).f25)).length));
      _nw291[(new BigNumber(4)).toNumber()] = (_this).f25;
      _nw291[(new BigNumber(5)).toNumber()] = (_this).f25;
      _nw291[(new BigNumber(6)).toNumber()] = (((_1843_v15).contains(false)) ? ((_1843_v15).get(false)) : ((_this).f25));
      _nw291[(new BigNumber(7)).toNumber()] = (_this).f25;
      _1845_v16 = _nw291;
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1845_v16).length))) {
        let _1846_i5 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1846_i5)) && ((_1846_i5).isLessThan(new BigNumber((_1845_v16).length))))) {
          (_1845_v16)[(_1846_i5)] = _module.__default.safeDivisionInt(_1846_i5, _module.__default.safeModuloInt((_this).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(475)), ((_1847_v1) => function (_1848_i7) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1847_v1,(_this).f25);
          })(_1824_v1))).length)));
        }
      }
      let _1849_v17;
      _1849_v17 = new _dafny.CodePoint('d'.codePointAt(0));
      let _1850_v18;
      _1850_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1849_v17,(_this).f25);
      let _1851_v19;
      _1851_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1850_v18,_1822_v0);
      _1851_v19 = ((_1851_v19).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1850_v18,_1822_v0))).Merge(_1851_v19);
      let _1852_v20;
      _1852_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))],_1822_v0);
      let _1853_v21;
      let _1854_v22;
      let _out31;
      let _out32;
      let _outcollector11 = (_this).m11(_1845_v16, _1824_v1, _1852_v20, globalState);
      _out31 = _outcollector11[0];
      _out32 = _outcollector11[1];
      _1853_v21 = _out31;
      _1854_v22 = _out32;
      r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(690)), ((_1855_v17, _1856_v0) => function (_1857_i8) {
        return (_module.D1.create_DC6(_1855_v17, (_1856_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1856_v0).length))], _dafny.Seq.UnicodeFromString("wubujscf"))).dtor_cf7;
      })(_1849_v17, _1822_v0));
      let _1858_v23;
      _1858_v23 = _dafny.Seq.of((_1822_v0)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_1822_v0).length))]);
      let _1859_v24;
      _1859_v24 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_1854_v22), _1858_v23);
      let _1860_v25;
      _1860_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1854_v22,_1858_v23);
      let _1861_v26;
      _1861_v26 = _dafny.MultiSet.fromElements((((_1859_v24).contains((((_1860_v25).contains(_1854_v22)) ? ((_1860_v25).get(_1854_v22)) : (_1858_v23)))) ? ((_1859_v24).get((((_1860_v25).contains(_1854_v22)) ? ((_1860_v25).get(_1854_v22)) : (_1858_v23)))) : ((_this).f25)));
      r1 = _1861_v26;
      let _1862_v27;
      _1862_v27 = _dafny.Seq.UnicodeFromString("y");
      let _1863_v28;
      _1863_v28 = _module.D1.create_DC6(_1849_v17, false, _1862_v27);
      r2 = (_1863_v28).dtor_cf8;
      return [r0, r1, r2];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f34 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f34) {
      let _this = this;
      (_this)._f34 = f34;
      return;
    }
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_this).f34).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qrq"), _dafny.Seq.UnicodeFromString("gdlw"))).length));
    };
    m10(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1864_v0;
      let _nw292 = new _module.C2();
      _nw292.__ctor();
      _1864_v0 = _nw292;
      let _1865_v1;
      let _nw293 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _1865_v1 = _nw293;
      let _1866_v2;
      _1866_v2 = false;
      let _1867_v3;
      _1867_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,!(!(_1866_v2)));
      let _index333 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length));
      (_1865_v1)[_index333] = (((((_1867_v3).contains((_this).f34)) ? ((_1867_v3).get((_this).f34)) : (true))) ? ((_this).f34) : (p1));
      let _index334 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length));
      let _rhs326 = _1866_v2;
      let _rhs327 = (_this).f34;
      let _rhs328 = _1865_v1;
      let _rhs329 = p1;
      let _lhs241 = globalState;
      let _lhs242 = _1865_v1;
      let _lhs243 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length));
      _lhs241.f12 = _rhs326;
      r0 = _rhs327;
      _1865_v1 = _rhs328;
      _lhs242[_lhs243] = _rhs329;
      let _1868_v4;
      _1868_v4 = _dafny.Map.Empty.slice().updateUnsafe((_1866_v2) && (false),new BigNumber(522));
      let _1869_v5;
      _1869_v5 = _dafny.Set.fromElements(_1866_v2);
      let _1870_v6;
      _1870_v6 = _dafny.MultiSet.fromElements(_1869_v5, _1869_v5);
      r0 = (((_1868_v4).contains((_1870_v6).IsProperSubsetOf(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_1866_v2, _1866_v2), _1869_v5, _1869_v5, _1869_v5)))) ? ((_1868_v4).get((_1870_v6).IsProperSubsetOf(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_1866_v2, _1866_v2), _1869_v5, _1869_v5, _1869_v5)))) : ((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))]));
      (globalState).f12 = _1866_v2;
      if (_1866_v2) {
        (globalState).f12 = _module.__default.fm0((_this).f34, globalState);
        let _index335 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length));
        (_1865_v1)[_index335] = p1;
        let _1871_v7;
        _1871_v7 = _dafny.Seq.of(_1866_v2);
        let _1872_v8;
        _1872_v8 = _dafny.Seq.of((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))], (_this).f34, ((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))]).plus(new BigNumber((_1871_v7).length)));
        r0 = (_dafny.ZERO).minus((_1872_v8)[_module.__default.safeIndex((_this).f34, new BigNumber((_1872_v8).length))]);
        r0 = (_this).f34;
        let _1873_v9;
        _1873_v9 = _module.D0.create_DC2(_1866_v2);
        let _1874_v10;
        let _nw294 = Array((new BigNumber(14)).toNumber());
        _nw294[(_dafny.ZERO).toNumber()] = _1866_v2;
        _nw294[(_dafny.ONE).toNumber()] = !(_1866_v2);
        _nw294[(new BigNumber(2)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(3)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(4)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(5)).toNumber()] = !(_1866_v2);
        _nw294[(new BigNumber(6)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(7)).toNumber()] = (_1873_v9).dtor_cf3;
        _nw294[(new BigNumber(8)).toNumber()] = true;
        _nw294[(new BigNumber(9)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(10)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(11)).toNumber()] = _1866_v2;
        _nw294[(new BigNumber(12)).toNumber()] = true;
        _nw294[(new BigNumber(13)).toNumber()] = _1866_v2;
        _1874_v10 = _nw294;
        let _1875_v11;
        _1875_v11 = _dafny.Map.Empty.slice().updateUnsafe(true,_1874_v10);
        _1875_v11 = (_1875_v11).Merge(_1875_v11);
      } else {
        let _1876_v12;
        _1876_v12 = _dafny.MultiSet.fromElements(_module.__default.fm0((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))], globalState));
        _1867_v3 = (_1867_v3).update((((_1876_v12).contains(_1866_v2)) ? ((_1876_v12).get(_1866_v2)) : (p1)), (_1866_v2) || (_1866_v2));
        let _index336 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length));
        (_1865_v1)[_index336] = (_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))];
        let _1877_v13;
        let _1878_v14;
        let _1879_v15;
        let _1880_v16;
        let _out33;
        let _out34;
        let _out35;
        let _out36;
        let _outcollector12 = (_1864_v0).m16(globalState);
        _out33 = _outcollector12[0];
        _out34 = _outcollector12[1];
        _out35 = _outcollector12[2];
        _out36 = _outcollector12[3];
        _1877_v13 = _out33;
        _1878_v14 = _out34;
        _1879_v15 = _out35;
        _1880_v16 = _out36;
        let _1881_v17;
        let _nw295 = new _module.C0();
        _nw295.__ctor();
        _1881_v17 = _nw295;
        _1877_v13 = (_this).f34;
      }
      let _1882_v18;
      let _nw296 = Array((new BigNumber(4)).toNumber());
      _nw296[(_dafny.ZERO).toNumber()] = ((_1866_v2) ? (_1865_v1) : (_1865_v1));
      _nw296[(_dafny.ONE).toNumber()] = _1865_v1;
      _nw296[(new BigNumber(2)).toNumber()] = _1865_v1;
      _nw296[(new BigNumber(3)).toNumber()] = _1865_v1;
      _1882_v18 = _nw296;
      let _index337 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1882_v18).length));
      (_1882_v18)[_index337] = _1865_v1;
      let _index338 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1882_v18).length));
      (_1882_v18)[_index338] = _1865_v1;
      let _1883_v19;
      _1883_v19 = _dafny.Seq.of(new BigNumber(77), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(903)), function (_1884_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)));
      let _1885_v20;
      _1885_v20 = _dafny.Set.fromElements(_1883_v19, _dafny.Seq.of((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))]));
      let _1886_v21;
      _1886_v21 = _module.D0.create_DC1(_1866_v2, _1866_v2);
      r0 = ((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))]).minus((_this).fm17((_1865_v1)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_1865_v1).length))], _1885_v20, _1886_v21, p0, globalState));
      return r0;
    }
    get f34() {
      let _this = this;
      return _this._f34;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
      this._f32 = false;
      this._f33 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f32, f33, f25, f26) {
      let _this = this;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm13(globalState) {
      let _this = this;
      let _source25 = _module.D0.create_DC3((_this).f25, (_this).f33);
      if (_source25.is_DC1) {
        let _1887___mcc_h0 = (_source25).cf1;
        let _1888___mcc_h1 = (_source25).cf2;
        let _1889_cf2 = _1888___mcc_h1;
        let _1890_cf1 = _1887___mcc_h0;
        return (_this).f32;
      } else if (_source25.is_DC2) {
        let _1891___mcc_h2 = (_source25).cf3;
        let _1892_cf3 = _1891___mcc_h2;
        return _1892_cf3;
      } else if (_source25.is_DC3) {
        let _1893___mcc_h3 = (_source25).cf4;
        let _1894___mcc_h4 = (_source25).cf5;
        let _1895_cf5 = _1894___mcc_h4;
        let _1896_cf4 = _1893___mcc_h3;
        return (_this).f32;
      } else {
        let _1897___mcc_h5 = (_source25).cf0;
        let _1898_cf0 = _1897___mcc_h5;
        return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("tgkclqti"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), function (_1899_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })));
      }
    };
    fm14(p0, globalState) {
      let _this = this;
      return (_module.D0.create_DC3((_this).f25, false)).dtor_cf4;
    };
    m4(p0, globalState) {
      let _this = this;
      let _hi9 = (_this).f25;
      for (let _1900_i0 = (_this).f25; _1900_i0.isLessThan(_hi9); _1900_i0 = _1900_i0.plus(_dafny.ONE)) {
        let _1901_v0;
        _1901_v0 = _dafny.Seq.of((_this).f33);
        let _1902_v1;
        _1902_v1 = _dafny.MultiSet.fromElements((_this).f32);
        let _1903_v2;
        _1903_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1900_i0,new BigNumber((_1902_v1).cardinality()));
        let _1904_v3;
        let _init68 = function (_1905_i1) {
          return _module.D1.create_DC7(_module.D1.create_DC6(new _dafny.CodePoint('t'.codePointAt(0)), (_this).f32, _dafny.Seq.UnicodeFromString("vwpkpvo")));
        };
        let _nw297 = Array((new BigNumber(2)).toNumber());
        for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw297.length); _i0_68++) {
          _nw297[_i0_68] = _init68(new BigNumber(_i0_68));
        }
        _1904_v3 = _nw297;
        let _1906_v4;
        let _nw298 = Array((new BigNumber(27)).toNumber());
        _nw298[(_dafny.ZERO).toNumber()] = new BigNumber((_1901_v0).length);
        _nw298[(_dafny.ONE).toNumber()] = new BigNumber((_module.__default.fm1((_this).f33, globalState)).length);
        _nw298[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_this).f25);
        _nw298[(new BigNumber(3)).toNumber()] = new BigNumber(866);
        _nw298[(new BigNumber(4)).toNumber()] = _1900_i0;
        _nw298[(new BigNumber(5)).toNumber()] = _1900_i0;
        _nw298[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_1900_i0);
        _nw298[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("gcib")).length);
        _nw298[(new BigNumber(8)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(9)).toNumber()] = (((_1903_v2).contains(_1900_i0)) ? ((_1903_v2).get(_1900_i0)) : ((_this).f25));
        _nw298[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(11)).toNumber()] = new BigNumber(-58);
        _nw298[(new BigNumber(12)).toNumber()] = (new BigNumber(882)).minus((_this).fm14((_this).f33, globalState));
        _nw298[(new BigNumber(13)).toNumber()] = _1900_i0;
        _nw298[(new BigNumber(14)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(15)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(16)).toNumber()] = (_this).fm14(_module.__default.fm0(new BigNumber((_1901_v0).length), globalState), globalState);
        _nw298[(new BigNumber(17)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(18)).toNumber()] = _1900_i0;
        _nw298[(new BigNumber(19)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1904_v3,(_this).f25)).length);
        _nw298[(new BigNumber(21)).toNumber()] = (_1900_i0).multipliedBy((_this).f25);
        _nw298[(new BigNumber(22)).toNumber()] = new BigNumber(575);
        _nw298[(new BigNumber(23)).toNumber()] = _1900_i0;
        _nw298[(new BigNumber(24)).toNumber()] = (_this).f25;
        _nw298[(new BigNumber(25)).toNumber()] = new BigNumber(276);
        _nw298[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f32, true, (_this).f32, (_this).f32, (_this).f33)).length);
        _1906_v4 = _nw298;
        let _index339 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        (_1906_v4)[_index339] = new BigNumber(424);
        let _1907_v5;
        _1907_v5 = _dafny.Seq.UnicodeFromString("vsr");
        let _index340 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        (_1906_v4)[_index340] = new BigNumber((_1907_v5).length);
        let _1908_v6;
        _1908_v6 = _dafny.Seq.of((_1906_v4)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length))]);
        let _1909_v7;
        _1909_v7 = _dafny.MultiSet.fromElements(new BigNumber((_1908_v6).length));
        (globalState).f12 = (_1909_v7).IsDisjointFrom(_module.__default.fm15(globalState));
        let _index341 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        let _rhs330 = (_dafny.ZERO).minus((new BigNumber(391)).minus((new BigNumber((_1908_v6).length)).minus((_1906_v4)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length))])));
        let _rhs331 = (_this).f32;
        let _lhs244 = _1906_v4;
        let _lhs245 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        let _lhs246 = globalState;
        _lhs244[_lhs245] = _rhs330;
        _lhs246.f12 = _rhs331;
        let _1910_v8;
        _1910_v8 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1907_v5).length)));
        let _1911_v9;
        _1911_v9 = _module.D0.create_DC2(!((_this).f33));
        let _1912_v10;
        _1912_v10 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1913_v11;
        _1913_v11 = _module.D1.create_DC6(_1912_v10, (_this).f33, _1907_v5);
        let _1914_v12;
        _1914_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1912_v10,(_1906_v4)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length))]);
        let _1915_v13;
        _1915_v13 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), ((_1916_i0) => function (_1917_i2) {
          return _1916_i0;
        })(_1900_i0)));
        let _1918_v14;
        _1918_v14 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(_1907_v5, _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1907_v5).length)), _1912_v10), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_1906_v4)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length))])).length), new BigNumber((_dafny.Seq.update(_1907_v5, _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1907_v5).length)), _1912_v10)).length)), _1912_v10));
        let _pat_let_tv38 = _1907_v5;
        let _index342 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        let _index343 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        let _rhs332 = _module.__default.fm16(_1911_v9, (_this).f32, (_this).f32, (function (_pat_let41_0) {
          return function (_1919_dt__update__tmp_h0) {
            return function (_pat_let42_0) {
              return function (_1920_dt__update_hcf9_h0) {
                return _module.D1.create_DC6((_1919_dt__update__tmp_h0).dtor_cf7, (_1919_dt__update__tmp_h0).dtor_cf8, _1920_dt__update_hcf9_h0);
              }(_pat_let42_0);
            }(_pat_let_tv38);
          }(_pat_let41_0);
        }(_1913_v11)).dtor_cf7, globalState);
        let _rhs333 = _dafny.Seq.contains(_1901_v0, !(!(_1900_i0).isEqualTo((((_1914_v12).contains(_1912_v10)) ? ((_1914_v12).get(_1912_v10)) : (new BigNumber(((_1915_v13)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_1915_v13).length))]).length))))));
        let _rhs334 = _dafny.Seq.of((_this).f32, !(!((_this).f32) || ((_this).f33)), (_this).f33, (_this).f32);
        let _rhs335 = _module.__default.safeDivisionInt(new BigNumber(31), (_dafny.ZERO).minus((_dafny.ZERO).minus((_1900_i0).minus((_this).f25))));
        let _rhs336 = new BigNumber(((_1918_v14)[_module.__default.safeIndex((new BigNumber((_1907_v5).length)).minus(new BigNumber((_1907_v5).length)), new BigNumber((_1918_v14).length))]).length);
        let _lhs247 = globalState;
        let _lhs248 = _1906_v4;
        let _lhs249 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        let _lhs250 = _1906_v4;
        let _lhs251 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1906_v4).length));
        _1910_v8 = _rhs332;
        _lhs247.f12 = _rhs333;
        _1901_v0 = _rhs334;
        _lhs248[_lhs249] = _rhs335;
        _lhs250[_lhs251] = _rhs336;
      }
      let _1921_v15;
      _1921_v15 = _dafny.Set.fromElements((_this).f25, (_this).f25);
      let _hi10 = (_this).f25;
      for (let _1922_i3 = (new BigNumber((_1921_v15).length)).minus(new BigNumber(609)); _1922_i3.isLessThan(_hi10); _1922_i3 = _1922_i3.plus(_dafny.ONE)) {
        let _1923_v16;
        let _nw299 = new _module.C5();
        _nw299.__ctor((_this).f32);
        _1923_v16 = _nw299;
        let _1924_v17;
        _1924_v17 = _dafny.Set.fromElements((_this).f32);
        let _1925_v18;
        _1925_v18 = _dafny.Seq.UnicodeFromString("yh");
        let _1926_v19;
        _1926_v19 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f25).minus(new BigNumber((_1924_v17).length)),new BigNumber((_1925_v18).length));
        _1926_v19 = (_1926_v19).update(_1922_i3, (_this).fm14((_this).f32, globalState));
        let _1927_v20;
        _1927_v20 = new BigNumber(308);
        let _1928_v21;
        _1928_v21 = _dafny.MultiSet.fromElements(!(true));
        _1927_v20 = _module.__default.fm44((_1928_v21).IsProperSubsetOf(_1928_v21), (_this).f33, (_this).fm13(globalState), new BigNumber(-445), globalState);
        let _1929_v22;
        let _nw300 = Array((new BigNumber(2)).toNumber());
        _nw300[(_dafny.ZERO).toNumber()] = _1922_i3;
        _nw300[(_dafny.ONE).toNumber()] = (_this).f25;
        _1929_v22 = _nw300;
        _1929_v22 = _1929_v22;
      }
      if (((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(396), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(157)), function (_1930_i4) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length)))).isLessThanOrEqualTo(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f25), (_this).f25))) {
        let _1931_v23;
        _1931_v23 = new BigNumber(56);
        _1931_v23 = (_this).fm14((_this).f33, globalState);
        let _1932_v24;
        _1932_v24 = _dafny.Seq.of(false);
        let _1933_v25;
        _1933_v25 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1932_v24).length)).multipliedBy(new BigNumber(753)),_1931_v23);
        let _1934_v26;
        _1934_v26 = _dafny.MultiSet.fromElements((_this).fm14(true, globalState));
        let _1935_v27;
        _1935_v27 = _dafny.MultiSet.fromElements(_1934_v26);
        _1933_v25 = (_1933_v25).update(_1931_v23, new BigNumber((_module.__default.fm37(_1935_v27, globalState)).length));
        let _1936_v28;
        _1936_v28 = _dafny.Seq.of(_1931_v23, _1931_v23, (_this).f25, (_this).f25, new BigNumber((_1932_v24).length));
        (globalState).f12 = _dafny.Seq.contains(_1936_v28, (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25)));
        let _1937_v29;
        let _nw301 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1937_v29 = _nw301;
        let _1938_v30;
        _1938_v30 = _module.D4.create_DC13((_this).f25, (_this).f33, (_this).f33, (_this).f33);
        let _1939_v31;
        _1939_v31 = _dafny.Map.Empty.slice().updateUnsafe((_1938_v30).dtor_cf20,!((_this).f33));
        let _1940_v32;
        _1940_v32 = _1939_v31;
        let _1941_v33;
        _1941_v33 = _dafny.Seq.of(_1940_v32, _1940_v32);
        let _index344 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length));
        (_1937_v29)[_index344] = new BigNumber((_1941_v33).length);
        let _index345 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length));
        (_1937_v29)[_index345] = _1931_v23;
        if ((function () {
          let _coll70 = new _dafny.Set();
          for (const _compr_70 of (_dafny.Set.fromElements((((_1933_v25).contains((_dafny.ZERO).minus((_this).f25))) ? ((_1933_v25).get((_dafny.ZERO).minus((_this).f25))) : (_1931_v23)))).Elements) {
            let _1942_v34 = _compr_70;
            if ((_dafny.Set.fromElements((((_1933_v25).contains((_dafny.ZERO).minus((_this).f25))) ? ((_1933_v25).get((_dafny.ZERO).minus((_this).f25))) : (_1931_v23)))).contains(_1942_v34)) {
              _coll70.add((_1942_v34).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(391), new BigNumber(821))).length)));
            }
          }
          return _coll70;
        }()).IsSubsetOf(_1921_v15)) {
          _1931_v23 = (_1937_v29)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length))];
          let _index346 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length));
          (_1937_v29)[_index346] = _module.__default.safeModuloInt(new BigNumber(302), (_1937_v29)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length))]);
          let _1943_v35;
          _1943_v35 = _module.D5.create_DC15(_1934_v26);
          let _1944_v36;
          let _nw302 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1944_v36 = _nw302;
          let _index347 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1944_v36).length));
          (_1944_v36)[_index347] = (_this).f33;
          let _1945_v37;
          _1945_v37 = _dafny.Seq.UnicodeFromString("bsjsoh");
          let _index348 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1944_v36).length));
          let _rhs337 = _1943_v35;
          let _rhs338 = _1945_v37;
          let _rhs339 = (_1931_v23).isEqualTo((_1937_v29)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length))]);
          let _rhs340 = !(!((_this).f33));
          let _lhs252 = globalState;
          let _lhs253 = globalState;
          let _lhs254 = _1944_v36;
          let _lhs255 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1944_v36).length));
          _1943_v35 = _rhs337;
          _lhs252.f10 = _rhs338;
          _lhs253.f12 = _rhs339;
          _lhs254[_lhs255] = _rhs340;
          (globalState).f12 = _module.__default.fm0(((_this).f25).minus((_this).f25), globalState);
          let _1946_v38;
          let _nw303 = new _module.C2();
          _nw303.__ctor();
          _1946_v38 = _nw303;
        } else {
          let _1947_v39;
          _1947_v39 = _dafny.MultiSet.fromElements((_this).f32);
          _1931_v23 = ((_dafny.ZERO).minus(new BigNumber((_1947_v39).cardinality()))).minus(_1931_v23);
          _1931_v23 = (_1937_v29)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length))];
          let _1948_v40;
          let _nw304 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _1948_v40 = _nw304;
          let _1949_v41;
          _1949_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_1948_v40);
          let _1950_v42;
          _1950_v42 = _dafny.Seq.of((((_1949_v41).contains((_this).f33)) ? ((_1949_v41).get((_this).f33)) : (_1948_v40)), _1948_v40, _1948_v40);
          _1948_v40 = (_1950_v42)[_module.__default.safeIndex(((_1937_v29)[_module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length))]).plus(new BigNumber((_1932_v24).length)), new BigNumber((_1950_v42).length))];
          let _index349 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1937_v29).length));
          (_1937_v29)[_index349] = new BigNumber(744);
          _1937_v29 = _1937_v29;
        }
      } else {
        let _1951_v43;
        _1951_v43 = _module.D9.create_DC27((_this).f25, _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), function (_1952_i5) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length), (_this).f25));
        let _source26 = _1951_v43;
        if (_source26.is_DC27) {
          let _1953___mcc_h0 = (_source26).cf39;
          let _1954___mcc_h1 = (_source26).cf40;
          let _1955_cf40 = _1954___mcc_h1;
          let _1956_cf39 = _1953___mcc_h0;
          let _1957_v44;
          _1957_v44 = _dafny.Seq.UnicodeFromString("jgikbaco");
          _1956_cf39 = (((true) ? (_1955_cf40) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f33),(_this).f33)).length))))).plus((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-569)), function (_1958_i6) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1957_v44).length))));
          let _1959_v45;
          let _init69 = function (_1960_i7) {
            return (new BigNumber(492)).isLessThanOrEqualTo(new BigNumber(215));
          };
          let _nw305 = Array((new BigNumber(23)).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw305.length); _i0_69++) {
            _nw305[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _1959_v45 = _nw305;
          let _index350 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1959_v45).length));
          (_1959_v45)[_index350] = (_this).f33;
          let _1961_v46;
          let _nw306 = Array((new BigNumber(19)).toNumber()).fill([]);
          _1961_v46 = _nw306;
          let _1962_v47;
          let _init70 = ((_1963_cf39) => function (_1964_i8) {
            return _module.__default.safeDivisionInt(_1964_i8, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1963_cf39,_dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_this).f33))).length));
          })(_1956_cf39);
          let _nw307 = Array((new BigNumber(7)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw307.length); _i0_70++) {
            _nw307[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _1962_v47 = _nw307;
          let _index351 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_1961_v46).length));
          (_1961_v46)[_index351] = _1962_v47;
          let _1965_v48;
          _1965_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1955_cf40,_1956_cf39);
          let _1966_v49;
          _1966_v49 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1967_v50;
          _1967_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1966_v49,_1962_v47);
          let _1968_v51;
          _1968_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1956_cf39,false);
          let _1969_v52;
          _1969_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,!((((_1968_v51).contains((_this).f25)) ? ((_1968_v51).get((_this).f25)) : ((_this).f32))) || ((_this).f33));
          let _1970_v53;
          _1970_v53 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f32);
          let _1971_v54;
          _1971_v54 = _dafny.Seq.of((_this).f33, (_this).f32);
          let _1972_v55;
          _1972_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1971_v54, _1971_v54)).length),_1962_v47);
          let _index352 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1959_v45).length));
          let _index353 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_1961_v46).length));
          let _rhs341 = ((_this).f25).isLessThanOrEqualTo(new BigNumber((_1967_v50).length));
          let _rhs342 = (((_1968_v51).contains(_module.__default.safeModuloInt(new BigNumber((_1970_v53).length), _1955_cf40))) ? ((_1968_v51).get(_module.__default.safeModuloInt(new BigNumber((_1970_v53).length), _1955_cf40))) : (((_this).f32) || (!((_this).f32))));
          let _rhs343 = _module.__default.safeModuloInt(new BigNumber((_1971_v54).length), new BigNumber(-823));
          let _rhs344 = (((_1972_v55).contains(new BigNumber((_1965_v48).length))) ? ((_1972_v55).get(new BigNumber((_1965_v48).length))) : (_1962_v47));
          let _rhs345 = (_1965_v48).Merge(function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of _dafny.IntegerRange(new BigNumber(104), new BigNumber(-568))) {
              let _1973_v56 = _compr_71;
              if (((new BigNumber(104)).isLessThanOrEqualTo(_1973_v56)) && ((_1973_v56).isLessThan(new BigNumber(-568)))) {
                _coll71.push([_module.__default.safeModuloInt(_1973_v56, (_this).f25),new BigNumber(-573)]);
              }
            }
            return _coll71;
          }());
          let _lhs256 = globalState;
          let _lhs257 = _1959_v45;
          let _lhs258 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1959_v45).length));
          let _lhs259 = _1961_v46;
          let _lhs260 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_1961_v46).length));
          _lhs256.f12 = _rhs341;
          _lhs257[_lhs258] = _rhs342;
          _1956_cf39 = _rhs343;
          _lhs259[_lhs260] = _rhs344;
          _1965_v48 = _rhs345;
          _1965_v48 = (_1965_v48).update(_1955_cf40, _1955_cf40);
          _1955_cf40 = new BigNumber(822);
        } else {
          let _1974___mcc_h2 = (_source26).cf38;
          let _1975_cf38 = _1974___mcc_h2;
          let _1976_v57;
          let _nw308 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1976_v57 = _nw308;
          let _index354 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1976_v57).length));
          (_1976_v57)[_index354] = true;
          let _index355 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_1976_v57).length));
          (_1976_v57)[_index355] = !(!((_this).f32));
          let _1977_v58;
          let _nw309 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
          _1977_v58 = _nw309;
          let _index356 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1977_v58).length));
          (_1977_v58)[_index356] = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f33);
          let _1978_v59;
          _1978_v59 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f25),(_1976_v57)[_module.__default.safeIndex(new BigNumber(113), new BigNumber((_1976_v57).length))]);
          let _index357 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1977_v58).length));
          (_1977_v58)[_index357] = (_1978_v59).Merge(_1978_v59);
          let _1979_v60;
          _1979_v60 = new BigNumber(238);
          _1979_v60 = (_this).f25;
          _1975_cf38 = (_1975_cf38).update((_this).fm13(globalState), (_this).f25);
        }
        let _1980_v61;
        _1980_v61 = new BigNumber(-292);
        let _1981_v62;
        _1981_v62 = _dafny.Seq.of((_this).f32, (_this).f32, (_this).f33);
        _1980_v61 = _module.__default.safeModuloInt((_this).f25, new BigNumber((_dafny.Seq.Concat(_1981_v62, _1981_v62)).length));
        let _1982_v63;
        _1982_v63 = new _dafny.CodePoint('w'.codePointAt(0));
        (globalState).f23 = _1982_v63;
        (globalState).f12 = (_this).f33;
        let _1983_v64;
        let _nw310 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1983_v64 = _nw310;
        _1983_v64 = _1983_v64;
      }
      let _1984_v65;
      _1984_v65 = _dafny.Seq.of((_this).f33);
      let _1985_v66;
      _1985_v66 = _dafny.MultiSet.fromElements((_this).f33, (_this).f33, !((_this).f33));
      let _1986_v69;
      _1986_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm64((_this).f25, (_this).f32, new _dafny.CodePoint('q'.codePointAt(0)), (_this).f32, globalState)).length),new BigNumber((function () {
        let _coll72 = new _dafny.Set();
        for (const _compr_72 of (_dafny.MultiSet.fromElements((_this).f25, (_this).f25)).Elements) {
          let _1987_v67 = _compr_72;
          if ((_dafny.MultiSet.fromElements((_this).f25, (_this).f25)).contains(_1987_v67)) {
            _coll72.add(_module.__default.safeModuloInt(_1987_v67, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
              let _coll73 = new _dafny.Map();
              for (const _compr_73 of _dafny.IntegerRange(new BigNumber(64), new BigNumber(-997))) {
                let _1988_v68 = _compr_73;
                if (((new BigNumber(64)).isLessThanOrEqualTo(_1988_v68)) && ((_1988_v68).isLessThan(new BigNumber(-997)))) {
                  _coll73.push([_module.__default.safeModuloInt(_1988_v68, new BigNumber(258)),new BigNumber(365)]);
                }
              }
              return _coll73;
            }()).length),new BigNumber((_dafny.Seq.UnicodeFromString("ksgs")).length))).length)));
          }
        }
        return _coll72;
      }()).length));
      let _1989_v70;
      _1989_v70 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.FromArray(_1984_v65)).IsDisjointFrom(_1985_v66),((_dafny.ZERO).minus((((_1986_v69).contains((_this).f25)) ? ((_1986_v69).get((_this).f25)) : ((_this).f25)))).plus((_this).f25));
      _1989_v70 = (_1989_v70).update((_this).f33, (_dafny.ZERO).minus((new BigNumber(-165)).plus((_this).f25)));
      if ((_this).f32) {
        let _1990_v71;
        _1990_v71 = _dafny.Set.fromElements((_this).f32);
        (globalState).f12 = (_1990_v71).IsSubsetOf(_1990_v71);
        let _1991_v72;
        _1991_v72 = new BigNumber(75);
        _1991_v72 = (((_this).f32) ? (_module.__default.safeDivisionInt(_1991_v72, new BigNumber(877))) : (_1991_v72));
        let _1992_v73;
        let _nw311 = new _module.C2();
        _nw311.__ctor();
        _1992_v73 = _nw311;
        _1990_v71 = (_dafny.Set.fromElements(_module.__default.fm0(_1991_v72, globalState))).Intersect(_1990_v71);
        let _1993_v74;
        let _init71 = ((_1994_v72) => function (_1995_i9) {
          return (_1995_i9).multipliedBy(_1994_v72);
        })(_1991_v72);
        let _nw312 = Array((new BigNumber(3)).toNumber());
        for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw312.length); _i0_71++) {
          _nw312[_i0_71] = _init71(new BigNumber(_i0_71));
        }
        _1993_v74 = _nw312;
        let _1996_v75;
        let _nw313 = Array((new BigNumber(16)).toNumber());
        _nw313[(_dafny.ZERO).toNumber()] = (_this).f33;
        _nw313[(_dafny.ONE).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(2)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(3)).toNumber()] = !((_this).f33);
        _nw313[(new BigNumber(4)).toNumber()] = (_this).f33;
        _nw313[(new BigNumber(5)).toNumber()] = (_this).f33;
        _nw313[(new BigNumber(6)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(7)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(8)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(9)).toNumber()] = (_this).f33;
        _nw313[(new BigNumber(10)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(11)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(12)).toNumber()] = true;
        _nw313[(new BigNumber(13)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(14)).toNumber()] = (_this).f32;
        _nw313[(new BigNumber(15)).toNumber()] = false;
        _1996_v75 = _nw313;
        let _1997_v76;
        let _nw314 = new _module.C5();
        _nw314.__ctor(false);
        _1997_v76 = _nw314;
        let _1998_v77;
        let _nw315 = Array((new BigNumber(19)).toNumber());
        _nw315[(_dafny.ZERO).toNumber()] = _1997_v76;
        _nw315[(_dafny.ONE).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(2)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(3)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(4)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(5)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(6)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(7)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(8)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(9)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(10)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(11)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(12)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(13)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(14)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(15)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(16)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(17)).toNumber()] = _1997_v76;
        _nw315[(new BigNumber(18)).toNumber()] = _1997_v76;
        _1998_v77 = _nw315;
        let _1999_v78;
        _1999_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v75,_1998_v77);
        let _index358 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_1993_v74).length));
        (_1993_v74)[_index358] = new BigNumber((_1999_v78).length);
        let _index359 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_1993_v74).length));
        (_1993_v74)[_index359] = _module.__default.safeModuloInt(_1991_v72, _1991_v72);
      } else {
        let _2000_v79;
        _2000_v79 = new BigNumber(-719);
        let _2001_v80;
        _2001_v80 = _module.D12.create_DC32((_this).f32);
        let _2002_v82;
        _2002_v82 = _dafny.Seq.of(_2000_v79, new BigNumber((function () {
          let _coll74 = new _dafny.Set();
          for (const _compr_74 of _dafny.IntegerRange(new BigNumber(448), new BigNumber(957))) {
            let _2003_v81 = _compr_74;
            if (((new BigNumber(448)).isLessThanOrEqualTo(_2003_v81)) && ((_2003_v81).isLessThan(new BigNumber(957)))) {
              _coll74.add(_module.__default.safeDivisionInt(_2003_v81, (_this).f25));
            }
          }
          return _coll74;
        }()).length), _2000_v79);
        let _2004_v83;
        _2004_v83 = _dafny.MultiSet.fromElements((_this).f25, new BigNumber((_2002_v82).length));
        let _rhs346 = (_2004_v83).equals(_2004_v83);
        let _rhs347 = _module.__default.safeDivisionInt((_this).f25, _2000_v79);
        let _rhs348 = _2001_v80;
        let _lhs261 = globalState;
        _lhs261.f12 = _rhs346;
        _2000_v79 = _rhs347;
        _2001_v80 = _rhs348;
        (globalState).f6 = _dafny.Seq.update(_2002_v82, _module.__default.safeIndex(new BigNumber(-234), new BigNumber((_2002_v82).length)), _2000_v79);
        let _2005_v84;
        _2005_v84 = _dafny.Set.fromElements((_this).f33, (_this).f32);
        _1986_v69 = (_1986_v69).update(new BigNumber((_2002_v82).length), (new BigNumber((_2005_v84).length)).plus(new BigNumber(289)));
        let _2006_v85;
        _2006_v85 = _module.D10.create_DC29((_this).f32, (_this).f33, new BigNumber((_1921_v15).length));
        let _2007_v86;
        let _nw316 = Array((_dafny.ONE).toNumber());
        _nw316[(_dafny.ZERO).toNumber()] = function (_pat_let43_0) {
          return function (_2008_dt__update__tmp_h1) {
            return function (_pat_let44_0) {
              return function (_2009_dt__update_hcf43_h0) {
                return _module.D10.create_DC29((_2008_dt__update__tmp_h1).dtor_cf42, _2009_dt__update_hcf43_h0, (_2008_dt__update__tmp_h1).dtor_cf44);
              }(_pat_let44_0);
            }((_this).f33);
          }(_pat_let43_0);
        }(_2006_v85);
        _2007_v86 = _nw316;
        let _index360 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_2007_v86).length));
        (_2007_v86)[_index360] = _2006_v85;
        let _2010_v87;
        let _nw317 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
        _2010_v87 = _nw317;
        let _index361 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_2010_v87).length));
        (_2010_v87)[_index361] = (((_this).f32) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-828)), function (_2011_i10) {
          return (_this).f25;
        })) : (_2002_v82));
        let _index362 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_2007_v86).length));
        let _index363 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_2010_v87).length));
        let _rhs349 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_2000_v79, (_this).f25), (_dafny.ZERO).minus(new BigNumber(-478)));
        let _rhs350 = _2006_v85;
        let _rhs351 = _module.__default.fm19(((_this).f32) === (true), globalState);
        let _lhs262 = _2007_v86;
        let _lhs263 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_2007_v86).length));
        let _lhs264 = _2010_v87;
        let _lhs265 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_2010_v87).length));
        _2000_v79 = _rhs349;
        _lhs262[_lhs263] = _rhs350;
        _lhs264[_lhs265] = _rhs351;
        _2004_v83 = (_dafny.MultiSet.fromElements((_this).f25)).Intersect(_2004_v83);
      }
      let _2012_v88;
      _2012_v88 = _dafny.Seq.UnicodeFromString("hnyeciygy");
      let _2013_v89;
      _2013_v89 = _dafny.MultiSet.fromElements(new BigNumber((_2012_v88).length));
      let _2014_v90;
      _2014_v90 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_2013_v89)).length), new BigNumber((_2012_v88).length), (_this).f25, new BigNumber(969));
      let _2015_v91;
      _2015_v91 = _dafny.Seq.of((_this).f25, new BigNumber(578), (_this).f25);
      let _2016_v92;
      _2016_v92 = _dafny.Set.fromElements((_this).f33);
      let _2017_v93;
      let _nw318 = Array((new BigNumber(25)).toNumber());
      _nw318[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw318[(_dafny.ONE).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(2)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(3)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(4)).toNumber()] = new BigNumber((_2013_v89).cardinality());
      _nw318[(new BigNumber(5)).toNumber()] = (_2015_v91)[_module.__default.safeIndex((_this).f25, new BigNumber((_2015_v91).length))];
      _nw318[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(994)), function (_2018_i11) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length);
      _nw318[(new BigNumber(7)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(8)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f33, (_this).f33)).length);
      _nw318[(new BigNumber(10)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(11)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(12)).toNumber()] = new BigNumber((_2016_v92).length);
      _nw318[(new BigNumber(13)).toNumber()] = new BigNumber((_2012_v88).length);
      _nw318[(new BigNumber(14)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(15)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(16)).toNumber()] = new BigNumber(317);
      _nw318[(new BigNumber(17)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(18)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(19)).toNumber()] = new BigNumber(874);
      _nw318[(new BigNumber(20)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(21)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(22)).toNumber()] = new BigNumber(-466);
      _nw318[(new BigNumber(23)).toNumber()] = (_this).f25;
      _nw318[(new BigNumber(24)).toNumber()] = new BigNumber(57);
      _2017_v93 = _nw318;
      let _2019_v94;
      _2019_v94 = _dafny.Seq.of(_2017_v93, _2017_v93);
      (globalState).f23 = _module.__default.fm33((_this).f25, (_this).f32, new BigNumber((_2019_v94).length), globalState);
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _2020_v0;
      let _init72 = ((_2021_p1, _2022_p2) => function (_2023_i1) {
        return (_2023_i1).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), ((_2024_p1, _2025_p2) => function (_2026_i2) {
          return _dafny.MultiSet.fromElements(_2024_p1, _2025_p2, (_this).f33);
        })(_2021_p1, _2022_p2))).length));
      })(p1, p2);
      let _nw319 = Array((new BigNumber(20)).toNumber());
      for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw319.length); _i0_72++) {
        _nw319[_i0_72] = _init72(new BigNumber(_i0_72));
      }
      _2020_v0 = _nw319;
      for (const _guard_loop_10 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2020_v0).length))) {
        let _2027_i0 = _guard_loop_10;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2027_i0)) && ((_2027_i0).isLessThan(new BigNumber((_2020_v0).length))))) {
          (_2020_v0)[(_2027_i0)] = _module.__default.safeDivisionInt(_2027_i0, _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_2028_i3) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          })).length), (_this).f25));
        }
      }
      let _2029_v1;
      let _nw320 = new _module.C2();
      _nw320.__ctor();
      _2029_v1 = _nw320;
      let _2030_v2;
      _2030_v2 = _dafny.Set.fromElements(p1);
      let _rhs352 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(-506)), ((_this).f25).plus(new BigNumber((_2030_v2).length)));
      let _rhs353 = (_module.__default.fm20(p2, globalState)).IsProperSubsetOf(_dafny.Set.fromElements(p1, (_this).f32));
      let _rhs354 = _2029_v1;
      let _lhs266 = globalState;
      r0 = _rhs352;
      _lhs266.f12 = _rhs353;
      _2029_v1 = _rhs354;
      let _2031_v3;
      _2031_v3 = _module.D10.create_DC29(p1, true, (_this).f25);
      let _hi11 = (_this).f25;
      for (let _2032_i4 = (_dafny.ZERO).minus((_2031_v3).dtor_cf44); _2032_i4.isLessThan(_hi11); _2032_i4 = _2032_i4.plus(_dafny.ONE)) {
        let _2033_v4;
        let _nw321 = Array((new BigNumber(29)).toNumber()).fill(false);
        _2033_v4 = _nw321;
        let _rhs355 = _2033_v4;
        let _rhs356 = p2;
        let _lhs267 = globalState;
        _2033_v4 = _rhs355;
        _lhs267.f12 = _rhs356;
        let _index364 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        (p3)[_index364] = _2032_i4;
        let _2034_v5;
        _2034_v5 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).length));
        let _index365 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        (p3)[_index365] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_2032_i4, (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f25, (_2034_v5)[_module.__default.safeIndex(_2032_i4, new BigNumber((_2034_v5).length))]))));
        let _index366 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_2033_v4).length));
        (_2033_v4)[_index366] = (_this).f33;
        let _index367 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_2033_v4).length));
        let _index368 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _index369 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _rhs357 = true;
        let _rhs358 = (p3)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length))];
        let _rhs359 = (_this).f32;
        let _rhs360 = (p3)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length))];
        let _lhs268 = _2033_v4;
        let _lhs269 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_2033_v4).length));
        let _lhs270 = p3;
        let _lhs271 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _lhs272 = globalState;
        let _lhs273 = p3;
        let _lhs274 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        _lhs268[_lhs269] = _rhs357;
        _lhs270[_lhs271] = _rhs358;
        _lhs272.f12 = _rhs359;
        _lhs273[_lhs274] = _rhs360;
        let _2035_v6;
        _2035_v6 = _dafny.Seq.of(_dafny.Seq.of(_2032_i4), _dafny.Seq.Concat(_2034_v5, _2034_v5), _2034_v5, _2034_v5, _dafny.Seq.of(_2032_i4));
        let _2036_v7;
        _2036_v7 = _dafny.Seq.of((_this).f32, p1);
        let _index370 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _index371 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _index372 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _rhs361 = (_2035_v6)[_module.__default.safeIndex(new BigNumber(((((_this).f33) ? (_dafny.Map.Empty.slice().updateUnsafe(_2033_v4,(_this).f25)) : (_dafny.Map.Empty.slice().updateUnsafe(_2033_v4,_2032_i4)))).length), new BigNumber((_2035_v6).length))];
        let _rhs362 = _2033_v4;
        let _rhs363 = new BigNumber((_2036_v7).length);
        let _rhs364 = (_2032_i4).plus((p3)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length))]);
        let _rhs365 = _2032_i4;
        let _lhs275 = p3;
        let _lhs276 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _lhs277 = p3;
        let _lhs278 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        let _lhs279 = p3;
        let _lhs280 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p3).length));
        _2034_v5 = _rhs361;
        _2033_v4 = _rhs362;
        _lhs275[_lhs276] = _rhs363;
        _lhs277[_lhs278] = _rhs364;
        _lhs279[_lhs280] = _rhs365;
      }
      let _2037_v8;
      _2037_v8 = _module.D7.create_DC22((new BigNumber(-783)).plus((_this).f25), (_this).f25);
      _2037_v8 = _2037_v8;
      if (p2) {
        let _2038_v9;
        _2038_v9 = new _dafny.CodePoint('s'.codePointAt(0));
        let _2039_v10;
        _2039_v10 = _dafny.Seq.UnicodeFromString("yxkxcmnk");
        let _2040_v11;
        _2040_v11 = _module.D1.create_DC6(_2038_v9, (_this).f32, _2039_v10);
        let _2041_v12;
        _2041_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2039_v10,p2);
        let _2042_v13;
        _2042_v13 = _2041_v12;
        let _2043_v14;
        _2043_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2040_v11).dtor_cf9).length),(_2042_v13));
        let _2044_v15;
        _2044_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((((_2043_v14).contains((_this).f25)) ? ((_2043_v14).get((_this).f25)) : (_2041_v12))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), (_this).f32, _2039_v10)).dtor_cf9,(_this).f32))).length),_2039_v10);
        let _2045_v16;
        _2045_v16 = _dafny.Seq.of(_2044_v15, _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_2039_v10));
        let _2046_v17;
        _2046_v17 = _dafny.MultiSet.fromElements(true);
        _2044_v15 = (_2045_v16)[_module.__default.safeIndex(new BigNumber(((_2046_v17).Intersect(_2046_v17)).cardinality()), new BigNumber((_2045_v16).length))];
        let _2047_v18;
        let _nw322 = Array((new BigNumber(19)).toNumber());
        _nw322[(_dafny.ZERO).toNumber()] = (_2039_v10)[_module.__default.safeIndex(new BigNumber((_2039_v10).length), new BigNumber((_2039_v10).length))];
        _nw322[(_dafny.ONE).toNumber()] = ((!((_this).f32)) ? (_module.__default.fm39(p2, globalState)) : (_2038_v9));
        _nw322[(new BigNumber(2)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(3)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(4)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(5)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(6)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(7)).toNumber()] = (((_this).f33) ? (_2038_v9) : (new _dafny.CodePoint('w'.codePointAt(0))));
        _nw322[(new BigNumber(8)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(9)).toNumber()] = _module.__default.fm39(p1, globalState);
        _nw322[(new BigNumber(10)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(11)).toNumber()] = _module.__default.fm39((_this).f33, globalState);
        _nw322[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw322[(new BigNumber(13)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(14)).toNumber()] = (_2039_v10)[_module.__default.safeIndex((_this).f25, new BigNumber((_2039_v10).length))];
        _nw322[(new BigNumber(15)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(16)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(17)).toNumber()] = _2038_v9;
        _nw322[(new BigNumber(18)).toNumber()] = _2038_v9;
        _2047_v18 = _nw322;
        let _index373 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2047_v18).length));
        (_2047_v18)[_index373] = _2038_v9;
        let _2048_v19;
        _2048_v19 = _module.D9.create_DC27((_dafny.ZERO).minus((_this).f25), (_this).f25);
        let _2049_v20;
        _2049_v20 = _dafny.Seq.of(_2048_v19);
        let _2050_v21;
        _2050_v21 = _dafny.Seq.of(_dafny.Seq.Concat(_2049_v20, _2049_v20), _dafny.Seq.Concat(_2049_v20, _2049_v20), _dafny.Seq.of(_2048_v19));
        let _index374 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2047_v18).length));
        let _rhs366 = _2038_v9;
        let _rhs367 = ((_this).f25).minus((_this).f25);
        let _rhs368 = new _dafny.CodePoint('w'.codePointAt(0));
        let _rhs369 = _dafny.Seq.Concat(_module.__default.fm2((_this).f25, (_this).f25, globalState), _2039_v10);
        let _rhs370 = (_2050_v21)[_module.__default.safeIndex((_this).f25, new BigNumber((_2050_v21).length))];
        let _lhs281 = _2047_v18;
        let _lhs282 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_2047_v18).length));
        let _lhs283 = globalState;
        _2038_v9 = _rhs366;
        r1 = _rhs367;
        _lhs281[_lhs282] = _rhs368;
        _lhs283.f10 = _rhs369;
        _2049_v20 = _rhs370;
        (globalState).f12 = (_this).f32;
        let _2051_v22;
        _2051_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2039_v10,(_2047_v18)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_2047_v18).length))]);
        let _2052_v23;
        let _nw323 = new _module.C1();
        _nw323.__ctor(new BigNumber((_2051_v22).length), (_this).f26);
        _2052_v23 = _nw323;
        let _2053_v24;
        _2053_v24 = _dafny.Seq.of((_this).f33);
        _2053_v24 = _2053_v24;
      } else {
        let _2054_v25;
        _2054_v25 = _module.D15.create_DC40();
        let _source27 = _2054_v25;
        if (_source27.is_DC40) {
          let _2055_v26;
          let _out37;
          _out37 = (_this).m8(p1, p3, _dafny.Seq.UnicodeFromString("maar"), globalState);
          _2055_v26 = _out37;
          let _2056_v27;
          _2056_v27 = new _dafny.CodePoint('d'.codePointAt(0));
          let _2057_v28;
          _2057_v28 = _dafny.Seq.of(_2056_v27);
          let _2058_v29;
          _2058_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
          let _index375 = _module.__default.safeIndex(new BigNumber(607), new BigNumber(((_2055_v26).f26).length));
          ((_2055_v26).f26)[_index375] = _dafny.Seq.Concat(_2057_v28, _dafny.Seq.update(_2057_v28, _module.__default.safeIndex(new BigNumber((_2058_v29).length), new BigNumber((_2057_v28).length)), _2056_v27));
          let _2059_v30;
          _2059_v30 = _dafny.MultiSet.fromElements(_2057_v28);
          let _2060_v31;
          _2060_v31 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_2061_i5) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length));
          let _2062_v32;
          _2062_v32 = _module.D1.create_DC6((_2057_v28)[_module.__default.safeIndex(new BigNumber((_2060_v31).length), new BigNumber((_2057_v28).length))], (_this).f33, _2057_v28);
          let _2063_v33;
          _2063_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2057_v28).length),(_this).f32);
          let _2064_v34;
          _2064_v34 = _dafny.Seq.of(true);
          let _index376 = _module.__default.safeIndex(new BigNumber(607), new BigNumber(((_2055_v26).f26).length));
          let _rhs371 = !(!(!((_dafny.ZERO).minus((((_2059_v30).contains((_2062_v32).dtor_cf9)) ? ((_2059_v30).get((_2062_v32).dtor_cf9)) : (new BigNumber((_2057_v28).length))))).isEqualTo((_this).f25)) || ((_this).f32));
          let _rhs372 = (((_this).f33) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), ((_2065_v27) => function (_2066_i6) {
            return _2065_v27;
          })(_2056_v27))) : (_dafny.Seq.update(_2057_v28, _module.__default.safeIndex(new BigNumber((_2058_v29).length), new BigNumber((_2057_v28).length)), _module.__default.fm39(false, globalState))));
          let _rhs373 = (((_this).f33) ? ((((_2063_v33).contains((_this).f25)) ? ((_2063_v33).get((_this).f25)) : ((_this).f32))) : ((_2064_v34)[_module.__default.safeIndex((_this).f25, new BigNumber((_2064_v34).length))]));
          let _lhs284 = globalState;
          let _lhs285 = (_2055_v26).f26;
          let _lhs286 = _module.__default.safeIndex(new BigNumber(607), new BigNumber(((_2055_v26).f26).length));
          let _lhs287 = globalState;
          _lhs284.f12 = _rhs371;
          _lhs285[_lhs286] = _rhs372;
          _lhs287.f12 = _rhs373;
          let _2067_v35;
          let _init73 = ((_2068_v29) => function (_2069_i7) {
            return _dafny.MultiSet.fromElements(new BigNumber((_2068_v29).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f32)).cardinality()));
          })(_2058_v29);
          let _nw324 = Array((new BigNumber(29)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw324.length); _i0_73++) {
            _nw324[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2067_v35 = _nw324;
          _2067_v35 = _2067_v35;
          let _2070_v36;
          _2070_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2057_v28).length),_2056_v27);
          (globalState).f12 = ((p2) ? (!_dafny.Seq.contains(_module.__default.fm2(new BigNumber((_2070_v36).length), new BigNumber(425), globalState), _2056_v27)) : (!(!(p2)) || (p1)));
        } else {
          let _2071___mcc_h0 = (_source27).cf59;
          let _2072_cf59 = _2071___mcc_h0;
          r1 = (_dafny.ZERO).minus((_this).f25);
          r0 = (_this).f25;
          (globalState).f12 = p2;
          r0 = ((!(!((_this).f32))) ? ((_dafny.ZERO).minus(_module.__default.fm31(_dafny.Seq.UnicodeFromString("lkfqgj"), globalState))) : ((_this).f25));
        }
        let _2073_v37;
        _2073_v37 = new _dafny.CodePoint('a'.codePointAt(0));
        let _2074_v38;
        let _nw325 = Array((new BigNumber(22)).toNumber());
        _nw325[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw325[(_dafny.ONE).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(2)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(3)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(4)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw325[(new BigNumber(6)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(7)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(8)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(9)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(10)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(11)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(12)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(13)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(14)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(15)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(16)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(17)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(18)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(19)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(20)).toNumber()] = _2073_v37;
        _nw325[(new BigNumber(21)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _2074_v38 = _nw325;
        let _2075_v39;
        _2075_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("b"),_2074_v38);
        let _2076_v40;
        _2076_v40 = _dafny.Map.Empty.slice().updateUnsafe((_2075_v39).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(223)), ((_2077_v37) => function (_2078_i8) {
          return _2077_v37;
        })(_2073_v37)),_2074_v38)),!((_this).f25).isEqualTo((_this).f25));
        _2076_v40 = (_2076_v40).update(_2075_v39, !(new BigNumber((_dafny.Seq.UnicodeFromString("wya")).length)).isEqualTo((_this).f25));
        (globalState).f12 = (_this).fm13(globalState);
        let _2079_v41;
        _2079_v41 = _module.D0.create_DC0(p1);
        let _2080_v42;
        _2080_v42 = _dafny.Seq.of(_2079_v41);
        let _2081_v43;
        _2081_v43 = _dafny.Set.fromElements(_2080_v42);
        if ((_2081_v43).IsSubsetOf((_2081_v43).Union(_2081_v43))) {
          let _2082_v44;
          let _nw326 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _2082_v44 = _nw326;
          let _2083_v45;
          _2083_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_2084_i9) {
            return (_this).f25;
          })).length),(_this).f25);
          let _index377 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_2082_v44).length));
          (_2082_v44)[_index377] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2083_v45).length),(_this).f33);
          let _2085_v46;
          _2085_v46 = _dafny.Seq.UnicodeFromString("fpmh");
          let _2086_v47;
          _2086_v47 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f25), (_this).f25);
          let _2087_v48;
          _2087_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
          let _index378 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_2082_v44).length));
          let _rhs374 = (_this).f33;
          let _rhs375 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pdjfim"), _2085_v46);
          let _rhs376 = ((((_2086_v47).contains(new BigNumber((_2085_v46).length))) ? ((_2086_v47).get(new BigNumber((_2085_v46).length))) : (_module.__default.fm44((_this).f33, p1, (_this).f32, (_this).f25, globalState)))).minus((_this).f25);
          let _rhs377 = (_2087_v48).Merge((function () {
            let _coll75 = new _dafny.Map();
            for (const _compr_75 of _dafny.IntegerRange(new BigNumber(349), new BigNumber(-119))) {
              let _2088_v49 = _compr_75;
              if (((new BigNumber(349)).isLessThanOrEqualTo(_2088_v49)) && ((_2088_v49).isLessThan(new BigNumber(-119)))) {
                _coll75.push([(_2088_v49).multipliedBy((_dafny.ZERO).minus((_this).f25)),p1]);
              }
            }
            return _coll75;
          }()).Merge(_2087_v48));
          let _lhs288 = globalState;
          let _lhs289 = globalState;
          let _lhs290 = _2082_v44;
          let _lhs291 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_2082_v44).length));
          _lhs288.f12 = _rhs374;
          _lhs289.f10 = _rhs375;
          r1 = _rhs376;
          _lhs290[_lhs291] = _rhs377;
          let _2089_v50;
          let _nw327 = new _module.C5();
          _nw327.__ctor(false);
          _2089_v50 = _nw327;
          let _2090_v51;
          let _init74 = ((_2091_v47) => function (_2092_i10) {
            return _2091_v47;
          })(_2086_v47);
          let _nw328 = Array((new BigNumber(9)).toNumber());
          for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw328.length); _i0_74++) {
            _nw328[_i0_74] = _init74(new BigNumber(_i0_74));
          }
          _2090_v51 = _nw328;
          let _index379 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2090_v51).length));
          (_2090_v51)[_index379] = _2086_v47;
          let _2093_v52;
          _2093_v52 = _module.D7.create_DC21(_2020_v0);
          let _2094_v53;
          let _nw329 = Array((new BigNumber(24)).toNumber());
          _nw329[(_dafny.ZERO).toNumber()] = p3;
          _nw329[(_dafny.ONE).toNumber()] = p3;
          _nw329[(new BigNumber(2)).toNumber()] = p3;
          _nw329[(new BigNumber(3)).toNumber()] = p3;
          _nw329[(new BigNumber(4)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(5)).toNumber()] = p3;
          _nw329[(new BigNumber(6)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(7)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(8)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(9)).toNumber()] = p3;
          _nw329[(new BigNumber(10)).toNumber()] = (_2093_v52).dtor_cf32;
          _nw329[(new BigNumber(11)).toNumber()] = p3;
          _nw329[(new BigNumber(12)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(13)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(14)).toNumber()] = (_2093_v52).dtor_cf32;
          _nw329[(new BigNumber(15)).toNumber()] = p3;
          _nw329[(new BigNumber(16)).toNumber()] = p3;
          _nw329[(new BigNumber(17)).toNumber()] = p3;
          _nw329[(new BigNumber(18)).toNumber()] = p3;
          _nw329[(new BigNumber(19)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(20)).toNumber()] = _2020_v0;
          _nw329[(new BigNumber(21)).toNumber()] = p3;
          _nw329[(new BigNumber(22)).toNumber()] = p3;
          _nw329[(new BigNumber(23)).toNumber()] = _2020_v0;
          _2094_v53 = _nw329;
          let _index380 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2094_v53).length));
          (_2094_v53)[_index380] = _2020_v0;
          let _2095_v54;
          _2095_v54 = _dafny.Seq.of(_2030_v2, (_2030_v2).Intersect(_dafny.Set.fromElements(p1, p1)), _2030_v2);
          let _2096_v55;
          _2096_v55 = _dafny.Seq.of(_2020_v0, p3, _2020_v0, _2020_v0, _2020_v0);
          let _index381 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2090_v51).length));
          let _index382 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2094_v53).length));
          let _rhs378 = new BigNumber((_2095_v54).length);
          let _rhs379 = _2086_v47;
          let _rhs380 = (_2096_v55)[_module.__default.safeIndex(_module.__default.fm44(p2, p1, p1, (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25)), globalState), new BigNumber((_2096_v55).length))];
          let _lhs292 = _2090_v51;
          let _lhs293 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2090_v51).length));
          let _lhs294 = _2094_v53;
          let _lhs295 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2094_v53).length));
          r1 = _rhs378;
          _lhs292[_lhs293] = _rhs379;
          _lhs294[_lhs295] = _rhs380;
          (globalState).f10 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-673)), ((_2097_v37) => function (_2098_i11) {
            return _2097_v37;
          })(_2073_v37));
          (globalState).f10 = _2085_v46;
        } else {
          let _2099_v56;
          let _nw330 = Array((new BigNumber(27)).toNumber()).fill(false);
          _2099_v56 = _nw330;
          let _2100_v57;
          _2100_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f32);
          let _index383 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2099_v56).length));
          (_2099_v56)[_index383] = (_module.D13.create_DC36(_2100_v57, (_this).fm13(globalState), (_this).f25)).dtor_cf55;
          let _index384 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2099_v56).length));
          (_2099_v56)[_index384] = p1;
          let _2101_v58;
          _2101_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2073_v37,(_dafny.ZERO).minus(((_this).f25).minus((_this).f25)));
          _2101_v58 = (_2101_v58).update(_2073_v37, (_this).f25);
          r0 = (new BigNumber(22)).plus((_this).f25);
          _2100_v57 = (_2100_v57).update(p1, !((_2099_v56)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_2099_v56).length))]));
          (globalState).f23 = (((_2099_v56)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_2099_v56).length))]) ? (_2073_v37) : (_2073_v37));
        }
        (globalState).f12 = !(p2) || ((_this).f32);
      }
      let _2102_v59;
      let _nw331 = Array((new BigNumber(2)).toNumber()).fill([]);
      _2102_v59 = _nw331;
      let _index385 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_2102_v59).length));
      (_2102_v59)[_index385] = p3;
      let _2103_v60;
      _2103_v60 = new _dafny.CodePoint('b'.codePointAt(0));
      let _2104_v61;
      let _nw332 = Array((new BigNumber(3)).toNumber());
      _nw332[(_dafny.ZERO).toNumber()] = ((p1) ? (new _dafny.CodePoint('f'.codePointAt(0))) : (_2103_v60));
      _nw332[(_dafny.ONE).toNumber()] = _2103_v60;
      _nw332[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
      _2104_v61 = _nw332;
      let _index386 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2104_v61).length));
      (_2104_v61)[_index386] = _2103_v60;
      let _2105_v62;
      _2105_v62 = _dafny.Seq.UnicodeFromString("qpkq");
      let _index387 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_2102_v59).length));
      let _index388 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2104_v61).length));
      let _rhs381 = p3;
      let _rhs382 = (_2105_v62)[_module.__default.safeIndex(new BigNumber((_2105_v62).length), new BigNumber((_2105_v62).length))];
      let _lhs296 = _2102_v59;
      let _lhs297 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_2102_v59).length));
      let _lhs298 = _2104_v61;
      let _lhs299 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2104_v61).length));
      _lhs296[_lhs297] = _rhs381;
      _lhs298[_lhs299] = _rhs382;
      r0 = new BigNumber(457);
      let _2106_v63;
      _2106_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f25);
      r1 = (new BigNumber((_2106_v63).length)).minus((new BigNumber((_2106_v63).length)).plus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25))));
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      (globalState).f12 = true;
      let _2107_v1;
      _2107_v1 = new _dafny.CodePoint('r'.codePointAt(0));
      let _2108_v2;
      _2108_v2 = _dafny.MultiSet.fromElements(_2107_v1);
      let _2109_v3;
      _2109_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2108_v2).cardinality()),!(false));
      let _2110_v4;
      _2110_v4 = _module.D7.create_DC22((_this).f25, new BigNumber((_2109_v3).length));
      let _2111_v5;
      _2111_v5 = _dafny.Seq.of((_this).f32, !((_this).f32), p1);
      let _2112_v6;
      _2112_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2110_v4,_2111_v5);
      let _2113_v7;
      let _nw333 = new _module.C4();
      _nw333.__ctor((_this).f26, new BigNumber(((function () {
        let _coll76 = new _dafny.Map();
        for (const _compr_76 of (_2112_v6).Keys.Elements) {
          let _2114_v0 = _compr_76;
          if ((_2112_v6).contains(_2114_v0)) {
            _coll76.push([_2114_v0,(_this).f32]);
          }
        }
        return _coll76;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2110_v4,(_this).f33))).length), (_this).f26);
      _2113_v7 = _nw333;
      let _2115_v8;
      _2115_v8 = _dafny.MultiSet.fromElements((_this).f25);
      (globalState).f12 = (_2115_v8).IsProperSubsetOf(_2115_v8);
      _2109_v3 = (_2109_v3).update((_this).f25, true);
      let _2116_v9;
      _2116_v9 = _dafny.Seq.of((_this).f25);
      let _2117_v10;
      let _nw334 = Array((new BigNumber(5)).toNumber());
      _nw334[(_dafny.ZERO).toNumber()] = (_2111_v5)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_2111_v5).length))];
      _nw334[(_dafny.ONE).toNumber()] = (_this).f33;
      _nw334[(new BigNumber(2)).toNumber()] = (_this).f32;
      _nw334[(new BigNumber(3)).toNumber()] = !_dafny.areEqual(_2116_v9, _2116_v9);
      _nw334[(new BigNumber(4)).toNumber()] = (_this).f33;
      _2117_v10 = _nw334;
      let _index389 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_2117_v10).length));
      (_2117_v10)[_index389] = !(_module.__default.fm0((_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2107_v1,(_this).f25)).update(_2107_v1, (_this).f25)).length)), globalState));
      let _index390 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_2117_v10).length));
      (_2117_v10)[_index390] = (_this).f33;
      (globalState).f6 = _dafny.Seq.Concat(_2116_v9, _2116_v9);
      let _2118_v11;
      let _nw335 = new _module.C2();
      _nw335.__ctor();
      _2118_v11 = _nw335;
      let _2119_v12;
      _2119_v12 = _dafny.MultiSet.fromElements(_2118_v11);
      let _2120_v13;
      _2120_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2119_v12).cardinality()),(_this).f25);
      let _2121_v14;
      _2121_v14 = _dafny.Set.fromElements((_this).f25);
      r0 = ((_this).f25).multipliedBy(((_2116_v9)[_module.__default.safeIndex((_this).f25, new BigNumber((_2116_v9).length))]).multipliedBy((((_2120_v13).contains((_this).f25)) ? ((_2120_v13).get((_this).f25)) : (new BigNumber((_2121_v14).length)))));
      return r0;
    }
    m8(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = undefined;
      let _2122_v0;
      _2122_v0 = new BigNumber(668);
      _2122_v0 = (_this).f25;
      (globalState).f12 = (_this).f33;
      let _2123_v1;
      _2123_v1 = _dafny.Seq.of(p2);
      let _2124_v2;
      _2124_v2 = new _dafny.CodePoint('l'.codePointAt(0));
      let _2125_v3;
      _2125_v3 = _dafny.MultiSet.fromElements(((p0) ? (new BigNumber(250)) : (new BigNumber((p2).length))), _2122_v0, _2122_v0, _module.__default.safeModuloInt(new BigNumber(((_2123_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gs"), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.UnicodeFromString("gs")).length)), _2124_v2)).length), new BigNumber((_2123_v1).length))]).length), _2122_v0), (_this).f25);
      _2125_v3 = _2125_v3;
      let _2126_v4;
      let _nw336 = Array((new BigNumber(8)).toNumber()).fill(false);
      _2126_v4 = _nw336;
      _2126_v4 = _2126_v4;
      let _2127_v5;
      _2127_v5 = _dafny.Set.fromElements(_2124_v2, _2124_v2);
      let _2128_i0;
      _2128_i0 = _dafny.ZERO;
      L15: {
        while ((_2127_v5).IsDisjointFrom(_2127_v5)) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2128_i0)) {
              break L15;
            }
            _2128_i0 = (_2128_i0).plus(_dafny.ONE);
            let _2129_v6;
            _2129_v6 = _dafny.Seq.of(!(p0), !((_this).f32));
            let _2130_v7;
            _2130_v7 = _dafny.Set.fromElements((_this).f32);
            let _rhs383 = (_this).f25;
            let _rhs384 = p0;
            let _rhs385 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f25, (_this).f25), _module.__default.safeModuloInt(_2122_v0, new BigNumber((_2129_v6).length)));
            let _rhs386 = _module.__default.safeModuloInt(new BigNumber(314), (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2130_v7,p2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2130_v7,_dafny.Seq.UnicodeFromString("hwy")))).length)));
            let _lhs300 = globalState;
            _2122_v0 = _rhs383;
            _lhs300.f12 = _rhs384;
            _2122_v0 = _rhs385;
            _2122_v0 = _rhs386;
            let _2131_v8;
            _2131_v8 = _dafny.MultiSet.fromElements((_this).f32, p0);
            (globalState).f12 = (_2131_v8).IsDisjointFrom((_2131_v8).Intersect(_2131_v8));
            if ((_this).f33) {
              let _index391 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((p1).length));
              (p1)[_index391] = (_dafny.ZERO).minus((_this).f25);
              let _index392 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((p1).length));
              (p1)[_index392] = (_this).f25;
              let _2132_v9;
              _2132_v9 = _dafny.Seq.of((p1)[_module.__default.safeIndex(new BigNumber(677), new BigNumber((p1).length))], new BigNumber((_dafny.Seq.UnicodeFromString("wfla")).length));
              let _2133_v10;
              _2133_v10 = _dafny.Set.fromElements(_2125_v3, (_2125_v3).Intersect((_2125_v3).update((p1)[_module.__default.safeIndex(new BigNumber(677), new BigNumber((p1).length))], _module.__default.abs(new BigNumber((_2132_v9).length)))));
              let _2134_v11;
              _2134_v11 = _module.D18.create_DC43(_2133_v10);
              _2133_v10 = (_2134_v11).dtor_cf62;
              let _2135_v12;
              let _nw337 = new _module.C3();
              _nw337.__ctor((p1)[_module.__default.safeIndex(new BigNumber(677), new BigNumber((p1).length))], p0, (_dafny.ZERO).minus((_this).f25), (_this).f26);
              _2135_v12 = _nw337;
              let _2136_v13;
              _2136_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2126_v4,(_2135_v12).f38);
              _2136_v13 = (_2136_v13).update(_2126_v4, (_this).f33);
              let _2137_v14;
              let _nw338 = new _module.C2();
              _nw338.__ctor();
              _2137_v14 = _nw338;
            } else {
              let _2138_v15;
              _2138_v15 = _dafny.Set.fromElements(_2122_v0, (_this).f25);
              let _2139_v16;
              _2139_v16 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new BigNumber(433))).Union(_2138_v15),_2129_v6);
              _2139_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2138_v15,_2129_v6);
              let _2140_v17;
              let _nw339 = Array((new BigNumber(20)).toNumber());
              _nw339[(_dafny.ZERO).toNumber()] = _2126_v4;
              _nw339[(_dafny.ONE).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(2)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(3)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(4)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(5)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(6)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(7)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(8)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(9)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(10)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(11)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(12)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(13)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(14)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(15)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(16)).toNumber()] = (((_this).f33) ? (_2126_v4) : (_2126_v4));
              _nw339[(new BigNumber(17)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(18)).toNumber()] = _2126_v4;
              _nw339[(new BigNumber(19)).toNumber()] = _2126_v4;
              _2140_v17 = _nw339;
              _2140_v17 = _2140_v17;
              let _2141_v18;
              _2141_v18 = _dafny.Seq.of(_2122_v0);
              let _2142_v19;
              _2142_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v2,_module.__default.fm0(new BigNumber((_2141_v18).length), globalState));
              let _2143_v20;
              _2143_v20 = _module.D10.create_DC28(_2142_v19);
              _2143_v20 = ((false) ? (_2143_v20) : (_module.D10.create_DC28(_2142_v19)));
              (globalState).f12 = _dafny.Seq.IsPrefixOf(p2, _dafny.Seq.UnicodeFromString("amuiefuy"));
              let _index393 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((p1).length));
              (p1)[_index393] = _2122_v0;
              let _index394 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((p1).length));
              (p1)[_index394] = (_dafny.ZERO).minus(_module.__default.fm44(p0, (_this).f32, _module.__default.fm0(new BigNumber((_2125_v3).cardinality()), globalState), _2122_v0, globalState));
            }
            if ((_this).f32) {
              (globalState).f12 = (_this).f32;
              let _index395 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_2126_v4).length));
              (_2126_v4)[_index395] = _module.__default.fm0(_2122_v0, globalState);
              let _index396 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_2126_v4).length));
              let _rhs387 = (_this).f33;
              let _rhs388 = (_this).f25;
              let _lhs301 = _2126_v4;
              let _lhs302 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_2126_v4).length));
              _lhs301[_lhs302] = _rhs387;
              _2122_v0 = _rhs388;
              let _2144_v21;
              _2144_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), ((_2145_v2) => function (_2146_i1) {
                return _2145_v2;
              })(_2124_v2)), p2), _module.__default.safeIndex(new BigNumber((_2123_v1).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), ((_2147_v2) => function (_2148_i1) {
                return _2147_v2;
              })(_2124_v2)), p2)).length)), _2124_v2),_dafny.Seq.Concat(_2129_v6, _dafny.Seq.of(!(p0))));
              _2129_v6 = (((_2144_v21).contains(p2)) ? ((_2144_v21).get(p2)) : (_dafny.Seq.update(_dafny.Seq.Concat(_2129_v6, _2129_v6), _module.__default.safeIndex(new BigNumber((_2130_v7).length), new BigNumber((_dafny.Seq.Concat(_2129_v6, _2129_v6)).length)), true)));
              (globalState).f12 = (_this).f32;
              let _2149_v22;
              let _nw340 = new _module.C6();
              _nw340.__ctor((_2122_v0).multipliedBy((_this).f25), (_this).f26);
              _2149_v22 = _nw340;
            } else {
              let _2150_v23;
              _2150_v23 = _module.D0.create_DC3((_this).f25, (_this).f33);
              let _2151_v24;
              let _init75 = ((_2152_v23, _2153_v8, _2154_v0, _2155_p2) => function (_2156_i2) {
                return _module.D6.create_DC19(_2152_v23, (((_2153_v8).contains((_this).f33)) ? ((_2153_v8).get((_this).f33)) : ((_module.D18.create_DC44(_2154_v0, false, _2155_p2, (_this).f33, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("boihs")).length), new BigNumber(-626))))).dtor_cf63)));
              })(_2150_v23, _2131_v8, _2122_v0, p2);
              let _nw341 = Array((new BigNumber(20)).toNumber());
              for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw341.length); _i0_75++) {
                _nw341[_i0_75] = _init75(new BigNumber(_i0_75));
              }
              _2151_v24 = _nw341;
              let _2157_v25;
              _2157_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC19(_2150_v23, (_this).f25),_2151_v24);
              let _2158_v26;
              _2158_v26 = _module.D12.create_DC31(_2157_v25);
              let _pat_let_tv39 = _2157_v25;
              _2158_v26 = function (_pat_let45_0) {
                return function (_2159_dt__update__tmp_h0) {
                  return function (_pat_let46_0) {
                    return function (_2160_dt__update_hcf46_h0) {
                      return _module.D12.create_DC31(_2160_dt__update_hcf46_h0);
                    }(_pat_let46_0);
                  }(_pat_let_tv39);
                }(_pat_let45_0);
              }(_2158_v26);
              (globalState).f12 = p0;
              (globalState).f10 = _module.__default.fm2(_2122_v0, new BigNumber(99), globalState);
              let _2161_v27;
              _2161_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_module.__default.safeDivisionInt(_2122_v0, _2122_v0));
              _2161_v27 = _2161_v27;
              _2122_v0 = (((_this).f25).plus((((_2131_v8).contains((_this).f32)) ? ((_2131_v8).get((_this).f32)) : (_2122_v0)))).minus(new BigNumber((p2).length));
            }
          }
        }
      }
      let _2162_v28;
      _2162_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_2122_v0).multipliedBy((_this).f25));
      _2162_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm13(globalState),_2122_v0);
      let _2163_v29;
      let _nw342 = new _module.C1();
      _nw342.__ctor((_this).f25, (_this).f26);
      _2163_v29 = _nw342;
      r0 = _2163_v29;
      return r0;
    }
    m9(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _2164_v0;
      _2164_v0 = _dafny.Seq.of((_this).f32, (_this).f33);
      let _2165_v1;
      let _nw343 = Array((new BigNumber(4)).toNumber());
      _nw343[(_dafny.ZERO).toNumber()] = ((!((_this).f32)) ? (_2164_v0) : (_module.__default.fm1(!((_this).f32), globalState)));
      _nw343[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2164_v0, _2164_v0);
      _nw343[(new BigNumber(2)).toNumber()] = _2164_v0;
      _nw343[(new BigNumber(3)).toNumber()] = _2164_v0;
      _2165_v1 = _nw343;
      for (const _guard_loop_11 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2165_v1).length))) {
        let _2166_i0 = _guard_loop_11;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2166_i0)) && ((_2166_i0).isLessThan(new BigNumber((_2165_v1).length))))) {
          (_2165_v1)[(_2166_i0)] = _dafny.Seq.Concat(_2164_v0, _2164_v0);
        }
      }
      let _2167_v2;
      let _nw344 = new _module.C3();
      _nw344.__ctor((_this).f25, ((_this).f33) === (false), p0, (_this).f26);
      _2167_v2 = _nw344;
      if (((_this).f25).isEqualTo((_this).f25)) {
        if (!(p0).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus((_2167_v2).f37)))) {
          (globalState).f12 = (_2167_v2).f38;
          let _2168_v3;
          let _nw345 = new _module.C0();
          _nw345.__ctor();
          _2168_v3 = _nw345;
          let _2169_v4;
          let _init76 = function (_2170_i1) {
            return (_this).f32;
          };
          let _nw346 = Array((new BigNumber(7)).toNumber());
          for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw346.length); _i0_76++) {
            _nw346[_i0_76] = _init76(new BigNumber(_i0_76));
          }
          _2169_v4 = _nw346;
          let _2171_v5;
          _2171_v5 = _dafny.Set.fromElements(_2169_v4, _2169_v4);
          let _2172_v6;
          _2172_v6 = _module.D3.create_DC9(_2171_v5);
          let _2173_v7;
          _2173_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2167_v2).f37,_2172_v6);
          let _2174_v8;
          _2174_v8 = _module.D19.create_DC46(_2173_v7);
          let _2175_v9;
          let _nw347 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _2175_v9 = _nw347;
          let _2176_v10;
          _2176_v10 = _dafny.Map.Empty.slice().updateUnsafe((_2174_v8).dtor_cf68,_2175_v9);
          let _2177_v11;
          _2177_v11 = _dafny.Seq.of(_2176_v10, _2176_v10, _2176_v10);
          let _rhs389 = (_2176_v10).Merge((_2177_v11)[_module.__default.safeIndex((_2167_v2).f37, new BigNumber((_2177_v11).length))]);
          let _rhs390 = (_this).f33;
          _2176_v10 = _rhs389;
          r0 = _rhs390;
          let _2178_v12;
          _2178_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2169_v4,(_this).f33);
          (globalState).f12 = !(_2178_v12).equals((_2178_v12).Merge(_2178_v12));
          let _2179_v13;
          _2179_v13 = _dafny.Map.Empty.slice().updateUnsafe((_2167_v2).f37,(_this).f32);
          let _2180_v14;
          _2180_v14 = _2179_v13;
          _2180_v14 = _module.__default.fm65(((_2167_v2).f37).plus((_2167_v2).f37), globalState);
        } else {
          r2 = (_this).f32;
          let _2181_v15;
          let _nw348 = Array((new BigNumber(9)).toNumber()).fill([]);
          _2181_v15 = _nw348;
          let _2182_v16;
          _2182_v16 = _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)));
          let _2183_v17;
          _2183_v17 = _dafny.Set.fromElements(new BigNumber((_2182_v16).length));
          let _2184_v18;
          _2184_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,(_2167_v2).f38);
          let _2185_v19;
          _2185_v19 = _module.D12.create_DC32((((_2184_v18).contains((_2167_v2).f38)) ? ((_2184_v18).get((_2167_v2).f38)) : ((_2167_v2).f38)));
          let _2186_v20;
          let _nw349 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _2186_v20 = _nw349;
          let _pat_let_tv40 = _2183_v17;
          let _2187_v21;
          let _2188_v22;
          let _out38;
          let _out39;
          let _outcollector13 = (_2167_v2).m2(_2181_v15, (_2183_v17).IsDisjointFrom((function (_pat_let47_0) {
            return function (_2189_dt__update__tmp_h0) {
              return function (_pat_let48_0) {
                return function (_2190_dt__update_hcf18_h0) {
                  return _module.D4.create_DC11(_2190_dt__update_hcf18_h0);
                }(_pat_let48_0);
              }(_pat_let_tv40);
            }(_pat_let47_0);
          }(_module.D4.create_DC11(_2183_v17))).dtor_cf18), (_2185_v19).dtor_cf47, _2186_v20, globalState);
          _out38 = _outcollector13[0];
          _out39 = _outcollector13[1];
          _2187_v21 = _out38;
          _2188_v22 = _out39;
          let _rhs391 = (_this).f33;
          let _rhs392 = (_2167_v2).f38;
          let _rhs393 = _module.__default.safeDivisionInt((new BigNumber(691)).minus((_2167_v2).f37), (_2167_v2).f37);
          let _lhs303 = globalState;
          r0 = _rhs391;
          _lhs303.f12 = _rhs392;
          _2187_v21 = _rhs393;
          let _2191_v23;
          _2191_v23 = _module.D19.create_DC48(_module.__default.safeModuloInt((_this).f25, new BigNumber(789)));
          _2191_v23 = _module.__default.fm66(globalState);
          let _2192_v24;
          let _init77 = ((_2193_v2) => function (_2194_i2) {
            return (_2193_v2).f38;
          })(_2167_v2);
          let _nw350 = Array((new BigNumber(13)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw350.length); _i0_77++) {
            _nw350[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _2192_v24 = _nw350;
          let _2195_v25;
          _2195_v25 = _module.D12.create_DC33(_2192_v24, _2192_v24, (_this).f33, (_2167_v2).f37);
          let _2196_v26;
          _2196_v26 = _module.D19.create_DC47(_2192_v24, (_this).f32);
          let _2197_v27;
          _2197_v27 = _dafny.MultiSet.fromElements((_2167_v2).f38);
          let _2198_v28;
          let _nw351 = Array((new BigNumber(16)).toNumber());
          _nw351[(_dafny.ZERO).toNumber()] = (_2167_v2).f38;
          _nw351[(_dafny.ONE).toNumber()] = true;
          _nw351[(new BigNumber(2)).toNumber()] = (_2167_v2).f38;
          _nw351[(new BigNumber(3)).toNumber()] = (_2167_v2).f38;
          _nw351[(new BigNumber(4)).toNumber()] = !(!((_2167_v2).f38));
          _nw351[(new BigNumber(5)).toNumber()] = (((_2167_v2).f38) ? ((_this).f33) : ((_2195_v25).dtor_cf50));
          _nw351[(new BigNumber(6)).toNumber()] = (((_2184_v18).contains((_2196_v26).dtor_cf70)) ? ((_2184_v18).get((_2196_v26).dtor_cf70)) : ((_2167_v2).f38));
          _nw351[(new BigNumber(7)).toNumber()] = (_2167_v2).f38;
          _nw351[(new BigNumber(8)).toNumber()] = (_this).f33;
          _nw351[(new BigNumber(9)).toNumber()] = (_2197_v27).IsProperSubsetOf(_2197_v27);
          _nw351[(new BigNumber(10)).toNumber()] = (_2167_v2).f38;
          _nw351[(new BigNumber(11)).toNumber()] = (_this).f33;
          _nw351[(new BigNumber(12)).toNumber()] = (_2188_v22).isLessThan(new BigNumber((_2164_v0).length));
          _nw351[(new BigNumber(13)).toNumber()] = (_2167_v2).f38;
          _nw351[(new BigNumber(14)).toNumber()] = (_2167_v2).f38;
          _nw351[(new BigNumber(15)).toNumber()] = (_2167_v2).f38;
          _2198_v28 = _nw351;
          let _index397 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_2198_v28).length));
          (_2198_v28)[_index397] = (_2167_v2).f38;
          let _index398 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_2198_v28).length));
          (_2198_v28)[_index398] = (_this).f32;
        }
        let _2199_v29;
        _2199_v29 = new _dafny.CodePoint('o'.codePointAt(0));
        let _2200_v30;
        _2200_v30 = _dafny.Seq.UnicodeFromString("gypjrxbv");
        let _2201_v31;
        let _nw352 = new _module.C5();
        _nw352.__ctor(_dafny.Seq.contains(_2200_v30, _2199_v29));
        _2201_v31 = _nw352;
        let _2202_v32;
        let _nw353 = new _module.C2();
        _nw353.__ctor();
        _2202_v32 = _nw353;
        let _2203_v33;
        _2203_v33 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_2200_v30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), ((_2204_v29) => function (_2205_i3) {
          return _2204_v29;
        })(_2199_v29))),_2200_v30);
        _2203_v33 = (_2203_v33).update((_2201_v31).f35, _2200_v30);
        let _2206_v34;
        let _nw354 = new _module.C2();
        _nw354.__ctor();
        _2206_v34 = _nw354;
      } else {
        let _2207_v35;
        _2207_v35 = new BigNumber(774);
        _2207_v35 = (_dafny.ZERO).minus(_2207_v35);
        let _2208_v36;
        let _nw355 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _2208_v36 = _nw355;
        let _index399 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2208_v36).length));
        (_2208_v36)[_index399] = (_2167_v2).f37;
        let _index400 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2208_v36).length));
        (_2208_v36)[_index400] = _2207_v35;
        let _2209_v37;
        _2209_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f32);
        let _2210_v38;
        _2210_v38 = _dafny.MultiSet.fromElements(true, _module.__default.fm0((_2167_v2).f37, globalState), (_2167_v2).f38);
        let _2211_v39;
        _2211_v39 = _dafny.Set.fromElements(_module.D14.create_DC38(p0));
        let _2212_v40;
        _2212_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC38(p0),(_2167_v2).f37);
        _2209_v37 = (_2209_v37).update(new BigNumber((_2210_v38).cardinality()), (function () {
          let _coll77 = new _dafny.Set();
          for (const _compr_77 of (_2212_v40).Keys.Elements) {
            let _2213_v41 = _compr_77;
            if ((_2212_v40).contains(_2213_v41)) {
              _coll77.add(_2213_v41);
            }
          }
          return _coll77;
        }()).IsProperSubsetOf(_2211_v39));
        let _2214_v42;
        _2214_v42 = new _dafny.CodePoint('m'.codePointAt(0));
        (globalState).f23 = _2214_v42;
        let _2215_v43;
        let _nw356 = Array((new BigNumber(6)).toNumber());
        _nw356[(_dafny.ZERO).toNumber()] = (_2167_v2).f38;
        _nw356[(_dafny.ONE).toNumber()] = !((_2167_v2).f38);
        _nw356[(new BigNumber(2)).toNumber()] = (_2167_v2).f38;
        _nw356[(new BigNumber(3)).toNumber()] = (_2167_v2).f38;
        _nw356[(new BigNumber(4)).toNumber()] = (_2167_v2).f38;
        _nw356[(new BigNumber(5)).toNumber()] = (_2167_v2).f38;
        _2215_v43 = _nw356;
        let _2216_v44;
        _2216_v44 = _dafny.Seq.of(_2215_v43);
        let _2217_v45;
        _2217_v45 = _module.D19.create_DC47((_2216_v44)[_module.__default.safeIndex(_2207_v35, new BigNumber((_2216_v44).length))], (_this).f32);
        let _pat_let_tv41 = _2215_v43;
        _2217_v45 = function (_pat_let49_0) {
          return function (_2218_dt__update__tmp_h1) {
            return function (_pat_let50_0) {
              return function (_2219_dt__update_hcf69_h0) {
                return _module.D19.create_DC47(_2219_dt__update_hcf69_h0, (_2218_dt__update__tmp_h1).dtor_cf70);
              }(_pat_let50_0);
            }(_pat_let_tv41);
          }(_pat_let49_0);
        }(_2217_v45);
      }
      let _2220_v46;
      _2220_v46 = new BigNumber(323);
      let _2221_v47;
      _2221_v47 = _module.D10.create_DC29((_2167_v2).f38, true, (_2167_v2).f37);
      let _2222_v48;
      _2222_v48 = _dafny.MultiSet.fromElements((_this).f25, _2220_v46);
      let _2223_v49;
      _2223_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2222_v48,_dafny.Set.fromElements((_this).f32));
      let _pat_let_tv42 = _2167_v2;
      let _2224_v50;
      _2224_v50 = _dafny.Set.fromElements(p0, (function (_pat_let51_0) {
        return function (_2225_dt__update__tmp_h2) {
          return function (_pat_let52_0) {
            return function (_2226_dt__update_hcf42_h0) {
              return function (_pat_let53_0) {
                return function (_2227_dt__update_hcf44_h0) {
                  return _module.D10.create_DC29(_2226_dt__update_hcf42_h0, (_2225_dt__update__tmp_h2).dtor_cf43, _2227_dt__update_hcf44_h0);
                }(_pat_let53_0);
              }((_pat_let_tv42).f37);
            }(_pat_let52_0);
          }(true);
        }(_pat_let51_0);
      }(_2221_v47)).dtor_cf44, new BigNumber((_module.__default.fm67((_this).f32, _2220_v46, _2223_v49, globalState)).length), _module.__default.safeModuloInt((_this).f25, p0));
      _2220_v46 = new BigNumber((_2224_v50).length);
      let _2228_v51;
      let _init78 = function (_2229_i4) {
        return (_this).f32;
      };
      let _nw357 = Array((new BigNumber(14)).toNumber());
      for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw357.length); _i0_78++) {
        _nw357[_i0_78] = _init78(new BigNumber(_i0_78));
      }
      _2228_v51 = _nw357;
      let _2230_v52;
      _2230_v52 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-578));
      let _2231_v53;
      _2231_v53 = _dafny.Seq.of(new BigNumber((_2230_v52).length), (_2167_v2).f37, p0);
      let _2232_v54;
      _2232_v54 = _module.D3.create_DC10((_this).f32, _2228_v51, false, (_2167_v2).f37, new BigNumber((_2231_v53).length));
      let _2233_v55;
      _2233_v55 = _module.D12.create_DC33(_2228_v51, (_2232_v54).dtor_cf14, (_2167_v2).f38, (_2167_v2).f37);
      let _source28 = _2233_v55;
      if (_source28.is_DC32) {
        let _2234___mcc_h0 = (_source28).cf47;
        let _2235_cf47 = _2234___mcc_h0;
        let _2236_v57;
        _2236_v57 = new _dafny.CodePoint('t'.codePointAt(0));
        let _2237_v58;
        _2237_v58 = _dafny.Seq.UnicodeFromString("kflkte");
        let _2238_v59;
        let _nw358 = Array((new BigNumber(19)).toNumber());
        _nw358[(_dafny.ZERO).toNumber()] = (_this).f25;
        _nw358[(_dafny.ONE).toNumber()] = (_this).f25;
        _nw358[(new BigNumber(2)).toNumber()] = p0;
        _nw358[(new BigNumber(3)).toNumber()] = new BigNumber((function () {
          let _coll78 = new _dafny.Map();
          for (const _compr_78 of (_dafny.Map.Empty.slice().updateUnsafe(_2236_v57,_2220_v46)).Keys.Elements) {
            let _2239_v56 = _compr_78;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_2236_v57,_2220_v46)).contains(_2239_v56)) {
              _coll78.push([_2239_v56,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2231_v53)).length),(_2167_v2).f37)).length)]);
            }
          }
          return _coll78;
        }()).length);
        _nw358[(new BigNumber(4)).toNumber()] = (_this).f25;
        _nw358[(new BigNumber(5)).toNumber()] = (_2167_v2).f37;
        _nw358[(new BigNumber(6)).toNumber()] = p0;
        _nw358[(new BigNumber(7)).toNumber()] = (_2167_v2).f37;
        _nw358[(new BigNumber(8)).toNumber()] = (_2167_v2).f37;
        _nw358[(new BigNumber(9)).toNumber()] = (_2167_v2).f37;
        _nw358[(new BigNumber(10)).toNumber()] = new BigNumber((_2237_v58).length);
        _nw358[(new BigNumber(11)).toNumber()] = (_module.D4.create_DC13((_2167_v2).f37, false, (_this).f33, (_this).f33)).dtor_cf19;
        _nw358[(new BigNumber(12)).toNumber()] = (_this).f25;
        _nw358[(new BigNumber(13)).toNumber()] = (_this).f25;
        _nw358[(new BigNumber(14)).toNumber()] = (_2167_v2).f37;
        _nw358[(new BigNumber(15)).toNumber()] = (_2167_v2).f37;
        _nw358[(new BigNumber(16)).toNumber()] = (_2233_v55).dtor_cf51;
        _nw358[(new BigNumber(17)).toNumber()] = (_this).f25;
        _nw358[(new BigNumber(18)).toNumber()] = (_this).f25;
        _2238_v59 = _nw358;
        let _2240_v60;
        _2240_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2238_v59,(_this).f25);
        let _2241_v61;
        let _nw359 = Array((new BigNumber(17)).toNumber());
        _nw359[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_2220_v46);
        _nw359[(_dafny.ONE).toNumber()] = new BigNumber((_2164_v0).length);
        _nw359[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_2167_v2).f37);
        _nw359[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_2167_v2).f37);
        _nw359[(new BigNumber(4)).toNumber()] = p0;
        _nw359[(new BigNumber(5)).toNumber()] = new BigNumber((_2240_v60).length);
        _nw359[(new BigNumber(6)).toNumber()] = (_this).f25;
        _nw359[(new BigNumber(7)).toNumber()] = new BigNumber(978);
        _nw359[(new BigNumber(8)).toNumber()] = (_this).f25;
        _nw359[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2237_v58, _2237_v58)).length);
        _nw359[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw359[(new BigNumber(11)).toNumber()] = (_this).f25;
        _nw359[(new BigNumber(12)).toNumber()] = _2220_v46;
        _nw359[(new BigNumber(13)).toNumber()] = (_2167_v2).f37;
        _nw359[(new BigNumber(14)).toNumber()] = (_2220_v46).minus(new BigNumber((_2237_v58).length));
        _nw359[(new BigNumber(15)).toNumber()] = new BigNumber(730);
        _nw359[(new BigNumber(16)).toNumber()] = (_2167_v2).f37;
        _2241_v61 = _nw359;
        let _index401 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_2238_v59).length));
        (_2238_v59)[_index401] = new BigNumber((_2237_v58).length);
        let _index402 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_2238_v59).length));
        (_2238_v59)[_index402] = (_2167_v2).f37;
        let _index403 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_2238_v59).length));
        (_2238_v59)[_index403] = (_2167_v2).f37;
        let _2242_v62;
        _2242_v62 = _dafny.MultiSet.fromElements(!(true));
        r2 = !((_2242_v62).IsDisjointFrom(_2242_v62));
        let _index404 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_2238_v59).length));
        (_2238_v59)[_index404] = new BigNumber((_2237_v58).length);
      } else if (_source28.is_DC33) {
        let _2243___mcc_h1 = (_source28).cf48;
        let _2244___mcc_h2 = (_source28).cf49;
        let _2245___mcc_h3 = (_source28).cf50;
        let _2246___mcc_h4 = (_source28).cf51;
        let _2247_cf51 = _2246___mcc_h4;
        let _2248_cf50 = _2245___mcc_h3;
        let _2249_cf49 = _2244___mcc_h2;
        let _2250_cf48 = _2243___mcc_h1;
        _2247_cf51 = (_module.__default.safeDivisionInt((_this).f25, new BigNumber((_2164_v0).length))).plus((_this).f25);
        let _2251_v63;
        _2251_v63 = _dafny.Map.Empty.slice().updateUnsafe((_2167_v2).f37,(_this).f25);
        let _2252_v64;
        _2252_v64 = _dafny.Seq.UnicodeFromString("lrmgmmsn");
        let _2253_v65;
        _2253_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2252_v64,true);
        let _2254_v66;
        let _nw360 = Array((new BigNumber(4)).toNumber());
        _nw360[(_dafny.ZERO).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2251_v63).length),_2252_v64)).update(new BigNumber((_2253_v65).length), _dafny.Seq.UnicodeFromString("dtftgurm"))).length);
        _nw360[(_dafny.ONE).toNumber()] = _2247_cf51;
        _nw360[(new BigNumber(2)).toNumber()] = _module.__default.fm44((_this).f32, !((_this).f33), (_this).f33, (_dafny.ZERO).minus((_dafny.ZERO).minus((_2167_v2).f37)), globalState);
        _nw360[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_this).f25, p0);
        _2254_v66 = _nw360;
        let _index405 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
        (_2254_v66)[_index405] = new BigNumber(788);
        let _index406 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
        (_2254_v66)[_index406] = (new BigNumber((_dafny.Seq.Concat(_2164_v0, _2164_v0)).length)).minus(_module.__default.fm18(globalState));
        let _index407 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
        (_2254_v66)[_index407] = (_2254_v66)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length))];
        if (true) {
          (globalState).f10 = _2252_v64;
          let _index408 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          (_2254_v66)[_index408] = ((_2167_v2).f37).minus((_2167_v2).f37);
          let _index409 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          (_2254_v66)[_index409] = (_this).f25;
          let _2255_v67;
          _2255_v67 = _dafny.Set.fromElements(_2222_v48, _2222_v48, _dafny.MultiSet.fromElements((_2167_v2).f37), _2222_v48);
          let _index410 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_2228_v51).length));
          (_2228_v51)[_index410] = (p0).isLessThanOrEqualTo(new BigNumber((_2255_v67).length));
          let _index411 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_2228_v51).length));
          (_2228_v51)[_index411] = !(((_2167_v2).f37).isLessThanOrEqualTo(_2220_v46));
          let _2256_v68;
          let _nw361 = new _module.C0();
          _nw361.__ctor();
          _2256_v68 = _nw361;
        } else {
          let _2257_v69;
          _2257_v69 = _module.D4.create_DC11(_2224_v50);
          let _2258_v70;
          _2258_v70 = _dafny.MultiSet.fromElements(_2257_v69);
          let _2259_v71;
          _2259_v71 = _dafny.Seq.of(_2258_v70);
          _2258_v70 = (_2259_v71)[_module.__default.safeIndex(_2220_v46, new BigNumber((_2259_v71).length))];
          _2248_cf50 = (_2167_v2).f38;
          let _2260_v72;
          _2260_v72 = new _dafny.CodePoint('s'.codePointAt(0));
          let _index412 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          let _index413 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          let _rhs394 = ((_2247_cf51).minus(_2247_cf51)).plus((_this).f25);
          let _rhs395 = _2260_v72;
          let _rhs396 = (new BigNumber((_module.__default.fm15(globalState)).cardinality())).plus((_2231_v53)[_module.__default.safeIndex((_2167_v2).f37, new BigNumber((_2231_v53).length))]);
          let _lhs304 = _2254_v66;
          let _lhs305 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          let _lhs306 = globalState;
          let _lhs307 = _2254_v66;
          let _lhs308 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          _lhs304[_lhs305] = _rhs394;
          _lhs306.f23 = _rhs395;
          _lhs307[_lhs308] = _rhs396;
          let _2261_v74;
          _2261_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2220_v46,(_2167_v2).f38);
          let _2262_v75;
          _2262_v75 = _dafny.Set.fromElements(function () {
            let _coll79 = new _dafny.Map();
            for (const _compr_79 of _dafny.IntegerRange(new BigNumber(470), new BigNumber(-568))) {
              let _2263_v73 = _compr_79;
              if (((new BigNumber(470)).isLessThanOrEqualTo(_2263_v73)) && ((_2263_v73).isLessThan(new BigNumber(-568)))) {
                _coll79.push([_module.__default.safeDivisionInt(_2263_v73, new BigNumber((_dafny.Seq.of(_2248_cf50)).length)),(_2167_v2).f38]);
              }
            }
            return _coll79;
          }(), _dafny.Map.Empty.slice().updateUnsafe((_2254_v66)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length))],(_this).f32), _2261_v74);
          let _index414 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          let _rhs397 = (_2262_v75).Difference(_2262_v75);
          let _rhs398 = (_2167_v2).f37;
          let _lhs309 = _2254_v66;
          let _lhs310 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          _2262_v75 = _rhs397;
          _lhs309[_lhs310] = _rhs398;
          let _index415 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2254_v66).length));
          (_2254_v66)[_index415] = (_2220_v46).plus(((_this).f25).plus(_2220_v46));
        }
      } else if (_source28.is_DC31) {
        let _2264___mcc_h5 = (_source28).cf46;
        let _2265_cf46 = _2264___mcc_h5;
        _2228_v51 = _2228_v51;
        let _2266_v77;
        _2266_v77 = _dafny.Seq.UnicodeFromString("erjupre");
        let _2267_v78;
        _2267_v78 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((function () {
          let _coll80 = new _dafny.Map();
          for (const _compr_80 of _dafny.IntegerRange(new BigNumber(329), new BigNumber(650))) {
            let _2268_v76 = _compr_80;
            if (((new BigNumber(329)).isLessThanOrEqualTo(_2268_v76)) && ((_2268_v76).isLessThan(new BigNumber(650)))) {
              _coll80.push([(_2268_v76).minus((_2167_v2).f37),p0]);
            }
          }
          return _coll80;
        }()).length), (_2167_v2).f37)),_dafny.areEqual(_2266_v77, _dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_2269_i5) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })));
        _2267_v78 = _2267_v78;
        _2220_v46 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        let _2270_v79;
        _2270_v79 = _module.D5.create_DC15((_2222_v48).Difference(_2222_v48));
        _2270_v79 = _2270_v79;
      } else {
        let _2271___mcc_h6 = (_source28).cf52;
        let _2272_cf52 = _2271___mcc_h6;
        let _2273_v80;
        let _init79 = function (_2274_i6) {
          return _module.__default.safeDivisionInt(_2274_i6, new BigNumber(175));
        };
        let _nw362 = Array((new BigNumber(17)).toNumber());
        for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw362.length); _i0_79++) {
          _nw362[_i0_79] = _init79(new BigNumber(_i0_79));
        }
        _2273_v80 = _nw362;
        let _index416 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_2228_v51).length));
        (_2228_v51)[_index416] = (_2167_v2).f38;
        let _index417 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
        (_2273_v80)[_index417] = (_this).f25;
        let _index418 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_2228_v51).length));
        let _index419 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
        let _rhs399 = (((_2164_v0)[_module.__default.safeIndex((_2167_v2).f37, new BigNumber((_2164_v0).length))]) ? (_2273_v80) : (_2273_v80));
        let _rhs400 = (_2167_v2).f38;
        let _rhs401 = _module.__default.safeDivisionInt((_2220_v46).multipliedBy(_2220_v46), p0);
        let _lhs311 = _2228_v51;
        let _lhs312 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_2228_v51).length));
        let _lhs313 = _2273_v80;
        let _lhs314 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
        _2273_v80 = _rhs399;
        _lhs311[_lhs312] = _rhs400;
        _lhs313[_lhs314] = _rhs401;
        let _index420 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
        (_2273_v80)[_index420] = _module.__default.safeModuloInt((_2167_v2).f37, p0);
        let _2275_v81;
        let _nw363 = Array((new BigNumber(4)).toNumber()).fill(false);
        _2275_v81 = _nw363;
        let _2276_v82;
        _2276_v82 = _module.D6.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_2275_v81,(_2273_v80)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length))]));
        let _source29 = _2276_v82;
        if (_source29.is_DC19) {
          let _2277___mcc_h7 = (_source29).cf29;
          let _2278___mcc_h8 = (_source29).cf30;
          let _2279_cf30 = _2278___mcc_h8;
          let _2280_cf29 = _2277___mcc_h7;
          let _rhs402 = (_2273_v80)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length))];
          let _rhs403 = (_dafny.ZERO).minus((_2167_v2).f37);
          _2220_v46 = _rhs402;
          _2279_cf30 = _rhs403;
          let _index421 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          (_2273_v80)[_index421] = (_dafny.ZERO).minus(((_2167_v2).f37).plus((new BigNumber(96)).minus(new BigNumber(-591))));
          _2279_cf30 = ((_2231_v53)[_module.__default.safeIndex(_2279_cf30, new BigNumber((_2231_v53).length))]).plus((new BigNumber(236)).plus(new BigNumber((_2164_v0).length)));
          _2220_v46 = _2220_v46;
        } else if (_source29.is_DC18) {
          let _2281___mcc_h9 = (_source29).cf28;
          let _2282_cf28 = _2281___mcc_h9;
          let _2283_v83;
          let _nw364 = Array((new BigNumber(16)).toNumber()).fill([]);
          _2283_v83 = _nw364;
          let _index422 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_2283_v83).length));
          (_2283_v83)[_index422] = _2273_v80;
          let _2284_v84;
          _2284_v84 = _dafny.Set.fromElements((_2167_v2).f38);
          let _2285_v85;
          _2285_v85 = _module.D7.create_DC21(_2273_v80);
          let _2286_v86;
          _2286_v86 = _dafny.MultiSet.fromElements((_this).f32, (_2167_v2).f38);
          let _index423 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_2283_v83).length));
          let _index424 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          let _index425 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          let _rhs404 = (_2285_v85).dtor_cf32;
          let _rhs405 = _2220_v46;
          let _rhs406 = new BigNumber((_dafny.Seq.Concat(_2164_v0, _2164_v0)).length);
          let _rhs407 = _2284_v84;
          let _rhs408 = (_module.__default.safeModuloInt(new BigNumber((_2286_v86).cardinality()), (_2167_v2).f37)).multipliedBy(new BigNumber(663));
          let _lhs315 = _2283_v83;
          let _lhs316 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_2283_v83).length));
          let _lhs317 = _2273_v80;
          let _lhs318 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          let _lhs319 = _2273_v80;
          let _lhs320 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          _lhs315[_lhs316] = _rhs404;
          _2220_v46 = _rhs405;
          _lhs317[_lhs318] = _rhs406;
          _2284_v84 = _rhs407;
          _lhs319[_lhs320] = _rhs408;
          _2220_v46 = (_dafny.ZERO).minus(p0);
          let _index426 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          (_2273_v80)[_index426] = (_2167_v2).f37;
          let _index427 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length));
          (_2273_v80)[_index427] = (_this).f25;
        } else {
          let _2287___mcc_h10 = (_source29).cf31;
          let _2288_cf31 = _2287___mcc_h10;
          _2220_v46 = (((_2167_v2).f37).plus((_2167_v2).f37)).multipliedBy((_2167_v2).f37);
          _2220_v46 = new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-635), (_this).f25)).cardinality());
          let _2289_v87;
          _2289_v87 = _dafny.Set.fromElements(_module.__default.fm39((_this).f32, globalState));
          let _2290_v88;
          _2290_v88 = new _dafny.CodePoint('k'.codePointAt(0));
          let _2291_v89;
          _2291_v89 = _dafny.Map.Empty.slice().updateUnsafe((_2167_v2).f38,_2290_v88);
          _2289_v87 = _module.__default.fm64(_module.__default.fm44(!((_2167_v2).f38), (_2167_v2).f38, (_2167_v2).f38, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2167_v2).f38,new BigNumber((_2291_v89).length))).length), globalState), (_2167_v2).f38, _2290_v88, (_2167_v2).f38, globalState);
          let _2292_v90;
          _2292_v90 = _dafny.Set.fromElements((_this).f32);
          let _2293_v91;
          _2293_v91 = _dafny.Map.Empty.slice().updateUnsafe((_2292_v90).Intersect(_dafny.Set.fromElements((_this).f32, (_2167_v2).f38, (_2167_v2).f38)),_2275_v81);
          let _2294_v92;
          _2294_v92 = _dafny.Map.Empty.slice().updateUnsafe((_this).f33,_2293_v91);
          let _2295_v93;
          _2295_v93 = _dafny.Seq.of(_2275_v81);
          let _2296_v94;
          let _nw365 = new _module.C6();
          _nw365.__ctor((_2273_v80)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_2273_v80).length))], (_this).f26);
          _2296_v94 = _nw365;
          let _2297_v95;
          _2297_v95 = _dafny.Map.Empty.slice().updateUnsafe(_2296_v94,(_this).f33);
          _2293_v91 = (_2293_v91).Merge((((_2294_v92).contains((_2167_v2).f38)) ? ((_2294_v92).get((_2167_v2).f38)) : (_dafny.Map.Empty.slice().updateUnsafe(_2292_v90,(_2295_v93)[_module.__default.safeIndex(new BigNumber((_2297_v95).length), new BigNumber((_2295_v93).length))]))));
        }
        let _2298_v96;
        _2298_v96 = _dafny.Seq.UnicodeFromString("tjlljdwo");
        let _2299_v97;
        _2299_v97 = _dafny.Set.fromElements(_2298_v96, _2298_v96);
        r2 = (_2299_v97).IsSubsetOf(_2299_v97);
      }
      let _2300_i7;
      _2300_i7 = _dafny.ZERO;
      L16: {
        while ((_this).f32) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2300_i7)) {
              break L16;
            }
            _2300_i7 = (_2300_i7).plus(_dafny.ONE);
            _2220_v46 = _2220_v46;
            let _2301_v98;
            let _nw366 = Array((new BigNumber(14)).toNumber()).fill(_module.D7.Default());
            _2301_v98 = _nw366;
            let _index428 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2301_v98).length));
            (_2301_v98)[_index428] = _module.D7.create_DC22((_2167_v2).f37, (_this).f25);
            let _2302_v99;
            _2302_v99 = _module.D7.create_DC22((_2167_v2).f37, new BigNumber(649));
            let _index429 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2301_v98).length));
            (_2301_v98)[_index429] = ((_module.__default.fm0(new BigNumber(392), globalState)) ? (_2302_v99) : (_2302_v99));
            _2220_v46 = _module.__default.safeModuloInt(new BigNumber((((_2230_v52).update((_2167_v2).f38, p0)).Merge(_2230_v52)).length), new BigNumber((_dafny.Seq.of(_2231_v53, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-814)), function (_2303_i8) {
              return (_this).f25;
            }), _dafny.Seq.of((_dafny.ZERO).minus((_this).f25)))).length));
            let _2304_v100;
            let _init80 = ((_2305_v2) => function (_2306_i9) {
              return (_2306_i9).multipliedBy(new BigNumber((_dafny.Seq.of((_2305_v2).f38)).length));
            })(_2167_v2);
            let _nw367 = Array((new BigNumber(28)).toNumber());
            for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw367.length); _i0_80++) {
              _nw367[_i0_80] = _init80(new BigNumber(_i0_80));
            }
            _2304_v100 = _nw367;
            let _index430 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_2304_v100).length));
            (_2304_v100)[_index430] = new BigNumber((_dafny.Seq.UnicodeFromString("w")).length);
            let _index431 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_2304_v100).length));
            (_2304_v100)[_index431] = (_2167_v2).f37;
          }
        }
      }
      let _2307_v101;
      _2307_v101 = _dafny.Map.Empty.slice().updateUnsafe(_2228_v51,(_this).f32);
      r0 = (((_2307_v101).contains(_2228_v51)) ? ((_2307_v101).get(_2228_v51)) : ((_this).f33));
      let _2308_v102;
      _2308_v102 = _dafny.Map.Empty.slice().updateUnsafe((_2167_v2).f38,(_this).f33);
      r1 = _2308_v102;
      let _2309_v103;
      _2309_v103 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2167_v2).f37),(_this).f32);
      let _2310_v104;
      _2310_v104 = _2309_v103;
      let _2311_v105;
      _2311_v105 = _dafny.Set.fromElements(_2310_v104);
      r2 = (_2311_v105).IsSubsetOf(_2311_v105);
      r3 = _dafny.Seq.Concat(_2164_v0, _dafny.Seq.Concat(_2164_v0, _dafny.Seq.update(_2164_v0, _module.__default.safeIndex((_2167_v2).f37, new BigNumber((_2164_v0).length)), (_this).f33)));
      return [r0, r1, r2, r3];
    }
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this.f30 = [];
      this._f31 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f30, f31) {
      let _this = this;
      (_this).f30 = f30;
      (_this)._f31 = f31;
      return;
    }
    fm11(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(((_this).f31).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(326))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("ce")).length));
    };
    fm12(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this).f31);
    };
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _2312_v0;
      _2312_v0 = true;
      (globalState).f12 = _2312_v0;
      let _2313_v1;
      let _nw368 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2313_v1 = _nw368;
      let _2314_v2;
      let _nw369 = new _module.C4();
      _nw369.__ctor(_2313_v1, p3, _2313_v1);
      _2314_v2 = _nw369;
      let _2315_v3;
      _2315_v3 = new BigNumber(717);
      _2315_v3 = new BigNumber(-527);
      let _2316_v4;
      _2316_v4 = _dafny.Set.fromElements(_2315_v3);
      let _2317_i0;
      _2317_i0 = _dafny.ZERO;
      L17: {
        while (!((_module.__default.safeDivisionInt(_2315_v3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2316_v4,_2315_v3)).length))).isLessThan((_this).f31))) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2317_i0)) {
              break L17;
            }
            _2317_i0 = (_2317_i0).plus(_dafny.ONE);
            (globalState).f12 = _module.__default.fm0(_2315_v3, globalState);
            let _2318_v5;
            _2318_v5 = _dafny.Map.Empty.slice().updateUnsafe(((_2312_v0) ? (p2) : (p2)),_2312_v0);
            let _2319_v6;
            _2319_v6 = _module.D10.create_DC28(_2318_v5);
            _2318_v5 = (_2319_v6).dtor_cf41;
            let _2320_v7;
            _2320_v7 = _dafny.Seq.of(_this.f30);
            let _2321_v8;
            let _out40;
            _out40 = (_2314_v2).m3(_dafny.Seq.Concat(_2320_v7, _2320_v7), _2312_v0, globalState);
            _2321_v8 = _out40;
            let _2322_v9;
            let _nw370 = new _module.C1();
            _nw370.__ctor(p3, (_2314_v2).f36);
            _2322_v9 = _nw370;
            _2322_v9 = _2322_v9;
          }
        }
      }
      let _2323_v10;
      _2323_v10 = _dafny.Seq.UnicodeFromString("bwksjhjgk");
      let _2324_v11;
      _2324_v11 = _dafny.Seq.of(_2323_v10, _2323_v10, _2323_v10, _2323_v10);
      let _2325_v12;
      _2325_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f30),_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lv"), (_2324_v11)[_module.__default.safeIndex(_2315_v3, new BigNumber((_2324_v11).length))]));
      let _2326_v13;
      _2326_v13 = _dafny.MultiSet.fromElements(_this.f30);
      let _2327_v14;
      _2327_v14 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("ntjlswu"));
      let _2328_v15;
      _2328_v15 = _dafny.MultiSet.fromElements((((_2327_v14).contains(p2)) ? ((_2327_v14).get(p2)) : (_2323_v10)), _dafny.Seq.update(_dafny.Seq.update(_2323_v10, _module.__default.safeIndex(p3, new BigNumber((_2323_v10).length)), p2), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.update(_2323_v10, _module.__default.safeIndex(p3, new BigNumber((_2323_v10).length)), p2)).length)), p2));
      (globalState).f16 = (((_2325_v12).contains(_2326_v13)) ? ((_2325_v12).get(_2326_v13)) : (_2328_v15));
      let _hi12 = (_dafny.ZERO).minus((_this).fm11(function () {
        let _coll81 = new _dafny.Set();
        for (const _compr_81 of (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_2323_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-153)), ((_2329_p2) => function (_2330_i2) {
          return _2329_p2;
        })(p2)))).cardinality())), p3)).Elements) {
          let _2331_v16 = _compr_81;
          if ((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_2323_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-153)), ((_2332_p2) => function (_2330_i2) {
            return _2332_p2;
          })(p2)))).cardinality())), p3)).contains(_2331_v16)) {
            _coll81.add((_2331_v16).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), function (_2333_i3) {
              return _dafny.Seq.UnicodeFromString("k");
            })).length)));
          }
        }
        return _coll81;
      }(), p2, globalState));
      for (let _2334_i1 = p3; _2334_i1.isLessThan(_hi12); _2334_i1 = _2334_i1.plus(_dafny.ONE)) {
        let _init81 = ((_2335_i1) => function (_2336_i4) {
          return (_2336_i4).minus(_2335_i1);
        })(_2334_i1);
        let _nw371 = Array((new BigNumber(4)).toNumber());
        for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw371.length); _i0_81++) {
          _nw371[_i0_81] = _init81(new BigNumber(_i0_81));
        }
        (_this).f30 = _nw371;
        let _2337_v17;
        let _nw372 = new _module.C2();
        _nw372.__ctor();
        _2337_v17 = _nw372;
        _2315_v3 = _2315_v3;
        (globalState).f12 = _dafny.areEqual(p1, p1);
      }
      r0 = _2312_v0;
      r1 = _2312_v0;
      return [r0, r1];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
      this._f29 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f29, f25, f26) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _2338_i0;
      _2338_i0 = _dafny.ZERO;
      L18: {
        while (p2) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2338_i0)) {
              break L18;
            }
            _2338_i0 = (_2338_i0).plus(_dafny.ONE);
            (globalState).f23 = (_this).f29;
            let _2339_v0;
            let _nw373 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
            _2339_v0 = _nw373;
            let _2340_v1;
            _2340_v1 = _dafny.Seq.UnicodeFromString("nnkgb");
            let _2341_v2;
            _2341_v2 = _dafny.Seq.of(_2340_v1);
            let _index432 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2339_v0).length));
            (_2339_v0)[_index432] = _2341_v2;
            let _index433 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
            (p3)[_index433] = _module.__default.safeModuloInt((_this).f25, (_this).f25);
            let _index434 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2339_v0).length));
            let _index435 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
            let _rhs409 = _2341_v2;
            let _rhs410 = (((p2) ? ((_this).f25) : ((_this).f25))).plus(new BigNumber(-425));
            let _lhs321 = _2339_v0;
            let _lhs322 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2339_v0).length));
            let _lhs323 = p3;
            let _lhs324 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
            _lhs321[_lhs322] = _rhs409;
            _lhs323[_lhs324] = _rhs410;
            if (p1) {
              let _2342_v3;
              let _nw374 = new _module.C0();
              _nw374.__ctor();
              _2342_v3 = _nw374;
              let _2343_v4;
              let _init82 = ((_2344_p2, _2345_p3) => function (_2346_i1) {
                return (_dafny.Set.fromElements((_this).f25)).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_module.D10.create_DC29(_2344_p2, _2344_p2, (_2345_p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2345_p3).length))])).dtor_cf44, (_2345_p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2345_p3).length))])).cardinality())));
              })(p2, p3);
              let _nw375 = Array((new BigNumber(21)).toNumber());
              for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw375.length); _i0_82++) {
                _nw375[_i0_82] = _init82(new BigNumber(_i0_82));
              }
              _2343_v4 = _nw375;
              _2343_v4 = _2343_v4;
              let _index436 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
              (p3)[_index436] = (_this).f25;
              let _2347_v5;
              let _nw376 = new _module.C5();
              _nw376.__ctor(false);
              _2347_v5 = _nw376;
              (globalState).f23 = (_this).f29;
            } else {
              let _2348_v6;
              _2348_v6 = _dafny.MultiSet.fromElements(p1, p2);
              let _2349_v7;
              _2349_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
              let _2350_v8;
              _2350_v8 = _module.D0.create_DC3(new BigNumber(((_2348_v6).update(p1, _module.__default.abs(new BigNumber((_2349_v7).length)))).cardinality()), p2);
              let _index437 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
              (p3)[_index437] = (_dafny.ZERO).minus((((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))]).multipliedBy((_2350_v8).dtor_cf4)).minus((_this).f25));
              let _index438 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
              (p3)[_index438] = ((true) ? ((_dafny.ZERO).minus((_this).f25)) : ((_this).f25));
              let _2351_v9;
              _2351_v9 = _dafny.Seq.of((_this).f25);
              let _index439 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
              let _rhs411 = !(!((_this).f25).isEqualTo((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))]));
              let _rhs412 = ((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))]).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), ((_2352_v1) => function (_2353_i2) {
                return new BigNumber((_2352_v1).length);
              })(_2340_v1))).length));
              let _rhs413 = ((new BigNumber((_2351_v9).length)).plus((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))])).plus((_this).f25);
              let _rhs414 = ((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))]).isLessThanOrEqualTo((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))]);
              let _lhs325 = globalState;
              let _lhs326 = p3;
              let _lhs327 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length));
              let _lhs328 = globalState;
              _lhs325.f12 = _rhs411;
              _lhs326[_lhs327] = _rhs412;
              r0 = _rhs413;
              _lhs328.f12 = _rhs414;
              (globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2351_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), ((_2354_p3) => function (_2355_i3) {
                return (_2354_p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2354_p3).length))];
              })(p3))), _2351_v9);
              r1 = ((_this).f25).multipliedBy((_dafny.ZERO).minus((p3)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((p3).length))]));
            }
            (globalState).f10 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_2356_i4) {
              return (_this).f29;
            }), _dafny.Seq.UnicodeFromString("kmuer"));
          }
        }
      }
      r1 = (_this).f25;
      let _2357_v10;
      let _nw377 = new _module.C6();
      _nw377.__ctor((_this).f25, (_this).f26);
      _2357_v10 = _nw377;
      let _hi13 = ((_this).f25).minus((_this).f25);
      for (let _2358_i5 = (_this).f25; _2358_i5.isLessThan(_hi13); _2358_i5 = _2358_i5.plus(_dafny.ONE)) {
        let _2359_v11;
        _2359_v11 = _dafny.Seq.of((_this).f26, (_this).f26, (_this).f26, (_this).f26, (_this).f26);
        let _2360_v12;
        let _nw378 = new _module.C3();
        _nw378.__ctor(_2358_i5, p1, (_this).f25, (_2359_v11)[_module.__default.safeIndex((_this).f25, new BigNumber((_2359_v11).length))]);
        _2360_v12 = _nw378;
        r0 = (_this).f25;
        let _2361_v13;
        _2361_v13 = _dafny.Seq.of((_2360_v12).f38, !(true), _module.__default.fm0((_dafny.ZERO).minus((_this).f25), globalState));
        let _2362_v14;
        let _nw379 = Array((new BigNumber(25)).toNumber()).fill(false);
        _2362_v14 = _nw379;
        let _2363_v15;
        _2363_v15 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f25, globalState),_2362_v14);
        let _rhs415 = _module.__default.fm1(p1, globalState);
        let _rhs416 = (_this).f25;
        let _rhs417 = _module.__default.fm0((_2360_v12).f37, globalState);
        let _rhs418 = ((p2) ? (_2362_v14) : ((((_2363_v15).contains(p2)) ? ((_2363_v15).get(p2)) : (_2362_v14))));
        let _rhs419 = (_2360_v12).f38;
        let _lhs329 = globalState;
        let _lhs330 = globalState;
        _2361_v13 = _rhs415;
        r0 = _rhs416;
        _lhs329.f12 = _rhs417;
        _2362_v14 = _rhs418;
        _lhs330.f12 = _rhs419;
        let _2364_v16;
        _2364_v16 = _dafny.Seq.UnicodeFromString("tqn");
        let _2365_v17;
        _2365_v17 = _module.D10.create_DC29((_2360_v12).f38, p2, (_this).f25);
        let _2366_v18;
        _2366_v18 = _dafny.MultiSet.fromElements((_2365_v17).dtor_cf42);
        if ((_2366_v18).IsProperSubsetOf((_dafny.MultiSet.FromArray(_2361_v13)).Intersect(_module.__default.fm50((_2360_v12).f37, _2364_v16, p1, new BigNumber(-333), globalState)))) {
          let _2367_v19;
          _2367_v19 = _dafny.Seq.of((_2360_v12).f37, (_2360_v12).f37);
          let _2368_v20;
          _2368_v20 = _dafny.Seq.of(_2367_v19, _2367_v19);
          _2368_v20 = _2368_v20;
          _2362_v14 = _2362_v14;
          let _2369_v21;
          let _nw380 = new _module.C0();
          _nw380.__ctor();
          _2369_v21 = _nw380;
          (globalState).f12 = p2;
          (globalState).f12 = p2;
        } else {
          let _2370_v22;
          _2370_v22 = _dafny.Set.fromElements(p2, !(p2));
          let _2371_v23;
          _2371_v23 = _dafny.MultiSet.fromElements(_module.__default.fm20(p2, globalState), _2370_v22);
          _2371_v23 = _2371_v23;
          let _2372_v24;
          _2372_v24 = _dafny.Map.Empty.slice().updateUnsafe(p2,(p2) && (p2));
          _2372_v24 = (_2372_v24).update(false, (_2360_v12).f38);
          let _2373_v25;
          let _nw381 = Array((new BigNumber(18)).toNumber()).fill([]);
          _2373_v25 = _nw381;
          _2373_v25 = _2373_v25;
          (globalState).f12 = true;
          let _2374_v26;
          let _2375_v27;
          let _out41;
          let _out42;
          let _outcollector14 = (_2357_v10).m11(p3, (_2360_v12).f38, (_2363_v15).Merge(_2363_v15), globalState);
          _out41 = _outcollector14[0];
          _out42 = _outcollector14[1];
          _2374_v26 = _out41;
          _2375_v27 = _out42;
        }
      }
      if (p1) {
        (globalState).f12 = p1;
        r0 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.fm44(p1, p2, false, (_this).f25, globalState), (_this).f25))).plus((_this).f25);
        (globalState).f12 = !(p2);
        if (p2) {
          let _2376_v28;
          _2376_v28 = _module.D13.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(p2,p2), p2, (_this).f25);
          let _2377_v29;
          _2377_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _2378_v30;
          _2378_v30 = _dafny.Seq.of(p2);
          let _pat_let_tv43 = _2377_v29;
          let _2379_v31;
          let _nw382 = Array((new BigNumber(13)).toNumber());
          _nw382[(_dafny.ZERO).toNumber()] = _2376_v28;
          _nw382[(_dafny.ONE).toNumber()] = _module.D13.create_DC36(_2377_v29, false, new BigNumber((_2378_v30).length));
          _nw382[(new BigNumber(2)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(3)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(4)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(5)).toNumber()] = function (_pat_let54_0) {
            return function (_2380_dt__update__tmp_h0) {
              return function (_pat_let55_0) {
                return function (_2381_dt__update_hcf54_h0) {
                  return _module.D13.create_DC36(_2381_dt__update_hcf54_h0, (_2380_dt__update__tmp_h0).dtor_cf55, (_2380_dt__update__tmp_h0).dtor_cf56);
                }(_pat_let55_0);
              }(_pat_let_tv43);
            }(_pat_let54_0);
          }(_module.D13.create_DC36(_2377_v29, p1, (_this).f25));
          _nw382[(new BigNumber(6)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(7)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(8)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(9)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(10)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(11)).toNumber()] = _2376_v28;
          _nw382[(new BigNumber(12)).toNumber()] = _module.D13.create_DC36(_2377_v29, p2, (_dafny.ZERO).minus((_this).f25));
          _2379_v31 = _nw382;
          let _index440 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2379_v31).length));
          (_2379_v31)[_index440] = _2376_v28;
          let _index441 = _module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f26).length));
          ((_this).f26)[_index441] = _dafny.Seq.of(_module.__default.fm39(p2, globalState), (_this).f29, (_this).f29, (_this).f29);
          let _2382_v32;
          _2382_v32 = _dafny.Seq.of((_this).f29);
          let _index442 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2379_v31).length));
          let _index443 = _module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f26).length));
          let _rhs420 = _2376_v28;
          let _rhs421 = _2382_v32;
          let _rhs422 = _dafny.areEqual(_2382_v32, _dafny.Seq.UnicodeFromString("pbgxd"));
          let _lhs331 = _2379_v31;
          let _lhs332 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2379_v31).length));
          let _lhs333 = (_this).f26;
          let _lhs334 = _module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f26).length));
          let _lhs335 = globalState;
          _lhs331[_lhs332] = _rhs420;
          _lhs333[_lhs334] = _rhs421;
          _lhs335.f12 = _rhs422;
          let _2383_v33;
          let _nw383 = Array((new BigNumber(16)).toNumber()).fill(false);
          _2383_v33 = _nw383;
          _2383_v33 = _2383_v33;
          let _2384_v34;
          _2384_v34 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f29), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.of((_this).f29)).length)), new _dafny.CodePoint('y'.codePointAt(0)))).length));
          let _2385_v35;
          _2385_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2384_v34).length),_2382_v32);
          let _index444 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((p3).length));
          (p3)[_index444] = new BigNumber((_dafny.Seq.Concat((((_2385_v35).contains((_this).f25)) ? ((_2385_v35).get((_this).f25)) : (_dafny.Seq.UnicodeFromString("v"))), _2382_v32)).length);
          let _index445 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((p3).length));
          (p3)[_index445] = (_this).f25;
          let _2386_v36;
          let _init83 = ((_2387_v30, _2388_p1, _2389_p2) => function (_2390_i6) {
            return (_dafny.MultiSet.FromArray(_2387_v30)).Union(_dafny.MultiSet.fromElements(!(_2388_p1), _2389_p2));
          })(_2378_v30, p1, p2);
          let _nw384 = Array((new BigNumber(25)).toNumber());
          for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw384.length); _i0_83++) {
            _nw384[_i0_83] = _init83(new BigNumber(_i0_83));
          }
          _2386_v36 = _nw384;
          let _2391_v37;
          _2391_v37 = _module.D12.create_DC32(p2);
          let _index446 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_2386_v36).length));
          (_2386_v36)[_index446] = _dafny.MultiSet.fromElements((_2391_v37).dtor_cf47);
          let _2392_v38;
          _2392_v38 = _dafny.MultiSet.fromElements(!(p1));
          let _2393_v39;
          _2393_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_2392_v38);
          let _2394_v40;
          _2394_v40 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(((_this).f26)[_module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f26).length))], _module.__default.safeIndex((p3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((p3).length))], new BigNumber((((_this).f26)[_module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f26).length))]).length)), (_this).f29)).length));
          let _index447 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_2386_v36).length));
          (_2386_v36)[_index447] = (((_2393_v39).contains((_dafny.ZERO).minus(((p3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((p3).length))]).multipliedBy((_2394_v40)[_module.__default.safeIndex((_this).f25, new BigNumber((_2394_v40).length))])))) ? ((_2393_v39).get((_dafny.ZERO).minus(((p3)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((p3).length))]).multipliedBy((_2394_v40)[_module.__default.safeIndex((_this).f25, new BigNumber((_2394_v40).length))])))) : (_2392_v38));
          (globalState).f12 = !_dafny.areEqual(_dafny.Seq.Concat(_2378_v30, _2378_v30), _dafny.Seq.Concat(_dafny.Seq.of(p1), _2378_v30));
        } else {
          let _2395_v41;
          let _nw385 = Array((new BigNumber(17)).toNumber()).fill(false);
          _2395_v41 = _nw385;
          let _2396_v42;
          _2396_v42 = _dafny.Set.fromElements(_2395_v41);
          let _2397_v43;
          _2397_v43 = _module.D3.create_DC9(_2396_v42);
          let _2398_v44;
          _2398_v44 = _dafny.MultiSet.fromElements(_2397_v43, _2397_v43, _2397_v43);
          let _2399_v45;
          _2399_v45 = _dafny.Seq.of(_2398_v44, (_2398_v44).Intersect(_2398_v44), _2398_v44, _2398_v44, (_2398_v44).Difference(_dafny.MultiSet.fromElements(_2397_v43)));
          let _2400_v46;
          _2400_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p2);
          let _2401_v47;
          _2401_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber((_2400_v46).length));
          let _2402_v48;
          _2402_v48 = _dafny.Seq.of(new BigNumber((_2401_v47).length));
          let _2403_v49;
          _2403_v49 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2402_v48);
          let _rhs423 = _2399_v45;
          let _rhs424 = p1;
          let _rhs425 = ((_this).f25).multipliedBy(new BigNumber(((((_2403_v49).contains(!(!(p1)))) ? ((_2403_v49).get(!(!(p1)))) : (_2402_v48))).length));
          let _lhs336 = globalState;
          _2399_v45 = _rhs423;
          _lhs336.f12 = _rhs424;
          r0 = _rhs425;
          let _2404_v50;
          let _nw386 = Array((new BigNumber(20)).toNumber());
          _nw386[(_dafny.ZERO).toNumber()] = p3;
          _nw386[(_dafny.ONE).toNumber()] = p3;
          _nw386[(new BigNumber(2)).toNumber()] = p3;
          _nw386[(new BigNumber(3)).toNumber()] = p3;
          _nw386[(new BigNumber(4)).toNumber()] = p3;
          _nw386[(new BigNumber(5)).toNumber()] = p3;
          _nw386[(new BigNumber(6)).toNumber()] = p3;
          _nw386[(new BigNumber(7)).toNumber()] = p3;
          _nw386[(new BigNumber(8)).toNumber()] = p3;
          _nw386[(new BigNumber(9)).toNumber()] = p3;
          _nw386[(new BigNumber(10)).toNumber()] = p3;
          _nw386[(new BigNumber(11)).toNumber()] = p3;
          _nw386[(new BigNumber(12)).toNumber()] = p3;
          _nw386[(new BigNumber(13)).toNumber()] = ((p2) ? (p3) : (p3));
          _nw386[(new BigNumber(14)).toNumber()] = p3;
          _nw386[(new BigNumber(15)).toNumber()] = p3;
          _nw386[(new BigNumber(16)).toNumber()] = p3;
          _nw386[(new BigNumber(17)).toNumber()] = p3;
          _nw386[(new BigNumber(18)).toNumber()] = p3;
          _nw386[(new BigNumber(19)).toNumber()] = p3;
          _2404_v50 = _nw386;
          _2404_v50 = _2404_v50;
          r0 = new BigNumber(551);
          (globalState).f12 = !((_this).f25).isEqualTo((_dafny.ZERO).minus((_this).f25));
          let _2405_v51;
          _2405_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_2402_v48);
          let _2406_v52;
          _2406_v52 = _dafny.MultiSet.fromElements((_this).f25);
          (globalState).f12 = ((p2) ? (!(p2)) : ((_dafny.MultiSet.FromArray((((_2405_v51).contains(new BigNumber((_2402_v48).length))) ? ((_2405_v51).get(new BigNumber((_2402_v48).length))) : (_2402_v48)))).IsSubsetOf(_2406_v52)));
        }
        let _index448 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((p3).length));
        (p3)[_index448] = _module.__default.safeModuloInt((_this).f25, (_this).f25);
        let _2407_v53;
        _2407_v53 = _dafny.MultiSet.fromElements(p1);
        let _2408_v54;
        _2408_v54 = _dafny.Seq.of(p1, !(p2));
        let _2409_v55;
        _2409_v55 = _module.D0.create_DC1(!(p2), p1);
        let _2410_v56;
        _2410_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2409_v55,new BigNumber(918));
        let _2411_v57;
        _2411_v57 = _dafny.Seq.UnicodeFromString("jveby");
        let _2412_v58;
        _2412_v58 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _2413_v59;
        _2413_v59 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus((_this).f25));
        let _2414_v60;
        let _nw387 = Array((new BigNumber(21)).toNumber());
        _nw387[(_dafny.ZERO).toNumber()] = _2407_v53;
        _nw387[(_dafny.ONE).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(p2, _module.__default.fm0((_this).f25, globalState)));
        _nw387[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray(_2408_v54);
        _nw387[(new BigNumber(3)).toNumber()] = (_2407_v53).update(p1, _module.__default.abs((_this).f25));
        _nw387[(new BigNumber(4)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(5)).toNumber()] = _module.__default.fm50(new BigNumber(((_2410_v56).update(_2409_v55, (_this).f25)).length), _2411_v57, p1, (_this).f25, globalState);
        _nw387[(new BigNumber(6)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(7)).toNumber()] = (_2407_v53).Difference(_dafny.MultiSet.FromArray(_2408_v54));
        _nw387[(new BigNumber(8)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(9)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(10)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements(p2)).update(p1, _module.__default.abs((_this).f25));
        _nw387[(new BigNumber(12)).toNumber()] = (_2407_v53).update(p2, _module.__default.abs(new BigNumber((_2412_v58).length)));
        _nw387[(new BigNumber(13)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(14)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(15)).toNumber()] = _module.__default.fm50(new BigNumber((_2413_v59).length), _2411_v57, true, (_this).f25, globalState);
        _nw387[(new BigNumber(16)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(17)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(18)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(19)).toNumber()] = _2407_v53;
        _nw387[(new BigNumber(20)).toNumber()] = (_2407_v53).Union(_2407_v53);
        _2414_v60 = _nw387;
        let _index449 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2414_v60).length));
        (_2414_v60)[_index449] = _2407_v53;
        let _index450 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((p3).length));
        let _index451 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2414_v60).length));
        let _rhs426 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p2)).length), (_this).f25));
        let _rhs427 = (_module.__default.fm18(globalState)).multipliedBy(new BigNumber(-367));
        let _rhs428 = _2407_v53;
        let _lhs337 = p3;
        let _lhs338 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((p3).length));
        let _lhs339 = _2414_v60;
        let _lhs340 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_2414_v60).length));
        r1 = _rhs426;
        _lhs337[_lhs338] = _rhs427;
        _lhs339[_lhs340] = _rhs428;
      } else {
        let _index452 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length));
        (p3)[_index452] = ((_this).f25).multipliedBy(new BigNumber((_dafny.Seq.of((_this).f25, new BigNumber(243))).length));
        let _2415_v61;
        let _init84 = ((_2416_p1, _2417_p2) => function (_2418_i7) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_2416_p1,_dafny.Seq.Create(_module.__default.abs(new BigNumber(163)), function (_2419_i8) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2417_p2,_dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), function (_2420_i9) {
            return (_this).f29;
          })));
        })(p1, p2);
        let _nw388 = Array((new BigNumber(14)).toNumber());
        for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw388.length); _i0_84++) {
          _nw388[_i0_84] = _init84(new BigNumber(_i0_84));
        }
        _2415_v61 = _nw388;
        let _2421_v62;
        _2421_v62 = _dafny.Seq.UnicodeFromString("is");
        let _2422_v63;
        _2422_v63 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2421_v62);
        let _index453 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2415_v61).length));
        (_2415_v61)[_index453] = _2422_v63;
        let _2423_v64;
        _2423_v64 = _dafny.MultiSet.fromElements(p1);
        let _2424_v65;
        _2424_v65 = _dafny.Set.fromElements(_module.__default.fm44(p1, p1, p2, (_this).f25, globalState), new BigNumber((_2423_v64).cardinality()), _module.__default.safeModuloInt((_this).f25, new BigNumber((_dafny.Seq.of((_this).f25, (_this).f25)).length)));
        let _index454 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length));
        let _index455 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2415_v61).length));
        let _rhs429 = new BigNumber((_2424_v65).length);
        let _rhs430 = _2422_v63;
        let _lhs341 = p3;
        let _lhs342 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length));
        let _lhs343 = _2415_v61;
        let _lhs344 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2415_v61).length));
        _lhs341[_lhs342] = _rhs429;
        _lhs343[_lhs344] = _rhs430;
        if (((p2) ? (p1) : (p2))) {
          let _2425_v66;
          _2425_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(77),_2421_v62);
          let _2426_v67;
          _2426_v67 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2425_v66);
          _2426_v67 = (_2426_v67).update(!(p2), (_2425_v66).Merge(_2425_v66));
          let _2427_v68;
          let _init85 = function (_2428_i10) {
            return (_this).f29;
          };
          let _nw389 = Array((new BigNumber(25)).toNumber());
          for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw389.length); _i0_85++) {
            _nw389[_i0_85] = _init85(new BigNumber(_i0_85));
          }
          _2427_v68 = _nw389;
          let _2429_v69;
          _2429_v69 = _module.D13.create_DC35(_2427_v68);
          let _2430_v70;
          _2430_v70 = _dafny.Set.fromElements(_2429_v69, _2429_v69);
          _2430_v70 = _2430_v70;
          let _index456 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length));
          (p3)[_index456] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25));
          (globalState).f12 = !(p2);
          let _2431_v71;
          _2431_v71 = _dafny.Seq.of(p1, p1, p2);
          let _2432_v72;
          _2432_v72 = _dafny.Seq.of(_2431_v71);
          _2432_v72 = _dafny.Seq.Concat(_2432_v72, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(598)), ((_2433_v71) => function (_2434_i11) {
            return (_module.D1.create_DC4(_2433_v71)).dtor_cf6;
          })(_2431_v71)), _2432_v72));
        } else {
          let _index457 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length));
          (p3)[_index457] = _module.__default.safeDivisionInt(new BigNumber((function () {
            let _coll82 = new _dafny.Map();
            for (const _compr_82 of _dafny.IntegerRange(new BigNumber(298), new BigNumber(891))) {
              let _2435_v73 = _compr_82;
              if (((new BigNumber(298)).isLessThanOrEqualTo(_2435_v73)) && ((_2435_v73).isLessThan(new BigNumber(891)))) {
                _coll82.push([(_2435_v73).plus((_this).f25),p1]);
              }
            }
            return _coll82;
          }()).length), new BigNumber(374));
          (globalState).f12 = p2;
          (globalState).f12 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), function (_2436_i12) {
            return (_this).f29;
          }), _dafny.Seq.update(_2421_v62, _module.__default.safeIndex((p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))], new BigNumber((_2421_v62).length)), (_this).f29)), _dafny.Seq.Concat(_2421_v62, _dafny.Seq.UnicodeFromString("xsanvurs")));
          let _2437_v74;
          _2437_v74 = _dafny.Seq.of(p2, p2);
          let _2438_v75;
          _2438_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_2421_v62);
          let _2439_v77;
          _2439_v77 = _dafny.Seq.of(function () {
            let _coll83 = new _dafny.Set();
            for (const _compr_83 of _dafny.IntegerRange(new BigNumber(362), new BigNumber(861))) {
              let _2440_v76 = _compr_83;
              if (((new BigNumber(362)).isLessThanOrEqualTo(_2440_v76)) && ((_2440_v76).isLessThan(new BigNumber(861)))) {
                _coll83.add((_2440_v76).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,false)).length)),new BigNumber(43))).length)));
              }
            }
            return _coll83;
          }(), _2424_v65);
          let _2441_v78;
          _2441_v78 = _dafny.Map.Empty.slice().updateUnsafe(_module.D18.create_DC44((_this).f25, p2, _2421_v62, p2, _2439_v77),(p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))]);
          let _2442_v80;
          _2442_v80 = _dafny.Set.fromElements(_module.D18.create_DC44((p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))], p1, _2421_v62, true, _2439_v77));
          let _2443_v81;
          _2443_v81 = _dafny.Map.Empty.slice().updateUnsafe(p2,(p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))]);
          let _2444_v82;
          _2444_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(489),_2443_v81);
          let _2445_v83;
          _2445_v83 = _module.D5.create_DC17(new BigNumber(((((_2444_v82).contains((_this).f25)) ? ((_2444_v82).get((_this).f25)) : (_2443_v81))).length), p1);
          let _2446_v84;
          let _nw390 = Array((new BigNumber(17)).toNumber());
          _nw390[(_dafny.ZERO).toNumber()] = _module.__default.fm0((p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))], globalState);
          _nw390[(_dafny.ONE).toNumber()] = p1;
          _nw390[(new BigNumber(2)).toNumber()] = (_2437_v74)[_module.__default.safeIndex(new BigNumber(((((_2438_v75).contains(new BigNumber(-237))) ? ((_2438_v75).get(new BigNumber(-237))) : (_2421_v62))).length), new BigNumber((_2437_v74).length))];
          _nw390[(new BigNumber(3)).toNumber()] = (_2442_v80).IsProperSubsetOf(function () {
            let _coll84 = new _dafny.Set();
            for (const _compr_84 of (_2441_v78).Keys.Elements) {
              let _2447_v79 = _compr_84;
              if ((_2441_v78).contains(_2447_v79)) {
                _coll84.add(_2447_v79);
              }
            }
            return _coll84;
          }());
          _nw390[(new BigNumber(4)).toNumber()] = (new BigNumber(848)).isLessThanOrEqualTo(new BigNumber(282));
          _nw390[(new BigNumber(5)).toNumber()] = (_2445_v83).dtor_cf27;
          _nw390[(new BigNumber(6)).toNumber()] = (_2437_v74)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_2437_v74).length))];
          _nw390[(new BigNumber(7)).toNumber()] = p1;
          _nw390[(new BigNumber(8)).toNumber()] = p2;
          _nw390[(new BigNumber(9)).toNumber()] = ((!(p1)) ? (p2) : (!(false)));
          _nw390[(new BigNumber(10)).toNumber()] = p2;
          _nw390[(new BigNumber(11)).toNumber()] = p1;
          _nw390[(new BigNumber(12)).toNumber()] = p1;
          _nw390[(new BigNumber(13)).toNumber()] = false;
          _nw390[(new BigNumber(14)).toNumber()] = !(p1);
          _nw390[(new BigNumber(15)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2437_v74, _2437_v74);
          _nw390[(new BigNumber(16)).toNumber()] = p2;
          _2446_v84 = _nw390;
          let _index458 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2446_v84).length));
          (_2446_v84)[_index458] = _dafny.areEqual(_2421_v62, _2421_v62);
          let _index459 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2446_v84).length));
          (_2446_v84)[_index459] = !(!_dafny.Seq.contains(_dafny.Seq.Concat(_2421_v62, _2421_v62), (_2421_v62)[_module.__default.safeIndex(new BigNumber(546), new BigNumber((_2421_v62).length))]));
          _2446_v84 = _2446_v84;
        }
        (globalState).f10 = _2421_v62;
        r0 = (p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))];
        let _2448_v85;
        let _nw391 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2448_v85 = _nw391;
        let _2449_v86;
        _2449_v86 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_2421_v62).length)),false);
        let _index460 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_2448_v85).length));
        (_2448_v85)[_index460] = _dafny.MultiSet.fromElements(_2449_v86);
        let _2450_v88;
        _2450_v88 = _dafny.MultiSet.fromElements(new BigNumber((_2421_v62).length), (p3)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((p3).length))]);
        let _2451_v89;
        _2451_v89 = _dafny.MultiSet.fromElements(function () {
          let _coll85 = new _dafny.Map();
          for (const _compr_85 of (_2450_v88).Elements) {
            let _2452_v87 = _compr_85;
            if ((_2450_v88).contains(_2452_v87)) {
              _coll85.push([(_2452_v87).minus((_this).f25),p2]);
            }
          }
          return _coll85;
        }(), _2449_v86, _2449_v86);
        let _index461 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_2448_v85).length));
        (_2448_v85)[_index461] = _2451_v89;
      }
      let _index462 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((p3).length));
      (p3)[_index462] = ((_this).f25).multipliedBy((_this).f25);
      let _index463 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((p3).length));
      (p3)[_index463] = (_this).f25;
      let _2453_v90;
      _2453_v90 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
      let _2454_v91;
      _2454_v91 = _dafny.Seq.UnicodeFromString("qinj");
      let _2455_v92;
      _2455_v92 = _dafny.Set.fromElements(new BigNumber(636));
      let _2456_v93;
      _2456_v93 = _module.D18.create_DC44(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(((_2453_v90).contains(p1)) ? ((_2453_v90).get(p1)) : (p2)))).length), true, _2454_v91, p2, _dafny.Seq.of(_2455_v92, _2455_v92));
      let _pat_let_tv44 = p3;
      let _pat_let_tv45 = p3;
      let _pat_let_tv46 = p2;
      let _pat_let_tv47 = p2;
      r0 = function (_source30) {
        if (_source30.is_DC44) {
          let _2457___mcc_h0 = (_source30).cf63;
          let _2458___mcc_h1 = (_source30).cf64;
          let _2459___mcc_h2 = (_source30).cf65;
          let _2460___mcc_h3 = (_source30).cf66;
          let _2461___mcc_h4 = (_source30).cf67;
          let _2462_cf67 = _2461___mcc_h4;
          let _2463_cf66 = _2460___mcc_h3;
          let _2464_cf65 = _2459___mcc_h2;
          let _2465_cf64 = _2458___mcc_h1;
          let _2466_cf63 = _2457___mcc_h0;
          return _2466_cf63;
        } else if (_source30.is_DC45) {
          return (_pat_let_tv45)[_module.__default.safeIndex(new BigNumber(251), new BigNumber((_pat_let_tv44).length))];
        } else {
          let _2467___mcc_h5 = (_source30).cf62;
          let _2468_cf62 = _2467___mcc_h5;
          return new BigNumber((_dafny.Set.fromElements(false, _pat_let_tv46, _pat_let_tv47)).length);
        }
      }(_2456_v93);
      r1 = (p3)[_module.__default.safeIndex(new BigNumber(251), new BigNumber((p3).length))];
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2469_i0;
      _2469_i0 = _dafny.ZERO;
      L19: {
        while (p1) {
          C19: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2469_i0)) {
              break L19;
            }
            _2469_i0 = (_2469_i0).plus(_dafny.ONE);
            (globalState).f23 = (_this).f29;
            let _2470_v0;
            _2470_v0 = _module.D0.create_DC2(p1);
            let _2471_v1;
            _2471_v1 = _dafny.MultiSet.fromElements(_2470_v0);
            let _2472_v2;
            _2472_v2 = _dafny.Seq.of(_2470_v0);
            let _2473_v3;
            _2473_v3 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_2472_v2), _2471_v1);
            let _2474_v4;
            _2474_v4 = _dafny.MultiSet.fromElements(p1);
            let _pat_let_tv48 = p1;
            let _2475_v5;
            let _nw392 = Array((new BigNumber(22)).toNumber());
            _nw392[(_dafny.ZERO).toNumber()] = _2471_v1;
            _nw392[(_dafny.ONE).toNumber()] = (_2471_v1).update(_2470_v0, _module.__default.abs((_this).f25));
            _nw392[(new BigNumber(2)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(3)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(4)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(5)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(6)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(7)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(8)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(9)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(function (_pat_let56_0) {
              return function (_2476_dt__update__tmp_h0) {
                return function (_pat_let57_0) {
                  return function (_2477_dt__update_hcf3_h0) {
                    return _module.D0.create_DC2(_2477_dt__update_hcf3_h0);
                  }(_pat_let57_0);
                }(_pat_let_tv48);
              }(_pat_let56_0);
            }(_2470_v0));
            _nw392[(new BigNumber(11)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(12)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(13)).toNumber()] = (_2473_v3)[_module.__default.safeIndex(new BigNumber(27), new BigNumber((_2473_v3).length))];
            _nw392[(new BigNumber(14)).toNumber()] = ((p1) ? (_2471_v1) : (_2471_v1));
            _nw392[(new BigNumber(15)).toNumber()] = (_2471_v1).Difference((_2471_v1).update(_2470_v0, _module.__default.abs((_this).f25)));
            _nw392[(new BigNumber(16)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(17)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(18)).toNumber()] = _2471_v1;
            _nw392[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements(_2470_v0, _2470_v0);
            _nw392[(new BigNumber(20)).toNumber()] = ((_dafny.MultiSet.FromArray(_2472_v2)).update(_2470_v0, _module.__default.abs((((_2474_v4).contains(p1)) ? ((_2474_v4).get(p1)) : ((_this).f25))))).Intersect(_2471_v1);
            _nw392[(new BigNumber(21)).toNumber()] = _2471_v1;
            _2475_v5 = _nw392;
            let _index464 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_2475_v5).length));
            (_2475_v5)[_index464] = _2471_v1;
            let _2478_v6;
            _2478_v6 = _dafny.Seq.UnicodeFromString("yupwkrstt");
            let _index465 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_2475_v5).length));
            let _rhs431 = _dafny.Seq.Concat(_2478_v6, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tmq"), _dafny.Seq.UnicodeFromString("in")));
            let _rhs432 = ((_this).f25).plus(((_this).f25).minus((((_2474_v4).contains(p1)) ? ((_2474_v4).get(p1)) : ((_this).f25))));
            let _rhs433 = (_2471_v1).update(_module.D0.create_DC2(p1), _module.__default.abs((_this).f25));
            let _lhs345 = globalState;
            let _lhs346 = _2475_v5;
            let _lhs347 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_2475_v5).length));
            _lhs345.f10 = _rhs431;
            r0 = _rhs432;
            _lhs346[_lhs347] = _rhs433;
            if (p1) {
              let _2479_v7;
              let _nw393 = new _module.C1();
              _nw393.__ctor((_this).f25, (_this).f26);
              _2479_v7 = _nw393;
              let _2480_v8;
              _2480_v8 = _dafny.Set.fromElements(_2479_v7);
              let _2481_v9;
              _2481_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2478_v6,new BigNumber((_2480_v8).length));
              _2481_v9 = _2481_v9;
              r0 = new BigNumber((_2478_v6).length);
              let _2482_v11;
              _2482_v11 = _dafny.MultiSet.fromElements((_this).f29);
              let _2483_v13;
              _2483_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm0(new BigNumber((function () {
                let _coll86 = new _dafny.Map();
                for (const _compr_86 of (function () {
                  let _coll87 = new _dafny.Set();
                  for (const _compr_87 of (_2482_v11).Elements) {
                    let _2484_v12 = _compr_87;
                    if ((_2482_v11).contains(_2484_v12)) {
                      _coll87.add(_2484_v12);
                    }
                  }
                  return _coll87;
                }()).Elements) {
                  let _2485_v10 = _compr_86;
                  if ((function () {
                    let _coll88 = new _dafny.Set();
                    for (const _compr_88 of (_2482_v11).Elements) {
                      let _2486_v12 = _compr_88;
                      if ((_2482_v11).contains(_2486_v12)) {
                        _coll88.add(_2486_v12);
                      }
                    }
                    return _coll88;
                  }()).contains(_2485_v10)) {
                    _coll86.push([_2485_v10,new BigNumber(165)]);
                  }
                }
                return _coll86;
              }()).length), globalState));
              let _2487_v14;
              let _nw394 = Array((new BigNumber(16)).toNumber());
              _nw394[(_dafny.ZERO).toNumber()] = (new BigNumber(-982)).isLessThanOrEqualTo((_this).f25);
              _nw394[(_dafny.ONE).toNumber()] = p1;
              _nw394[(new BigNumber(2)).toNumber()] = true;
              _nw394[(new BigNumber(3)).toNumber()] = (p1) === (p1);
              _nw394[(new BigNumber(4)).toNumber()] = !((((_2483_v13).contains(p1)) ? ((_2483_v13).get(p1)) : (p1)));
              _nw394[(new BigNumber(5)).toNumber()] = !(p1);
              _nw394[(new BigNumber(6)).toNumber()] = p1;
              _nw394[(new BigNumber(7)).toNumber()] = p1;
              _nw394[(new BigNumber(8)).toNumber()] = p1;
              _nw394[(new BigNumber(9)).toNumber()] = p1;
              _nw394[(new BigNumber(10)).toNumber()] = p1;
              _nw394[(new BigNumber(11)).toNumber()] = p1;
              _nw394[(new BigNumber(12)).toNumber()] = p1;
              _nw394[(new BigNumber(13)).toNumber()] = p1;
              _nw394[(new BigNumber(14)).toNumber()] = _module.__default.fm0((_this).f25, globalState);
              _nw394[(new BigNumber(15)).toNumber()] = p1;
              _2487_v14 = _nw394;
              let _index466 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_2487_v14).length));
              (_2487_v14)[_index466] = p1;
              let _index467 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_2487_v14).length));
              (_2487_v14)[_index467] = p1;
              let _2488_v15;
              _2488_v15 = _dafny.Set.fromElements((_this).f25, (_this).f25);
              let _2489_v16;
              let _nw395 = new _module.C4();
              _nw395.__ctor((_this).f26, new BigNumber(((_2488_v15).Intersect(_2488_v15)).length), (_this).f26);
              _2489_v16 = _nw395;
              _2487_v14 = _2487_v14;
            } else {
              r0 = (_this).f25;
              let _2490_v17;
              let _nw396 = Array((new BigNumber(25)).toNumber()).fill(false);
              _2490_v17 = _nw396;
              let _2491_v18;
              _2491_v18 = _dafny.Seq.of(new BigNumber(-886), (_dafny.ZERO).minus(new BigNumber((_2478_v6).length)), (_this).f25);
              let _2492_v19;
              _2492_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2491_v18, _2491_v18),_2478_v6);
              let _2493_v20;
              _2493_v20 = _module.D4.create_DC13((_this).f25, p1, _module.__default.fm0((_this).f25, globalState), p1);
              let _2494_v21;
              _2494_v21 = _dafny.Seq.of((_2493_v20).dtor_cf22);
              let _rhs434 = ((_this).f25).plus((_this).f25);
              let _rhs435 = new BigNumber((_2492_v19).length);
              let _rhs436 = _dafny.Seq.contains(_dafny.Seq.Concat(_module.__default.fm1(p1, globalState), _2494_v21), p1);
              let _rhs437 = _2490_v17;
              let _lhs348 = globalState;
              r0 = _rhs434;
              r0 = _rhs435;
              _lhs348.f12 = _rhs436;
              _2490_v17 = _rhs437;
              let _2495_v22;
              let _nw397 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
              _2495_v22 = _nw397;
              let _index468 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_2495_v22).length));
              (_2495_v22)[_index468] = new BigNumber(-11);
              let _2496_v23;
              _2496_v23 = _dafny.Set.fromElements(p1, p1, p1);
              let _index469 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_2495_v22).length));
              (_2495_v22)[_index469] = new BigNumber((((p1) ? (_2496_v23) : (_2496_v23))).length);
              let _index470 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_2490_v17).length));
              (_2490_v17)[_index470] = p1;
              let _index471 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_2490_v17).length));
              (_2490_v17)[_index471] = p1;
              let _2497_v24;
              _2497_v24 = _dafny.Set.fromElements((_dafny.ZERO).minus((_2495_v22)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_2495_v22).length))]));
              let _index472 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_2495_v22).length));
              let _rhs438 = (_2495_v22)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_2495_v22).length))];
              let _rhs439 = (_2490_v17)[_module.__default.safeIndex(new BigNumber(553), new BigNumber((_2490_v17).length))];
              let _rhs440 = (_2497_v24).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(801)), function (_2498_i1) {
                return (_this).f29;
              })).length));
              let _lhs349 = _2495_v22;
              let _lhs350 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_2495_v22).length));
              let _lhs351 = globalState;
              let _lhs352 = globalState;
              _lhs349[_lhs350] = _rhs438;
              _lhs351.f12 = _rhs439;
              _lhs352.f12 = _rhs440;
            }
            let _2499_v25;
            let _nw398 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
            _2499_v25 = _nw398;
            let _2500_v26;
            let _nw399 = Array((new BigNumber(11)).toNumber());
            _nw399[(_dafny.ZERO).toNumber()] = _2499_v25;
            _nw399[(_dafny.ONE).toNumber()] = ((p1) ? (_2499_v25) : (_2499_v25));
            _nw399[(new BigNumber(2)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(3)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(4)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(5)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(6)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(7)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(8)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(9)).toNumber()] = _2499_v25;
            _nw399[(new BigNumber(10)).toNumber()] = (p0)[_module.__default.safeIndex((_this).f25, new BigNumber((p0).length))];
            _2500_v26 = _nw399;
            _2500_v26 = _2500_v26;
          }
        }
      }
      let _2501_v27;
      let _nw400 = new _module.C6();
      _nw400.__ctor((_this).f25, (_this).f26);
      _2501_v27 = _nw400;
      _2501_v27 = _2501_v27;
      let _2502_v28;
      _2502_v28 = _module.D0.create_DC1(p1, (new BigNumber((_dafny.MultiSet.fromElements((_this).f25)).cardinality())).isLessThanOrEqualTo((_this).f25));
      let _2503_v29;
      _2503_v29 = _dafny.Seq.of(p1, p1);
      let _2504_v30;
      _2504_v30 = _dafny.Set.fromElements(new BigNumber((_2503_v29).length));
      let _2505_v31;
      _2505_v31 = _dafny.Set.fromElements((_this).f25, new BigNumber((_2504_v30).length));
      let _2506_v32;
      _2506_v32 = _module.D4.create_DC14(_module.D4.create_DC11(_2505_v31));
      let _2507_v33;
      _2507_v33 = _module.D4.create_DC12();
      let _pat_let_tv49 = _2502_v28;
      let _pat_let_tv50 = _2502_v28;
      let _pat_let_tv51 = _2502_v28;
      let _pat_let_tv52 = _2507_v33;
      _2502_v28 = function (_source31) {
        if (_source31.is_DC12) {
          return _pat_let_tv49;
        } else if (_source31.is_DC13) {
          let _2508___mcc_h0 = (_source31).cf19;
          let _2509___mcc_h1 = (_source31).cf20;
          let _2510___mcc_h2 = (_source31).cf21;
          let _2511___mcc_h3 = (_source31).cf22;
          let _2512_cf22 = _2511___mcc_h3;
          let _2513_cf21 = _2510___mcc_h2;
          let _2514_cf20 = _2509___mcc_h1;
          let _2515_cf19 = _2508___mcc_h0;
          return _module.D0.create_DC1(_2514_cf20, _2514_cf20);
        } else if (_source31.is_DC11) {
          let _2516___mcc_h4 = (_source31).cf18;
          let _2517_cf18 = _2516___mcc_h4;
          return _pat_let_tv50;
        } else {
          let _2518___mcc_h5 = (_source31).cf23;
          let _2519_cf23 = _2518___mcc_h5;
          return _pat_let_tv51;
        }
      }(function (_pat_let58_0) {
        return function (_2520_dt__update__tmp_h1) {
          return function (_pat_let59_0) {
            return function (_2521_dt__update_hcf23_h0) {
              return _module.D4.create_DC14(_2521_dt__update_hcf23_h0);
            }(_pat_let59_0);
          }(_pat_let_tv52);
        }(_pat_let58_0);
      }(_2506_v32));
      let _2522_v34;
      _2522_v34 = _dafny.MultiSet.fromElements(((_this).f25).isLessThanOrEqualTo((_this).f25), p1, p1);
      let _2523_v35;
      _2523_v35 = _dafny.Seq.UnicodeFromString("u");
      let _2524_v36;
      _2524_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      _2522_v34 = _module.__default.fm50((_this).f25, _2523_v35, (((_2524_v36).contains(true)) ? ((_2524_v36).get(true)) : (p1)), (_this).f25, globalState);
      let _2525_v37;
      let _nw401 = Array((new BigNumber(2)).toNumber()).fill(false);
      _2525_v37 = _nw401;
      let _2526_v38;
      _2526_v38 = _dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_this).f25);
      let _index473 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2525_v37).length));
      (_2525_v37)[_index473] = (_2526_v38).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f25));
      let _index474 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2525_v37).length));
      (_2525_v37)[_index474] = p1;
      let _hi14 = (_dafny.ZERO).minus((_this).f25);
      for (let _2527_i2 = (_this).f25; _2527_i2.isLessThan(_hi14); _2527_i2 = _2527_i2.plus(_dafny.ONE)) {
        let _2528_v39;
        let _init86 = ((_2529_v35) => function (_2530_i3) {
          return _2529_v35;
        })(_2523_v35);
        let _nw402 = Array((new BigNumber(23)).toNumber());
        for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw402.length); _i0_86++) {
          _nw402[_i0_86] = _init86(new BigNumber(_i0_86));
        }
        _2528_v39 = _nw402;
        let _rhs441 = _module.__default.fm18(globalState);
        let _rhs442 = _2528_v39;
        r0 = _rhs441;
        _2528_v39 = _rhs442;
        let _2531_v40;
        _2531_v40 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("isuhwhf")).length), (_this).f25);
        let _2532_v41;
        _2532_v41 = _dafny.Set.fromElements(_2531_v40);
        r0 = _module.__default.safeModuloInt(new BigNumber((_2532_v41).length), (_this).f25);
        let _2533_v42;
        let _2534_v43;
        let _2535_v44;
        let _out43;
        let _out44;
        let _out45;
        let _outcollector15 = (_2501_v27).m12(globalState);
        _out43 = _outcollector15[0];
        _out44 = _outcollector15[1];
        _out45 = _outcollector15[2];
        _2533_v42 = _out43;
        _2534_v43 = _out44;
        _2535_v44 = _out45;
        let _2536_v45;
        let _nw403 = new _module.C5();
        _nw403.__ctor((_2525_v37)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_2525_v37).length))]);
        _2536_v45 = _nw403;
      }
      r0 = _module.__default.fm18(globalState);
      return r0;
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm9(p0, p1, globalState) {
      let _this = this;
      return new _dafny.CodePoint('l'.codePointAt(0));
    };
    fm10(globalState) {
      let _this = this;
      return _module.D1.create_DC5();
    };
    m6(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _2537_i0;
      _2537_i0 = _dafny.ZERO;
      L20: {
        while (false) {
          C20: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2537_i0)) {
              break L20;
            }
            _2537_i0 = (_2537_i0).plus(_dafny.ONE);
            let _2538_v0;
            _2538_v0 = false;
            let _2539_v1;
            _2539_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2538_v0,_2538_v0);
            let _2540_v2;
            _2540_v2 = _dafny.Map.Empty.slice().updateUnsafe((((_2539_v1).contains(_2538_v0)) ? ((_2539_v1).get(_2538_v0)) : (_2538_v0)),_2538_v0);
            _2540_v2 = (_2540_v2).update(_2538_v0, _2538_v0);
            let _2541_v3;
            _2541_v3 = _dafny.Seq.of(((_2538_v0) ? (!(_2538_v0)) : (_2538_v0)));
            let _2542_v4;
            _2542_v4 = new BigNumber(-601);
            _2541_v3 = _dafny.Seq.update(p0, _module.__default.safeIndex(_2542_v4, new BigNumber((p0).length)), _2538_v0);
            let _2543_v5;
            let _nw404 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
            _2543_v5 = _nw404;
            let _2544_v6;
            let _nw405 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2544_v6 = _nw405;
            let _2545_v7;
            let _nw406 = new _module.C4();
            _nw406.__ctor(_2544_v6, _2542_v4, _2544_v6);
            _2545_v7 = _nw406;
            let _2546_v8;
            _2546_v8 = _dafny.Seq.of(_2545_v7);
            let _index475 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2543_v5).length));
            (_2543_v5)[_index475] = _dafny.Seq.Concat(_2546_v8, _dafny.Seq.update(_2546_v8, _module.__default.safeIndex(_2542_v4, new BigNumber((_2546_v8).length)), _2545_v7));
            let _index476 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2543_v5).length));
            (_2543_v5)[_index476] = _dafny.Seq.update(_dafny.Seq.of(_2545_v7), _module.__default.safeIndex((_2545_v7).f25, new BigNumber((_dafny.Seq.of(_2545_v7)).length)), _2545_v7);
            let _2547_v9;
            _2547_v9 = _dafny.Map.Empty.slice().updateUnsafe((_2545_v7).f25,_2538_v0);
            let _2548_v10;
            _2548_v10 = _dafny.Set.fromElements(_2538_v0);
            let _2549_v11;
            _2549_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2548_v10,_2538_v0);
            let _2550_v12;
            let _nw407 = Array((new BigNumber(19)).toNumber());
            _nw407[(_dafny.ZERO).toNumber()] = _2542_v4;
            _nw407[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("pb")).length);
            _nw407[(new BigNumber(2)).toNumber()] = _module.__default.fm44(_2538_v0, _2538_v0, _2538_v0, (_2545_v7).f25, globalState);
            _nw407[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_2545_v7).f25, _2542_v4);
            _nw407[(new BigNumber(4)).toNumber()] = (_2545_v7).f25;
            _nw407[(new BigNumber(5)).toNumber()] = ((_dafny.ZERO).minus((_2545_v7).f25)).plus((_2545_v7).f25);
            _nw407[(new BigNumber(6)).toNumber()] = (_2545_v7).f25;
            _nw407[(new BigNumber(7)).toNumber()] = new BigNumber((_2547_v9).length);
            _nw407[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("xeoqcuqdk")).length);
            _nw407[(new BigNumber(9)).toNumber()] = (_2545_v7).f25;
            _nw407[(new BigNumber(10)).toNumber()] = (_2545_v7).f25;
            _nw407[(new BigNumber(11)).toNumber()] = _2542_v4;
            _nw407[(new BigNumber(12)).toNumber()] = _module.__default.fm18(globalState);
            _nw407[(new BigNumber(13)).toNumber()] = (new BigNumber(378)).minus(_2542_v4);
            _nw407[(new BigNumber(14)).toNumber()] = _2542_v4;
            _nw407[(new BigNumber(15)).toNumber()] = new BigNumber((_2549_v11).length);
            _nw407[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_2542_v4);
            _nw407[(new BigNumber(17)).toNumber()] = (_2542_v4).minus(_2542_v4);
            _nw407[(new BigNumber(18)).toNumber()] = (_2545_v7).f25;
            _2550_v12 = _nw407;
            _2550_v12 = _2550_v12;
          }
        }
      }
      let _2551_v13;
      _2551_v13 = new BigNumber(-572);
      let _2552_v14;
      _2552_v14 = _dafny.Seq.UnicodeFromString("hkg");
      _2551_v13 = new BigNumber((_2552_v14).length);
      let _2553_v15;
      _2553_v15 = true;
      let _2554_v16;
      let _nw408 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2554_v16 = _nw408;
      let _2555_v17;
      let _nw409 = new _module.C3();
      _nw409.__ctor(_2551_v13, false, new BigNumber(637), ((_2553_v15) ? (_2554_v16) : (_2554_v16)));
      _2555_v17 = _nw409;
      let _2556_v18;
      let _nw410 = new _module.C1();
      _nw410.__ctor(_2551_v13, _2554_v16);
      _2556_v18 = _nw410;
      let _2557_v19;
      _2557_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2556_v18,_module.__default.safeDivisionInt((_2555_v17).f37, (_2555_v17).f37));
      _2557_v19 = (_2557_v19).update(_2556_v18, new BigNumber(610));
      if (_2553_v15) {
        if ((_2555_v17).f38) {
          _2553_v15 = _dafny.Seq.contains(_2552_v14, p1);
          let _2558_v20;
          let _nw411 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _2558_v20 = _nw411;
          let _2559_v21;
          _2559_v21 = _dafny.Set.fromElements(_2558_v20, _2558_v20, _2558_v20, _2558_v20);
          let _2560_v22;
          _2560_v22 = _dafny.Map.Empty.slice().updateUnsafe((_2559_v21).Intersect(_2559_v21),(_2555_v17).f37);
          _2560_v22 = (_2560_v22).update(_2559_v21, _module.__default.safeModuloInt(_2551_v13, _2551_v13));
          _2551_v13 = (_module.__default.safeModuloInt((_2555_v17).f37, (_2555_v17).f37)).minus(new BigNumber((_dafny.Seq.update(_2552_v14, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(956)), ((_2561_p1) => function (_2562_i1) {
            return _2561_p1;
          })(p1))).length), new BigNumber((_2552_v14).length)), p1)).length));
          let _2563_v23;
          _2563_v23 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2551_v13), (_2555_v17).f37, new BigNumber(254));
          _2551_v13 = ((_2555_v17).f37).plus((_module.__default.fm44(_2553_v15, false, (_2555_v17).f38, (_2555_v17).f37, globalState)).multipliedBy(new BigNumber((_2563_v23).cardinality())));
          let _2564_v24;
          _2564_v24 = _module.D0.create_DC0((_2555_v17).f38);
          let _2565_v25;
          _2565_v25 = _dafny.Map.Empty.slice().updateUnsafe((_2555_v17).fm38(_2564_v24, globalState),(_2555_v17).f38);
          let _index477 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2558_v20).length));
          (_2558_v20)[_index477] = new BigNumber((_2565_v25).length);
          let _index478 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2558_v20).length));
          (_2558_v20)[_index478] = _module.__default.safeDivisionInt(_module.__default.fm18(globalState), new BigNumber((_dafny.Set.fromElements(_module.__default.fm44((_2555_v17).f38, _2553_v15, (_2555_v17).f38, new BigNumber((_2563_v23).cardinality()), globalState), (_2555_v17).f37, (_2555_v17).f37)).length));
        } else {
          let _2566_v26;
          let _init87 = ((_2567_v17) => function (_2568_i2) {
            return (_2567_v17).f38;
          })(_2555_v17);
          let _nw412 = Array((new BigNumber(8)).toNumber());
          for (let _i0_87 = 0; _i0_87 < new BigNumber(_nw412.length); _i0_87++) {
            _nw412[_i0_87] = _init87(new BigNumber(_i0_87));
          }
          _2566_v26 = _nw412;
          let _index479 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2566_v26).length));
          (_2566_v26)[_index479] = (_2555_v17).f38;
          let _index480 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2566_v26).length));
          (_2566_v26)[_index480] = !((_2555_v17).f38);
          _2551_v13 = (_dafny.ZERO).minus((_2555_v17).f37);
          let _2569_v27;
          _2569_v27 = _dafny.Set.fromElements(new BigNumber(209));
          let _2570_v28;
          let _nw413 = Array((new BigNumber(9)).toNumber());
          _nw413[(_dafny.ZERO).toNumber()] = (_2555_v17).f37;
          _nw413[(_dafny.ONE).toNumber()] = (_2555_v17).f37;
          _nw413[(new BigNumber(2)).toNumber()] = (_2555_v17).f37;
          _nw413[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_2571_i3) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          })).length);
          _nw413[(new BigNumber(4)).toNumber()] = (_2555_v17).f37;
          _nw413[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("vbievbw"), _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_dafny.Seq.UnicodeFromString("vbievbw")).length)), p1)).length));
          _nw413[(new BigNumber(6)).toNumber()] = _2551_v13;
          _nw413[(new BigNumber(7)).toNumber()] = new BigNumber((_2569_v27).length);
          _nw413[(new BigNumber(8)).toNumber()] = new BigNumber(429);
          _2570_v28 = _nw413;
          let _2572_v29;
          _2572_v29 = _dafny.Map.Empty.slice().updateUnsafe((_2555_v17).f38,(_2555_v17).f37);
          let _2573_v30;
          let _nw414 = new _module.C7();
          _nw414.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2570_v28,new BigNumber(((_2572_v29).update((_2566_v26)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_2566_v26).length))], new BigNumber((_2552_v14).length))).length))).length));
          _2573_v30 = _nw414;
          let _2574_v31;
          _2574_v31 = _dafny.Set.fromElements(_2566_v26);
          let _2575_v32;
          _2575_v32 = _module.D3.create_DC9(_2574_v31);
          let _2576_v33;
          _2576_v33 = _dafny.Seq.of((_2555_v17).f37);
          let _2577_v34;
          let _nw415 = new _module.C4();
          _nw415.__ctor(_2554_v16, (_2573_v30).f34, _2554_v16);
          _2577_v34 = _nw415;
          let _2578_v35;
          _2578_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2576_v33,_2577_v34);
          let _rhs443 = _2573_v30;
          let _rhs444 = !((_2555_v17).f38) || (_2553_v15);
          let _rhs445 = (_2555_v17).f37;
          let _rhs446 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((new BigNumber((_2552_v14).length)).multipliedBy(new BigNumber((_2572_v29).length))), new BigNumber((_2578_v35).length));
          let _rhs447 = _2575_v32;
          _2573_v30 = _rhs443;
          _2553_v15 = _rhs444;
          _2551_v13 = _rhs445;
          _2551_v13 = _rhs446;
          _2575_v32 = _rhs447;
          _2551_v13 = _2551_v13;
          let _2579_v36;
          _2579_v36 = _dafny.Map.Empty.slice().updateUnsafe(((_2573_v30).f34).minus(new BigNumber((_2552_v14).length)),_2566_v26);
          let _2580_v37;
          let _nw416 = Array((new BigNumber(20)).toNumber()).fill(_module.D6.Default());
          _2580_v37 = _nw416;
          let _2581_v38;
          _2581_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2566_v26,new BigNumber(832));
          let _2582_v39;
          _2582_v39 = _module.D6.create_DC18(_2581_v38);
          let _index481 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_2580_v37).length));
          (_2580_v37)[_index481] = _2582_v39;
          let _index482 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2566_v26).length));
          let _index483 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_2580_v37).length));
          let _rhs448 = (_2555_v17).f38;
          let _rhs449 = ((_dafny.ZERO).minus((_2555_v17).f37)).isLessThanOrEqualTo(_2551_v13);
          let _rhs450 = _dafny.Map.Empty.slice().updateUnsafe((_2555_v17).f37,_2566_v26);
          let _rhs451 = _2582_v39;
          let _lhs353 = _2566_v26;
          let _lhs354 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_2566_v26).length));
          let _lhs355 = _2580_v37;
          let _lhs356 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_2580_v37).length));
          r0 = _rhs448;
          _lhs353[_lhs354] = _rhs449;
          _2579_v36 = _rhs450;
          _lhs355[_lhs356] = _rhs451;
        }
        let _2583_v40;
        let _init88 = ((_2584_v15) => function (_2585_i4) {
          return _module.__default.safeDivisionInt(_2585_i4, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2584_v15,new BigNumber(872))).length));
        })(_2553_v15);
        let _nw417 = Array((new BigNumber(11)).toNumber());
        for (let _i0_88 = 0; _i0_88 < new BigNumber(_nw417.length); _i0_88++) {
          _nw417[_i0_88] = _init88(new BigNumber(_i0_88));
        }
        _2583_v40 = _nw417;
        let _index484 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length));
        (_2583_v40)[_index484] = (_dafny.ZERO).minus(((_2555_v17).f37).plus((_2555_v17).f37));
        let _index485 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length));
        (_2583_v40)[_index485] = (((_dafny.ZERO).minus((_2555_v17).f37)).multipliedBy(_2551_v13)).multipliedBy(_module.__default.fm44((_2555_v17).f38, (_2555_v17).f38, (_2555_v17).f38, (_2555_v17).f37, globalState));
        (globalState).f12 = (_2555_v17).f38;
        if ((_2555_v17).f38) {
          let _2586_v41;
          _2586_v41 = _dafny.Map.Empty.slice().updateUnsafe((_2555_v17).f38,(_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))]);
          _2586_v41 = (_2586_v41).update((_2555_v17).f38, (_2555_v17).f37);
          let _2587_v42;
          _2587_v42 = _dafny.Seq.of(_2583_v40);
          _2583_v40 = (_2587_v42)[_module.__default.safeIndex((_2555_v17).f37, new BigNumber((_2587_v42).length))];
          _2551_v13 = (_2555_v17).f37;
          let _2588_v43;
          let _nw418 = Array((new BigNumber(18)).toNumber()).fill(false);
          _2588_v43 = _nw418;
          let _2589_v44;
          _2589_v44 = _module.D3.create_DC9(_dafny.Set.fromElements(_2588_v43));
          let _2590_v45;
          _2590_v45 = _dafny.Map.Empty.slice().updateUnsafe((_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))],_2589_v44);
          let _2591_v46;
          _2591_v46 = _module.D19.create_DC46((_2590_v45).Merge(_2590_v45));
          _2591_v46 = _2591_v46;
          let _2592_v47;
          _2592_v47 = _module.D1.create_DC4(p0);
          let _pat_let_tv53 = p0;
          let _2593_v48;
          _2593_v48 = _dafny.Seq.of(_2592_v47, function (_pat_let60_0) {
            return function (_2594_dt__update__tmp_h0) {
              return function (_pat_let61_0) {
                return function (_2595_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4(_2595_dt__update_hcf6_h0);
                }(_pat_let61_0);
              }(_pat_let_tv53);
            }(_pat_let60_0);
          }(_2592_v47), _2592_v47);
          let _2596_v49;
          let _nw419 = Array((new BigNumber(26)).toNumber());
          _nw419[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_2593_v48, _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_2593_v48).length)), _module.D1.create_DC4(p0));
          _nw419[(_dafny.ONE).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(2)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(3)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(4)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(5)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_2592_v47, _module.D1.create_DC4(p0));
          _nw419[(new BigNumber(7)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(8)).toNumber()] = _module.__default.fm68((_2555_v17).f38, (_2555_v17).f38, _2552_v14, (_this).fm9(_2553_v15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2555_v17).f38,_2553_v15)).length), globalState), globalState);
          _nw419[(new BigNumber(9)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_2593_v48, _dafny.Seq.update(_2593_v48, _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_2593_v48).length)), _2592_v47));
          _nw419[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_2592_v47);
          _nw419[(new BigNumber(12)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(13)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(869)), ((_2597_v47) => function (_2598_i5) {
            return _2597_v47;
          })(_2592_v47));
          _nw419[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_2593_v48, _2593_v48);
          _nw419[(new BigNumber(16)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(17)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_2593_v48, _module.__default.safeIndex((_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))], new BigNumber((_2593_v48).length)), _2592_v47);
          _nw419[(new BigNumber(19)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(20)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(21)).toNumber()] = _dafny.Seq.of(_module.D1.create_DC4(_dafny.Seq.of(_2553_v15)), _2592_v47);
          _nw419[(new BigNumber(22)).toNumber()] = _2593_v48;
          _nw419[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_2593_v48, _module.__default.safeIndex((_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))], new BigNumber((_2593_v48).length)), _2592_v47), _2593_v48);
          _nw419[(new BigNumber(24)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_2592_v47), _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_dafny.Seq.of(_2592_v47)).length)), _2592_v47), _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_2592_v47), _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_dafny.Seq.of(_2592_v47)).length)), _2592_v47)).length)), _2592_v47);
          _nw419[(new BigNumber(25)).toNumber()] = _dafny.Seq.of(_2592_v47);
          _2596_v49 = _nw419;
          _2596_v49 = _2596_v49;
        } else {
          let _2599_v50;
          let _init89 = ((_2600_v17) => function (_2601_i6) {
            return _dafny.Set.fromElements((_2600_v17).f38, (_2600_v17).f38, false);
          })(_2555_v17);
          let _nw420 = Array((new BigNumber(24)).toNumber());
          for (let _i0_89 = 0; _i0_89 < new BigNumber(_nw420.length); _i0_89++) {
            _nw420[_i0_89] = _init89(new BigNumber(_i0_89));
          }
          _2599_v50 = _nw420;
          _2599_v50 = _2599_v50;
          let _2602_v51;
          _2602_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2553_v15,(_2555_v17).f38);
          let _index486 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length));
          (_2583_v40)[_index486] = ((((_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))]).isLessThan((_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))])) ? ((_2555_v17).f37) : (_module.__default.safeDivisionInt(new BigNumber((_2602_v51).length), (_2555_v17).f37)));
          let _2603_v52;
          let _nw421 = Array((new BigNumber(26)).toNumber());
          _nw421[(_dafny.ZERO).toNumber()] = (_2555_v17).f38;
          _nw421[(_dafny.ONE).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(2)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(3)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(4)).toNumber()] = _2553_v15;
          _nw421[(new BigNumber(5)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(6)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(7)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(8)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(9)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(10)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(11)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(12)).toNumber()] = (p0)[_module.__default.safeIndex((_module.__default.fm66(globalState)).dtor_cf71, new BigNumber((p0).length))];
          _nw421[(new BigNumber(13)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(14)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(15)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(16)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(17)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(18)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(19)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(20)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(21)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(22)).toNumber()] = (_2555_v17).f38;
          _nw421[(new BigNumber(23)).toNumber()] = !(!(!(true)));
          _nw421[(new BigNumber(24)).toNumber()] = _2553_v15;
          _nw421[(new BigNumber(25)).toNumber()] = (_2555_v17).f38;
          _2603_v52 = _nw421;
          let _2604_v53;
          _2604_v53 = _dafny.Seq.of(_2603_v52, _2603_v52, _2603_v52);
          let _2605_v54;
          let _nw422 = Array((new BigNumber(26)).toNumber());
          _nw422[(_dafny.ZERO).toNumber()] = _2603_v52;
          _nw422[(_dafny.ONE).toNumber()] = (_2604_v53)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_2604_v53).length))];
          _nw422[(new BigNumber(2)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(3)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(4)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(5)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(6)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(7)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(8)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(9)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(10)).toNumber()] = (((_2555_v17).f38) ? (_2603_v52) : (_2603_v52));
          _nw422[(new BigNumber(11)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(12)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(13)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(14)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(15)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(16)).toNumber()] = ((false) ? (_2603_v52) : (_2603_v52));
          _nw422[(new BigNumber(17)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(18)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(19)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(20)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(21)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(22)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(23)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(24)).toNumber()] = _2603_v52;
          _nw422[(new BigNumber(25)).toNumber()] = _2603_v52;
          _2605_v54 = _nw422;
          let _2606_v55;
          _2606_v55 = _dafny.MultiSet.fromElements((_2555_v17).f38, false);
          let _2607_v56;
          _2607_v56 = _module.D20.create_DC49(_2605_v54);
          let _index487 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length));
          let _rhs452 = (_module.D14.create_DC38((_2583_v40)[_module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length))])).dtor_cf58;
          let _rhs453 = (_2607_v56).dtor_cf72;
          let _rhs454 = _2551_v13;
          let _rhs455 = _2606_v55;
          let _lhs357 = _2583_v40;
          let _lhs358 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_2583_v40).length));
          _2551_v13 = _rhs452;
          _2605_v54 = _rhs453;
          _lhs357[_lhs358] = _rhs454;
          _2606_v55 = _rhs455;
          let _2608_v57;
          _2608_v57 = _dafny.Seq.of(_2553_v15);
          _2608_v57 = _dafny.Seq.Concat(_2608_v57, _dafny.Seq.update(_2608_v57, _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_2608_v57).length)), _2553_v15));
          let _2609_v58;
          let _nw423 = new _module.C5();
          _nw423.__ctor((_2555_v17).f38);
          _2609_v58 = _nw423;
        }
        _2583_v40 = (((_2555_v17).f38) ? (_2583_v40) : (_2583_v40));
      } else {
        if (!((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_2551_v13, (_2555_v17).f37))).isEqualTo((_2555_v17).f37)) {
          let _2610_v59;
          _2610_v59 = _dafny.Set.fromElements((_2555_v17).f38, (_2555_v17).f38, _2553_v15, _module.__default.fm0((_2555_v17).f37, globalState), (_2555_v17).f38);
          _2610_v59 = _2610_v59;
          let _2611_v60;
          _2611_v60 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus((_2555_v17).f37)).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_2552_v14, _module.__default.safeIndex((_2555_v17).f37, new BigNumber((_2552_v14).length)), p1)).length), (_2555_v17).f37)).cardinality())),true);
          _2611_v60 = (_2611_v60).update((_2555_v17).f37, _2553_v15);
          let _2612_v61;
          _2612_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2551_v13,_module.__default.safeDivisionInt((_2555_v17).f37, (_2555_v17).f37));
          _2612_v61 = (_2612_v61).update(_2551_v13, new BigNumber(439));
          (globalState).f12 = !_dafny.Seq.contains(_2552_v14, p1);
          let _2613_v62;
          _2613_v62 = _module.D0.create_DC0((_2555_v17).f38);
          let _2614_v63;
          _2614_v63 = _dafny.Set.fromElements(_2613_v62, _2613_v62);
          let _2615_v64;
          let _nw424 = Array((new BigNumber(4)).toNumber());
          _nw424[(_dafny.ZERO).toNumber()] = (_2555_v17).f37;
          _nw424[(_dafny.ONE).toNumber()] = (_2555_v17).f37;
          _nw424[(new BigNumber(2)).toNumber()] = new BigNumber((_2614_v63).length);
          _nw424[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_2555_v17).f37);
          _2615_v64 = _nw424;
          let _2616_v65;
          _2616_v65 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-64)),_2615_v64);
          let _2617_v66;
          _2617_v66 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-182),_2615_v64), _2616_v65, _2616_v65, _dafny.Map.Empty.slice().updateUnsafe((_2555_v17).f37,_2615_v64));
          _2551_v13 = (((_2617_v66).contains(_dafny.Map.Empty.slice().updateUnsafe(_2551_v13,_2615_v64))) ? ((_2617_v66).get(_dafny.Map.Empty.slice().updateUnsafe(_2551_v13,_2615_v64))) : (new BigNumber(195)));
        } else {
          let _2618_v67;
          _2618_v67 = _dafny.Set.fromElements(_2553_v15, (_2555_v17).f38);
          let _2619_v68;
          _2619_v68 = _module.D5.create_DC16((_2618_v67).IsDisjointFrom(_2618_v67));
          _2619_v68 = _module.D5.create_DC16((_2555_v17).f38);
          _2553_v15 = (_2555_v17).f38;
          let _2620_v69;
          _2620_v69 = _dafny.Set.fromElements((_2555_v17).f37);
          let _rhs456 = _dafny.Set.fromElements((_2555_v17).f37, (_2555_v17).f37, (_2555_v17).f37);
          let _rhs457 = (_dafny.ZERO).minus(_2551_v13);
          let _rhs458 = (_2555_v17).f38;
          _2620_v69 = _rhs456;
          _2551_v13 = _rhs457;
          r0 = _rhs458;
          _2551_v13 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2555_v17).f37));
          let _2621_v70;
          let _init90 = ((_2622_v15) => function (_2623_i7) {
            return _2622_v15;
          })(_2553_v15);
          let _nw425 = Array((new BigNumber(25)).toNumber());
          for (let _i0_90 = 0; _i0_90 < new BigNumber(_nw425.length); _i0_90++) {
            _nw425[_i0_90] = _init90(new BigNumber(_i0_90));
          }
          _2621_v70 = _nw425;
          let _2624_v71;
          _2624_v71 = _dafny.Seq.of((_2555_v17).f37);
          let _2625_v72;
          _2625_v72 = _dafny.MultiSet.fromElements((_2555_v17).f38);
          let _2626_v73;
          _2626_v73 = _dafny.MultiSet.fromElements((_module.D3.create_DC10(_2553_v15, _2621_v70, (_2555_v17).f38, (_2624_v71)[_module.__default.safeIndex(new BigNumber((_2625_v72).cardinality()), new BigNumber((_2624_v71).length))], new BigNumber((_2618_v67).length))).dtor_cf13);
          let _2627_v74;
          _2627_v74 = _dafny.Seq.of(_2625_v72, _dafny.MultiSet.fromElements((_2555_v17).f38));
          let _2628_v75;
          _2628_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2626_v73,new BigNumber((_dafny.Seq.Concat(_module.__default.fm69((_2555_v17).f38, globalState), _2627_v74)).length));
          let _2629_v76;
          _2629_v76 = _module.D0.create_DC0(!((_2555_v17).f38));
          let _2630_v77;
          _2630_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2629_v76,_2553_v15);
          let _rhs459 = (_2628_v75).Merge((_dafny.Map.Empty.slice().updateUnsafe(_2625_v72,(_2555_v17).f37)).Merge(_2628_v75));
          let _rhs460 = _2621_v70;
          let _rhs461 = _module.__default.fm21((_dafny.Map.Empty.slice().updateUnsafe(_2629_v76,true)).Merge(_2630_v77), _2551_v13, !((_2555_v17).f38), _2552_v14, globalState);
          let _rhs462 = new BigNumber((_2552_v14).length);
          let _rhs463 = _2553_v15;
          let _lhs359 = globalState;
          _2628_v75 = _rhs459;
          _2621_v70 = _rhs460;
          _2551_v13 = _rhs461;
          _2551_v13 = _rhs462;
          _lhs359.f12 = _rhs463;
        }
        (globalState).f23 = p1;
        (globalState).f12 = (_2555_v17).f38;
        _2551_v13 = ((_2555_v17).f37).multipliedBy(_module.__default.fm18(globalState));
        _2551_v13 = (_2555_v17).f37;
      }
      let _2631_v78;
      _2631_v78 = _module.D1.create_DC4(p0);
      let _pat_let_tv54 = _2553_v15;
      let _pat_let_tv55 = _2553_v15;
      let _pat_let_tv56 = _2553_v15;
      if (function (_source32) {
        if (_source32.is_DC5) {
          return _pat_let_tv54;
        } else if (_source32.is_DC6) {
          let _2632___mcc_h0 = (_source32).cf7;
          let _2633___mcc_h1 = (_source32).cf8;
          let _2634___mcc_h2 = (_source32).cf9;
          let _2635_cf9 = _2634___mcc_h2;
          let _2636_cf8 = _2633___mcc_h1;
          let _2637_cf7 = _2632___mcc_h0;
          return (false) && (_2636_cf8);
        } else if (_source32.is_DC4) {
          let _2638___mcc_h3 = (_source32).cf6;
          let _2639_cf6 = _2638___mcc_h3;
          return _pat_let_tv55;
        } else {
          let _2640___mcc_h4 = (_source32).cf10;
          let _2641_cf10 = _2640___mcc_h4;
          return _pat_let_tv56;
        }
      }(_2631_v78)) {
        let _2642_v79;
        let _nw426 = new _module.C0();
        _nw426.__ctor();
        _2642_v79 = _nw426;
        let _2643_v80;
        _2643_v80 = _dafny.Seq.of(_2642_v79);
        let _2644_v81;
        let _nw427 = Array((new BigNumber(26)).toNumber());
        _nw427[(_dafny.ZERO).toNumber()] = _2642_v79;
        _nw427[(_dafny.ONE).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(2)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(3)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(4)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(5)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(6)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(7)).toNumber()] = (_2643_v80)[_module.__default.safeIndex(_2551_v13, new BigNumber((_2643_v80).length))];
        _nw427[(new BigNumber(8)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(9)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(10)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(11)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(12)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(13)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(14)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(15)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(16)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(17)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(18)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(19)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(20)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(21)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(22)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(23)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(24)).toNumber()] = _2642_v79;
        _nw427[(new BigNumber(25)).toNumber()] = _2642_v79;
        _2644_v81 = _nw427;
        let _index488 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_2644_v81).length));
        (_2644_v81)[_index488] = _2642_v79;
        let _index489 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_2644_v81).length));
        let _nw428 = new _module.C0();
        _nw428.__ctor();
        (_2644_v81)[_index489] = _nw428;
        let _2645_v82;
        _2645_v82 = _dafny.MultiSet.fromElements((_2555_v17).f38);
        let _2646_v83;
        _2646_v83 = _dafny.Set.fromElements((_2555_v17).f38);
        let _2647_v84;
        _2647_v84 = _dafny.Map.Empty.slice().updateUnsafe(((_2555_v17).f37).plus(_2551_v13),new BigNumber((_2552_v14).length));
        let _2648_v85;
        _2648_v85 = _dafny.Seq.of((_2555_v17).f37, _2551_v13, (_2555_v17).f37);
        let _rhs464 = _dafny.Seq.of((((_2555_v17).f38) ? (new BigNumber((_2645_v82).cardinality())) : (new BigNumber((_2646_v83).length))));
        let _rhs465 = (((_2647_v84).contains(_module.__default.safeDivisionInt(new BigNumber((_2648_v85).length), new BigNumber((_2552_v14).length)))) ? ((_2647_v84).get(_module.__default.safeDivisionInt(new BigNumber((_2648_v85).length), new BigNumber((_2552_v14).length)))) : ((_2555_v17).f37));
        let _rhs466 = (_2555_v17).f37;
        let _lhs360 = globalState;
        _lhs360.f6 = _rhs464;
        _2551_v13 = _rhs465;
        _2551_v13 = _rhs466;
        let _2649_v86;
        let _nw429 = Array((new BigNumber(21)).toNumber()).fill(false);
        _2649_v86 = _nw429;
        _2649_v86 = _2649_v86;
        (globalState).f10 = _dafny.Seq.Concat(_2552_v14, _2552_v14);
        let _2650_v87;
        let _nw430 = new _module.C5();
        _nw430.__ctor((_2555_v17).f38);
        _2650_v87 = _nw430;
      } else {
        let _2651_v88;
        _2651_v88 = _dafny.Seq.of(_2551_v13);
        let _2652_v89;
        _2652_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2651_v88,(((_2555_v17).f38) ? ((_2555_v17).f38) : (!(_2553_v15))));
        _2652_v89 = _2652_v89;
        let _2653_v90;
        let _nw431 = new _module.C4();
        _nw431.__ctor(_2554_v16, (_2555_v17).f37, _2554_v16);
        _2653_v90 = _nw431;
        let _2654_v91;
        _2654_v91 = _dafny.Seq.of(_2653_v90, _2653_v90, _2653_v90);
        _2653_v90 = (_2654_v91)[_module.__default.safeIndex(_2551_v13, new BigNumber((_2654_v91).length))];
        r0 = (_2555_v17).f38;
        let _2655_v92;
        let _nw432 = new _module.C6();
        _nw432.__ctor((new BigNumber((_2552_v14).length)).minus((_2555_v17).f37), _2554_v16);
        _2655_v92 = _nw432;
        _2655_v92 = _2655_v92;
        _2551_v13 = _2551_v13;
      }
      let _2656_v93;
      _2656_v93 = _module.D4.create_DC13(_2551_v13, _2553_v15, _module.__default.fm0((_2555_v17).f37, globalState), (_2555_v17).f38);
      let _pat_let_tv57 = _2555_v17;
      let _pat_let_tv58 = _2553_v15;
      let _pat_let_tv59 = _2555_v17;
      let _pat_let_tv60 = _2555_v17;
      let _pat_let_tv61 = _2553_v15;
      let _pat_let_tv62 = _2555_v17;
      r0 = function (_source33) {
        if (_source33.is_DC12) {
          return (_dafny.MultiSet.fromElements((_pat_let_tv57).f38)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_pat_let_tv58, !((_pat_let_tv59).f38), (_pat_let_tv60).f38));
        } else if (_source33.is_DC13) {
          let _2657___mcc_h5 = (_source33).cf19;
          let _2658___mcc_h6 = (_source33).cf20;
          let _2659___mcc_h7 = (_source33).cf21;
          let _2660___mcc_h8 = (_source33).cf22;
          let _2661_cf22 = _2660___mcc_h8;
          let _2662_cf21 = _2659___mcc_h7;
          let _2663_cf20 = _2658___mcc_h6;
          let _2664_cf19 = _2657___mcc_h5;
          return _2663_cf20;
        } else if (_source33.is_DC11) {
          let _2665___mcc_h9 = (_source33).cf18;
          let _2666_cf18 = _2665___mcc_h9;
          return _pat_let_tv61;
        } else {
          let _2667___mcc_h10 = (_source33).cf23;
          let _2668_cf23 = _2667___mcc_h10;
          return (_pat_let_tv62).f38;
        }
      }(_2656_v93);
      return r0;
    }
  };

  $module.C12 = class C12 {
    constructor () {
      this._tname = "_module.C12";
      this._f25 = _dafny.ZERO;
      this._f26 = [];
      this.f28 = false;
      this._f27 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    __ctor(f27, f28, f25, f26) {
      let _this = this;
      (_this)._f27 = f27;
      (_this).f28 = f28;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return new _dafny.CodePoint('q'.codePointAt(0));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return (_module.D0.create_DC3((_dafny.ZERO).minus((_this).f27), _this.f28)).dtor_cf4;
    };
    m4(p0, globalState) {
      let _this = this;
      let _2669_v0;
      _2669_v0 = _dafny.Seq.of((_this).fm8(_module.__default.fm0(new BigNumber((_dafny.Set.fromElements((_this).f27)).length), globalState), _this.f28, globalState));
      let _2670_v1;
      _2670_v1 = _dafny.Seq.UnicodeFromString("ntbeock");
      let _2671_v2;
      _2671_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_2670_v1);
      let _2672_v3;
      let _nw433 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _2672_v3 = _nw433;
      let _2673_v4;
      _2673_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2672_v3);
      let _2674_v5;
      _2674_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber((_2673_v4).length));
      let _2675_i0;
      _2675_i0 = _dafny.ZERO;
      L21: {
        while ((((_2669_v0)[_module.__default.safeIndex((_this).f25, new BigNumber((_2669_v0).length))]).multipliedBy(new BigNumber(((((_2671_v2).contains(new BigNumber((_2674_v5).length))) ? ((_2671_v2).get(new BigNumber((_2674_v5).length))) : (_2670_v1))).length))).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_this).f27, (_this).f27))) {
          C21: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2675_i0)) {
              break L21;
            }
            _2675_i0 = (_2675_i0).plus(_dafny.ONE);
            let _2676_v6;
            _2676_v6 = _dafny.Seq.of(_this.f28, _this.f28);
            let _2677_v7;
            let _nw434 = Array((new BigNumber(24)).toNumber());
            _nw434[(_dafny.ZERO).toNumber()] = _this.f28;
            _nw434[(_dafny.ONE).toNumber()] = _this.f28;
            _nw434[(new BigNumber(2)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(3)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(4)).toNumber()] = false;
            _nw434[(new BigNumber(5)).toNumber()] = _module.__default.fm0((_this).f25, globalState);
            _nw434[(new BigNumber(6)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(7)).toNumber()] = !(_this.f28);
            _nw434[(new BigNumber(8)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(9)).toNumber()] = (_2676_v6)[_module.__default.safeIndex((_this).f27, new BigNumber((_2676_v6).length))];
            _nw434[(new BigNumber(10)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(11)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(12)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(13)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(14)).toNumber()] = true;
            _nw434[(new BigNumber(15)).toNumber()] = false;
            _nw434[(new BigNumber(16)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(17)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(18)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(19)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(20)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(21)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(22)).toNumber()] = _this.f28;
            _nw434[(new BigNumber(23)).toNumber()] = _this.f28;
            _2677_v7 = _nw434;
            let _2678_v8;
            _2678_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_2677_v7);
            let _2679_v9;
            _2679_v9 = _dafny.Seq.of(_2678_v8);
            _2678_v8 = (((_2678_v8).update(_this.f28, _2677_v7)).Merge(_2678_v8)).Merge(((_2679_v9)[_module.__default.safeIndex((_this).f27, new BigNumber((_2679_v9).length))]).Merge(_2678_v8));
            let _2680_v10;
            _2680_v10 = new BigNumber(40);
            _2680_v10 = (_module.__default.safeDivisionInt((_this).f27, _2680_v10)).plus(_2680_v10);
            _2677_v7 = _2677_v7;
            let _2681_v11;
            _2681_v11 = _dafny.Set.fromElements(_2669_v0, _2669_v0);
            _2681_v11 = _2681_v11;
          }
        }
      }
      let _2682_v12;
      _2682_v12 = _dafny.MultiSet.fromElements(_this.f28, _this.f28);
      let _2683_v13;
      _2683_v13 = _dafny.Seq.of(_this.f28, _this.f28);
      let _2684_v14;
      let _nw435 = new _module.C8();
      _nw435.__ctor((_2682_v12).IsProperSubsetOf(_dafny.MultiSet.FromArray(_2683_v13)), !(_this.f28), (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f25, (_this).f25)), (_this).f26);
      _2684_v14 = _nw435;
      let _2685_v15;
      let _init91 = function (_2686_i2) {
        return !(_this.f28);
      };
      let _nw436 = Array((new BigNumber(24)).toNumber());
      for (let _i0_91 = 0; _i0_91 < new BigNumber(_nw436.length); _i0_91++) {
        _nw436[_i0_91] = _init91(new BigNumber(_i0_91));
      }
      _2685_v15 = _nw436;
      for (const _guard_loop_12 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2685_v15).length))) {
        let _2687_i1 = _guard_loop_12;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2687_i1)) && ((_2687_i1).isLessThan(new BigNumber((_2685_v15).length))))) {
          (_2685_v15)[(_2687_i1)] = (_module.D5.create_DC17((_this).f27, (_2684_v14).f32)).dtor_cf27;
        }
      }
      let _2688_v16;
      _2688_v16 = _dafny.MultiSet.fromElements((_this).f27);
      let _2689_v17;
      _2689_v17 = _dafny.Set.fromElements(_dafny.Seq.Concat(_2683_v13, _2683_v13), _dafny.Seq.Concat(_2683_v13, _dafny.Seq.of((_2684_v14).f33, (_2684_v14).f33, (_2684_v14).f33)), _2683_v13);
      let _2690_v18;
      let _nw437 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
      _2690_v18 = _nw437;
      let _2691_v19;
      _2691_v19 = _module.D3.create_DC10(!((_2684_v14).f33), _2685_v15, (_2684_v14).f32, (_2684_v14).fm14(_this.f28, globalState), (_this).f27);
      let _2692_v20;
      _2692_v20 = _dafny.MultiSet.fromElements(_2691_v19, _2691_v19);
      let _index490 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_2690_v18).length));
      (_2690_v18)[_index490] = (_2692_v20).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_2691_v19, _module.D3.create_DC10(_this.f28, _2685_v15, false, (_this).f27, (_this).f27), _2691_v19, _2691_v19, _2691_v19)));
      let _index491 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_2672_v3).length));
      (_2672_v3)[_index491] = (_this).f27;
      let _2693_v21;
      _2693_v21 = new _dafny.CodePoint('y'.codePointAt(0));
      let _index492 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_2690_v18).length));
      let _index493 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_2672_v3).length));
      let _rhs467 = ((_module.__default.fm0(new BigNumber(883), globalState)) ? (_module.__default.fm15(globalState)) : (((_this.f28) ? ((_2688_v16).update((_this).f27, _module.__default.abs((_this).f25))) : (_dafny.MultiSet.fromElements((_this).f25, (_this).f27)))));
      let _rhs468 = _2693_v21;
      let _rhs469 = (_2689_v17).Union(_2689_v17);
      let _rhs470 = (_dafny.MultiSet.fromElements(_2691_v19)).Difference(_2692_v20);
      let _rhs471 = (_this).f25;
      let _lhs361 = globalState;
      let _lhs362 = _2690_v18;
      let _lhs363 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_2690_v18).length));
      let _lhs364 = _2672_v3;
      let _lhs365 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_2672_v3).length));
      _2688_v16 = _rhs467;
      _lhs361.f23 = _rhs468;
      _2689_v17 = _rhs469;
      _lhs362[_lhs363] = _rhs470;
      _lhs364[_lhs365] = _rhs471;
      for (const _guard_loop_13 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2672_v3).length))) {
        let _2694_i3 = _guard_loop_13;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2694_i3)) && ((_2694_i3).isLessThan(new BigNumber((_2672_v3).length))))) {
          (_2672_v3)[(_2694_i3)] = (_2694_i3).minus(_module.__default.safeDivisionInt(new BigNumber(582), (_dafny.ZERO).minus((_this).f27)));
        }
      }
      let _2695_v22;
      let _nw438 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
      _2695_v22 = _nw438;
      let _2696_v23;
      _2696_v23 = _module.D4.create_DC13((_2672_v3)[_module.__default.safeIndex(new BigNumber(74), new BigNumber((_2672_v3).length))], (_2684_v14).f33, (_2684_v14).f33, _this.f28);
      let _2697_v24;
      _2697_v24 = _dafny.MultiSet.fromElements(_module.D4.create_DC13((_this).f25, _this.f28, (_2684_v14).f32, (_2684_v14).f32), _2696_v23);
      let _index494 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2695_v22).length));
      (_2695_v22)[_index494] = (_2697_v24).Union(_2697_v24);
      let _2698_v25;
      _2698_v25 = _2669_v0;
      let _2699_v26;
      _2699_v26 = _module.D14.create_DC38((_2684_v14).fm14(_this.f28, globalState));
      let _pat_let_tv63 = _2672_v3;
      let _pat_let_tv64 = _2672_v3;
      let _index495 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_2672_v3).length));
      let _index496 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2695_v22).length));
      let _rhs472 = function (_source34) {
        if (_source34.is_DC38) {
          let _2700___mcc_h0 = (_source34).cf58;
          let _2701_cf58 = _2700___mcc_h0;
          if (!(false)) {
            return (_pat_let_tv64)[_module.__default.safeIndex(new BigNumber(74), new BigNumber((_pat_let_tv63).length))];
          } else {
            return (_this).f25;
          }
        } else {
          let _2702___mcc_h1 = (_source34).cf57;
          let _2703_cf57 = _2702___mcc_h1;
          return (_this).f27;
        }
      }(_2699_v26);
      let _rhs473 = ((true) ? (_2697_v24) : ((_2697_v24).Union(_2697_v24)));
      let _rhs474 = _2672_v3;
      let _rhs475 = _2698_v25;
      let _lhs366 = _2672_v3;
      let _lhs367 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_2672_v3).length));
      let _lhs368 = _2695_v22;
      let _lhs369 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2695_v22).length));
      _lhs366[_lhs367] = _rhs472;
      _lhs368[_lhs369] = _rhs473;
      _2672_v3 = _rhs474;
      _2698_v25 = _rhs475;
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _2704_v0;
      _2704_v0 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2705_v1;
      _2705_v1 = _dafny.Seq.UnicodeFromString("evoys");
      r0 = ((_module.__default.fm0(new BigNumber(((_module.D1.create_DC6(_2704_v0, p1, _dafny.Seq.UnicodeFromString("xvrcenaw"))).dtor_cf9).length), globalState)) ? ((_dafny.ZERO).minus((_this).f27)) : (new BigNumber((_2705_v1).length)));
      _2704_v0 = new _dafny.CodePoint('b'.codePointAt(0));
      let _2706_i0;
      _2706_i0 = _dafny.ZERO;
      L22: {
        while (_this.f28) {
          C22: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2706_i0)) {
              break L22;
            }
            _2706_i0 = (_2706_i0).plus(_dafny.ONE);
            let _2707_v2;
            _2707_v2 = _dafny.Seq.of(p2, p1);
            let _2708_v3;
            _2708_v3 = _dafny.MultiSet.fromElements(_2707_v2);
            let _2709_v4;
            let _nw439 = new _module.C8();
            _nw439.__ctor((_2708_v3).IsSubsetOf(_2708_v3), p1, _module.__default.fm18(globalState), (_this).f26);
            _2709_v4 = _nw439;
            _2709_v4 = _2709_v4;
            (globalState).f10 = _dafny.Seq.Concat(_2705_v1, _2705_v1);
            let _index497 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((p3).length));
            (p3)[_index497] = _module.__default.safeModuloInt(new BigNumber(886), (_this).f25);
            let _index498 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((p3).length));
            (p3)[_index498] = (_this).fm8(p2, _this.f28, globalState);
            (_this).f28 = (_2709_v4).fm13(globalState);
          }
        }
      }
      let _2710_v5;
      _2710_v5 = _dafny.MultiSet.fromElements((_this).f27, (_this).f27);
      let _2711_v6;
      _2711_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p2, p2),p1);
      let _2712_v7;
      let _nw440 = new _module.C10();
      _nw440.__ctor(_2704_v0, (((_2710_v5).contains(_module.__default.fm44(p2, p2, p2, new BigNumber((_2711_v6).length), globalState))) ? ((_2710_v5).get(_module.__default.fm44(p2, p2, p2, new BigNumber((_2711_v6).length), globalState))) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,new _dafny.CodePoint('a'.codePointAt(0)))).length))), (_this).f26);
      _2712_v7 = _nw440;
      let _2713_v8;
      _2713_v8 = _dafny.Set.fromElements(_2712_v7);
      let _2714_v9;
      _2714_v9 = _dafny.Set.fromElements(_2713_v8, _2713_v8, _2713_v8);
      r1 = ((_this.f28) ? ((_this).f27) : (new BigNumber(((_2714_v9).Intersect(_2714_v9)).length)));
      let _2715_i1;
      _2715_i1 = _dafny.ZERO;
      L23: {
        while (p2) {
          C23: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2715_i1)) {
              break L23;
            }
            _2715_i1 = (_2715_i1).plus(_dafny.ONE);
            r1 = (_dafny.ZERO).minus((_this).f25);
            let _2716_v10;
            _2716_v10 = _dafny.Set.fromElements(p2);
            let _2717_v11;
            _2717_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2716_v10,_2710_v5);
            let _2718_v12;
            _2718_v12 = _dafny.Seq.of(new BigNumber(604), (_this).f27);
            _2717_v11 = (_2717_v11).update((_2716_v10).Union(_2716_v10), (_2710_v5).update(new BigNumber(-477), _module.__default.abs(new BigNumber((_2718_v12).length))));
            let _2719_v13;
            _2719_v13 = _module.D5.create_DC17((_this).f27, p2);
            let _2720_v14;
            _2720_v14 = _dafny.Map.Empty.slice().updateUnsafe((_2719_v13).dtor_cf27,(_this).f27);
            let _2721_v15;
            _2721_v15 = _module.D9.create_DC26(_2720_v14);
            let _source35 = _2721_v15;
            if (_source35.is_DC27) {
              let _2722___mcc_h0 = (_source35).cf39;
              let _2723___mcc_h1 = (_source35).cf40;
              let _2724_cf40 = _2723___mcc_h1;
              let _2725_cf39 = _2722___mcc_h0;
              _2705_v1 = _dafny.Seq.Concat(_dafny.Seq.update(_2705_v1, _module.__default.safeIndex(_2725_cf39, new BigNumber((_2705_v1).length)), _2704_v0), _dafny.Seq.update(_2705_v1, _module.__default.safeIndex((_this).f25, new BigNumber((_2705_v1).length)), new _dafny.CodePoint('u'.codePointAt(0))));
              let _2726_v16;
              _2726_v16 = _dafny.Seq.of(false);
              let _2727_v17;
              let _nw441 = new _module.C9();
              _nw441.__ctor(((true) ? (p3) : (p3)), ((_2718_v12)[_module.__default.safeIndex(new BigNumber((_2726_v16).length), new BigNumber((_2718_v12).length))]).plus(new BigNumber((_2718_v12).length)));
              _2727_v17 = _nw441;
              let _2728_v18;
              _2728_v18 = _dafny.Seq.of(_2719_v13);
              let _2729_v19;
              let _nw442 = new _module.C4();
              _nw442.__ctor((_this).f26, (new BigNumber((_2718_v12).length)).multipliedBy(new BigNumber((_2728_v18).length)), (_this).f26);
              _2729_v19 = _nw442;
              (_2727_v17).f30 = _2727_v17.f30;
            } else {
              let _2730___mcc_h2 = (_source35).cf38;
              let _2731_cf38 = _2730___mcc_h2;
              (globalState).f6 = ((!(p2)) ? (_2718_v12) : (_2718_v12));
              let _2732_v20;
              _2732_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), ((_2733_v7) => function (_2734_i2) {
                return (_2733_v7).f29;
              })(_2712_v7)));
              (globalState).f6 = _dafny.Seq.of(new BigNumber((_2718_v12).length), ((_this).f27).plus(new BigNumber((_module.__default.fm27(p2, _2732_v20, globalState)).length)), (_this).f25);
              let _2735_v21;
              let _nw443 = Array((new BigNumber(10)).toNumber()).fill(false);
              _2735_v21 = _nw443;
              let _2736_v22;
              _2736_v22 = _dafny.Seq.of(_2705_v1, _2705_v1);
              let _index499 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2735_v21).length));
              (_2735_v21)[_index499] = !_dafny.areEqual((_2736_v22)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("uuaexm")).length), new BigNumber((_2736_v22).length))], _2705_v1);
              let _2737_v23;
              _2737_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2716_v10,(_this).f25);
              let _index500 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2735_v21).length));
              (_2735_v21)[_index500] = (_2737_v23).equals(_2737_v23);
              r1 = new BigNumber((_2705_v1).length);
            }
            let _2738_v24;
            _2738_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_2739_i3) {
              return new _dafny.CodePoint('n'.codePointAt(0));
            }),_2720_v14);
            _2738_v24 = (function () {
              let _coll89 = new _dafny.Map();
              for (const _compr_89 of (_dafny.Map.Empty.slice().updateUnsafe(_2705_v1,_2705_v1)).Keys.Elements) {
                let _2740_v25 = _compr_89;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_2705_v1,_2705_v1)).contains(_2740_v25)) {
                  _coll89.push([_2740_v25,_2720_v14]);
                }
              }
              return _coll89;
            }()).update(_2705_v1, _2720_v14);
          }
        }
      }
      let _2741_v26;
      _2741_v26 = _dafny.Seq.of(p2, _module.__default.fm0((_this).f25, globalState));
      let _2742_v27;
      _2742_v27 = _dafny.Map.Empty.slice().updateUnsafe((_2741_v26)[_module.__default.safeIndex((_this).f27, new BigNumber((_2741_v26).length))],p3);
      _2742_v27 = (_2742_v27).update(p2, ((!(p1)) ? (p3) : (p3)));
      r0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f27, ((_this).f27).minus((_this).f27)));
      r1 = _module.__default.safeDivisionInt((_this).fm8(p1, !(_this.f28), globalState), new BigNumber(405));
      return [r0, r1];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2743_v0;
      _2743_v0 = _dafny.Seq.of(p1);
      let _2744_v1;
      _2744_v1 = _module.D0.create_DC0(!((_2743_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_2743_v0).length))]));
      let _source36 = _2744_v1;
      if (_source36.is_DC1) {
        let _2745___mcc_h0 = (_source36).cf1;
        let _2746___mcc_h1 = (_source36).cf2;
        let _2747_cf2 = _2746___mcc_h1;
        let _2748_cf1 = _2745___mcc_h0;
        let _2749_v2;
        _2749_v2 = _dafny.Set.fromElements(_2747_cf2, _2747_cf2, _this.f28, p1, _this.f28);
        if (((_this.f28) ? ((_2749_v2).IsDisjointFrom(_2749_v2)) : (true))) {
          let _2750_v3;
          let _init92 = function (_2751_i0) {
            return (_2751_i0).plus((_this).f27);
          };
          let _nw444 = Array((new BigNumber(4)).toNumber());
          for (let _i0_92 = 0; _i0_92 < new BigNumber(_nw444.length); _i0_92++) {
            _nw444[_i0_92] = _init92(new BigNumber(_i0_92));
          }
          _2750_v3 = _nw444;
          let _2752_v4;
          _2752_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_this.f28);
          let _2753_v5;
          _2753_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,new BigNumber((_2752_v4).length));
          let _2754_v6;
          _2754_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_2753_v5);
          let _2755_v7;
          let _nw445 = new _module.C9();
          _nw445.__ctor(_2750_v3, new BigNumber((_2754_v6).length));
          _2755_v7 = _nw445;
          let _2756_v8;
          _2756_v8 = _dafny.Seq.UnicodeFromString("p");
          let _2757_v9;
          _2757_v9 = _dafny.Seq.of(new BigNumber(679), (_dafny.ZERO).minus(new BigNumber((_2756_v8).length)), (_this).f25, (_this).f25, (_this).f25);
          _2747_cf2 = _dafny.Seq.IsProperPrefixOf(_2757_v9, _dafny.Seq.update(_2757_v9, _module.__default.safeIndex((_2755_v7).f31, new BigNumber((_2757_v9).length)), (_this).f27));
          let _2758_v10;
          let _nw446 = new _module.C0();
          _nw446.__ctor();
          _2758_v10 = _nw446;
          (globalState).f12 = !_dafny.Seq.contains(_dafny.Seq.of((_2755_v7).f31, (_2755_v7).f31, (_this).f27), _module.__default.fm44(_2748_cf1, _2748_cf1, _this.f28, new BigNumber((function () {
            let _coll90 = new _dafny.Map();
            for (const _compr_90 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2748_cf1)).Keys.Elements) {
              let _2759_v11 = _compr_90;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2748_cf1)).contains(_2759_v11)) {
                _coll90.push([(_2759_v11).plus((_2755_v7).f31),new _dafny.CodePoint('s'.codePointAt(0))]);
              }
            }
            return _coll90;
          }()).length), globalState));
          let _2760_v12;
          _2760_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2748_cf1,false);
          _2760_v12 = (_2760_v12).update(!_dafny.Seq.contains(_2743_v0, (_module.__default.fm30(_2752_v4, globalState)).dtor_cf8), _2747_cf2);
        } else {
          let _2761_v13;
          _2761_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,true);
          _2761_v13 = (_2761_v13).update((_this).f27, _2747_cf2);
          let _2762_v14;
          _2762_v14 = new _dafny.CodePoint('u'.codePointAt(0));
          let _2763_v15;
          _2763_v15 = _dafny.MultiSet.fromElements(_2762_v14, _2762_v14);
          let _2764_v16;
          _2764_v16 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2763_v15).Intersect(_2763_v15));
          let _2765_v17;
          _2765_v17 = _dafny.MultiSet.fromElements((_this).f25);
          let _2766_v18;
          _2766_v18 = _dafny.MultiSet.fromElements(p1);
          _2764_v16 = _module.__default.fm48(_2748_cf1, _module.__default.safeDivisionInt(new BigNumber((_2765_v17).cardinality()), new BigNumber(-146)), (_this).f25, new BigNumber(((_2766_v18).Intersect(_2766_v18)).cardinality()), globalState);
          r0 = ((p1) ? ((_this).f25) : ((_this).f25));
          (globalState).f23 = _2762_v14;
          _2747_cf2 = (new BigNumber(442)).isLessThanOrEqualTo(new BigNumber((_2749_v2).length));
        }
        let _2767_v19;
        let _nw447 = new _module.C11();
        _nw447.__ctor();
        _2767_v19 = _nw447;
        let _2768_v20;
        _2768_v20 = _dafny.Seq.of(_2767_v19, _2767_v19);
        let _2769_v22;
        _2769_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2748_cf1,new BigNumber((function () {
          let _coll91 = new _dafny.Map();
          for (const _compr_91 of _dafny.IntegerRange(new BigNumber(-774), new BigNumber(769))) {
            let _2770_v21 = _compr_91;
            if (((new BigNumber(-774)).isLessThanOrEqualTo(_2770_v21)) && ((_2770_v21).isLessThan(new BigNumber(769)))) {
              _coll91.push([(_2770_v21).minus((_this).f25),(_this).f27]);
            }
          }
          return _coll91;
        }()).length));
        let _2771_v23;
        _2771_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2768_v20)[_module.__default.safeIndex(new BigNumber((_2769_v22).length), new BigNumber((_2768_v20).length))]);
        let _2772_v24;
        _2772_v24 = _dafny.Map.Empty.slice().updateUnsafe(!(_2747_cf2),(((_2771_v23).contains(_2748_cf1)) ? ((_2771_v23).get(_2748_cf1)) : (_2767_v19)));
        _2772_v24 = (_2772_v24).update(!(p1), _2767_v19);
        let _2773_v25;
        _2773_v25 = _dafny.Seq.of((_this).f25, (_dafny.ZERO).minus((_this).f27), new BigNumber(849));
        let _2774_v26;
        _2774_v26 = _module.D5.create_DC17((_dafny.ZERO).minus(((_dafny.ZERO).minus((_this).f25)).minus((_this).f27)), ((_this).f27).isLessThan((_2773_v25)[_module.__default.safeIndex((_this).f25, new BigNumber((_2773_v25).length))]));
        let _2775_v27;
        _2775_v27 = _dafny.Seq.UnicodeFromString("lfdfdgsi");
        let _2776_v28;
        _2776_v28 = _dafny.Seq.of(_2775_v27);
        let _2777_v29;
        let _nw448 = Array((new BigNumber(4)).toNumber());
        _nw448[(_dafny.ZERO).toNumber()] = !_dafny.areEqual((_2776_v28)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_2776_v28).length))], _2775_v27);
        _nw448[(_dafny.ONE).toNumber()] = true;
        _nw448[(new BigNumber(2)).toNumber()] = !(!(p1));
        _nw448[(new BigNumber(3)).toNumber()] = ((_this).f27).isLessThan((_this).f27);
        _2777_v29 = _nw448;
        let _index501 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_2777_v29).length));
        (_2777_v29)[_index501] = _this.f28;
        let _2778_v30;
        _2778_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2773_v25).length),_2775_v27);
        let _pat_let_tv65 = _2748_cf1;
        let _index502 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_2777_v29).length));
        let _rhs476 = (_this).f25;
        let _rhs477 = _dafny.Seq.IsPrefixOf(_2775_v27, (((_2778_v30).contains((_this).f27)) ? ((_2778_v30).get((_this).f27)) : (_2775_v27)));
        let _rhs478 = function (_pat_let62_0) {
          return function (_2779_dt__update__tmp_h0) {
            return function (_pat_let63_0) {
              return function (_2780_dt__update_hcf27_h0) {
                return _module.D5.create_DC17((_2779_dt__update__tmp_h0).dtor_cf26, _2780_dt__update_hcf27_h0);
              }(_pat_let63_0);
            }(_pat_let_tv65);
          }(_pat_let62_0);
        }(_2774_v26);
        let _rhs479 = ((_this).f27).isLessThan((_this).f25);
        let _lhs370 = _2777_v29;
        let _lhs371 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_2777_v29).length));
        r0 = _rhs476;
        _2748_cf1 = _rhs477;
        _2774_v26 = _rhs478;
        _lhs370[_lhs371] = _rhs479;
        let _index503 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_2777_v29).length));
        (_2777_v29)[_index503] = _2747_cf2;
      } else if (_source36.is_DC2) {
        let _2781___mcc_h2 = (_source36).cf3;
        let _2782_cf3 = _2781___mcc_h2;
        let _2783_v31;
        _2783_v31 = _dafny.Set.fromElements((_this).f27);
        _2782_cf3 = ((_2783_v31).Intersect(_dafny.Set.fromElements((_this).f25))).IsSubsetOf((_2783_v31).Difference(_2783_v31));
        let _2784_v32;
        _2784_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,((_dafny.ZERO).minus((_this).f27)).multipliedBy(new BigNumber(716)));
        _2784_v32 = (_2784_v32).update((_this).f27, (_this).f25);
        r0 = (_this).f25;
        let _2785_v33;
        _2785_v33 = _dafny.Set.fromElements(_this.f28);
        let _2786_v34;
        _2786_v34 = _dafny.MultiSet.fromElements(_2785_v33, _dafny.Set.fromElements(_2782_cf3));
        let _2787_v35;
        _2787_v35 = _dafny.Seq.of(_2785_v33);
        let _2788_v36;
        _2788_v36 = _dafny.MultiSet.fromElements(_2786_v34, _dafny.MultiSet.FromArray(_2787_v35));
        let _2789_v37;
        _2789_v37 = _dafny.Seq.UnicodeFromString("cjamaf");
        let _2790_v38;
        _2790_v38 = new _dafny.CodePoint('n'.codePointAt(0));
        let _2791_v39;
        _2791_v39 = _dafny.MultiSet.fromElements(_2790_v38, _2790_v38, _2790_v38, new _dafny.CodePoint('n'.codePointAt(0)), _2790_v38);
        let _2792_v40;
        _2792_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_this.f28);
        let _2793_v42;
        _2793_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2744_v1,p1);
        let _2794_v43;
        let _nw449 = Array((new BigNumber(21)).toNumber());
        _nw449[(_dafny.ZERO).toNumber()] = (_this).f25;
        _nw449[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f27);
        _nw449[(new BigNumber(2)).toNumber()] = (((_2788_v36).contains(_2786_v34)) ? ((_2788_v36).get(_2786_v34)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_2782_cf3)).length)));
        _nw449[(new BigNumber(3)).toNumber()] = (_this).f25;
        _nw449[(new BigNumber(4)).toNumber()] = _module.__default.fm31(_2789_v37, globalState);
        _nw449[(new BigNumber(5)).toNumber()] = (_this).f27;
        _nw449[(new BigNumber(6)).toNumber()] = (_this).f25;
        _nw449[(new BigNumber(7)).toNumber()] = new BigNumber(89);
        _nw449[(new BigNumber(8)).toNumber()] = (_this).f27;
        _nw449[(new BigNumber(9)).toNumber()] = new BigNumber((_2791_v39).cardinality());
        _nw449[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw449[(new BigNumber(11)).toNumber()] = (_this).f25;
        _nw449[(new BigNumber(12)).toNumber()] = new BigNumber((function () {
          let _coll92 = new _dafny.Set();
          for (const _compr_92 of (_2792_v40).Keys.Elements) {
            let _2795_v41 = _compr_92;
            if ((_2792_v40).contains(_2795_v41)) {
              _coll92.add(_module.__default.safeModuloInt(_2795_v41, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, !(true))).length))));
            }
          }
          return _coll92;
        }()).length);
        _nw449[(new BigNumber(13)).toNumber()] = (((_2784_v32).contains(_module.__default.fm21(_2793_v42, new BigNumber((p0).length), _2782_cf3, _2789_v37, globalState))) ? ((_2784_v32).get(_module.__default.fm21(_2793_v42, new BigNumber((p0).length), _2782_cf3, _2789_v37, globalState))) : ((_this).f25));
        _nw449[(new BigNumber(14)).toNumber()] = new BigNumber((_2784_v32).length);
        _nw449[(new BigNumber(15)).toNumber()] = (_this).f27;
        _nw449[(new BigNumber(16)).toNumber()] = (_this).f27;
        _nw449[(new BigNumber(17)).toNumber()] = new BigNumber((_2785_v33).length);
        _nw449[(new BigNumber(18)).toNumber()] = new BigNumber((_2783_v31).length);
        _nw449[(new BigNumber(19)).toNumber()] = (_this).f27;
        _nw449[(new BigNumber(20)).toNumber()] = (_this).f27;
        _2794_v43 = _nw449;
        let _2796_v44;
        _2796_v44 = _dafny.Seq.of((_this).f27, new BigNumber(551));
        let _2797_v45;
        _2797_v45 = _dafny.MultiSet.fromElements((_2796_v44)[_module.__default.safeIndex((_this).f27, new BigNumber((_2796_v44).length))], (_this).f27, (_this).f27);
        let _2798_v46;
        _2798_v46 = _module.D5.create_DC15(_2797_v45);
        let _2799_v47;
        _2799_v47 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f25, (_this).f25, new BigNumber((_2792_v40).length)), _2797_v45, (_2798_v46).dtor_cf24);
        let _2800_v48;
        _2800_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2794_v43,(_2799_v47)[_module.__default.safeIndex((_this).f27, new BigNumber((_2799_v47).length))]);
        _2800_v48 = _2800_v48;
      } else if (_source36.is_DC3) {
        let _2801___mcc_h3 = (_source36).cf4;
        let _2802___mcc_h4 = (_source36).cf5;
        let _2803_cf5 = _2802___mcc_h4;
        let _2804_cf4 = _2801___mcc_h3;
        let _2805_v49;
        let _nw450 = Array((new BigNumber(23)).toNumber()).fill(false);
        _2805_v49 = _nw450;
        let _2806_v50;
        _2806_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_2804_cf4, globalState),_2805_v49);
        _2806_v50 = (_2806_v50).update(_2803_cf5, _2805_v49);
        let _2807_v51;
        _2807_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), function (_2808_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }));
        (globalState).f10 = (((_2807_v51).contains(p1)) ? ((_2807_v51).get(p1)) : (_module.__default.fm2(new BigNumber(267), (_this).f27, globalState)));
        let _index504 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_2805_v49).length));
        (_2805_v49)[_index504] = (_2743_v0)[_module.__default.safeIndex(_2804_cf4, new BigNumber((_2743_v0).length))];
        let _index505 = _module.__default.safeIndex(new BigNumber(619), new BigNumber((_2805_v49).length));
        (_2805_v49)[_index505] = !(!(_2803_cf5));
        (globalState).f23 = _module.__default.fm33(_2804_cf4, !(_2803_cf5), (_this).fm8((_2805_v49)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_2805_v49).length))], (_2805_v49)[_module.__default.safeIndex(new BigNumber(619), new BigNumber((_2805_v49).length))], globalState), globalState);
      } else {
        let _2809___mcc_h5 = (_source36).cf0;
        let _2810_cf0 = _2809___mcc_h5;
        let _2811_v52;
        _2811_v52 = _dafny.Set.fromElements(p1, !(_this.f28), _2810_cf0);
        let _2812_v53;
        _2812_v53 = _dafny.Seq.of(_2811_v52, _2811_v52, _2811_v52, _2811_v52, _2811_v52);
        _2812_v53 = _2812_v53;
        let _2813_v54;
        let _nw451 = Array((new BigNumber(3)).toNumber()).fill(false);
        _2813_v54 = _nw451;
        let _index506 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2813_v54).length));
        (_2813_v54)[_index506] = !(p1);
        let _index507 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2813_v54).length));
        (_2813_v54)[_index507] = _2810_cf0;
        if (((_this).f27).isLessThanOrEqualTo((_this).f25)) {
          let _index508 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2813_v54).length));
          (_2813_v54)[_index508] = !(((new BigNumber(302)).multipliedBy((_this).f27)).isLessThanOrEqualTo((_this).f25));
          let _rhs480 = (_2813_v54)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_2813_v54).length))];
          let _rhs481 = !(p1);
          let _rhs482 = (_this).f27;
          let _rhs483 = ((_this).f25).minus((_this).f25);
          let _lhs372 = globalState;
          let _lhs373 = globalState;
          _lhs372.f12 = _rhs480;
          _lhs373.f12 = _rhs481;
          r0 = _rhs482;
          r0 = _rhs483;
          r0 = new BigNumber(454);
          (_this).f28 = (new BigNumber(506)).isLessThan(_module.__default.safeModuloInt((_this).f25, (_this).f27));
          let _2814_v55;
          let _init93 = ((_2815_v0) => function (_2816_i2) {
            return (_2816_i2).plus(new BigNumber((_2815_v0).length));
          })(_2743_v0);
          let _nw452 = Array((new BigNumber(2)).toNumber());
          for (let _i0_93 = 0; _i0_93 < new BigNumber(_nw452.length); _i0_93++) {
            _nw452[_i0_93] = _init93(new BigNumber(_i0_93));
          }
          _2814_v55 = _nw452;
          let _2817_v56;
          _2817_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f25);
          let _rhs484 = _this.f28;
          let _rhs485 = (((new BigNumber(859)).isLessThanOrEqualTo(new BigNumber((_2817_v56).length))) ? (_2814_v55) : (_2814_v55));
          let _rhs486 = (_dafny.ZERO).minus((_this).f25);
          let _rhs487 = p1;
          let _rhs488 = (_2743_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_2743_v0).length))];
          let _lhs374 = _this;
          let _lhs375 = _this;
          let _lhs376 = _this;
          _lhs374.f28 = _rhs484;
          _2814_v55 = _rhs485;
          r0 = _rhs486;
          _lhs375.f28 = _rhs487;
          _lhs376.f28 = _rhs488;
        } else {
          r0 = (_this).f25;
          let _2818_v57;
          _2818_v57 = _dafny.Map.Empty.slice().updateUnsafe((_2743_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_2743_v0).length))],(_this).f25);
          _2818_v57 = (_2818_v57).Merge((_2818_v57).Merge(_2818_v57));
          let _2819_v58;
          let _nw453 = Array((new BigNumber(7)).toNumber()).fill([]);
          _2819_v58 = _nw453;
          let _2820_v59;
          let _nw454 = Array((new BigNumber(8)).toNumber());
          _nw454[(_dafny.ZERO).toNumber()] = _2819_v58;
          _nw454[(_dafny.ONE).toNumber()] = _2819_v58;
          _nw454[(new BigNumber(2)).toNumber()] = _2819_v58;
          _nw454[(new BigNumber(3)).toNumber()] = _2819_v58;
          _nw454[(new BigNumber(4)).toNumber()] = _2819_v58;
          _nw454[(new BigNumber(5)).toNumber()] = _2819_v58;
          _nw454[(new BigNumber(6)).toNumber()] = _2819_v58;
          _nw454[(new BigNumber(7)).toNumber()] = _2819_v58;
          _2820_v59 = _nw454;
          let _index509 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_2820_v59).length));
          (_2820_v59)[_index509] = _2819_v58;
          let _2821_v60;
          _2821_v60 = new _dafny.CodePoint('o'.codePointAt(0));
          let _2822_v61;
          _2822_v61 = _dafny.Seq.UnicodeFromString("m");
          let _2823_v62;
          _2823_v62 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(701)), ((_2824_v60) => function (_2825_i3) {
            return _2824_v60;
          })(_2821_v60)),_2819_v58);
          let _index510 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_2820_v59).length));
          (_2820_v59)[_index510] = ((_dafny.Seq.contains(_2822_v61, _2821_v60)) ? (_2819_v58) : ((((_2823_v62).contains(_2822_v61)) ? ((_2823_v62).get(_2822_v61)) : (_2819_v58))));
          let _2826_v63;
          _2826_v63 = _module.D14.create_DC38((_this).f25);
          let _2827_v64;
          _2827_v64 = _dafny.MultiSet.fromElements(_2810_cf0);
          _2826_v63 = _module.__default.fm70((_2813_v54)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_2813_v54).length))], ((_this).f25).minus(new BigNumber((_2827_v64).cardinality())), (_2813_v54)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_2813_v54).length))], globalState);
          (globalState).f12 = (((p1) ? ((_this).f27) : (new BigNumber(-488)))).isLessThanOrEqualTo((_this).f25);
        }
        let _2828_v65;
        let _nw455 = new _module.C11();
        _nw455.__ctor();
        _2828_v65 = _nw455;
        _2828_v65 = _2828_v65;
      }
      let _2829_v66;
      _2829_v66 = _dafny.MultiSet.fromElements(p1);
      (_this).f28 = (_2829_v66).equals((_2829_v66).Difference(_dafny.MultiSet.fromElements(!(_this.f28))));
      let _2830_v67;
      _2830_v67 = _dafny.Seq.UnicodeFromString("igxf");
      let _2831_v68;
      _2831_v68 = _dafny.Set.fromElements(_2830_v67, _2830_v67);
      let _2832_v69;
      _2832_v69 = _dafny.Seq.of(new BigNumber((_2830_v67).length), (_this).f27, new BigNumber((_2831_v68).length));
      let _2833_v70;
      let _nw456 = new _module.C6();
      _nw456.__ctor(new BigNumber((_dafny.Seq.Concat(_2832_v69, _2832_v69)).length), (_this).f26);
      _2833_v70 = _nw456;
      let _pat_let_tv66 = _2830_v67;
      let _pat_let_tv67 = _2830_v67;
      let _pat_let_tv68 = _2830_v67;
      let _pat_let_tv69 = p1;
      let _pat_let_tv70 = _2830_v67;
      let _pat_let_tv71 = _2830_v67;
      let _source37 = function (_source38) {
        if (_source38.is_DC5) {
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv66,_this.f28);
        } else if (_source38.is_DC6) {
          let _2834___mcc_h7 = (_source38).cf7;
          let _2835___mcc_h8 = (_source38).cf8;
          let _2836___mcc_h9 = (_source38).cf9;
          let _2837_cf9 = _2836___mcc_h9;
          let _2838_cf8 = _2835___mcc_h8;
          let _2839_cf7 = _2834___mcc_h7;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv67,_this.f28);
        } else if (_source38.is_DC4) {
          let _2840___mcc_h10 = (_source38).cf6;
          let _2841_cf6 = _2840___mcc_h10;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv68,false);
        } else {
          let _2842___mcc_h11 = (_source38).cf10;
          let _2843_cf10 = _2842___mcc_h11;
          if (_pat_let_tv69) {
            return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv70,_this.f28);
          } else {
            return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv71,false);
          }
        }
      }(_module.D1.create_DC4(_2743_v0));
      let _2844___mcc_h6 = _source37;
      let _2845_cf61 = _2844___mcc_h6;
      let _2846_v71;
      _2846_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(p1) || (p1),_module.__default.safeDivisionInt((_this).f27, (_this).f27));
      _2846_v71 = _2846_v71;
      let _2847_v72;
      let _nw457 = new _module.C4();
      _nw457.__ctor((_this).f26, (_this).f25, (_this).f26);
      _2847_v72 = _nw457;
      let _2848_v73;
      _2848_v73 = _dafny.Seq.of(_2847_v72);
      let _2849_v74;
      _2849_v74 = _dafny.Set.fromElements(_2848_v73, _2848_v73, _2848_v73);
      let _2850_v75;
      _2850_v75 = _dafny.Seq.of(_2848_v73);
      let _2851_v76;
      let _init94 = ((_2852_p1) => function (_2853_i4) {
        return _2852_p1;
      })(p1);
      let _nw458 = Array((new BigNumber(22)).toNumber());
      for (let _i0_94 = 0; _i0_94 < new BigNumber(_nw458.length); _i0_94++) {
        _nw458[_i0_94] = _init94(new BigNumber(_i0_94));
      }
      _2851_v76 = _nw458;
      let _2854_v77;
      let _nw459 = new _module.C5();
      _nw459.__ctor(p1);
      _2854_v77 = _nw459;
      let _2855_v78;
      _2855_v78 = _dafny.Map.Empty.slice().updateUnsafe(_2851_v76,_2854_v77);
      let _2856_v79;
      let _nw460 = Array((new BigNumber(16)).toNumber());
      _nw460[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_2848_v73);
      _nw460[(_dafny.ONE).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(2)).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_dafny.Seq.of(_2847_v72));
      _nw460[(new BigNumber(4)).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(5)).toNumber()] = (_2849_v74).Union(_2849_v74);
      _nw460[(new BigNumber(6)).toNumber()] = (_2849_v74).Difference(_2849_v74);
      _nw460[(new BigNumber(7)).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(8)).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(9)).toNumber()] = (_2849_v74).Intersect(_2849_v74);
      _nw460[(new BigNumber(10)).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(11)).toNumber()] = (_dafny.Set.fromElements(_2848_v73, _2848_v73)).Difference(_2849_v74);
      _nw460[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(_2848_v73, (_2850_v75)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_2850_v75).length))], _2848_v73, _dafny.Seq.update(_dafny.Seq.of(_2847_v72), _module.__default.safeIndex(new BigNumber(((_2855_v78).update(_2851_v76, _2854_v77)).length), new BigNumber((_dafny.Seq.of(_2847_v72)).length)), _2847_v72), _2848_v73);
      _nw460[(new BigNumber(13)).toNumber()] = _2849_v74;
      _nw460[(new BigNumber(14)).toNumber()] = (_2849_v74).Intersect(_2849_v74);
      _nw460[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(_2848_v73, _2848_v73);
      _2856_v79 = _nw460;
      let _index511 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_2856_v79).length));
      (_2856_v79)[_index511] = ((_this.f28) ? (_dafny.Set.fromElements(_dafny.Seq.of(_2847_v72, _2847_v72, _2847_v72))) : (_2849_v74));
      let _2857_v80;
      _2857_v80 = _module.D4.create_DC13(new BigNumber(777), false, (_2854_v77).f35, true);
      let _2858_v81;
      _2858_v81 = _dafny.Map.Empty.slice().updateUnsafe((_2857_v80).dtor_cf19,_2851_v76);
      let _2859_v82;
      _2859_v82 = new _dafny.CodePoint('n'.codePointAt(0));
      let _2860_v83;
      _2860_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2859_v82);
      let _2861_v84;
      _2861_v84 = _dafny.Set.fromElements(_2851_v76);
      let _pat_let_tv72 = _2849_v74;
      let _index512 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_2856_v79).length));
      let _rhs489 = new BigNumber((((_2858_v81).update((_2847_v72).f25, _2851_v76)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2860_v83).length),_2851_v76))).length);
      let _rhs490 = _module.__default.safeDivisionInt((_this).f25, new BigNumber((_2861_v84).length));
      let _rhs491 = (function (_pat_let64_0) {
        return function (_2862_dt__update__tmp_h1) {
          return function (_pat_let65_0) {
            return function (_2863_dt__update_hcf76_h0) {
              return _module.D21.create_DC51(_2863_dt__update_hcf76_h0);
            }(_pat_let65_0);
          }(_pat_let_tv72);
        }(_pat_let64_0);
      }(_module.D21.create_DC51(_2849_v74))).dtor_cf76;
      let _lhs377 = _2856_v79;
      let _lhs378 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_2856_v79).length));
      r0 = _rhs489;
      r0 = _rhs490;
      _lhs377[_lhs378] = _rhs491;
      (globalState).f6 = _2832_v69;
      let _2864_v85;
      _2864_v85 = _module.D3.create_DC9(_2861_v84);
      let _2865_v86;
      _2865_v86 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2864_v85);
      let _2866_v87;
      _2866_v87 = _module.D19.create_DC46(_2865_v86);
      let _source39 = _2866_v87;
      if (_source39.is_DC47) {
        let _2867___mcc_h12 = (_source39).cf69;
        let _2868___mcc_h13 = (_source39).cf70;
        let _2869_cf70 = _2868___mcc_h13;
        let _2870_cf69 = _2867___mcc_h12;
        let _2871_v88;
        _2871_v88 = _dafny.Set.fromElements(_dafny.Seq.update(_2832_v69, _module.__default.safeIndex((_this).f25, new BigNumber((_2832_v69).length)), (_this).f25));
        (globalState).f12 = (((_2871_v88).IsDisjointFrom(_dafny.Set.fromElements(_2832_v69, _dafny.Seq.Create(_module.__default.abs(new BigNumber(504)), ((_2872_v67) => function (_2873_i5) {
          return new BigNumber((_2872_v67).length);
        })(_2830_v67)), _2832_v69, _2832_v69))) ? (!(((_2847_v72).f25).isLessThan((_2847_v72).f25))) : (_this.f28));
        let _index513 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_2870_cf69).length));
        (_2870_cf69)[_index513] = (_2854_v77).f35;
        let _index514 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_2870_cf69).length));
        (_2870_cf69)[_index514] = _module.__default.fm0((_2847_v72).f25, globalState);
        (globalState).f12 = (_2854_v77).f35;
        _2869_cf70 = _module.__default.fm0(((_2847_v72).f25).minus((_2847_v72).f25), globalState);
      } else if (_source39.is_DC48) {
        let _2874___mcc_h14 = (_source39).cf71;
        let _2875_cf71 = _2874___mcc_h14;
        _2743_v0 = _2743_v0;
        let _index515 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_2851_v76).length));
        (_2851_v76)[_index515] = ((_this.f28) ? ((_2854_v77).f35) : (p1));
        let _2876_v89;
        _2876_v89 = _dafny.Set.fromElements((_this).f25, _2875_cf71);
        let _index516 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_2851_v76).length));
        (_2851_v76)[_index516] = (_2876_v89).IsProperSubsetOf(_2876_v89);
        (globalState).f12 = (_2851_v76)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_2851_v76).length))];
        let _2877_v90;
        let _init95 = function (_2878_i6) {
          return _module.D4.create_DC12();
        };
        let _nw461 = Array((new BigNumber(28)).toNumber());
        for (let _i0_95 = 0; _i0_95 < new BigNumber(_nw461.length); _i0_95++) {
          _nw461[_i0_95] = _init95(new BigNumber(_i0_95));
        }
        _2877_v90 = _nw461;
        let _2879_v91;
        _2879_v91 = _module.D4.create_DC12();
        let _index517 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2877_v90).length));
        (_2877_v90)[_index517] = _2879_v91;
        let _2880_v92;
        _2880_v92 = _dafny.Map.Empty.slice().updateUnsafe(_2743_v0,(_2847_v72).f25);
        let _2881_v93;
        let _nw462 = Array((new BigNumber(3)).toNumber());
        _nw462[(_dafny.ZERO).toNumber()] = (_2854_v77).f35;
        _nw462[(_dafny.ONE).toNumber()] = true;
        _nw462[(new BigNumber(2)).toNumber()] = (_2854_v77).f35;
        _2881_v93 = _nw462;
        let _2882_v94;
        let _nw463 = Array((new BigNumber(16)).toNumber());
        _nw463[(_dafny.ZERO).toNumber()] = (_this).f25;
        _nw463[(_dafny.ONE).toNumber()] = (_this).f25;
        _nw463[(new BigNumber(2)).toNumber()] = (_2847_v72).f25;
        _nw463[(new BigNumber(3)).toNumber()] = _module.__default.fm18(globalState);
        _nw463[(new BigNumber(4)).toNumber()] = _2875_cf71;
        _nw463[(new BigNumber(5)).toNumber()] = (_2847_v72).f25;
        _nw463[(new BigNumber(6)).toNumber()] = new BigNumber((_2830_v67).length);
        _nw463[(new BigNumber(7)).toNumber()] = ((_dafny.ZERO).minus((_this).f25)).plus((_2847_v72).f25);
        _nw463[(new BigNumber(8)).toNumber()] = _2875_cf71;
        _nw463[(new BigNumber(9)).toNumber()] = (_2847_v72).f25;
        _nw463[(new BigNumber(10)).toNumber()] = _2875_cf71;
        _nw463[(new BigNumber(11)).toNumber()] = new BigNumber((_2829_v66).cardinality());
        _nw463[(new BigNumber(12)).toNumber()] = ((_2847_v72).f25).multipliedBy((_this).f27);
        _nw463[(new BigNumber(13)).toNumber()] = (((_2880_v92).contains(_2743_v0)) ? ((_2880_v92).get(_2743_v0)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2881_v93,(_dafny.ZERO).minus(_2875_cf71))).length)));
        _nw463[(new BigNumber(14)).toNumber()] = new BigNumber(79);
        _nw463[(new BigNumber(15)).toNumber()] = new BigNumber(((_2876_v89).Intersect(_2876_v89)).length);
        _2882_v94 = _nw463;
        let _index518 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2877_v90).length));
        let _rhs492 = _2859_v82;
        let _rhs493 = _module.D4.create_DC12();
        let _rhs494 = (p0)[_module.__default.safeIndex((_2847_v72).f25, new BigNumber((p0).length))];
        let _lhs379 = globalState;
        let _lhs380 = _2877_v90;
        let _lhs381 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_2877_v90).length));
        _lhs379.f23 = _rhs492;
        _lhs380[_lhs381] = _rhs493;
        _2882_v94 = _rhs494;
      } else {
        let _2883___mcc_h15 = (_source39).cf68;
        let _2884_cf68 = _2883___mcc_h15;
        let _2885_v95;
        _2885_v95 = _dafny.Map.Empty.slice().updateUnsafe((((_2846_v71).contains((_2854_v77).f35)) ? ((_2846_v71).get((_2854_v77).f35)) : ((_this).f25)),_dafny.Seq.of((_2854_v77).f35));
        let _index519 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2851_v76).length));
        (_2851_v76)[_index519] = _dafny.Seq.IsPrefixOf((((_2885_v95).contains((_2847_v72).f25)) ? ((_2885_v95).get((_2847_v72).f25)) : (_dafny.Seq.of(p1, _this.f28))), _2743_v0);
        let _index520 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2851_v76).length));
        (_2851_v76)[_index520] = _this.f28;
        let _index521 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2851_v76).length));
        (_2851_v76)[_index521] = p1;
        r0 = (((_2854_v77).f35) ? (_module.__default.safeDivisionInt(new BigNumber(-200), new BigNumber((_2830_v67).length))) : ((_2847_v72).f25));
        let _2886_v96;
        let _nw464 = new _module.C0();
        _nw464.__ctor();
        _2886_v96 = _nw464;
      }
      if (!(_dafny.Seq.IsPrefixOf(_2830_v67, ((p1) ? (_2830_v67) : (_module.__default.fm2(new BigNumber(551), (_this).f27, globalState)))))) {
        let _2887_v97;
        _2887_v97 = new _dafny.CodePoint('y'.codePointAt(0));
        (globalState).f10 = (_module.D1.create_DC6(_2887_v97, (_2743_v0)[_module.__default.safeIndex((_this).f27, new BigNumber((_2743_v0).length))], _2830_v67)).dtor_cf9;
        (globalState).f12 = !(p1);
        let _2888_v98;
        _2888_v98 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f28);
        let _2889_v99;
        _2889_v99 = _module.D21.create_DC52(_module.__default.safeDivisionInt(new BigNumber(164), new BigNumber((_2888_v98).length)), ((_dafny.ZERO).minus((_this).f25)).multipliedBy((_this).f25), _2830_v67, _this.f28, (_this).f27);
        let _pat_let_tv73 = _2830_v67;
        let _pat_let_tv74 = _2830_v67;
        _2889_v99 = function (_pat_let66_0) {
          return function (_2890_dt__update__tmp_h2) {
            return function (_pat_let67_0) {
              return function (_2891_dt__update_hcf80_h0) {
                return function (_pat_let68_0) {
                  return function (_2892_dt__update_hcf79_h0) {
                    return function (_pat_let69_0) {
                      return function (_2893_dt__update_hcf81_h0) {
                        return _module.D21.create_DC52((_2890_dt__update__tmp_h2).dtor_cf77, (_2890_dt__update__tmp_h2).dtor_cf78, _2892_dt__update_hcf79_h0, _2891_dt__update_hcf80_h0, _2893_dt__update_hcf81_h0);
                      }(_pat_let69_0);
                    }(_module.__default.safeDivisionInt((_this).f27, (_this).f25));
                  }(_pat_let68_0);
                }(_dafny.Seq.Concat(_pat_let_tv73, _pat_let_tv74));
              }(_pat_let67_0);
            }(!(false));
          }(_pat_let66_0);
        }(_2889_v99);
        let _2894_v100;
        _2894_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length),_dafny.Set.fromElements(_this.f28));
        let _2895_v101;
        _2895_v101 = _dafny.Set.fromElements(_this.f28, _this.f28, false);
        let _2896_v102;
        _2896_v102 = _dafny.Map.Empty.slice().updateUnsafe(_2887_v97,p1);
        let _2897_v103;
        _2897_v103 = _dafny.Set.fromElements((_this).f25, (_this).f25, (_dafny.ZERO).minus(new BigNumber(((_2894_v100).update((_this).f25, _2895_v101)).length)), new BigNumber((_2896_v102).length), (_this).f25);
        let _2898_v104;
        _2898_v104 = _module.D4.create_DC11(_2897_v103);
        _2898_v104 = _2898_v104;
        (globalState).f12 = p1;
      } else {
        let _2899_v105;
        let _nw465 = new _module.C7();
        _nw465.__ctor((_this).f27);
        _2899_v105 = _nw465;
        let _2900_v106;
        _2900_v106 = _dafny.Map.Empty.slice().updateUnsafe((_2899_v105).f34,true);
        let _2901_v107;
        _2901_v107 = _dafny.Set.fromElements(!(false), p1, (((_2900_v106).contains((_2899_v105).f34)) ? ((_2900_v106).get((_2899_v105).f34)) : (p1)), _this.f28, false);
        let _2902_v108;
        _2902_v108 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f25, (_2899_v105).f34),(new BigNumber((_module.__default.fm2((_this).f25, new BigNumber((_2743_v0).length), globalState)).length)).multipliedBy(new BigNumber((_2901_v107).length)));
        _2902_v108 = (_2902_v108).update((new BigNumber(190)).multipliedBy((_this).f27), _module.__default.safeDivisionInt((_this).f25, (_2899_v105).f34));
        let _2903_v109;
        let _nw466 = Array((new BigNumber(4)).toNumber()).fill([]);
        _2903_v109 = _nw466;
        let _2904_v110;
        let _init96 = ((_2905_p1, _2906_v105) => function (_2907_i7) {
          return (_module.D13.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(false,_2905_p1), _this.f28, (_2906_v105).f34)).dtor_cf55;
        })(p1, _2899_v105);
        let _nw467 = Array((new BigNumber(23)).toNumber());
        for (let _i0_96 = 0; _i0_96 < new BigNumber(_nw467.length); _i0_96++) {
          _nw467[_i0_96] = _init96(new BigNumber(_i0_96));
        }
        _2904_v110 = _nw467;
        let _index522 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_2903_v109).length));
        (_2903_v109)[_index522] = _2904_v110;
        let _2908_v111;
        _2908_v111 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _index523 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_2903_v109).length));
        (_2903_v109)[_index523] = (_module.D3.create_DC10(p1, _2904_v110, (((_2908_v111).contains(p1)) ? ((_2908_v111).get(p1)) : (_this.f28)), (_2899_v105).f34, (_this).f27)).dtor_cf14;
        let _2909_v112;
        let _nw468 = new _module.C4();
        _nw468.__ctor((_this).f26, new BigNumber((_2901_v107).length), (_this).f26);
        _2909_v112 = _nw468;
        let _2910_v113;
        _2910_v113 = _dafny.Set.fromElements(_2902_v108, (_2902_v108).update(new BigNumber(-244), (_dafny.ZERO).minus(new BigNumber((_2830_v67).length))));
        let _2911_v114;
        let _nw469 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _2911_v114 = _nw469;
        let _2912_v115;
        _2912_v115 = _module.D21.create_DC53(_2910_v113, p1, _2911_v114);
        let _2913_v116;
        _2913_v116 = _dafny.Map.Empty.slice().updateUnsafe(false,_2912_v115);
        let _2914_v118;
        _2914_v118 = new _dafny.CodePoint('u'.codePointAt(0));
        let _2915_v119;
        _2915_v119 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2914_v118);
        let _2916_v120;
        let _nw470 = Array((new BigNumber(17)).toNumber());
        _nw470[(_dafny.ZERO).toNumber()] = _2902_v108;
        _nw470[(_dafny.ONE).toNumber()] = (_2902_v108).Merge(_module.__default.fm57(p1, p1, (_this).f25, globalState));
        _nw470[(new BigNumber(2)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(3)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sqkjki")).length)),(_this).f27);
        _nw470[(new BigNumber(5)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(6)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2830_v67).length),new BigNumber((_2913_v116).length));
        _nw470[(new BigNumber(8)).toNumber()] = _module.__default.fm57(false, p1, (_this).f25, globalState);
        _nw470[(new BigNumber(9)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(10)).toNumber()] = (_2902_v108).Merge(_2902_v108);
        _nw470[(new BigNumber(11)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(12)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(13)).toNumber()] = function () {
          let _coll93 = new _dafny.Map();
          for (const _compr_93 of _dafny.IntegerRange(new BigNumber(231), new BigNumber(498))) {
            let _2917_v117 = _compr_93;
            if (((new BigNumber(231)).isLessThanOrEqualTo(_2917_v117)) && ((_2917_v117).isLessThan(new BigNumber(498)))) {
              _coll93.push([(_2917_v117).minus(new BigNumber(-424)),new BigNumber(398)]);
            }
          }
          return _coll93;
        }();
        _nw470[(new BigNumber(14)).toNumber()] = _2902_v108;
        _nw470[(new BigNumber(15)).toNumber()] = ((_2902_v108).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2915_v119).length),p1)).length), new BigNumber((_2829_v66).cardinality()))).Merge(_2902_v108);
        _nw470[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_2899_v105).f34);
        _2916_v120 = _nw470;
        let _index524 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_2916_v120).length));
        (_2916_v120)[_index524] = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f27);
        let _2918_v121;
        _2918_v121 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2899_v105).f34),_2830_v67);
        let _2919_v122;
        _2919_v122 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm44(_this.f28, _this.f28, false, (_2899_v105).f34, globalState),_2904_v110);
        let _2920_v123;
        _2920_v123 = _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe((((_2919_v122).contains((_this).f27)) ? ((_2919_v122).get((_this).f27)) : ((_2903_v109)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_2903_v109).length))])),p1));
        let _2921_v124;
        _2921_v124 = _module.D8.create_DC25(_2920_v123);
        let _2922_v125;
        _2922_v125 = _module.D8.create_DC25(_2921_v124);
        let _pat_let_tv75 = _2920_v123;
        let _2923_v126;
        _2923_v126 = _dafny.Seq.of(_module.D8.create_DC25(_2920_v123), _2922_v125, _2922_v125, function (_pat_let70_0) {
          return function (_2924_dt__update__tmp_h3) {
            return function (_pat_let71_0) {
              return function (_2925_dt__update_hcf37_h0) {
                return _module.D8.create_DC25(_2925_dt__update_hcf37_h0);
              }(_pat_let71_0);
            }(_pat_let_tv75);
          }(_pat_let70_0);
        }(_2922_v125));
        let _2926_v128;
        _2926_v128 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_2830_v67, _module.__default.safeIndex((_2899_v105).f34, new BigNumber((_2830_v67).length)), _2914_v118)).length), (_this).f27);
        let _2927_v129;
        _2927_v129 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(942),_2926_v128);
        let _index525 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_2916_v120).length));
        let _rhs495 = function () {
          let _coll94 = new _dafny.Map();
          for (const _compr_94 of _dafny.IntegerRange(new BigNumber(937), new BigNumber(61))) {
            let _2928_v127 = _compr_94;
            if (((new BigNumber(937)).isLessThanOrEqualTo(_2928_v127)) && ((_2928_v127).isLessThan(new BigNumber(61)))) {
              _coll94.push([(_2928_v127).minus((_2899_v105).f34),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_2743_v0).length))).length)]);
            }
          }
          return _coll94;
        }();
        let _rhs496 = (_this).f27;
        let _rhs497 = ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm8(p1, true, globalState)),_dafny.Seq.UnicodeFromString("lf"))).Merge(_2918_v121)).update((_dafny.ZERO).minus((_this).f27), _dafny.Seq.Concat(_2830_v67, _dafny.Seq.UnicodeFromString("h")));
        let _rhs498 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2923_v126, _2923_v126), _2923_v126);
        let _rhs499 = !(p1) || ((new BigNumber(((_module.__default.fm71((((_2927_v129).contains((_this).f27)) ? ((_2927_v129).get((_this).f27)) : (_dafny.MultiSet.FromArray(_2832_v69))), globalState)).dtor_cf6).length)).isLessThanOrEqualTo((_this).f25));
        let _lhs382 = _2916_v120;
        let _lhs383 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_2916_v120).length));
        let _lhs384 = globalState;
        _lhs382[_lhs383] = _rhs495;
        r0 = _rhs496;
        _2918_v121 = _rhs497;
        _2923_v126 = _rhs498;
        _lhs384.f12 = _rhs499;
      }
      let _2929_v130;
      let _nw471 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2929_v130 = _nw471;
      let _2930_v131;
      _2930_v131 = _dafny.Map.Empty.slice().updateUnsafe(_2929_v130,((_this).f27).isLessThan((_this).f27));
      _2930_v131 = (_2930_v131).update(_2929_v130, (true) && (false));
      r0 = new BigNumber(-857);
      return r0;
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _2931_v0;
      _2931_v0 = new _dafny.CodePoint('u'.codePointAt(0));
      (globalState).f23 = _2931_v0;
      let _2932_v1;
      let _init97 = function (_2933_i0) {
        return _this.f28;
      };
      let _nw472 = Array((new BigNumber(14)).toNumber());
      for (let _i0_97 = 0; _i0_97 < new BigNumber(_nw472.length); _i0_97++) {
        _nw472[_i0_97] = _init97(new BigNumber(_i0_97));
      }
      _2932_v1 = _nw472;
      let _2934_v2;
      _2934_v2 = _dafny.Set.fromElements(_2931_v0);
      let _index526 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
      (_2932_v1)[_index526] = !((_2934_v2).IsSubsetOf(_2934_v2));
      let _2935_v3;
      _2935_v3 = new BigNumber(17);
      let _2936_v4;
      _2936_v4 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f28);
      let _2937_v5;
      _2937_v5 = _module.D4.create_DC13(_2935_v3, p3, p3, _this.f28);
      let _2938_v6;
      _2938_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,false);
      let _2939_v7;
      _2939_v7 = _dafny.Seq.of(new BigNumber((_2938_v6).length), _2935_v3);
      let _2940_v8;
      _2940_v8 = _dafny.MultiSet.fromElements(_2931_v0);
      let _2941_v9;
      _2941_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2939_v7).length),_2940_v8);
      let _2942_v10;
      _2942_v10 = _dafny.MultiSet.fromElements(new BigNumber(((((_2941_v9).contains(_2935_v3)) ? ((_2941_v9).get(_2935_v3)) : (_2940_v8))).cardinality()), _2935_v3, (_this).f25);
      let _2943_v11;
      _2943_v11 = _dafny.Seq.of((((_2936_v4).contains((_2937_v5).dtor_cf21)) ? ((_2936_v4).get((_2937_v5).dtor_cf21)) : (p1)), !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-932)), function (_2944_i1) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _2931_v0), (_2942_v10).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(-893))));
      let _index527 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
      let _rhs500 = ((_this).f27).isLessThan(_2935_v3);
      let _rhs501 = (_this).f27;
      let _rhs502 = (_2943_v11)[_module.__default.safeIndex(new BigNumber(802), new BigNumber((_2943_v11).length))];
      let _rhs503 = (_dafny.ZERO).minus(_2935_v3);
      let _lhs385 = _2932_v1;
      let _lhs386 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
      _lhs385[_lhs386] = _rhs500;
      _2935_v3 = _rhs501;
      r1 = _rhs502;
      _2935_v3 = _rhs503;
      if ((p1) || (((_this).f25).isLessThan((_dafny.ZERO).minus(_2935_v3)))) {
        let _2945_v12;
        _2945_v12 = _dafny.MultiSet.fromElements(p3, p1, _this.f28);
        _2936_v4 = (_2936_v4).update((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))], (_dafny.MultiSet.fromElements((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))], p1)).IsProperSubsetOf(_2945_v12));
        (globalState).f12 = !(p1);
        let _index528 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
        (_2932_v1)[_index528] = !((p3) || (((p3) ? ((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))]) : (_this.f28))));
        let _index529 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((p0).length));
        (p0)[_index529] = _2935_v3;
        let _2946_v13;
        _2946_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), function (_2947_i2) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }));
        let _2948_v14;
        _2948_v14 = _dafny.Set.fromElements((_this).f27);
        let _2949_v15;
        _2949_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_dafny.ZERO).minus((_dafny.ZERO).minus(_2935_v3)));
        let _2950_v16;
        _2950_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm44((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))], !(p3), _module.__default.fm0(new BigNumber((_dafny.Seq.UnicodeFromString("cojfkt")).length), globalState), _2935_v3, globalState),(_dafny.ZERO).minus((((_2949_v15).contains(new BigNumber(482))) ? ((_2949_v15).get(new BigNumber(482))) : (new BigNumber(799)))));
        let _2951_v17;
        _2951_v17 = _dafny.Seq.UnicodeFromString("cqj");
        let _2952_v18;
        _2952_v18 = _dafny.Set.fromElements(p3);
        let _2953_v19;
        _2953_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2952_v18,_module.__default.fm44(true, (_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))], _this.f28, _2935_v3, globalState));
        let _index530 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((p0).length));
        let _rhs504 = (_this).f25;
        let _rhs505 = (_2948_v14).IsSubsetOf(_module.__default.fm27(p1, _2946_v13, globalState));
        let _rhs506 = (((_2949_v15).contains(_module.__default.safeModuloInt(new BigNumber((_2951_v17).length), (_this).f25))) ? ((_2949_v15).get(_module.__default.safeModuloInt(new BigNumber((_2951_v17).length), (_this).f25))) : (new BigNumber(((_2953_v19).Merge(_2953_v19)).length)));
        let _rhs507 = true;
        let _lhs387 = p0;
        let _lhs388 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((p0).length));
        let _lhs389 = globalState;
        let _lhs390 = _this;
        _lhs387[_lhs388] = _rhs504;
        _lhs389.f12 = _rhs505;
        _2935_v3 = _rhs506;
        _lhs390.f28 = _rhs507;
        let _2954_v20;
        _2954_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2931_v0,(_this).f27);
        _2954_v20 = (_2954_v20).update(_2931_v0, (_this).f25);
      } else {
        let _2955_v21;
        _2955_v21 = _dafny.Seq.of(p0);
        let _2956_v22;
        _2956_v22 = _dafny.MultiSet.fromElements(p3, (new BigNumber((_dafny.Set.fromElements((_2955_v21)[_module.__default.safeIndex((_this).f27, new BigNumber((_2955_v21).length))])).length)).isEqualTo((_this).f27), true);
        _2956_v22 = (_2956_v22).Difference((_2956_v22).Difference(_2956_v22));
        let _2957_v23;
        let _init98 = ((_2958_v0) => function (_2959_i3) {
          return _module.__default.safeModuloInt(_2959_i3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(499)), ((_2960_v0) => function (_2961_i4) {
            return _2960_v0;
          })(_2958_v0))).length));
        })(_2931_v0);
        let _nw473 = Array((new BigNumber(12)).toNumber());
        for (let _i0_98 = 0; _i0_98 < new BigNumber(_nw473.length); _i0_98++) {
          _nw473[_i0_98] = _init98(new BigNumber(_i0_98));
        }
        _2957_v23 = _nw473;
        _2957_v23 = p0;
        if ((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))]) {
          let _2962_v24;
          _2962_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2931_v0);
          let _2963_v25;
          let _nw474 = new _module.C9();
          _nw474.__ctor(p0, _2935_v3);
          _2963_v25 = _nw474;
          let _2964_v26;
          _2964_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))],_2963_v25);
          let _2965_v27;
          _2965_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2962_v24,_2964_v26);
          let _2966_v29;
          _2966_v29 = _dafny.Map.Empty.slice().updateUnsafe((_2965_v27).update(_2962_v24, _2964_v26),function () {
            let _coll95 = new _dafny.Set();
            for (const _compr_95 of (_module.__default.fm72(p3, p3, (_2963_v25).f31, _2935_v3, globalState)).Elements) {
              let _2967_v28 = _compr_95;
              if (_dafny.Seq.contains(_module.__default.fm72(p3, p3, (_2963_v25).f31, _2935_v3, globalState), _2967_v28)) {
                _coll95.add(_2967_v28);
              }
            }
            return _coll95;
          }());
          let _2968_v30;
          _2968_v30 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-868)), function (_2969_i5) {
            return (_this).f27;
          }));
          _2966_v29 = (_2966_v29).update((_dafny.Map.Empty.slice().updateUnsafe(_2962_v24,_2964_v26)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2962_v24,_2964_v26)), _2968_v30);
          (globalState).f12 = p1;
          (_this).f28 = !(!((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))]));
          let _2970_v31;
          _2970_v31 = _dafny.Seq.UnicodeFromString("fjfg");
          let _rhs508 = (((_2943_v11)[_module.__default.safeIndex(_2935_v3, new BigNumber((_2943_v11).length))]) ? (_2970_v31) : (_2970_v31));
          let _rhs509 = (_2963_v25).f31;
          let _rhs510 = new BigNumber((_dafny.Seq.Concat(_2970_v31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-750)), ((_2971_v0) => function (_2972_i6) {
            return _2971_v0;
          })(_2931_v0)))).length);
          r2 = _rhs508;
          _2935_v3 = _rhs509;
          _2935_v3 = _rhs510;
          let _index531 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2957_v23).length));
          (_2957_v23)[_index531] = new BigNumber(-852);
          let _2973_v32;
          _2973_v32 = _dafny.Map.Empty.slice().updateUnsafe((_2963_v25).f31,(_this).f25);
          let _index532 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_2957_v23).length));
          (_2957_v23)[_index532] = ((_this).f27).multipliedBy(((_2963_v25).f31).plus(new BigNumber((_2973_v32).length)));
        } else {
          let _index533 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
          (_2932_v1)[_index533] = !(!(_this.f28));
          _2935_v3 = new BigNumber((_module.__default.fm2((_this).f27, (_this).f25, globalState)).length);
          _2935_v3 = _2935_v3;
          r0 = !(!(_this.f28));
          let _2974_v33;
          _2974_v33 = _module.D3.create_DC10((_2943_v11)[_module.__default.safeIndex(_2935_v3, new BigNumber((_2943_v11).length))], _2932_v1, _this.f28, (_this).f25, _2935_v3);
          (globalState).f12 = (_2974_v33).dtor_cf15;
        }
        if (true) {
          let _2975_v34;
          let _init99 = ((_2976_v22) => function (_2977_i7) {
            return _module.D15.create_DC39(_2976_v22);
          })(_2956_v22);
          let _nw475 = Array((new BigNumber(24)).toNumber());
          for (let _i0_99 = 0; _i0_99 < new BigNumber(_nw475.length); _i0_99++) {
            _nw475[_i0_99] = _init99(new BigNumber(_i0_99));
          }
          _2975_v34 = _nw475;
          let _2978_v35;
          _2978_v35 = _dafny.Seq.UnicodeFromString("w");
          let _rhs511 = _2975_v34;
          let _rhs512 = _2935_v3;
          let _rhs513 = _2935_v3;
          let _rhs514 = _dafny.Seq.contains(_2978_v35, _2931_v0);
          let _rhs515 = (_this).f25;
          _2975_v34 = _rhs511;
          _2935_v3 = _rhs512;
          _2935_v3 = _rhs513;
          r0 = _rhs514;
          _2935_v3 = _rhs515;
          let _2979_v36;
          _2979_v36 = _module.D14.create_DC38(((_this).f27).plus(new BigNumber(868)));
          let _2980_v37;
          _2980_v37 = _dafny.MultiSet.fromElements(_2942_v10, _dafny.MultiSet.FromArray(_2939_v7), _2942_v10);
          let _2981_v38;
          _2981_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2980_v37);
          let _pat_let_tv76 = _2936_v4;
          let _rhs516 = function (_pat_let72_0) {
            return function (_2984_dt__update__tmp_h1) {
              return function (_pat_let75_0) {
                return function (_2985_dt__update_hcf58_h1) {
                  return _module.D14.create_DC38(_2985_dt__update_hcf58_h1);
                }(_pat_let75_0);
              }((_this).f27);
            }(_pat_let72_0);
          }(function (_pat_let73_0) {
            return function (_2982_dt__update__tmp_h0) {
              return function (_pat_let74_0) {
                return function (_2983_dt__update_hcf58_h0) {
                  return _module.D14.create_DC38(_2983_dt__update_hcf58_h0);
                }(_pat_let74_0);
              }(new BigNumber((_pat_let_tv76).length));
            }(_pat_let73_0);
          }(_2979_v36));
          let _rhs517 = (_2935_v3).plus((_this).f25);
          let _rhs518 = (new BigNumber(((((_2981_v38).contains((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))])) ? ((_2981_v38).get((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))])) : (_dafny.MultiSet.fromElements(_2942_v10, _dafny.MultiSet.FromArray(_2939_v7), _2942_v10)))).cardinality())).isEqualTo((_2935_v3).minus(new BigNumber(810)));
          let _rhs519 = ((_this).f25).minus((_this).f27);
          let _rhs520 = (_2955_v21)[_module.__default.safeIndex(((_2939_v7)[_module.__default.safeIndex(_2935_v3, new BigNumber((_2939_v7).length))]).plus(_2935_v3), new BigNumber((_2955_v21).length))];
          _2979_v36 = _rhs516;
          _2935_v3 = _rhs517;
          r0 = _rhs518;
          _2935_v3 = _rhs519;
          _2957_v23 = _rhs520;
          _2935_v3 = new BigNumber((_dafny.Seq.Concat(_2943_v11, _2943_v11)).length);
          let _2986_v39;
          _2986_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(601),(_this).f27);
          let _2987_v40;
          _2987_v40 = _module.D14.create_DC37(_2986_v39);
          let _2988_v41;
          _2988_v41 = _dafny.Set.fromElements(p1);
          let _pat_let_tv77 = _2988_v41;
          _2987_v40 = function (_pat_let76_0) {
            return function (_2989_dt__update__tmp_h2) {
              return function (_pat_let77_0) {
                return function (_2990_dt__update_hcf57_h0) {
                  return _module.D14.create_DC37(_2990_dt__update_hcf57_h0);
                }(_pat_let77_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_pat_let_tv77).length),(_this).f25));
            }(_pat_let76_0);
          }(_module.D14.create_DC37(_2986_v39));
          (globalState).f12 = (_module.__default.fm73(p1, globalState)).dtor_cf47;
        } else {
          let _2991_v42;
          _2991_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2931_v0);
          let _2992_v43;
          _2992_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2991_v42).length),(_this).f25);
          let _2993_v44;
          _2993_v44 = _dafny.Set.fromElements(_2992_v43);
          let _2994_v45;
          _2994_v45 = _module.D21.create_DC53(_2993_v44, _module.__default.fm0((_this).f25, globalState), p0);
          let _2995_v46;
          _2995_v46 = _module.D7.create_DC21(_2957_v23);
          let _2996_v47;
          let _nw476 = Array((new BigNumber(25)).toNumber());
          _nw476[(_dafny.ZERO).toNumber()] = _2957_v23;
          _nw476[(_dafny.ONE).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(2)).toNumber()] = (_2994_v45).dtor_cf84;
          _nw476[(new BigNumber(3)).toNumber()] = p0;
          _nw476[(new BigNumber(4)).toNumber()] = p0;
          _nw476[(new BigNumber(5)).toNumber()] = p0;
          _nw476[(new BigNumber(6)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(7)).toNumber()] = (_2995_v46).dtor_cf32;
          _nw476[(new BigNumber(8)).toNumber()] = p0;
          _nw476[(new BigNumber(9)).toNumber()] = p0;
          _nw476[(new BigNumber(10)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(11)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(12)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(13)).toNumber()] = p0;
          _nw476[(new BigNumber(14)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(15)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(16)).toNumber()] = p0;
          _nw476[(new BigNumber(17)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(18)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(19)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(20)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(21)).toNumber()] = p0;
          _nw476[(new BigNumber(22)).toNumber()] = _2957_v23;
          _nw476[(new BigNumber(23)).toNumber()] = p0;
          _nw476[(new BigNumber(24)).toNumber()] = _2957_v23;
          _2996_v47 = _nw476;
          let _2997_v48;
          _2997_v48 = _module.D15.create_DC39(_dafny.MultiSet.fromElements((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))], (_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))]));
          let _2998_v49;
          _2998_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2997_v48,_2996_v47);
          _2996_v47 = (((_2998_v49).contains(_module.D15.create_DC39(_2956_v22))) ? ((_2998_v49).get(_module.D15.create_DC39(_2956_v22))) : (_2996_v47));
          let _2999_v50;
          let _3000_v51;
          let _3001_v52;
          let _out46;
          let _out47;
          let _out48;
          let _outcollector16 = _module.__default.m0(globalState);
          _out46 = _outcollector16[0];
          _out47 = _outcollector16[1];
          _out48 = _outcollector16[2];
          _2999_v50 = _out46;
          _3000_v51 = _out47;
          _3001_v52 = _out48;
          let _index534 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p0).length));
          (p0)[_index534] = new BigNumber((_2943_v11).length);
          let _index535 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p0).length));
          (p0)[_index535] = (_this).f25;
          let _index536 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p0).length));
          let _rhs521 = _module.__default.fm44(false, !(_this.f28), false, _module.__default.fm18(globalState), globalState);
          let _rhs522 = (_this).f25;
          let _rhs523 = _2943_v11;
          let _lhs391 = p0;
          let _lhs392 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((p0).length));
          _lhs391[_lhs392] = _rhs521;
          _2935_v3 = _rhs522;
          _2943_v11 = _rhs523;
          (_this).f28 = _module.__default.fm0(((p0)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((p0).length))]).plus((_this).f27), globalState);
        }
        if (p1) {
          let _3002_v53;
          let _nw477 = new _module.C7();
          _nw477.__ctor(_module.__default.safeModuloInt(_2935_v3, new BigNumber(576)));
          _3002_v53 = _nw477;
          let _3003_v54;
          _3003_v54 = _dafny.Seq.UnicodeFromString("gmclbyo");
          let _index537 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
          let _rhs524 = ((((_2956_v22).contains(!(p1))) ? ((_2956_v22).get(!(p1))) : ((_this).f27))).isLessThanOrEqualTo((_3002_v53).f34);
          let _rhs525 = p3;
          let _rhs526 = p0;
          let _rhs527 = _3003_v54;
          let _rhs528 = (_dafny.ZERO).minus((_this).f27);
          let _lhs393 = globalState;
          let _lhs394 = _2932_v1;
          let _lhs395 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
          _lhs393.f12 = _rhs524;
          _lhs394[_lhs395] = _rhs525;
          _2957_v23 = _rhs526;
          r2 = _rhs527;
          _2935_v3 = _rhs528;
          _2935_v3 = (_3002_v53).f34;
          _2935_v3 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_2935_v3)), (_this).f25);
          _2939_v7 = _dafny.Seq.Concat(_dafny.Seq.update(_2939_v7, _module.__default.safeIndex((_this).f27, new BigNumber((_2939_v7).length)), new BigNumber(155)), _2939_v7);
        } else {
          let _3004_v55;
          let _3005_v56;
          let _3006_v57;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector17 = _module.__default.m0(globalState);
          _out49 = _outcollector17[0];
          _out50 = _outcollector17[1];
          _out51 = _outcollector17[2];
          _3004_v55 = _out49;
          _3005_v56 = _out50;
          _3006_v57 = _out51;
          let _3007_v58;
          _3007_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_2956_v22);
          _3007_v58 = _3007_v58;
          let _3008_v59;
          _3008_v59 = _dafny.Set.fromElements(_2957_v23, _3005_v56, _2957_v23, p0, _2957_v23);
          let _3009_v60;
          _3009_v60 = _module.D1.create_DC4(_2943_v11);
          let _3010_v61;
          _3010_v61 = _dafny.Seq.of(_module.D1.create_DC7(_3009_v60));
          let _rhs529 = ((_this).f25).plus(new BigNumber((_3010_v61).length));
          let _rhs530 = _3008_v59;
          _2935_v3 = _rhs529;
          _3008_v59 = _rhs530;
          let _3011_v62;
          _3011_v62 = _dafny.Seq.UnicodeFromString("gjomdhuxy");
          r2 = _dafny.Seq.Concat(_module.__default.fm2((_dafny.ZERO).minus(_2935_v3), (_this).f27, globalState), _dafny.Seq.Concat(_3011_v62, _3011_v62));
          let _3012_v63;
          _3012_v63 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), function (_3013_i8) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }));
          _3012_v63 = (_3012_v63).update(_this.f28, _3011_v62);
        }
      }
      let _ingredients0 = [];
      for (const _guard_loop_14 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2932_v1).length))) {
        let _3014_i9 = _guard_loop_14;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_3014_i9)) && ((_3014_i9).isLessThan(new BigNumber((_2932_v1).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_2932_v1, (_3014_i9).toNumber(), ((!(_this.f28)) ? ((_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))]) : (true))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      r1 = (new BigNumber((_dafny.Seq.of((_2939_v7)[_module.__default.safeIndex((_this).f27, new BigNumber((_2939_v7).length))], (_this).f27)).length)).isLessThanOrEqualTo((_this).f27);
      let _index538 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length));
      (_2932_v1)[_index538] = (_2932_v1)[_module.__default.safeIndex(new BigNumber(187), new BigNumber((_2932_v1).length))];
      r0 = (!(false) || (p1)) === (_this.f28);
      r1 = ((_this).f27).isLessThan((_this).f25);
      let _3015_v64;
      _3015_v64 = _dafny.Seq.UnicodeFromString("bowcxhc");
      r2 = _3015_v64;
      return [r0, r1, r2];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C13 = class C13 {
    constructor () {
      this._tname = "_module.C13";
      this._f24 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f24) {
      let _this = this;
      (_this)._f24 = f24;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f24;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = false;
      let r3 = _dafny.Seq.of();
      let _3016_v0;
      _3016_v0 = _dafny.Seq.of(false);
      let _3017_v1;
      _3017_v1 = _module.D1.create_DC4(_3016_v0);
      let _3018_v2;
      _3018_v2 = _module.D1.create_DC7(_3017_v1);
      let _3019_v3;
      _3019_v3 = _dafny.Set.fromElements(_3018_v2, _3018_v2, _3018_v2, _module.D1.create_DC7(_module.D1.create_DC7(_3017_v1)));
      let _3020_v4;
      _3020_v4 = _dafny.Seq.of(_3019_v3);
      let _3021_v5;
      _3021_v5 = _dafny.Seq.of(_module.D1.create_DC4(_3016_v0), _module.D1.create_DC5());
      let _3022_v6;
      _3022_v6 = _dafny.MultiSet.fromElements(_3018_v2, _3018_v2, _module.D1.create_DC7((_3021_v5)[_module.__default.safeIndex(p0, new BigNumber((_3021_v5).length))]), _3018_v2);
      r3 = _dafny.Seq.of((function () {
        let _coll96 = new _dafny.Set();
        for (const _compr_96 of (_3022_v6).Elements) {
          let _3023_v7 = _compr_96;
          if ((_3022_v6).contains(_3023_v7)) {
            _coll96.add(_3023_v7);
          }
        }
        return _coll96;
      }()).IsSubsetOf((_3020_v4)[_module.__default.safeIndex(p0, new BigNumber((_3020_v4).length))]), false);
      let _3024_v8;
      _3024_v8 = new _dafny.CodePoint('c'.codePointAt(0));
      (globalState).f23 = _3024_v8;
      let _3025_v9;
      _3025_v9 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
      let _hi15 = (_dafny.ZERO).minus(new BigNumber(((_module.__default.fm4((_this).f24, globalState)).Difference(_3025_v9)).cardinality()));
      for (let _3026_i0 = p0; _3026_i0.isLessThan(_hi15); _3026_i0 = _3026_i0.plus(_dafny.ONE)) {
        if (((_this).f24) === (p1)) {
          let _3027_v10;
          _3027_v10 = _module.D0.create_DC2(p1);
          let _3028_v11;
          _3028_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),_3027_v10);
          _3028_v11 = (_3028_v11).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length), _3027_v10);
          let _3029_v12;
          _3029_v12 = new BigNumber(701);
          _3029_v12 = (_3026_i0).minus((_3029_v12).plus(_3026_i0));
          let _3030_v13;
          _3030_v13 = _module.D0.create_DC3(_3029_v12, (_this).f24);
          let _3031_v14;
          _3031_v14 = _dafny.Map.Empty.slice().updateUnsafe((_3030_v13).dtor_cf4,!(_dafny.Set.fromElements((_this).f24, (_this).f24, (_this).f24)).contains((_this).f24));
          (globalState).f12 = (((_3031_v14).contains((_dafny.ZERO).minus(p0))) ? ((_3031_v14).get((_dafny.ZERO).minus(p0))) : (((_this).f24) || ((_this).f24)));
          let _3032_v15;
          let _nw478 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _3032_v15 = _nw478;
          let _index539 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_3032_v15).length));
          (_3032_v15)[_index539] = p0;
          let _3033_v16;
          let _nw479 = Array((new BigNumber(27)).toNumber()).fill(false);
          _3033_v16 = _nw479;
          let _index540 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_3033_v16).length));
          (_3033_v16)[_index540] = p1;
          let _index541 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_3033_v16).length));
          (_3033_v16)[_index541] = p1;
          let _3034_v17;
          _3034_v17 = _dafny.Map.Empty.slice().updateUnsafe(_3032_v15,new BigNumber((_3031_v14).length));
          let _3035_v18;
          _3035_v18 = _dafny.Seq.of(new BigNumber(484));
          let _3036_v19;
          _3036_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_3016_v0)[_module.__default.safeIndex((_3030_v13).dtor_cf4, new BigNumber((_3016_v0).length))]);
          let _index542 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_3032_v15).length));
          let _index543 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_3033_v16).length));
          let _index544 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_3033_v16).length));
          let _rhs531 = (new BigNumber((_3034_v17).length)).multipliedBy(new BigNumber((_3035_v18).length));
          let _rhs532 = _3024_v8;
          let _rhs533 = (_this).f24;
          let _rhs534 = !((((_3036_v19).contains(p1)) ? ((_3036_v19).get(p1)) : (p1))) || (p1);
          let _lhs396 = _3032_v15;
          let _lhs397 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_3032_v15).length));
          let _lhs398 = globalState;
          let _lhs399 = _3033_v16;
          let _lhs400 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_3033_v16).length));
          let _lhs401 = _3033_v16;
          let _lhs402 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_3033_v16).length));
          _lhs396[_lhs397] = _rhs531;
          _lhs398.f23 = _rhs532;
          _lhs399[_lhs400] = _rhs533;
          _lhs401[_lhs402] = _rhs534;
          let _index545 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_3032_v15).length));
          let _rhs535 = _3031_v14;
          let _rhs536 = p0;
          let _rhs537 = false;
          let _rhs538 = _module.__default.fm5((_this).f24, globalState);
          let _lhs403 = _3032_v15;
          let _lhs404 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_3032_v15).length));
          let _lhs405 = globalState;
          _3031_v14 = _rhs535;
          _lhs403[_lhs404] = _rhs536;
          _lhs405.f12 = _rhs537;
          _3029_v12 = _rhs538;
        } else {
          let _3037_v20;
          _3037_v20 = new BigNumber(472);
          let _3038_v21;
          _3038_v21 = _dafny.Seq.of(p0, (_dafny.ZERO).minus(p0));
          let _3039_v22;
          _3039_v22 = _dafny.MultiSet.fromElements(p0, _3026_i0, _module.__default.safeModuloInt(_3037_v20, p0), _module.__default.fm5(!((_this).f24), globalState), _module.__default.safeModuloInt((_3038_v21)[_module.__default.safeIndex(p0, new BigNumber((_3038_v21).length))], p0));
          let _3040_v23;
          _3040_v23 = _module.D0.create_DC2((_this).f24);
          let _3041_v24;
          _3041_v24 = _module.D0.create_DC1(p1, p1);
          let _3042_v25;
          _3042_v25 = _dafny.Seq.UnicodeFromString("gdwocu");
          let _3043_v26;
          _3043_v26 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          let _3044_v27;
          _3044_v27 = _dafny.Set.fromElements(_3026_i0, p0);
          let _3045_v28;
          let _nw480 = Array((new BigNumber(16)).toNumber());
          _nw480[(_dafny.ZERO).toNumber()] = (_3040_v23).dtor_cf3;
          _nw480[(_dafny.ONE).toNumber()] = (_this).f24;
          _nw480[(new BigNumber(2)).toNumber()] = false;
          _nw480[(new BigNumber(3)).toNumber()] = (p1) || ((_this).f24);
          _nw480[(new BigNumber(4)).toNumber()] = !(p1);
          _nw480[(new BigNumber(5)).toNumber()] = (_this).f24;
          _nw480[(new BigNumber(6)).toNumber()] = ((_this).fm3(p0, _3038_v21, new BigNumber((_dafny.Seq.of(_dafny.Seq.of(p1, (_3041_v24).dtor_cf2), _3016_v0)).length), p1, globalState)) && ((_this).f24);
          _nw480[(new BigNumber(7)).toNumber()] = p1;
          _nw480[(new BigNumber(8)).toNumber()] = (_this).f24;
          _nw480[(new BigNumber(9)).toNumber()] = p1;
          _nw480[(new BigNumber(10)).toNumber()] = (p1) && (!(p1));
          _nw480[(new BigNumber(11)).toNumber()] = (_3037_v20).isEqualTo(_3037_v20);
          _nw480[(new BigNumber(12)).toNumber()] = p1;
          _nw480[(new BigNumber(13)).toNumber()] = (_3044_v27).IsProperSubsetOf(_dafny.Set.fromElements((((_3025_v9).contains((_this).f24)) ? ((_3025_v9).get((_this).f24)) : (p0)), new BigNumber((_3042_v25).length), new BigNumber((_3043_v26).length), _3026_i0));
          _nw480[(new BigNumber(14)).toNumber()] = p1;
          _nw480[(new BigNumber(15)).toNumber()] = p1;
          _3045_v28 = _nw480;
          let _index546 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3045_v28).length));
          (_3045_v28)[_index546] = p1;
          let _3046_v29;
          _3046_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,_3045_v28);
          let _3047_v30;
          _3047_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_3046_v29).update(p1, _3045_v28)).length),_3037_v20);
          let _3048_v31;
          _3048_v31 = _dafny.MultiSet.fromElements(_3047_v30);
          let _3049_v32;
          _3049_v32 = _3038_v21;
          let _index547 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3045_v28).length));
          let _rhs539 = (_this).f24;
          let _rhs540 = p0;
          let _rhs541 = ((_3039_v22).Union(_3039_v22)).Intersect(_dafny.MultiSet.fromElements(_3026_i0, _3037_v20, new BigNumber((_3048_v31).cardinality()), _3037_v20, new BigNumber(((_3049_v32)).length)));
          let _rhs542 = !((_3016_v0)[_module.__default.safeIndex(_3026_i0, new BigNumber((_3016_v0).length))]);
          let _rhs543 = _module.__default.fm5(p1, globalState);
          let _lhs406 = globalState;
          let _lhs407 = _3045_v28;
          let _lhs408 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_3045_v28).length));
          _lhs406.f12 = _rhs539;
          _3037_v20 = _rhs540;
          _3039_v22 = _rhs541;
          _lhs407[_lhs408] = _rhs542;
          _3037_v20 = _rhs543;
          let _3050_v33;
          let _nw481 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _3050_v33 = _nw481;
          let _3051_v34;
          _3051_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(920),(_this).f24);
          let _3052_v35;
          _3052_v35 = _dafny.Map.Empty.slice().updateUnsafe(_3050_v33,_3051_v34);
          _3052_v35 = (_3052_v35).update(_3050_v33, _3051_v34);
          let _3053_v36;
          _3053_v36 = _dafny.MultiSet.fromElements(_3040_v23);
          _3053_v36 = (_3053_v36).Union(_3053_v36);
          let _3054_v37;
          _3054_v37 = _module.D0.create_DC3(p0, !((_this).f24));
          (globalState).f12 = ((_3026_i0).multipliedBy((_3054_v37).dtor_cf4)).isLessThan(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_3026_i0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(_3026_i0)).length)), _3026_i0), _3038_v21)).length));
          r2 = (_3045_v28)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_3045_v28).length))];
        }
        let _3055_v38;
        let _init100 = ((_3056_i0) => function (_3057_i1) {
          return _module.__default.safeModuloInt(_3057_i1, _3056_i0);
        })(_3026_i0);
        let _nw482 = Array((new BigNumber(4)).toNumber());
        for (let _i0_100 = 0; _i0_100 < new BigNumber(_nw482.length); _i0_100++) {
          _nw482[_i0_100] = _init100(new BigNumber(_i0_100));
        }
        _3055_v38 = _nw482;
        let _index548 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length));
        (_3055_v38)[_index548] = p0;
        let _index549 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length));
        (_3055_v38)[_index549] = (_module.__default.safeModuloInt(p0, _3026_i0)).minus((new BigNumber((_3016_v0).length)).multipliedBy(_3026_i0));
        let _3058_v39;
        _3058_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3016_v0).length),(_this).f24);
        let _3059_v40;
        _3059_v40 = _dafny.Seq.of(_3058_v39, _3058_v39);
        let _3060_v41;
        _3060_v41 = _dafny.Map.Empty.slice().updateUnsafe(((_3055_v38)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length))]).isLessThan(_3026_i0),_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_3016_v0, _module.__default.safeIndex((_3055_v38)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length))], new BigNumber((_3016_v0).length)), (_this).f24), _module.__default.safeIndex(new BigNumber(((_3059_v40)[_module.__default.safeIndex(_3026_i0, new BigNumber((_3059_v40).length))]).length), new BigNumber((_dafny.Seq.update(_3016_v0, _module.__default.safeIndex((_3055_v38)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length))], new BigNumber((_3016_v0).length)), (_this).f24)).length)), (_this).f24)));
        _3060_v41 = (_3060_v41).update((_this).f24, _3025_v9);
        let _3061_v42;
        _3061_v42 = _dafny.Seq.of(_3026_i0, new BigNumber((_3025_v9).cardinality()));
        let _source40 = _3061_v42;
        let _3062___mcc_h0 = _source40;
        let _3063_cf11 = _3062___mcc_h0;
        let _3064_v43;
        let _nw483 = Array((new BigNumber(27)).toNumber()).fill(false);
        _3064_v43 = _nw483;
        let _3065_v44;
        _3065_v44 = _dafny.MultiSet.fromElements(p0);
        let _index550 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_3064_v43).length));
        (_3064_v43)[_index550] = !(!((((_3058_v39).contains((((_3065_v44).contains(p0)) ? ((_3065_v44).get(p0)) : ((_3055_v38)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length))])))) ? ((_3058_v39).get((((_3065_v44).contains(p0)) ? ((_3065_v44).get(p0)) : ((_3055_v38)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_3055_v38).length))])))) : (false))));
        let _index551 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_3064_v43).length));
        (_3064_v43)[_index551] = p1;
        let _index552 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_3064_v43).length));
        (_3064_v43)[_index552] = ((_this).f24) && ((_this).f24);
        let _3066_v45;
        _3066_v45 = _dafny.Map.Empty.slice().updateUnsafe((_3064_v43)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_3064_v43).length))],p1);
        let _3067_v46;
        _3067_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), ((_3068_i0) => function (_3069_i2) {
          return _3068_i0;
        })(_3026_i0))).length),new BigNumber((_3066_v45).length));
        let _rhs544 = _3067_v46;
        let _rhs545 = (((_3066_v45).update((_3064_v43)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_3064_v43).length))], p1)).Merge(_3066_v45)).Merge(_3066_v45);
        let _rhs546 = _3064_v43;
        _3067_v46 = _rhs544;
        _3066_v45 = _rhs545;
        _3064_v43 = _rhs546;
        let _3070_v47;
        _3070_v47 = _module.D1.create_DC4(_3016_v0);
        _3070_v47 = _3070_v47;
      }
      r0 = !(p0).isEqualTo((_dafny.ZERO).minus(p0));
      let _3071_v48;
      _3071_v48 = _module.D0.create_DC0(p1);
      let _pat_let_tv78 = _3024_v8;
      let _pat_let_tv79 = _3024_v8;
      let _pat_let_tv80 = _3024_v8;
      let _pat_let_tv81 = _3024_v8;
      (globalState).f23 = function (_source41) {
        if (_source41.is_DC1) {
          let _3072___mcc_h1 = (_source41).cf1;
          let _3073___mcc_h2 = (_source41).cf2;
          let _3074_cf2 = _3073___mcc_h2;
          let _3075_cf1 = _3072___mcc_h1;
          return _pat_let_tv78;
        } else if (_source41.is_DC2) {
          let _3076___mcc_h3 = (_source41).cf3;
          let _3077_cf3 = _3076___mcc_h3;
          return _pat_let_tv79;
        } else if (_source41.is_DC3) {
          let _3078___mcc_h4 = (_source41).cf4;
          let _3079___mcc_h5 = (_source41).cf5;
          let _3080_cf5 = _3079___mcc_h5;
          let _3081_cf4 = _3078___mcc_h4;
          return _pat_let_tv80;
        } else {
          let _3082___mcc_h6 = (_source41).cf0;
          let _3083_cf0 = _3082___mcc_h6;
          return _pat_let_tv81;
        }
      }(_3071_v48);
      let _3084_v49;
      _3084_v49 = _module.D0.create_DC3((p0).multipliedBy(p0), false);
      let _source42 = _3084_v49;
      if (_source42.is_DC1) {
        let _3085___mcc_h7 = (_source42).cf1;
        let _3086___mcc_h8 = (_source42).cf2;
        let _3087_cf2 = _3086___mcc_h8;
        let _3088_cf1 = _3085___mcc_h7;
        let _3089_v50;
        let _init101 = function (_3090_i3) {
          return (_this).f24;
        };
        let _nw484 = Array((new BigNumber(5)).toNumber());
        for (let _i0_101 = 0; _i0_101 < new BigNumber(_nw484.length); _i0_101++) {
          _nw484[_i0_101] = _init101(new BigNumber(_i0_101));
        }
        _3089_v50 = _nw484;
        let _3091_v51;
        let _init102 = ((_3092_p0) => function (_3093_i4) {
          return _module.__default.safeModuloInt(_3093_i4, _3092_p0);
        })(p0);
        let _nw485 = Array((new BigNumber(8)).toNumber());
        for (let _i0_102 = 0; _i0_102 < new BigNumber(_nw485.length); _i0_102++) {
          _nw485[_i0_102] = _init102(new BigNumber(_i0_102));
        }
        _3091_v51 = _nw485;
        let _index553 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_3091_v51).length));
        (_3091_v51)[_index553] = p0;
        let _index554 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_3091_v51).length));
        let _rhs547 = _3088_cf1;
        let _rhs548 = _3089_v50;
        let _rhs549 = new BigNumber(285);
        let _lhs409 = _3091_v51;
        let _lhs410 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_3091_v51).length));
        _3087_cf2 = _rhs547;
        _3089_v50 = _rhs548;
        _lhs409[_lhs410] = _rhs549;
        let _3094_v52;
        _3094_v52 = _dafny.Seq.UnicodeFromString("ugn");
        let _3095_v53;
        _3095_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_3087_cf2) ? (_dafny.Seq.UnicodeFromString("qfodmffx")) : (_3094_v52)));
        _3095_v53 = (_3095_v53).update(new BigNumber(463), _dafny.Seq.UnicodeFromString("aqq"));
        let _3096_v54;
        _3096_v54 = _dafny.Set.fromElements((p0).minus(p0), (_3091_v51)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_3091_v51).length))]);
        _3096_v54 = _3096_v54;
        let _3097_v55;
        _3097_v55 = _dafny.Map.Empty.slice().updateUnsafe((_3091_v51)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_3091_v51).length))],(_dafny.ZERO).minus(((_3087_cf2) ? (p0) : (p0))));
        _3097_v55 = _3097_v55;
      } else if (_source42.is_DC2) {
        let _3098___mcc_h9 = (_source42).cf3;
        let _3099_cf3 = _3098___mcc_h9;
        if ((_this).f24) {
          let _3100_v56;
          let _3101_v57;
          let _3102_v58;
          let _out52;
          let _out53;
          let _out54;
          let _outcollector18 = _module.__default.m0(globalState);
          _out52 = _outcollector18[0];
          _out53 = _outcollector18[1];
          _out54 = _outcollector18[2];
          _3100_v56 = _out52;
          _3101_v57 = _out53;
          _3102_v58 = _out54;
          let _3103_v59;
          _3103_v59 = new BigNumber(337);
          _3103_v59 = _3103_v59;
          let _3104_v60;
          _3104_v60 = _dafny.Set.fromElements(_3100_v56, _3100_v56, _3100_v56, _3100_v56);
          let _3105_v61;
          _3105_v61 = _module.D3.create_DC9(_dafny.Set.fromElements(_3100_v56, _3100_v56));
          let _3106_v62;
          let _nw486 = Array((new BigNumber(12)).toNumber());
          _nw486[(_dafny.ZERO).toNumber()] = _3104_v60;
          _nw486[(_dafny.ONE).toNumber()] = _3104_v60;
          _nw486[(new BigNumber(2)).toNumber()] = (_3104_v60).Difference(_dafny.Set.fromElements(_3100_v56, _3100_v56));
          _nw486[(new BigNumber(3)).toNumber()] = (_3104_v60).Difference(_3104_v60);
          _nw486[(new BigNumber(4)).toNumber()] = _3104_v60;
          _nw486[(new BigNumber(5)).toNumber()] = (_3104_v60).Union(_3104_v60);
          _nw486[(new BigNumber(6)).toNumber()] = _3104_v60;
          _nw486[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_3100_v56, _3100_v56, _3100_v56, _3100_v56);
          _nw486[(new BigNumber(8)).toNumber()] = (_3105_v61).dtor_cf12;
          _nw486[(new BigNumber(9)).toNumber()] = (_3104_v60).Union(_3104_v60);
          _nw486[(new BigNumber(10)).toNumber()] = (_3104_v60).Union(_3104_v60);
          _nw486[(new BigNumber(11)).toNumber()] = _3104_v60;
          _3106_v62 = _nw486;
          let _index555 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_3106_v62).length));
          (_3106_v62)[_index555] = _3104_v60;
          let _index556 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_3106_v62).length));
          (_3106_v62)[_index556] = _3104_v60;
          let _3107_v63;
          _3107_v63 = _dafny.Set.fromElements(p0);
          let _index557 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_3100_v56).length));
          (_3100_v56)[_index557] = (_3107_v63).IsDisjointFrom(_dafny.Set.fromElements(p0, p0));
          let _3108_v64;
          _3108_v64 = _dafny.Set.fromElements(false, _3099_cf3, p1, !(_3102_v58), (_this).f24);
          let _index558 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_3100_v56).length));
          (_3100_v56)[_index558] = ((_3108_v64).Intersect(_dafny.Set.fromElements(_3099_cf3))).IsSubsetOf(_3108_v64);
          let _3109_v65;
          _3109_v65 = _dafny.Seq.UnicodeFromString("teauuxeu");
          _3103_v59 = new BigNumber((_3109_v65).length);
        } else {
          let _3110_v66;
          _3110_v66 = new BigNumber(135);
          _3110_v66 = p0;
          let _3111_v67;
          _3111_v67 = _dafny.Seq.UnicodeFromString("fo");
          let _3112_v68;
          _3112_v68 = _dafny.Seq.of(_3111_v67);
          let _3113_v69;
          _3113_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(_3111_v67));
          let _3114_v70;
          let _nw487 = Array((new BigNumber(11)).toNumber());
          _nw487[(_dafny.ZERO).toNumber()] = _3112_v68;
          _nw487[(_dafny.ONE).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(2)).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(3)).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_3111_v67), (((_3113_v69).contains(new BigNumber(690))) ? ((_3113_v69).get(new BigNumber(690))) : (_3112_v68)));
          _nw487[(new BigNumber(5)).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(6)).toNumber()] = _module.__default.fm6(globalState);
          _nw487[(new BigNumber(7)).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(8)).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(9)).toNumber()] = _3112_v68;
          _nw487[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_3112_v68, _dafny.Seq.of(_3111_v67));
          _3114_v70 = _nw487;
          let _index559 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_3114_v70).length));
          (_3114_v70)[_index559] = _3112_v68;
          let _3115_v71;
          let _nw488 = Array((new BigNumber(27)).toNumber()).fill(false);
          _3115_v71 = _nw488;
          let _3116_v72;
          _3116_v72 = _dafny.Seq.of(new BigNumber(-52));
          let _index560 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_3115_v71).length));
          (_3115_v71)[_index560] = _dafny.Seq.IsPrefixOf(_3116_v72, _dafny.Seq.update(_3116_v72, _module.__default.safeIndex(new BigNumber(((_3112_v68)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_3016_v0)).cardinality()), new BigNumber((_3112_v68).length))]).length), new BigNumber((_3116_v72).length)), p0));
          let _index561 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_3114_v70).length));
          let _index562 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_3115_v71).length));
          let _rhs550 = _3112_v68;
          let _rhs551 = p1;
          let _lhs411 = _3114_v70;
          let _lhs412 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_3114_v70).length));
          let _lhs413 = _3115_v71;
          let _lhs414 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_3115_v71).length));
          _lhs411[_lhs412] = _rhs550;
          _lhs413[_lhs414] = _rhs551;
          let _3117_v73;
          let _3118_v74;
          let _3119_v75;
          let _out55;
          let _out56;
          let _out57;
          let _outcollector19 = _module.__default.m0(globalState);
          _out55 = _outcollector19[0];
          _out56 = _outcollector19[1];
          _out57 = _outcollector19[2];
          _3117_v73 = _out55;
          _3118_v74 = _out56;
          _3119_v75 = _out57;
          let _3120_v76;
          _3120_v76 = _dafny.MultiSet.fromElements(_3110_v66);
          (globalState).f12 = (_3016_v0)[_module.__default.safeIndex(new BigNumber(((_3120_v76).Intersect(_3120_v76)).cardinality()), new BigNumber((_3016_v0).length))];
          _3120_v76 = (_3120_v76).Difference(_3120_v76);
        }
        let _3121_v77;
        let _nw489 = new _module.C0();
        _nw489.__ctor();
        _3121_v77 = _nw489;
        let _3122_v78;
        _3122_v78 = new BigNumber(513);
        let _rhs552 = p0;
        let _rhs553 = (((p1) ? (_3122_v78) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-272)), ((_3123_v78) => function (_3124_i5) {
          return _3123_v78;
        })(_3122_v78))).length)))).minus(_3122_v78);
        _3122_v78 = _rhs552;
        _3122_v78 = _rhs553;
        let _3125_v79;
        let _nw490 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _3125_v79 = _nw490;
        let _index563 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_3125_v79).length));
        (_3125_v79)[_index563] = new BigNumber((_3016_v0).length);
        let _index564 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_3125_v79).length));
        (_3125_v79)[_index564] = _3122_v78;
      } else if (_source42.is_DC3) {
        let _3126___mcc_h10 = (_source42).cf4;
        let _3127___mcc_h11 = (_source42).cf5;
        let _3128_cf5 = _3127___mcc_h11;
        let _3129_cf4 = _3126___mcc_h10;
        let _rhs554 = (_this).f24;
        let _rhs555 = (_this).f24;
        r0 = _rhs554;
        r2 = _rhs555;
        _3025_v9 = _3025_v9;
        let _3130_v80;
        _3130_v80 = _dafny.MultiSet.fromElements(p0);
        let _3131_v81;
        let _nw491 = Array((new BigNumber(3)).toNumber());
        _nw491[(_dafny.ZERO).toNumber()] = _3024_v8;
        _nw491[(_dafny.ONE).toNumber()] = _3024_v8;
        _nw491[(new BigNumber(2)).toNumber()] = _module.__default.fm33(new BigNumber(13), _3128_cf5, new BigNumber((_3130_v80).cardinality()), globalState);
        _3131_v81 = _nw491;
        let _index565 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_3131_v81).length));
        (_3131_v81)[_index565] = _3024_v8;
        let _index566 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_3131_v81).length));
        (_3131_v81)[_index566] = _module.__default.fm39((_this).f24, globalState);
        let _3132_v82;
        _3132_v82 = _module.D1.create_DC5();
        let _3133_v83;
        _3133_v83 = _module.D8.create_DC24(_3132_v82);
        let _source43 = _3133_v83;
        if (_source43.is_DC24) {
          let _3134___mcc_h13 = (_source43).cf36;
          let _3135_cf36 = _3134___mcc_h13;
          let _3136_v84;
          let _nw492 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _3136_v84 = _nw492;
          let _index567 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_3136_v84).length));
          (_3136_v84)[_index567] = _module.__default.safeModuloInt(_3129_cf4, _3129_cf4);
          let _index568 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_3136_v84).length));
          (_3136_v84)[_index568] = (p0).multipliedBy((_dafny.ZERO).minus(p0));
          let _3137_v85;
          _3137_v85 = _dafny.Set.fromElements((_this).f24);
          let _3138_v86;
          _3138_v86 = _dafny.Seq.of(_3137_v85);
          let _3139_v87;
          _3139_v87 = _dafny.Seq.of((_3136_v84)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_3136_v84).length))]);
          let _3140_v88;
          _3140_v88 = _module.D9.create_DC27(new BigNumber(((_3138_v86)[_module.__default.safeIndex(_3129_cf4, new BigNumber((_3138_v86).length))]).length), ((_3139_v87)[_module.__default.safeIndex(new BigNumber(-873), new BigNumber((_3139_v87).length))]).minus(_3129_cf4));
          _3140_v88 = _module.D9.create_DC27(p0, p0);
          let _3141_v89;
          let _nw493 = new _module.C5();
          _nw493.__ctor((p0).isLessThan(p0));
          _3141_v89 = _nw493;
          _3129_cf4 = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(_3129_cf4));
        } else if (_source43.is_DC23) {
          let _3142___mcc_h14 = (_source43).cf35;
          let _3143_cf35 = _3142___mcc_h14;
          _3025_v9 = (((_this).f24) ? (_dafny.MultiSet.fromElements(true, p1)) : (_3025_v9));
          let _3144_v90;
          let _3145_v91;
          let _3146_v92;
          let _out58;
          let _out59;
          let _out60;
          let _outcollector20 = _module.__default.m0(globalState);
          _out58 = _outcollector20[0];
          _out59 = _outcollector20[1];
          _out60 = _outcollector20[2];
          _3144_v90 = _out58;
          _3145_v91 = _out59;
          _3146_v92 = _out60;
          let _3147_v93;
          _3147_v93 = _module.D7.create_DC21(_3145_v91);
          let _3148_v94;
          let _nw494 = new _module.C9();
          _nw494.__ctor((_3147_v93).dtor_cf32, (new BigNumber(-268)).minus(new BigNumber(401)));
          _3148_v94 = _nw494;
          let _3149_v95;
          _3149_v95 = _dafny.Seq.UnicodeFromString("grvwvats");
          let _3150_v96;
          _3150_v96 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_3149_v95).length));
          let _3151_v97;
          _3151_v97 = _dafny.Set.fromElements(new BigNumber((_3150_v96).length), p0);
          let _3152_v98;
          _3152_v98 = _dafny.Map.Empty.slice().updateUnsafe((_3151_v97).IsProperSubsetOf(_3151_v97),(_3016_v0)[_module.__default.safeIndex(new BigNumber((_module.__default.fm2((_dafny.ZERO).minus(p0), p0, globalState)).length), new BigNumber((_3016_v0).length))]);
          let _3153_v99;
          _3153_v99 = _dafny.Map.Empty.slice().updateUnsafe((_3131_v81)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_3131_v81).length))],(_3148_v94).f31);
          _3152_v98 = (_3152_v98).update(!((((_3153_v99).contains((_3131_v81)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_3131_v81).length))])) ? ((_3153_v99).get((_3131_v81)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_3131_v81).length))])) : (new BigNumber((_3151_v97).length)))).isEqualTo(_module.__default.fm31(_3149_v95, globalState)), true);
        } else {
          let _3154___mcc_h15 = (_source43).cf37;
          let _3155_cf37 = _3154___mcc_h15;
          let _3156_v100;
          _3156_v100 = _dafny.Map.Empty.slice().updateUnsafe(_3129_cf4,_3129_cf4);
          let _3157_v101;
          let _nw495 = Array((new BigNumber(7)).toNumber()).fill(false);
          _3157_v101 = _nw495;
          let _3158_v102;
          _3158_v102 = _dafny.Map.Empty.slice().updateUnsafe(_3157_v101,!(p1));
          let _3159_v103;
          _3159_v103 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_3128_cf5),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), ((_3160_v8) => function (_3161_i6) {
            return _3160_v8;
          })(_3024_v8))).length))).length), p0);
          _3156_v100 = (_3156_v100).update(_3129_cf4, (new BigNumber((_3158_v102).length)).multipliedBy(new BigNumber((_3159_v103).length)));
          let _3162_v104;
          _3162_v104 = _dafny.Set.fromElements(_3071_v48, _3071_v48);
          let _rhs556 = _3162_v104;
          let _rhs557 = _3129_cf4;
          _3162_v104 = _rhs556;
          _3129_cf4 = _rhs557;
          let _3163_v105;
          _3163_v105 = _dafny.Map.Empty.slice().updateUnsafe(true,_3129_cf4);
          _3129_cf4 = (p0).minus((_dafny.ZERO).minus(new BigNumber((_3163_v105).length)));
          _3163_v105 = (_3163_v105).update(!(!(p0).isEqualTo(_3129_cf4)), new BigNumber((_3016_v0).length));
        }
      } else {
        let _3164___mcc_h12 = (_source42).cf0;
        let _3165_cf0 = _3164___mcc_h12;
        _3165_cf0 = _3165_cf0;
        _3016_v0 = _dafny.Seq.Concat(_3016_v0, _3016_v0);
        let _3166_v106;
        _3166_v106 = new BigNumber(461);
        _3166_v106 = p0;
        let _3167_v107;
        _3167_v107 = _dafny.Set.fromElements(_3165_cf0);
        let _3168_v108;
        _3168_v108 = _dafny.Map.Empty.slice().updateUnsafe(true,_3165_cf0);
        let _3169_v109;
        _3169_v109 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_3168_v108);
        let _rhs558 = ((_3167_v107).Union(_3167_v107)).Union(_3167_v107);
        let _rhs559 = p1;
        let _rhs560 = _3169_v109;
        let _rhs561 = p0;
        _3167_v107 = _rhs558;
        r2 = _rhs559;
        _3169_v109 = _rhs560;
        _3166_v106 = _rhs561;
      }
      let _3170_v110;
      let _nw496 = new _module.C11();
      _nw496.__ctor();
      _3170_v110 = _nw496;
      let _3171_v111;
      _3171_v111 = _dafny.Map.Empty.slice().updateUnsafe(_3024_v8,_3170_v110);
      let _3172_v112;
      _3172_v112 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      r0 = (_module.__default.safeModuloInt(p0, p0)).isLessThan(_module.__default.safeDivisionInt(new BigNumber((_3171_v111).length), new BigNumber((_3172_v112).length)));
      let _3173_v113;
      _3173_v113 = _dafny.Seq.UnicodeFromString("jtjgpuljk");
      let _3174_v114;
      _3174_v114 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, p0),_3173_v113);
      r1 = _3174_v114;
      r2 = p1;
      let _3175_v115;
      _3175_v115 = _module.D9.create_DC27(p0, p0);
      let _pat_let_tv82 = _3016_v0;
      let _pat_let_tv83 = _3016_v0;
      r3 = function (_source44) {
        if (_source44.is_DC27) {
          let _3176___mcc_h16 = (_source44).cf39;
          let _3177___mcc_h17 = (_source44).cf40;
          let _3178_cf40 = _3177___mcc_h17;
          let _3179_cf39 = _3176___mcc_h16;
          return _pat_let_tv82;
        } else {
          let _3180___mcc_h18 = (_source44).cf38;
          let _3181_cf38 = _3180___mcc_h18;
          return _pat_let_tv83;
        }
      }(_3175_v115);
      return [r0, r1, r2, r3];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
