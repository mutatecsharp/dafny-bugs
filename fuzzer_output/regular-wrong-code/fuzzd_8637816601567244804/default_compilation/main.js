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
    static fm0(p0, p1, p2, p3, globalState) {
      return true;
    };
    static fm2(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), function (_0_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      });
    };
    static fm3(p0, p1, p2, globalState) {
      return (new BigNumber(658)).minus(_module.__default.safeModuloInt(new BigNumber(788), new BigNumber((_dafny.Seq.of(new BigNumber(161))).length)));
    };
    static fm4(p0, globalState) {
      return _dafny.Seq.of(!(!(!(false))), (_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(971))).contains(false));
    };
    static fm5(p0, globalState) {
      let _source0 = _module.D0.create_DC1(true, _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(435)), false, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), function (_1_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length)), new BigNumber(848));
      if (_source0.is_DC1) {
        let _2___mcc_h0 = (_source0).cf1;
        let _3___mcc_h1 = (_source0).cf2;
        let _4___mcc_h2 = (_source0).cf3;
        let _5___mcc_h3 = (_source0).cf4;
        let _6___mcc_h4 = (_source0).cf5;
        let _7_cf5 = _6___mcc_h4;
        let _8_cf4 = _5___mcc_h3;
        let _9_cf3 = _4___mcc_h2;
        let _10_cf2 = _3___mcc_h1;
        let _11_cf1 = _2___mcc_h0;
        return new _dafny.CodePoint('p'.codePointAt(0));
      } else if (_source0.is_DC2) {
        let _12___mcc_h5 = (_source0).cf6;
        let _13___mcc_h6 = (_source0).cf7;
        let _14___mcc_h7 = (_source0).cf8;
        let _15___mcc_h8 = (_source0).cf9;
        let _16_cf9 = _15___mcc_h8;
        let _17_cf8 = _14___mcc_h7;
        let _18_cf7 = _13___mcc_h6;
        let _19_cf6 = _12___mcc_h5;
        return new _dafny.CodePoint('e'.codePointAt(0));
      } else if (_source0.is_DC0) {
        let _20___mcc_h9 = (_source0).cf0;
        let _21_cf0 = _20___mcc_h9;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else {
        let _22___mcc_h10 = (_source0).cf10;
        let _23_cf10 = _22___mcc_h10;
        return new _dafny.CodePoint('i'.codePointAt(0));
      }
    };
    static m0(p0, globalState) {
      let _24_v0;
      _24_v0 = _dafny.Seq.UnicodeFromString("i");
      _24_v0 = _24_v0;
      let _25_v2;
      _25_v2 = new BigNumber(795);
      let _26_v3;
      _26_v3 = _dafny.Map.Empty.slice().updateUnsafe(_25_v2,_25_v2);
      let _27_v4;
      _27_v4 = _dafny.Map.Empty.slice().updateUnsafe(_26_v3,true);
      let _28_v5;
      _28_v5 = _dafny.Set.fromElements(_25_v2, new BigNumber((_27_v4).length));
      let _29_v6;
      _29_v6 = _dafny.MultiSet.fromElements(_25_v2, _25_v2, _module.__default.fm3(_25_v2, p0, p0, globalState), new BigNumber((_24_v0).length));
      _24_v0 = _dafny.Seq.Concat(_module.__default.fm2(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_28_v5).Elements) {
          let _30_v1 = _compr_0;
          if ((_28_v5).contains(_30_v1)) {
            _coll0.push([(_30_v1).plus(_25_v2),p0]);
          }
        }
        return _coll0;
      }()).length), _29_v6, globalState), _24_v0);
      let _31_v7;
      let _nw0 = Array((new BigNumber(11)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = p0;
      _nw0[(_dafny.ONE).toNumber()] = p0;
      _nw0[(new BigNumber(2)).toNumber()] = !(p0);
      _nw0[(new BigNumber(3)).toNumber()] = p0;
      _nw0[(new BigNumber(4)).toNumber()] = p0;
      _nw0[(new BigNumber(5)).toNumber()] = p0;
      _nw0[(new BigNumber(6)).toNumber()] = p0;
      _nw0[(new BigNumber(7)).toNumber()] = (p0) === (p0);
      _nw0[(new BigNumber(8)).toNumber()] = p0;
      _nw0[(new BigNumber(9)).toNumber()] = p0;
      _nw0[(new BigNumber(10)).toNumber()] = p0;
      _31_v7 = _nw0;
      let _index0 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length));
      (_31_v7)[_index0] = p0;
      let _index1 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length));
      (_31_v7)[_index1] = !(((_25_v2).isLessThan(_25_v2)) || (_dafny.areEqual(_dafny.Seq.of(p0), _module.__default.fm4(_25_v2, globalState))));
      let _32_v8;
      let _nw1 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _32_v8 = _nw1;
      let _33_v9;
      _33_v9 = _dafny.Map.Empty.slice().updateUnsafe(_25_v2,(_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))]);
      let _34_v10;
      _34_v10 = _dafny.Seq.of(new BigNumber((_33_v9).length));
      let _35_v11;
      _35_v11 = _dafny.Seq.of((_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))]);
      let _index2 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length));
      (_32_v8)[_index2] = (_34_v10)[_module.__default.safeIndex(new BigNumber((_35_v11).length), new BigNumber((_34_v10).length))];
      let _36_v12;
      _36_v12 = _module.D0.create_DC2(new BigNumber((_24_v0).length), (_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))], p0, new BigNumber((_33_v9).length));
      let _index3 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length));
      (_32_v8)[_index3] = (_36_v12).dtor_cf6;
      let _37_i0;
      _37_i0 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_37_i0)) {
              break L0;
            }
            _37_i0 = (_37_i0).plus(_dafny.ONE);
            let _index4 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length));
            (_32_v8)[_index4] = (_dafny.ZERO).minus((((_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))]).plus(new BigNumber((_24_v0).length))).plus(_25_v2));
            let _38_v13;
            _38_v13 = new _dafny.CodePoint('b'.codePointAt(0));
            let _index5 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length));
            (_31_v7)[_index5] = !(_module.__default.fm0(p0, _module.__default.fm5((_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))], globalState), _38_v13, _25_v2, globalState));
            (globalState).f5 = _25_v2;
            let _39_v14;
            let _nw2 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
            _39_v14 = _nw2;
            let _40_v15;
            _40_v15 = _module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), function (_41_i1) {
  return new _dafny.CodePoint('o'.codePointAt(0));
}));
            let _42_v16;
            _42_v16 = _module.D1.create_DC6(_40_v15);
            let _43_v17;
            _43_v17 = _module.D1.create_DC6(_42_v16);
            let _44_v18;
            _44_v18 = _module.D1.create_DC6(_42_v16);
            let _45_v19;
            _45_v19 = _module.D1.create_DC6(_42_v16);
            let _46_v20;
            _46_v20 = _dafny.Map.Empty.slice().updateUnsafe(_45_v19,(_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))]);
            let _47_v21;
            _47_v21 = _dafny.Map.Empty.slice().updateUnsafe(_46_v20,_module.__default.fm3(new BigNumber(172), false, p0, globalState));
            let _index6 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_39_v14).length));
            (_39_v14)[_index6] = (_47_v21).Merge(_47_v21);
            let _index7 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_39_v14).length));
            (_39_v14)[_index7] = _47_v21;
          }
        }
      }
      let _48_v22;
      _48_v22 = new _dafny.CodePoint('b'.codePointAt(0));
      let _49_i2;
      _49_i2 = _dafny.ZERO;
      L1: {
        while (_dafny.Seq.IsProperPrefixOf(_24_v0, _dafny.Seq.update(_dafny.Seq.Concat(_24_v0, _24_v0), _module.__default.safeIndex((_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))], new BigNumber((_dafny.Seq.Concat(_24_v0, _24_v0)).length)), _48_v22))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_49_i2)) {
              break L1;
            }
            _49_i2 = (_49_i2).plus(_dafny.ONE);
            let _50_v24;
            let _nw3 = Array((new BigNumber(16)).toNumber());
            _nw3[(_dafny.ZERO).toNumber()] = _33_v9;
            _nw3[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))],(_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))]);
            _nw3[(new BigNumber(2)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(3)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(4)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(5)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(6)).toNumber()] = function () {
              let _coll1 = new _dafny.Map();
              for (const _compr_1 of (_26_v3).Keys.Elements) {
                let _51_v23 = _compr_1;
                if ((_26_v3).contains(_51_v23)) {
                  _coll1.push([(_51_v23).minus((_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))]),true]);
                }
              }
              return _coll1;
            }();
            _nw3[(new BigNumber(7)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(8)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(9)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(10)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(11)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(12)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(13)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(14)).toNumber()] = _33_v9;
            _nw3[(new BigNumber(15)).toNumber()] = _33_v9;
            _50_v24 = _nw3;
            let _52_v25;
            _52_v25 = _dafny.Map.Empty.slice().updateUnsafe(_50_v24,_dafny.Seq.update(_dafny.Seq.of(p0, p0), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_34_v10, _module.__default.safeIndex(new BigNumber((_34_v10).length), new BigNumber((_34_v10).length)), (_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))])).length), new BigNumber((_dafny.Seq.of(p0, p0)).length)), (_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))]));
            _52_v25 = (_52_v25).update(_50_v24, _35_v11);
            let _53_v26;
            _53_v26 = _dafny.Map.Empty.slice().updateUnsafe((_32_v8)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_32_v8).length))],_24_v0);
            let _54_v27;
            _54_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_53_v26).length));
            _54_v27 = _54_v27;
            (globalState).f0 = !(p0);
            let _55_v28;
            let _nw4 = new _module.C0();
            _nw4.__ctor(_25_v2, _dafny.Seq.UnicodeFromString("bxfg"));
            _55_v28 = _nw4;
            let _56_v29;
            _56_v29 = _module.D1.create_DC5((_31_v7)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_31_v7).length))], _55_v28, p0);
            (globalState).f7 = (_56_v29).dtor_cf14;
          }
        }
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _57_v0;
      _57_v0 = _dafny.Seq.UnicodeFromString("sraobfhjd");
      let _58_v1;
      _58_v1 = new BigNumber(178);
      let _59_v2;
      _59_v2 = new _dafny.CodePoint('q'.codePointAt(0));
      let _60_v3;
      let _init0 = ((_61_v1) => function (_62_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC2(_61_v1, true, true, new BigNumber((_dafny.Seq.of(false)).length))).dtor_cf7,_61_v1);
      })(_58_v1);
      let _nw5 = Array((new BigNumber(8)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw5.length); _i0_0++) {
        _nw5[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _60_v3 = _nw5;
      let _63_v4;
      let _init1 = ((_64_v1) => function (_65_i1) {
        return _dafny.Set.fromElements(_64_v1);
      })(_58_v1);
      let _nw6 = Array((new BigNumber(19)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
        _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _63_v4 = _nw6;
      let _66_v5;
      let _init2 = function (_67_i2) {
        return _module.__default.safeModuloInt(_67_i2, new BigNumber(855));
      };
      let _nw7 = Array((new BigNumber(23)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw7.length); _i0_2++) {
        _nw7[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _66_v5 = _nw7;
      let _68_v6;
      let _nw8 = Array((new BigNumber(17)).toNumber()).fill([]);
      _68_v6 = _nw8;
      let _69_v7;
      _69_v7 = false;
      let _70_v8;
      _70_v8 = _dafny.Set.fromElements(_59_v2, _59_v2, new _dafny.CodePoint('u'.codePointAt(0)), _59_v2);
      let _71_globalState;
      let _nw9 = new _module.GlobalState();
      _nw9.__ctor(false, new BigNumber(237), _dafny.Seq.update(_57_v0, _module.__default.safeIndex(_58_v1, new BigNumber((_57_v0).length)), _59_v2), _60_v3, new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber(-46), _63_v4, false, _66_v5, _68_v6, _66_v5, _dafny.Map.Empty.slice().updateUnsafe(_69_v7,(_dafny.ZERO).minus(_58_v1)), false, true, _70_v8, true);
      _71_globalState = _nw9;
      _59_v2 = _59_v2;
      let _72_v9;
      _72_v9 = _dafny.MultiSet.fromElements(_69_v7);
      let _hi0 = (((_72_v9).contains(_69_v7)) ? ((_72_v9).get(_69_v7)) : (_58_v1));
      for (let _73_i3 = _58_v1; _73_i3.isLessThan(_hi0); _73_i3 = _73_i3.plus(_dafny.ONE)) {
        let _74_v10;
        _74_v10 = _module.D0.create_DC2(new BigNumber(998), _69_v7, _69_v7, new BigNumber((_57_v0).length));
        let _pat_let_tv0 = _57_v0;
        if (((_69_v7) ? ((_58_v1).isLessThanOrEqualTo(_73_i3)) : ((function (_pat_let0_0) {
          return function (_75_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_76_dt__update_hcf9_h0) {
                return _module.D0.create_DC2((_75_dt__update__tmp_h0).dtor_cf6, (_75_dt__update__tmp_h0).dtor_cf7, (_75_dt__update__tmp_h0).dtor_cf8, _76_dt__update_hcf9_h0);
              }(_pat_let1_0);
            }(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_73_i3,new BigNumber((_pat_let_tv0).length))).length));
          }(_pat_let0_0);
        }(_74_v10)).dtor_cf8))) {
          let _77_v11;
          let _init3 = function (_78_i4) {
            return !(!(!(true)));
          };
          let _nw10 = Array((new BigNumber(3)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw10.length); _i0_3++) {
            _nw10[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _77_v11 = _nw10;
          let _index8 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_77_v11).length));
          (_77_v11)[_index8] = (_73_i3).isLessThan(new BigNumber((_72_v9).cardinality()));
          let _index9 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_77_v11).length));
          (_77_v11)[_index9] = !(_69_v7) || (_69_v7);
          let _79_v12;
          _79_v12 = _dafny.Seq.of(_66_v5);
          let _80_v13;
          _80_v13 = _dafny.Map.Empty.slice().updateUnsafe((_77_v11)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_77_v11).length))],_69_v7);
          let _81_v14;
          let _nw11 = Array((new BigNumber(5)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _66_v5;
          _nw11[(_dafny.ONE).toNumber()] = _66_v5;
          _nw11[(new BigNumber(2)).toNumber()] = _66_v5;
          _nw11[(new BigNumber(3)).toNumber()] = (_79_v12)[_module.__default.safeIndex(new BigNumber((_80_v13).length), new BigNumber((_79_v12).length))];
          _nw11[(new BigNumber(4)).toNumber()] = _66_v5;
          _81_v14 = _nw11;
          let _index10 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_81_v14).length));
          (_81_v14)[_index10] = _66_v5;
          let _index11 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_81_v14).length));
          (_81_v14)[_index11] = _66_v5;
          let _82_v15;
          _82_v15 = _dafny.Map.Empty.slice().updateUnsafe(_66_v5,!(true) || (_69_v7));
          let _83_v16;
          _83_v16 = _dafny.Map.Empty.slice().updateUnsafe(_69_v7,_66_v5);
          _82_v15 = (_82_v15).update((((_83_v16).contains((_77_v11)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_77_v11).length))])) ? ((_83_v16).get((_77_v11)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_77_v11).length))])) : (_66_v5)), (_73_i3).isLessThan(_73_i3));
          (_71_globalState).f5 = (((_77_v11)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_77_v11).length))]) ? (_58_v1) : (new BigNumber(823)));
          (_71_globalState).f5 = (_58_v1).minus(_73_i3);
        } else {
          _module.__default.m0(_69_v7, _71_globalState);
          (_71_globalState).f12 = _module.__default.fm0((_69_v7) || (_69_v7), _59_v2, _59_v2, _58_v1, _71_globalState);
          _59_v2 = _59_v2;
          _58_v1 = new BigNumber((_57_v0).length);
          (_71_globalState).f0 = true;
        }
        let _84_v17;
        _84_v17 = _dafny.Map.Empty.slice().updateUnsafe(true,_69_v7);
        let _85_v18;
        _85_v18 = _dafny.Map.Empty.slice().updateUnsafe(_84_v17,_57_v0);
        (_71_globalState).f5 = new BigNumber((_dafny.Seq.Concat((((_85_v18).contains(_84_v17)) ? ((_85_v18).get(_84_v17)) : (_57_v0)), ((_69_v7) ? (_57_v0) : (_57_v0)))).length);
        let _86_v19;
        let _nw12 = new _module.C0();
        _nw12.__ctor(_73_i3, _57_v0);
        _86_v19 = _nw12;
        (_71_globalState).f0 = ((_69_v7) ? (_69_v7) : (_69_v7));
      }
      _module.__default.m0(_69_v7, _71_globalState);
      let _hi1 = _58_v1;
      for (let _87_i5 = (_dafny.ZERO).minus((_58_v1).plus(_58_v1)); _87_i5.isLessThan(_hi1); _87_i5 = _87_i5.plus(_dafny.ONE)) {
        let _rhs0 = _66_v5;
        let _rhs1 = _59_v2;
        let _rhs2 = !(_69_v7) || (_69_v7);
        let _rhs3 = _69_v7;
        let _lhs0 = _71_globalState;
        let _lhs1 = _71_globalState;
        let _lhs2 = _71_globalState;
        _lhs0.f10 = _rhs0;
        _59_v2 = _rhs1;
        _lhs1.f12 = _rhs2;
        _lhs2.f12 = _rhs3;
        let _88_v20;
        let _nw13 = new _module.C0();
        _nw13.__ctor(_87_i5, _dafny.Seq.Concat(_57_v0, _57_v0));
        _88_v20 = _nw13;
        let _89_v21;
        _89_v21 = _dafny.Seq.of((_88_v20).f16);
        _89_v21 = ((_69_v7) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(346)), ((_90_v20) => function (_91_i6) {
          return (_90_v20).f16;
        })(_88_v20))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(801)), ((_92_i5) => function (_93_i7) {
          return _92_i5;
        })(_87_i5))));
        let _94_v22;
        _94_v22 = _module.D0.create_DC0(_69_v7);
        let _95_v23;
        _95_v23 = _dafny.Map.Empty.slice().updateUnsafe(_94_v22,(_88_v20).f17);
        let _96_v24;
        _96_v24 = _module.D1.create_DC4((_88_v20).f17);
        _95_v23 = (_95_v23).update(_module.D0.create_DC0(!(_69_v7)), _dafny.Seq.Concat((_96_v24).dtor_cf11, _57_v0));
      }
      let _97_v25;
      _97_v25 = _dafny.MultiSet.fromElements(_58_v1);
      (_71_globalState).f7 = _dafny.Seq.IsPrefixOf(_module.__default.fm2(_58_v1, _97_v25, _71_globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-49)), ((_98_v2) => function (_99_i8) {
        return _98_v2;
      })(_59_v2)));
      let _100_v26;
      _100_v26 = _dafny.Set.fromElements(_69_v7, _69_v7, _69_v7);
      let _101_v27;
      _101_v27 = _dafny.Map.Empty.slice().updateUnsafe(_100_v26,_module.__default.fm0(_69_v7, _59_v2, _59_v2, (_dafny.ZERO).minus(_58_v1), _71_globalState));
      let _hi2 = new BigNumber(103);
      for (let _102_i9 = new BigNumber((_101_v27).length); _102_i9.isLessThan(_hi2); _102_i9 = _102_i9.plus(_dafny.ONE)) {
        let _103_v28;
        let _nw14 = Array((new BigNumber(11)).toNumber());
        _103_v28 = _nw14;
        let _104_v29;
        let _nw15 = new _module.C0();
        _nw15.__ctor(_102_i9, _dafny.Seq.UnicodeFromString("hrcnf"));
        _104_v29 = _nw15;
        let _105_v30;
        _105_v30 = _module.D1.create_DC5(_69_v7, _104_v29, _69_v7);
        let _pat_let_tv1 = _69_v7;
        let _pat_let_tv2 = _69_v7;
        let _index12 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_103_v28).length));
        (_103_v28)[_index12] = function (_pat_let2_0) {
          return function (_106_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_107_dt__update_hcf14_h0) {
                return function (_pat_let4_0) {
                  return function (_108_dt__update_hcf12_h0) {
                    return _module.D1.create_DC5(_108_dt__update_hcf12_h0, (_106_dt__update__tmp_h1).dtor_cf13, _107_dt__update_hcf14_h0);
                  }(_pat_let4_0);
                }(_pat_let_tv2);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_105_v30);
        let _index13 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_103_v28).length));
        (_103_v28)[_index13] = _105_v30;
        let _109_v31;
        _109_v31 = _dafny.Set.fromElements(_104_v29);
        let _rhs4 = (_dafny.ZERO).minus((_104_v29).f16);
        let _rhs5 = new BigNumber(((_109_v31).Union(_109_v31)).length);
        let _lhs3 = _71_globalState;
        let _lhs4 = _71_globalState;
        _lhs3.f5 = _rhs4;
        _lhs4.f5 = _rhs5;
        let _110_v32;
        _110_v32 = _module.D1.create_DC4((_104_v29).f17);
        let _111_v33;
        _111_v33 = _module.D1.create_DC6(_110_v32);
        _111_v33 = _module.D1.create_DC6(_110_v32);
        _module.__default.m0(_69_v7, _71_globalState);
      }
      let _112_v34;
      _112_v34 = _dafny.Map.Empty.slice().updateUnsafe(false,_69_v7);
      _112_v34 = (_112_v34).update(false, _69_v7);
      let _113_v35;
      let _nw16 = new _module.C0();
      _nw16.__ctor(_58_v1, _57_v0);
      _113_v35 = _nw16;
      (_71_globalState).f5 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_58_v1,_113_v35)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_58_v1,_113_v35))).length);
      let _114_v36;
      _114_v36 = _dafny.Map.Empty.slice().updateUnsafe((_113_v35).f16,_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), ((_115_v2) => function (_116_i10) {
        return _115_v2;
      })(_59_v2)), _dafny.Seq.UnicodeFromString("fal")));
      _114_v36 = (_114_v36).update(new BigNumber(-282), (_113_v35).f17);
      let _117_v37;
      let _nw17 = new _module.C0();
      _nw17.__ctor(_module.__default.safeModuloInt((_113_v35).f16, (_113_v35).f16), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_118_v2) => function (_119_i11) {
        return _118_v2;
      })(_59_v2)), (_113_v35).f17));
      _117_v37 = _nw17;
      _module.__default.m0(!((_117_v37).f16).isEqualTo((_113_v35).f16), _71_globalState);
      let _120_v38;
      _120_v38 = _dafny.Map.Empty.slice().updateUnsafe(_59_v2,new BigNumber((_57_v0).length));
      _120_v38 = (function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.of(_59_v2)).Elements) {
          let _121_v39 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(_59_v2), _121_v39)) {
            _coll2.push([_121_v39,new BigNumber((_dafny.Set.fromElements((_117_v37).f16, (_113_v35).f16, (_117_v37).f16, new BigNumber((_112_v34).length), _58_v1)).length)]);
          }
        }
        return _coll2;
      }()).Merge(_120_v38);
      _69_v7 = _69_v7;
      (_71_globalState).f5 = (_117_v37).f16;
      _module.__default.m0(_69_v7, _71_globalState);
      let _122_v40;
      _122_v40 = _dafny.Map.Empty.slice().updateUnsafe(_69_v7,_113_v35);
      (_71_globalState).f5 = (((_117_v37).f16).multipliedBy(new BigNumber(((_module.D2.create_DC7(_122_v40)).dtor_cf16).length))).plus(new BigNumber(255));
      process.stdout.write((_57_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_58_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_59_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v3)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(6)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(7)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(8)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(9)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(10)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(11)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(12)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(13)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(14)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(15)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(16)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(17)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_63_v4)[new BigNumber(18)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_69_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_70_v8).equals(_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_71_globalState).f2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_71_globalState).f3)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(6)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(7)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(8)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(9)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(10)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(11)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(12)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(13)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(14)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(15)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(16)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(17)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState.f6)[new BigNumber(18)]).equals(_dafny.Set.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f8)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f10)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f11).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState.f14).equals(_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_72_v9).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v25).equals(_dafny.MultiSet.fromElements(new BigNumber(178)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v26).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_112_v34).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v35).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_113_v35).f17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v36).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(178),_dafny.Seq.UnicodeFromString("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfal")).updateUnsafe(new BigNumber(-282),_dafny.Seq.UnicodeFromString("sraobfhjd")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_117_v37).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_117_v37).f17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v38).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_122_v40).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4, cf5) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf6, cf7, cf8, cf9) {
      let $dt = new D0(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf10) {
      let $dt = new D0(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.Map.Empty, false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC5(cf12, cf13, cf14) {
      let $dt = new D1(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC4(cf11) {
      let $dt = new D1(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6(cf15) {
      let $dt = new D1(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + this.cf11.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12 && this.cf13 === other.cf13 && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false, null, false);
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
    static create_DC8(cf17, cf18, cf19) {
      let $dt = new D2(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC7(cf16) {
      let $dt = new D2(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.ZERO, _dafny.ZERO, null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f5 = _dafny.ZERO;
      this.f6 = [];
      this.f7 = false;
      this.f10 = [];
      this.f12 = false;
      this.f14 = _dafny.Set.Empty;
      this._f1 = _dafny.ZERO;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f3 = [];
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = [];
      this._f9 = [];
      this._f11 = _dafny.Map.Empty;
      this._f13 = false;
      this._f15 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      return;
    }
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
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f16 = _dafny.ZERO;
      this._f17 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f16, f17) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((((!((_dafny.MultiSet.fromElements((_module.D0.create_DC1(true, _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f16), true, (_this).f16, (_this).f16)).dtor_cf1)).IsSubsetOf(_dafny.MultiSet.fromElements(false, true)))) ? ((_this).f17) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(552)), function (_123_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })))).length);
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
