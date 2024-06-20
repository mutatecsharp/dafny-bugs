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
    static fm4(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-771),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(523),false));
    };
    static fm5(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("lvxthx");
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return new BigNumber((((_dafny.MultiSet.fromElements(new BigNumber(693))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(533), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).Elements) {
          let _0_v0 = _compr_0;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).contains(_0_v0)) {
            _coll0.push([_0_v0,(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vnemwohk")).length), new BigNumber((_dafny.Seq.of(new BigNumber(762))).length))).length)))]);
          }
        }
        return _coll0;
      }()).length)))).Union(_dafny.MultiSet.fromElements(new BigNumber(((_module.D1.create_DC4(!((_module.D1.create_DC4(false, true, _dafny.Seq.UnicodeFromString("rookixm"), new BigNumber(276), new BigNumber((_dafny.Seq.of(false)).length))).dtor_cf9), true, _dafny.Seq.UnicodeFromString("kowjt"), new BigNumber(344), new BigNumber(452))).dtor_cf10).length), new BigNumber(-249), new BigNumber(293), (_module.D1.create_DC4(false, true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(299)), function (_1_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}), new BigNumber(361), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-739)), function (_2_i1) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length))).dtor_cf12))).cardinality());
    };
    static fm9(p0, globalState) {
      return _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("khaefkx"), _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(false, false, false, false, (_module.D9.create_DC23(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_module.D10.create_DC25(_dafny.Set.fromElements(false, false))).dtor_cf49).length)),false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)))).cardinality())), true, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(934),_dafny.Set.fromElements(true)), new BigNumber(-844))).dtor_cf45)).cardinality()), new BigNumber(112)), _dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(117)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(724)), _dafny.Seq.of(new BigNumber(-729))));
    };
    static fm10(p0, p1, globalState) {
      return ((_module.D5.create_DC15(_dafny.MultiSet.fromElements(false, !(!(true)), false, !(!(false))))).dtor_cf31).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false, false))));
    };
    static fm11(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm12(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false)).length)), _dafny.Seq.of(new BigNumber(678))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ucej")).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(66))).length))));
    };
    static fm13(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Set.fromElements(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.UnicodeFromString("nmj")).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(718))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-50))));
    };
    static fm14(p0, globalState) {
      if (((true) ? (false) : (true))) {
        return _module.D1.create_DC7(_module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(true, true)).length))).length),false)));
      } else {
        return _module.D1.create_DC7(_module.D1.create_DC5(true, false, new BigNumber((_dafny.Seq.UnicodeFromString("gc")).length), new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)));
      }
    };
    static fm15(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-246)).isLessThan(new BigNumber(36)),_dafny.Seq.of(new BigNumber(-678)));
    };
    static fm16(globalState) {
      return _dafny.Seq.of((_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))).IsDisjointFrom(_dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)))));
    };
    static fm17(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(631)), function (_3_i0) {
        return _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)));
      }), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), function (_4_i1) {
        return _dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)));
      })));
    };
    static fm18(globalState) {
      if ((_dafny.MultiSet.fromElements(true)).IsDisjointFrom(_dafny.MultiSet.fromElements(!(!(!(!(true)))), false, !(false), false, false))) {
        return _module.D6.create_DC17(_dafny.Seq.of(true));
      } else {
        return _module.D6.create_DC17(_dafny.Seq.of(!(true)));
      }
    };
    static fm19(p0, globalState) {
      return (_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), function (_5_i0) {
        return new BigNumber(208);
      })), _dafny.MultiSet.fromElements(new BigNumber(-582), new BigNumber(809)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), function (_6_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length)))).Union((_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("seyfsdjs")).length), new BigNumber(136))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("pwmmx")).length))).length)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true, true)).length)))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("dnlhtu")).length))))));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D4.create_DC13((_module.D4.create_DC13(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)))).Elements) {
    let _7_v0 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0))), _7_v0)) {
      _coll1.push([_7_v0,new BigNumber((_dafny.Seq.UnicodeFromString("penvbw")).length)]);
    }
  }
  return _coll1;
}())).dtor_cf30);
      if (_source0.is_DC14) {
        return _module.D1.create_DC6(true, new BigNumber(19), new BigNumber(-603), false);
      } else {
        let _8___mcc_h0 = (_source0).cf30;
        let _9_cf30 = _8___mcc_h0;
        return _module.D1.create_DC6(false, (_module.D1.create_DC5(false, !(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-945))).length), new BigNumber(284))).dtor_cf15, new BigNumber(915), !(!(true)));
      }
    };
    static fm21(globalState) {
      return function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-357), new BigNumber(743))) {
          let _10_v0 = _compr_2;
          if (((new BigNumber(-357)).isLessThanOrEqualTo(_10_v0)) && ((_10_v0).isLessThan(new BigNumber(743)))) {
            _coll2.add(_module.__default.safeModuloInt(_10_v0, new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)), _dafny.Map.Empty.slice().updateUnsafe(true,true))).length)));
          }
        }
        return _coll2;
      }();
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return _module.D5.create_DC15((_dafny.MultiSet.fromElements(!(!(true)))).Intersect(_dafny.MultiSet.fromElements(false)));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      if (true) {
        return _module.D2.create_DC8(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-344))));
      } else {
        return _module.D2.create_DC8(_dafny.MultiSet.fromElements(new BigNumber(852)));
      }
    };
    static fm24(globalState) {
      return (!((_dafny.Set.fromElements(new BigNumber(-500))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(458), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),false)).length)))))) === ((new BigNumber(53)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("bjjecxhfo")).length)));
    };
    static fm25(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber(796))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Set.fromElements(true)).length))).length)));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _11_v0;
      _11_v0 = true;
      let _12_v1;
      _12_v1 = _dafny.Seq.of(_11_v0, _11_v0, _11_v0, _11_v0, _11_v0);
      let _hi0 = new BigNumber((_12_v1).length);
      for (let _13_i0 = p1; _13_i0.isLessThan(_hi0); _13_i0 = _13_i0.plus(_dafny.ONE)) {
        let _14_v3;
        _14_v3 = _dafny.Set.fromElements(new BigNumber(-156), p0);
        let _15_v4;
        _15_v4 = _dafny.Map.Empty.slice().updateUnsafe(_13_i0,new BigNumber(983));
        _11_v0 = !(function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_14_v3).Elements) {
            let _16_v2 = _compr_3;
            if ((_14_v3).contains(_16_v2)) {
              _coll3.push([_module.__default.safeDivisionInt(_16_v2, _13_i0),p0]);
            }
          }
          return _coll3;
        }()).equals(_15_v4);
        let _17_v5;
        let _init0 = ((_18_p0) => function (_19_i1) {
          return _module.__default.safeDivisionInt(_19_i1, _18_p0);
        })(p0);
        let _nw0 = Array((new BigNumber(25)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _17_v5 = _nw0;
        let _index0 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_17_v5).length));
        (_17_v5)[_index0] = p0;
        let _20_v6;
        _20_v6 = _dafny.Seq.UnicodeFromString("riotpp");
        let _21_v7;
        _21_v7 = _module.D1.create_DC4(_11_v0, _11_v0, _20_v6, p1, p0);
        let _22_v8;
        _22_v8 = _dafny.Seq.of(_12_v1);
        let _23_v9;
        let _nw1 = new _module.C3();
        _nw1.__ctor(new BigNumber((_dafny.Seq.update((_22_v8)[_module.__default.safeIndex(p0, new BigNumber((_22_v8).length))], _module.__default.safeIndex(new BigNumber(364), new BigNumber(((_22_v8)[_module.__default.safeIndex(p0, new BigNumber((_22_v8).length))]).length)), _11_v0)).length), _13_i0, _11_v0);
        _23_v9 = _nw1;
        let _24_v10;
        _24_v10 = _dafny.Seq.of(_23_v9, _23_v9, _23_v9);
        let _25_v11;
        _25_v11 = _dafny.MultiSet.fromElements(_13_i0, p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_24_v10).length))), _13_i0, new BigNumber((_20_v6).length));
        let _index1 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_17_v5).length));
        let _rhs0 = (p1).minus((_21_v7).dtor_cf12);
        let _rhs1 = new BigNumber(((_module.D2.create_DC8(_25_v11)).dtor_cf22).cardinality());
        let _rhs2 = _13_i0;
        let _lhs0 = _17_v5;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_17_v5).length));
        _lhs0[_lhs1] = _rhs0;
        r2 = _rhs1;
        r3 = _rhs2;
        r2 = ((p1).multipliedBy(new BigNumber((_20_v6).length))).plus(p0);
        let _26_v12;
        let _nw2 = new _module.C2();
        _nw2.__ctor(_module.__default.safeDivisionInt(p1, (_17_v5)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_17_v5).length))]));
        _26_v12 = _nw2;
      }
      r2 = p0;
      let _27_v13;
      _27_v13 = new _dafny.CodePoint('q'.codePointAt(0));
      _27_v13 = function (_source1) {
        let _28___mcc_h0 = _source1;
        let _29_cf40 = _28___mcc_h0;
        return new _dafny.CodePoint('c'.codePointAt(0));
      }(_27_v13);
      let _30_v14;
      let _init1 = function (_31_i2) {
        return (_31_i2).plus(new BigNumber(834));
      };
      let _nw3 = Array((new BigNumber(4)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
        _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _30_v14 = _nw3;
      let _32_v15;
      let _nw4 = new _module.C3();
      _nw4.__ctor(p0, p0, _11_v0);
      _32_v15 = _nw4;
      let _33_v16;
      _33_v16 = _dafny.MultiSet.fromElements(_32_v15);
      let _index2 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_30_v14).length));
      (_30_v14)[_index2] = new BigNumber((_33_v16).cardinality());
      let _index3 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_30_v14).length));
      (_30_v14)[_index3] = p0;
      let _34_v17;
      _34_v17 = _module.D4.create_DC13(_module.__default.fm25(globalState));
      let _pat_let_tv0 = _32_v15;
      let _pat_let_tv1 = _11_v0;
      let _pat_let_tv2 = _32_v15;
      _11_v0 = function (_source2) {
        if (_source2.is_DC14) {
          return (_pat_let_tv0).f6;
        } else {
          let _35___mcc_h1 = (_source2).cf30;
          let _36_cf30 = _35___mcc_h1;
          return !(_pat_let_tv1) || ((_pat_let_tv2).f6);
        }
      }(_34_v17);
      let _37_v18;
      let _nw5 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
      _37_v18 = _nw5;
      let _38_v19;
      _38_v19 = _dafny.MultiSet.fromElements((_32_v15).f6, (_32_v15).f6);
      let _39_v20;
      let _nw6 = new _module.C0();
      _nw6.__ctor(new BigNumber(867), (((_38_v19).contains(_11_v0)) ? ((_38_v19).get(_11_v0)) : ((_30_v14)[_module.__default.safeIndex(new BigNumber(835), new BigNumber((_30_v14).length))])));
      _39_v20 = _nw6;
      let _40_v21;
      _40_v21 = _dafny.Map.Empty.slice().updateUnsafe(_39_v20,(_32_v15).f6);
      let _41_v22;
      _41_v22 = _dafny.Map.Empty.slice().updateUnsafe(((_11_v0) ? (_37_v18) : (_37_v18)),(new BigNumber(((_40_v21).update(_39_v20, (_32_v15).f6)).length)).isLessThanOrEqualTo(new BigNumber(831)));
      _41_v22 = (_41_v22).update(_37_v18, _11_v0);
      let _42_v23;
      _42_v23 = _dafny.Seq.UnicodeFromString("tjkieshnd");
      r0 = _42_v23;
      let _43_v24;
      _43_v24 = _dafny.Map.Empty.slice().updateUnsafe((_32_v15).f6,_11_v0);
      let _44_v25;
      _44_v25 = _dafny.Seq.of(p0, new BigNumber((_38_v19).cardinality()), new BigNumber((_dafny.Seq.of(_42_v23, _42_v23, _42_v23, _dafny.Seq.UnicodeFromString("eao"), _dafny.Seq.UnicodeFromString("iolhb"))).length), new BigNumber((((_43_v24).update((_32_v15).f6, _11_v0)).update(_11_v0, _11_v0)).length));
      let _nw7 = Array((new BigNumber(10)).toNumber());
      _nw7[(_dafny.ZERO).toNumber()] = (_30_v14)[_module.__default.safeIndex(new BigNumber(835), new BigNumber((_30_v14).length))];
      _nw7[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw7[(new BigNumber(2)).toNumber()] = new BigNumber(112);
      _nw7[(new BigNumber(3)).toNumber()] = p0;
      _nw7[(new BigNumber(4)).toNumber()] = (p1).multipliedBy(_module.__default.fm6((_39_v20).f12, (_32_v15).f5, (_32_v15).f6, (_39_v20).f13, globalState));
      _nw7[(new BigNumber(5)).toNumber()] = p1;
      _nw7[(new BigNumber(6)).toNumber()] = (_44_v25)[_module.__default.safeIndex((_39_v20).f12, new BigNumber((_44_v25).length))];
      _nw7[(new BigNumber(7)).toNumber()] = (_39_v20).f12;
      _nw7[(new BigNumber(8)).toNumber()] = ((_30_v14)[_module.__default.safeIndex(new BigNumber(835), new BigNumber((_30_v14).length))]).plus((_39_v20).f13);
      _nw7[(new BigNumber(9)).toNumber()] = (_39_v20).f12;
      r1 = _nw7;
      r2 = (_30_v14)[_module.__default.safeIndex(new BigNumber(835), new BigNumber((_30_v14).length))];
      r3 = new BigNumber((_dafny.Seq.of(_30_v14)).length);
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _45_v0;
      _45_v0 = false;
      let _46_v1;
      _46_v1 = new BigNumber(82);
      let _47_v2;
      _47_v2 = new _dafny.CodePoint('d'.codePointAt(0));
      let _48_v3;
      _48_v3 = _dafny.MultiSet.fromElements(_47_v2);
      let _49_v4;
      _49_v4 = _dafny.Map.Empty.slice().updateUnsafe(_45_v0,_46_v1);
      let _50_v5;
      _50_v5 = _dafny.Map.Empty.slice().updateUnsafe(_45_v0,_45_v0);
      let _51_v6;
      _51_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_50_v5).length),_45_v0);
      let _52_v7;
      _52_v7 = _dafny.Seq.UnicodeFromString("oxasa");
      let _53_v8;
      _53_v8 = _dafny.Seq.of(_45_v0);
      let _54_v9;
      let _nw8 = Array((new BigNumber(29)).toNumber());
      _nw8[(_dafny.ZERO).toNumber()] = _46_v1;
      _nw8[(_dafny.ONE).toNumber()] = (((_48_v3).contains(_47_v2)) ? ((_48_v3).get(_47_v2)) : (new BigNumber(87)));
      _nw8[(new BigNumber(2)).toNumber()] = new BigNumber(872);
      _nw8[(new BigNumber(3)).toNumber()] = new BigNumber((_49_v4).length);
      _nw8[(new BigNumber(4)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(5)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(6)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_46_v1);
      _nw8[(new BigNumber(8)).toNumber()] = new BigNumber((_51_v6).length);
      _nw8[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.of(_46_v1, _46_v1)).length);
      _nw8[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_46_v1);
      _nw8[(new BigNumber(11)).toNumber()] = new BigNumber((_52_v7).length);
      _nw8[(new BigNumber(12)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(13)).toNumber()] = new BigNumber((_52_v7).length);
      _nw8[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_55_v2) => function (_56_i0) {
        return _55_v2;
      })(_47_v2))).length);
      _nw8[(new BigNumber(15)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(16)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(17)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(18)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(19)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(20)).toNumber()] = new BigNumber((_53_v8).length);
      _nw8[(new BigNumber(21)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), ((_57_v2) => function (_58_i1) {
        return _57_v2;
      })(_47_v2))).length);
      _nw8[(new BigNumber(23)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(24)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(25)).toNumber()] = (_dafny.ZERO).minus(_46_v1);
      _nw8[(new BigNumber(26)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(27)).toNumber()] = _46_v1;
      _nw8[(new BigNumber(28)).toNumber()] = _46_v1;
      _54_v9 = _nw8;
      let _59_v10;
      _59_v10 = _dafny.Map.Empty.slice().updateUnsafe(_45_v0,_54_v9);
      let _60_v11;
      _60_v11 = _dafny.Set.fromElements(_46_v1);
      let _61_v12;
      _61_v12 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(475)), ((_62_v2) => function (_63_i2) {
        return _62_v2;
      })(_47_v2)));
      let _64_globalState;
      let _nw9 = new _module.GlobalState();
      _nw9.__ctor(new BigNumber(766), _59_v10, _dafny.Seq.of(_60_v11), _61_v12, true);
      _64_globalState = _nw9;
      let _65_v13;
      let _66_v14;
      let _67_v15;
      let _68_v16;
      let _out0;
      let _out1;
      let _out2;
      let _out3;
      let _outcollector0 = _module.__default.m0((new BigNumber(861)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), ((_69_v2) => function (_70_i3) {
        return _69_v2;
      })(_47_v2))).length)), _46_v1, _64_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _out3 = _outcollector0[3];
      _65_v13 = _out0;
      _66_v14 = _out1;
      _67_v15 = _out2;
      _68_v16 = _out3;
      let _71_v17;
      _71_v17 = _dafny.MultiSet.fromElements(_68_v16, _67_v15, _67_v15);
      _68_v16 = _module.__default.safeModuloInt(_68_v16, new BigNumber(((_71_v17).Intersect(_71_v17)).cardinality()));
      let _72_v18;
      let _nw10 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
      _72_v18 = _nw10;
      let _index4 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length));
      (_72_v18)[_index4] = _71_v17;
      let _73_v19;
      _73_v19 = _dafny.Seq.of(_71_v17);
      let _74_v20;
      _74_v20 = _dafny.Seq.of(new BigNumber(720));
      let _pat_let_tv3 = _52_v7;
      let _pat_let_tv4 = _49_v4;
      let _index5 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length));
      (_72_v18)[_index5] = ((_71_v17).Intersect((_73_v19)[_module.__default.safeIndex(new BigNumber((_65_v13).length), new BigNumber((_73_v19).length))])).Intersect((_dafny.MultiSet.FromArray(_74_v20)).Union(_dafny.MultiSet.FromArray((function (_pat_let0_0) {
        return function (_75_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_76_dt__update_hcf1_h0) {
              return function (_pat_let2_0) {
                return function (_77_dt__update_hcf3_h0) {
                  return _module.D0.create_DC1(_76_dt__update_hcf1_h0, (_75_dt__update__tmp_h0).dtor_cf2, _77_dt__update_hcf3_h0, (_75_dt__update__tmp_h0).dtor_cf4);
                }(_pat_let2_0);
              }(_pat_let_tv4);
            }(_pat_let1_0);
          }(_pat_let_tv3);
        }(_pat_let0_0);
      }(_module.D0.create_DC1(_65_v13, _67_v15, _49_v4, _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_60_v11).length), _46_v1), _module.__default.safeIndex(_67_v15, new BigNumber((_dafny.Seq.of(new BigNumber((_60_v11).length), _46_v1)).length)), _67_v15)))).dtor_cf4)));
      let _78_v21;
      _78_v21 = _dafny.Map.Empty.slice().updateUnsafe(_68_v16,_module.__default.safeModuloInt(new BigNumber(125), _68_v16));
      let _79_v22;
      _79_v22 = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("vewema"), _46_v1, _49_v4, _74_v20);
      _78_v21 = (_78_v21).update((_67_v15).minus((_79_v22).dtor_cf2), (_dafny.ZERO).minus((new BigNumber(741)).multipliedBy(_68_v16)));
      _67_v15 = _68_v16;
      let _80_v23;
      let _init2 = ((_81_v5) => function (_82_i4) {
        return _81_v5;
      })(_50_v5);
      let _nw11 = Array((new BigNumber(27)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw11.length); _i0_2++) {
        _nw11[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _80_v23 = _nw11;
      let _index6 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_80_v23).length));
      (_80_v23)[_index6] = (_50_v5).Merge(_50_v5);
      let _index7 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_80_v23).length));
      (_80_v23)[_index7] = (_50_v5).Merge((_50_v5).Merge(_50_v5));
      _47_v2 = new _dafny.CodePoint('y'.codePointAt(0));
      let _83_v24;
      let _nw12 = Array((new BigNumber(2)).toNumber()).fill(false);
      _83_v24 = _nw12;
      let _84_v25;
      _84_v25 = _dafny.Map.Empty.slice().updateUnsafe(_83_v24,!(_45_v0));
      _84_v25 = (_84_v25).update(_83_v24, !(!(_45_v0) || (_45_v0)));
      if ((_46_v1).isLessThanOrEqualTo(new BigNumber((_60_v11).length))) {
        let _85_v26;
        _85_v26 = _dafny.Map.Empty.slice().updateUnsafe(_45_v0,_74_v20);
        let _86_v27;
        _86_v27 = _module.D0.create_DC2(_85_v26, !(_45_v0));
        let _source3 = _86_v27;
        if (_source3.is_DC1) {
          let _87___mcc_h0 = (_source3).cf1;
          let _88___mcc_h1 = (_source3).cf2;
          let _89___mcc_h2 = (_source3).cf3;
          let _90___mcc_h3 = (_source3).cf4;
          let _91_cf4 = _90___mcc_h3;
          let _92_cf3 = _89___mcc_h2;
          let _93_cf2 = _88___mcc_h1;
          let _94_cf1 = _87___mcc_h0;
          _45_v0 = _dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-964)), ((_95_v2) => function (_96_i5) {
            return _95_v2;
          })(_47_v2)), _module.__default.safeIndex(_68_v16, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-964)), ((_97_v2) => function (_98_i5) {
            return _97_v2;
          })(_47_v2))).length)), _47_v2), _dafny.Seq.UnicodeFromString("l")), _module.__default.safeIndex(_67_v15, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-964)), ((_99_v2) => function (_100_i5) {
            return _99_v2;
          })(_47_v2)), _module.__default.safeIndex(_68_v16, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-964)), ((_101_v2) => function (_102_i5) {
            return _101_v2;
          })(_47_v2))).length)), _47_v2), _dafny.Seq.UnicodeFromString("l"))).length)), _47_v2), _52_v7);
          let _103_v28;
          let _nw13 = new _module.C0();
          _nw13.__ctor(_93_cf2, (_dafny.ZERO).minus(_46_v1));
          _103_v28 = _nw13;
          _93_cf2 = _module.__default.safeDivisionInt((_103_v28).f13, _93_cf2);
          _67_v15 = (_68_v16).minus((_103_v28).f13);
        } else if (_source3.is_DC2) {
          let _104___mcc_h4 = (_source3).cf5;
          let _105___mcc_h5 = (_source3).cf6;
          let _106_cf6 = _105___mcc_h5;
          let _107_cf5 = _104___mcc_h4;
          _67_v15 = _68_v16;
          let _index8 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_83_v24).length));
          (_83_v24)[_index8] = _106_cf6;
          let _index9 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_54_v9).length));
          (_54_v9)[_index9] = _67_v15;
          let _108_v29;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_46_v1, _67_v15);
          _108_v29 = _nw14;
          let _109_v30;
          _109_v30 = _dafny.Map.Empty.slice().updateUnsafe(_108_v29,true);
          let _110_v31;
          _110_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_109_v30).length),_49_v4);
          let _index10 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_83_v24).length));
          (_83_v24)[_index10] = (new BigNumber((_module.__default.fm5(_49_v4, _110_v31, _64_globalState)).length)).isLessThan(_67_v15);
          let _111_v32;
          _111_v32 = _dafny.Set.fromElements(true);
          let _index11 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_83_v24).length));
          let _index12 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length));
          let _index13 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_54_v9).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_83_v24).length));
          let _rhs3 = (_45_v0) === (false);
          let _rhs4 = (_71_v17).Union((_72_v18)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length))]);
          let _rhs5 = (_108_v29).f12;
          let _rhs6 = _83_v24;
          let _rhs7 = !(((false) ? ((_111_v32).IsSubsetOf(_111_v32)) : (_45_v0)));
          let _lhs2 = _83_v24;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_83_v24).length));
          let _lhs4 = _72_v18;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length));
          let _lhs6 = _54_v9;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_54_v9).length));
          let _lhs8 = _83_v24;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_83_v24).length));
          _lhs2[_lhs3] = _rhs3;
          _lhs4[_lhs5] = _rhs4;
          _lhs6[_lhs7] = _rhs5;
          _83_v24 = _rhs6;
          _lhs8[_lhs9] = _rhs7;
          _47_v2 = _module.__default.fm11(_47_v2, (!((_83_v24)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_83_v24).length))])) || ((_83_v24)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_83_v24).length))]), (_108_v29).fm8((_83_v24)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_83_v24).length))], _52_v7, _47_v2, _64_globalState), _64_globalState);
          let _112_v33;
          let _nw15 = new _module.C4();
          _nw15.__ctor((_108_v29).f12, (_54_v9)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_54_v9).length))], _106_cf6);
          _112_v33 = _nw15;
          _112_v33 = _112_v33;
        } else {
          let _113___mcc_h6 = (_source3).cf0;
          let _114_cf0 = _113___mcc_h6;
          let _115_v34;
          let _116_v35;
          let _117_v36;
          let _118_v37;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(_67_v15, _46_v1, _64_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _115_v34 = _out4;
          _116_v35 = _out5;
          _117_v36 = _out6;
          _118_v37 = _out7;
          let _119_v38;
          let _nw16 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _119_v38 = _nw16;
          let _index15 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_119_v38).length));
          (_119_v38)[_index15] = _65_v13;
          let _index16 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_119_v38).length));
          (_119_v38)[_index16] = _65_v13;
          let _120_v39;
          let _init3 = function (_121_i6) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          };
          let _nw17 = Array((new BigNumber(20)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw17.length); _i0_3++) {
            _nw17[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _120_v39 = _nw17;
          let _index17 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_120_v39).length));
          (_120_v39)[_index17] = _47_v2;
          let _index18 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_120_v39).length));
          (_120_v39)[_index18] = _47_v2;
          _67_v15 = _46_v1;
        }
        let _122_v40;
        _122_v40 = _47_v2;
        _122_v40 = _47_v2;
        let _nw18 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _66_v14 = _nw18;
        _47_v2 = _module.__default.fm11(_47_v2, _45_v0, !(_45_v0) || (_module.__default.fm24(_64_globalState)), _64_globalState);
        let _123_v41;
        let _nw19 = new _module.C1();
        _nw19.__ctor(_67_v15, _67_v15, _67_v15, _45_v0);
        _123_v41 = _nw19;
        let _124_v42;
        let _nw20 = Array((new BigNumber(3)).toNumber());
        _nw20[(_dafny.ZERO).toNumber()] = ((_45_v0) ? (_123_v41) : (_123_v41));
        _nw20[(_dafny.ONE).toNumber()] = _123_v41;
        _nw20[(new BigNumber(2)).toNumber()] = _123_v41;
        _124_v42 = _nw20;
        let _index19 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_124_v42).length));
        (_124_v42)[_index19] = _123_v41;
        let _index20 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_124_v42).length));
        (_124_v42)[_index20] = (((_123_v41).f6) ? (_123_v41) : (_123_v41));
      } else {
        let _125_v43;
        let _nw21 = new _module.C3();
        _nw21.__ctor(_68_v16, ((_dafny.ZERO).minus(new BigNumber(-469))).minus(_46_v1), _45_v0);
        _125_v43 = _nw21;
        _68_v16 = (_125_v43).f8;
        let _126_v44;
        _126_v44 = _dafny.Set.fromElements(_45_v0);
        let _127_v45;
        _127_v45 = _dafny.Map.Empty.slice().updateUnsafe(_126_v44,_52_v7);
        let _128_v46;
        let _nw22 = new _module.C1();
        _nw22.__ctor((_125_v43).f8, new BigNumber(431), _67_v15, _45_v0);
        _128_v46 = _nw22;
        let _129_v47;
        _129_v47 = _dafny.Set.fromElements(_128_v46, _128_v46);
        _45_v0 = !_dafny.Seq.contains(_dafny.Seq.update((((_127_v45).contains(_126_v44)) ? ((_127_v45).get(_126_v44)) : (_65_v13)), _module.__default.safeIndex(new BigNumber((_129_v47).length), new BigNumber(((((_127_v45).contains(_126_v44)) ? ((_127_v45).get(_126_v44)) : (_65_v13))).length)), _47_v2), _47_v2);
        let _rhs8 = ((_45_v0) ? (true) : ((_128_v46).f6));
        let _rhs9 = ((_67_v15).plus(_46_v1)).multipliedBy((_dafny.ZERO).minus((new BigNumber((_53_v8).length)).minus(new BigNumber((_53_v8).length))));
        _45_v0 = _rhs8;
        _68_v16 = _rhs9;
        let _rhs10 = _47_v2;
        let _rhs11 = (_dafny.ZERO).minus(_68_v16);
        _47_v2 = _rhs10;
        _68_v16 = _rhs11;
      }
      _46_v1 = _module.__default.safeModuloInt((_74_v20)[_module.__default.safeIndex(_67_v15, new BigNumber((_74_v20).length))], _67_v15);
      if (_module.__default.fm24(_64_globalState)) {
        _45_v0 = _45_v0;
        let _130_v49;
        _130_v49 = _module.D1.create_DC3(function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of (_dafny.Set.fromElements(_67_v15)).Elements) {
    let _131_v48 = _compr_4;
    if ((_dafny.Set.fromElements(_67_v15)).contains(_131_v48)) {
      _coll4.push([_module.__default.safeModuloInt(_131_v48, new BigNumber(236)),_45_v0]);
    }
  }
  return _coll4;
}());
        let _132_v50;
        let _nw23 = new _module.C3();
        _nw23.__ctor(new BigNumber(-302), new BigNumber((_74_v20).length), _45_v0);
        _132_v50 = _nw23;
        let _133_v51;
        let _nw24 = new _module.C1();
        _nw24.__ctor(_68_v16, _67_v15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_130_v49,_132_v50)).length), _45_v0);
        _133_v51 = _nw24;
        let _134_v52;
        _134_v52 = _module.D5.create_DC16(_133_v51, (_133_v51).f6, _68_v16);
        let _135_v53;
        let _nw25 = new _module.C0();
        _nw25.__ctor((_module.__default.fm6(new BigNumber((_dafny.MultiSet.fromElements(_45_v0)).cardinality()), _46_v1, _45_v0, _46_v1, _64_globalState)).plus(_module.__default.fm6(new BigNumber((_74_v20).length), _68_v16, _45_v0, new BigNumber((_dafny.Seq.UnicodeFromString("puoy")).length), _64_globalState)), _module.__default.safeModuloInt((_dafny.ZERO).minus(_68_v16), (_134_v52).dtor_cf34));
        _135_v53 = _nw25;
        _47_v2 = _module.__default.fm11(_47_v2, (_45_v0) && ((_133_v51).f6), (_133_v51).f6, _64_globalState);
        _47_v2 = _47_v2;
        let _index21 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_54_v9).length));
        (_54_v9)[_index21] = ((_135_v53).f13).minus((_135_v53).f13);
        let _index22 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_54_v9).length));
        let _rhs12 = _65_v13;
        let _rhs13 = ((_135_v53).f12).plus(_67_v15);
        let _rhs14 = _46_v1;
        let _lhs10 = _54_v9;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_54_v9).length));
        _65_v13 = _rhs12;
        _46_v1 = _rhs13;
        _lhs10[_lhs11] = _rhs14;
      } else {
        let _136_v54;
        let _nw26 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
        _136_v54 = _nw26;
        _136_v54 = _136_v54;
        _83_v24 = _83_v24;
        _45_v0 = _45_v0;
        let _137_v55;
        _137_v55 = _dafny.Map.Empty.slice().updateUnsafe(true,_74_v20);
        let _pat_let_tv5 = _45_v0;
        let _source4 = function (_pat_let3_0) {
          return function (_138_dt__update__tmp_h1) {
            return function (_pat_let4_0) {
              return function (_139_dt__update_hcf28_h0) {
                return _module.D3.create_DC12(_139_dt__update_hcf28_h0, (_138_dt__update__tmp_h1).dtor_cf29);
              }(_pat_let4_0);
            }(_pat_let_tv5);
          }(_pat_let3_0);
        }(_module.D3.create_DC12((_module.D0.create_DC2(_137_v55, _module.__default.fm24(_64_globalState))).dtor_cf6, new BigNumber((_71_v17).cardinality())));
        if (_source4.is_DC12) {
          let _140___mcc_h7 = (_source4).cf28;
          let _141___mcc_h8 = (_source4).cf29;
          let _142_cf29 = _141___mcc_h8;
          let _143_cf28 = _140___mcc_h7;
          let _144_v56;
          _144_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("skjvslqs"),_67_v15);
          let _145_v57;
          _145_v57 = _dafny.Map.Empty.slice().updateUnsafe(_144_v56,_83_v24);
          let _146_v58;
          _146_v58 = _dafny.Seq.of(_83_v24);
          _145_v57 = (_145_v57).update(_144_v56, (_146_v58)[_module.__default.safeIndex(_142_cf29, new BigNumber((_146_v58).length))]);
          let _147_v59;
          let _148_v60;
          let _149_v61;
          let _150_v62;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0(_67_v15, _46_v1, _64_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _147_v59 = _out8;
          _148_v60 = _out9;
          _149_v61 = _out10;
          _150_v62 = _out11;
          let _151_v63;
          let _nw27 = new _module.C0();
          _nw27.__ctor(_46_v1, _68_v16);
          _151_v63 = _nw27;
          let _152_v64;
          _152_v64 = _dafny.Map.Empty.slice().updateUnsafe(_151_v63,new BigNumber(((_72_v18)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length))]).cardinality()));
          let _153_v65;
          _153_v65 = _dafny.MultiSet.fromElements(_152_v64);
          _51_v6 = ((_module.__default.fm24(_64_globalState)) ? ((_51_v6).update(_142_cf29, _45_v0)) : (_dafny.Map.Empty.slice().updateUnsafe((((_153_v65).contains(_152_v64)) ? ((_153_v65).get(_152_v64)) : (_149_v61)),_45_v0)));
          let _index23 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_54_v9).length));
          (_54_v9)[_index23] = _142_cf29;
          let _154_v66;
          _154_v66 = _dafny.Set.fromElements(_45_v0);
          let _index24 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_54_v9).length));
          (_54_v9)[_index24] = new BigNumber((_154_v66).length);
        } else {
          let _155___mcc_h9 = (_source4).cf27;
          let _156_cf27 = _155___mcc_h9;
          let _157_v67;
          _157_v67 = _dafny.Map.Empty.slice().updateUnsafe((_60_v11).Intersect(_60_v11),!(_module.__default.fm24(_64_globalState)));
          _157_v67 = (_157_v67).update(_60_v11, _45_v0);
          _46_v1 = new BigNumber((((_79_v22).dtor_cf3).Merge(_module.__default.fm13(_dafny.Seq.UnicodeFromString("ptcfinu"), _68_v16, (_73_v19)[_module.__default.safeIndex(_67_v15, new BigNumber((_73_v19).length))], _64_globalState))).length);
          let _158_v68;
          _158_v68 = _dafny.Map.Empty.slice().updateUnsafe(_46_v1,_dafny.Map.Empty.slice().updateUnsafe(_45_v0,_68_v16));
          let _159_v69;
          let _nw28 = new _module.C0();
          _nw28.__ctor((_67_v15).plus(new BigNumber(722)), new BigNumber((_module.__default.fm5(_49_v4, _158_v68, _64_globalState)).length));
          _159_v69 = _nw28;
          let _160_v70;
          let _nw29 = new _module.C3();
          _nw29.__ctor((_159_v69).f12, _68_v16, _45_v0);
          _160_v70 = _nw29;
          let _rhs15 = _159_v69;
          let _rhs16 = _160_v70;
          _159_v69 = _rhs15;
          _160_v70 = _rhs16;
          let _161_v71;
          let _nw30 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _161_v71 = _nw30;
          let _index25 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_161_v71).length));
          (_161_v71)[_index25] = _dafny.Seq.Concat(_52_v7, _52_v7);
          let _index26 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_161_v71).length));
          (_161_v71)[_index26] = _52_v7;
        }
        let _162_v72;
        let _163_v73;
        let _164_v74;
        let _165_v75;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector3 = _module.__default.m0(_module.__default.safeModuloInt(_46_v1, _67_v15), _68_v16, _64_globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _out15 = _outcollector3[3];
        _162_v72 = _out12;
        _163_v73 = _out13;
        _164_v74 = _out14;
        _165_v75 = _out15;
      }
      _54_v9 = _66_v14;
      let _166_v76;
      let _nw31 = new _module.C4();
      _nw31.__ctor(_68_v16, _67_v15, _45_v0);
      _166_v76 = _nw31;
      let _167_v77;
      _167_v77 = _dafny.Seq.of(_166_v76);
      let _168_v78;
      _168_v78 = _module.D9.create_DC22(_167_v77);
      let _169_v79;
      _169_v79 = _dafny.MultiSet.fromElements(_166_v76);
      let _170_v81;
      _170_v81 = _module.D2.create_DC10(_module.D2.create_DC8((_71_v17).update(_46_v1, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of (_52_v7).Elements) {
    let _171_v80 = _compr_5;
    if (_dafny.Seq.contains(_52_v7, _171_v80)) {
      _coll5.push([_171_v80,!(_45_v0)]);
    }
  }
  return _coll5;
}()).length))))));
      let _172_v82;
      _172_v82 = _module.D2.create_DC8((_72_v18)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length))]);
      let _pat_let_tv6 = _167_v77;
      let _source5 = (((_169_v79).IsProperSubsetOf(_dafny.MultiSet.FromArray((function (_pat_let5_0) {
        return function (_173_dt__update__tmp_h2) {
          return function (_pat_let6_0) {
            return function (_174_dt__update_hcf42_h0) {
              return _module.D9.create_DC22(_174_dt__update_hcf42_h0);
            }(_pat_let6_0);
          }(_pat_let_tv6);
        }(_pat_let5_0);
      }(_168_v78)).dtor_cf42))) ? (_170_v81) : (_module.D2.create_DC10(_172_v82)));
      if (_source5.is_DC9) {
        let _175___mcc_h10 = (_source5).cf23;
        let _176___mcc_h11 = (_source5).cf24;
        let _177___mcc_h12 = (_source5).cf25;
        let _178_cf25 = _177___mcc_h12;
        let _179_cf24 = _176___mcc_h11;
        let _180_cf23 = _175___mcc_h10;
        _46_v1 = _67_v15;
        _65_v13 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-45)), ((_181_v7, _182_v76) => function (_183_i7) {
          return (_181_v7)[_module.__default.safeIndex(_182_v76.f7, new BigNumber((_181_v7).length))];
        })(_52_v7, _166_v76));
        _180_cf23 = _180_cf23;
        let _184_v83;
        let _nw32 = new _module.C0();
        _nw32.__ctor(_46_v1, _67_v15);
        _184_v83 = _nw32;
        let _185_v84;
        _185_v84 = _dafny.Map.Empty.slice().updateUnsafe(_45_v0,_184_v83);
        let _186_v85;
        let _nw33 = new _module.C2();
        _nw33.__ctor(new BigNumber((_185_v84).length));
        _186_v85 = _nw33;
      } else if (_source5.is_DC8) {
        let _187___mcc_h13 = (_source5).cf22;
        let _188_cf22 = _187___mcc_h13;
        let _index27 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_66_v14).length));
        (_66_v14)[_index27] = (_166_v76.f7).plus((_dafny.ZERO).minus(_166_v76.f7));
        let _index28 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_66_v14).length));
        (_66_v14)[_index28] = _166_v76.f7;
        (_166_v76).m1(_166_v76.f7, _64_globalState);
        let _189_v86;
        let _nw34 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _189_v86 = _nw34;
        let _index29 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_189_v86).length));
        (_189_v86)[_index29] = _dafny.Seq.update(_53_v8, _module.__default.safeIndex(new BigNumber((_53_v8).length), new BigNumber((_53_v8).length)), _45_v0);
        let _index30 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_189_v86).length));
        (_189_v86)[_index30] = _dafny.Seq.Concat(((_45_v0) ? (_dafny.Seq.of(_45_v0, true)) : (_53_v8)), _53_v8);
        let _190_v87;
        let _nw35 = new _module.C0();
        _nw35.__ctor(_46_v1, new BigNumber((_53_v8).length));
        _190_v87 = _nw35;
        let _191_v88;
        _191_v88 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_190_v87),(_190_v87).f13);
        let _192_v89;
        _192_v89 = _dafny.Set.fromElements(_190_v87);
        let _193_v90;
        _193_v90 = _dafny.Map.Empty.slice().updateUnsafe(((_191_v88).update(_192_v89, new BigNumber((_78_v21).length))).update(_192_v89, new BigNumber((_60_v11).length)),!(!(_45_v0)));
        _193_v90 = (_193_v90).update(_191_v88, !((_190_v87).fm8(_45_v0, (_79_v22).dtor_cf1, _47_v2, _64_globalState)));
      } else {
        let _194___mcc_h14 = (_source5).cf26;
        let _195_cf26 = _194___mcc_h14;
        _67_v15 = _module.__default.fm6(_67_v15, _module.__default.safeModuloInt(_166_v76.f7, _166_v76.f7), !((_166_v76).fm0(false, _64_globalState)), new BigNumber(((_72_v18)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length))]).cardinality()), _64_globalState);
        let _196_v91;
        let _nw36 = Array((new BigNumber(4)).toNumber());
        _nw36[(_dafny.ZERO).toNumber()] = _166_v76;
        _nw36[(_dafny.ONE).toNumber()] = _166_v76;
        _nw36[(new BigNumber(2)).toNumber()] = _166_v76;
        _nw36[(new BigNumber(3)).toNumber()] = _166_v76;
        _196_v91 = _nw36;
        let _197_v92;
        _197_v92 = _module.D3.create_DC11(_66_v14);
        let _198_v93;
        _198_v93 = _dafny.Map.Empty.slice().updateUnsafe(_196_v91,_197_v92);
        _198_v93 = (_198_v93).update(_196_v91, _197_v92);
        _78_v21 = (_78_v21).update(_67_v15, ((true) ? (_166_v76.f7) : (_67_v15)));
        let _199_v94;
        _199_v94 = _dafny.Map.Empty.slice().updateUnsafe(_45_v0,_83_v24);
        let _200_v95;
        let _nw37 = new _module.C3();
        _nw37.__ctor(_68_v16, _46_v1, _45_v0);
        _200_v95 = _nw37;
        let _201_v96;
        _201_v96 = _dafny.Set.fromElements(_200_v95);
        let _202_v97;
        _202_v97 = _dafny.Map.Empty.slice().updateUnsafe(_199_v94,_201_v96);
        _202_v97 = (_202_v97).update(_199_v94, _201_v96);
      }
      let _203_i8;
      _203_i8 = _dafny.ZERO;
      L0: {
        while ((_166_v76).fm0((_46_v1).isLessThanOrEqualTo(new BigNumber(383)), _64_globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_203_i8)) {
              break L0;
            }
            _203_i8 = (_203_i8).plus(_dafny.ONE);
            let _204_v98;
            _204_v98 = _dafny.Set.fromElements(_83_v24, _83_v24);
            let _205_v99;
            let _206_v100;
            let _207_v101;
            let _208_v102;
            let _out16;
            let _out17;
            let _out18;
            let _out19;
            let _outcollector4 = _module.__default.m0(_module.__default.safeModuloInt(new BigNumber((_204_v98).length), _166_v76.f7), ((_45_v0) ? (_67_v15) : (_166_v76.f7)), _64_globalState);
            _out16 = _outcollector4[0];
            _out17 = _outcollector4[1];
            _out18 = _outcollector4[2];
            _out19 = _outcollector4[3];
            _205_v99 = _out16;
            _206_v100 = _out17;
            _207_v101 = _out18;
            _208_v102 = _out19;
            _45_v0 = ((_dafny.Seq.IsProperPrefixOf(_205_v99, _65_v13)) ? (_module.__default.fm24(_64_globalState)) : (_45_v0));
            let _209_v103;
            let _nw38 = new _module.C0();
            _nw38.__ctor(_207_v101, _67_v15);
            _209_v103 = _nw38;
            let _210_v104;
            let _nw39 = new _module.C1();
            _nw39.__ctor(_67_v15, _module.__default.fm6(_68_v16, _208_v102, _45_v0, _207_v101, _64_globalState), (_209_v103).f12, _45_v0);
            _210_v104 = _nw39;
            let _211_v105;
            _211_v105 = _module.D4.create_DC14();
            let _212_v106;
            _212_v106 = _dafny.Map.Empty.slice().updateUnsafe(_210_v104,_211_v105);
            (_166_v76).f7 = _module.__default.fm6(_module.__default.safeDivisionInt(new BigNumber(148), _166_v76.f7), new BigNumber((_212_v106).length), !(_68_v16).isEqualTo(_166_v76.f7), (_68_v16).multipliedBy(_46_v1), _64_globalState);
          }
        }
      }
      let _213_v107;
      let _214_v108;
      let _215_v109;
      let _216_v110;
      let _out20;
      let _out21;
      let _out22;
      let _out23;
      let _outcollector5 = _module.__default.m0(((((_72_v18)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length))]).contains(_68_v16)) ? (((_72_v18)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_72_v18).length))]).get(_68_v16)) : (_68_v16)), _67_v15, _64_globalState);
      _out20 = _outcollector5[0];
      _out21 = _outcollector5[1];
      _out22 = _outcollector5[2];
      _out23 = _outcollector5[3];
      _213_v107 = _out20;
      _214_v108 = _out21;
      _215_v109 = _out22;
      _216_v110 = _out23;
      let _217_v111;
      _217_v111 = _dafny.MultiSet.fromElements(_45_v0, !(_45_v0));
      _45_v0 = (!(_217_v111).contains(_45_v0)) === (_45_v0);
      process.stdout.write(_dafny.toString(_45_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_46_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_47_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_48_v3).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_49_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(82)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_50_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_51_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_52_v7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_53_v8, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_54_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_59_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_v11).equals(_dafny.Set.fromElements(new BigNumber(82)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_61_v12, _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_64_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_64_globalState).f1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_64_globalState).f2, _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(82))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_64_globalState).f3, _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_64_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_65_v13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v14)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_67_v15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_68_v16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v17).equals(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(1352), new BigNumber(1352)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v18)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(1352)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_73_v19, _dafny.Seq.of(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(1352), new BigNumber(1352))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_74_v20, _dafny.Seq.of(new BigNumber(720)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_78_v21).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO).updateUnsafe(new BigNumber(1270),new BigNumber(-741)).updateUnsafe(new BigNumber(4),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_79_v22).dtor_cf1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_79_v22).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_79_v22).dtor_cf3).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(82)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_79_v22).dtor_cf4, _dafny.Seq.of(new BigNumber(720)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(25)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_80_v23)[new BigNumber(26)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_84_v25).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_166_v76.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_167_v77).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_168_v78).dtor_cf42).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_169_v79).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_170_v81).dtor_cf26).dtor_cf22).equals(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(1352), new BigNumber(1352), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_172_v82).dtor_cf22).equals(_dafny.MultiSet.fromElements(new BigNumber(1352)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_203_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_213_v107).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v108)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_215_v109));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_216_v110));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_217_v111).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf5, cf6) {
      let $dt = new D0(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + this.cf1.toVerbatimString(true) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.Map.Empty, _dafny.Seq.of());
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
    static create_DC4(cf8, cf9, cf10, cf11, cf12) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC5(cf13, cf14, cf15, cf16) {
      let $dt = new D1(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC6(cf17, cf18, cf19, cf20) {
      let $dt = new D1(2);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D1(3);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC7(cf21) {
      let $dt = new D1(4);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get is_DC7() { return this.$tag === 4; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + this.cf10.toVerbatimString(true) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 4) {
        return "D1.DC7" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, false, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC9(cf23, cf24, cf25) {
      let $dt = new D2(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC8(cf22) {
      let $dt = new D2(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC10(cf26) {
      let $dt = new D2(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC10" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC9(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC12(cf28, cf29) {
      let $dt = new D3(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC11(cf27) {
      let $dt = new D3(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf27 === other.cf27;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC12(false, _dafny.ZERO);
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
    static create_DC14() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC13(cf30) {
      let $dt = new D4(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf30) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14();
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
    static create_DC16(cf32, cf33, cf34) {
      let $dt = new D5(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC15(cf31) {
      let $dt = new D5(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf32 === other.cf32 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(null, false, _dafny.ZERO);
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
    static create_DC18(cf36) {
      let $dt = new D6(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC19(cf37, cf38, cf39) {
      let $dt = new D6(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC17(cf35) {
      let $dt = new D6(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(false);
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
    static create_DC20(cf40) {
      let $dt = new D7(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC21(cf41) {
      let $dt = new D8(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
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
    static create_DC23(cf43, cf44, cf45, cf46, cf47) {
      let $dt = new D9(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC22(cf42) {
      let $dt = new D9(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC24(cf48) {
      let $dt = new D9(2);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44) && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(_dafny.ZERO, _dafny.ZERO, false, _dafny.Map.Empty, _dafny.ZERO);
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
    static create_DC26(cf50, cf51, cf52) {
      let $dt = new D10(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC25(cf49) {
      let $dt = new D10(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC26(false, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.of();
      this._f3 = _dafny.Seq.of();
      this._f4 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(925), new BigNumber(394))) {
          let _218_v0 = _compr_6;
          if (((new BigNumber(925)).isLessThanOrEqualTo(_218_v0)) && ((_218_v0).isLessThan(new BigNumber(394)))) {
            _coll6.push([(_218_v0).multipliedBy((_this).f12),true]);
          }
        }
        return _coll6;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,true)).Merge((_module.D1.create_DC3(function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(397), new BigNumber(116))) {
    let _219_v1 = _compr_7;
    if (((new BigNumber(397)).isLessThanOrEqualTo(_219_v1)) && ((_219_v1).isLessThan(new BigNumber(116)))) {
      _coll7.push([(_219_v1).plus((_this).f13),false]);
    }
  }
  return _coll7;
}())).dtor_cf7));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (!(true)) || (false);
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this.f10 = _dafny.ZERO;
      this._f11 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f10, f11, f5, f6) {
      let _this = this;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f11).isEqualTo((_dafny.ZERO).minus(((_this).f5).plus((_this).f11)));
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f10);
    };
    m1(p0, globalState) {
      let _this = this;
      let _220_v0;
      _220_v0 = _dafny.Seq.UnicodeFromString("ju");
      let _221_v1;
      _221_v1 = _dafny.MultiSet.fromElements(((false) ? (_220_v0) : (_dafny.Seq.UnicodeFromString("kcobgiw"))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_222_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }), _dafny.Seq.Concat(_220_v0, _dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), function (_223_i1) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })));
      (_this).f10 = (((_221_v1).contains(_220_v0)) ? ((_221_v1).get(_220_v0)) : (p0));
      let _224_v2;
      let _nw40 = Array((new BigNumber(15)).toNumber()).fill(false);
      _224_v2 = _nw40;
      let _225_v3;
      _225_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p0);
      let _226_v4;
      _226_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_225_v3).length),_this.f10);
      let _227_v5;
      _227_v5 = _dafny.Set.fromElements(true, (_this).f6);
      let _228_v6;
      _228_v6 = _dafny.Seq.of(new BigNumber((_227_v5).length));
      let _229_v7;
      _229_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_228_v6);
      let _230_v8;
      _230_v8 = _module.D0.create_DC2(_229_v7, !((_this).f6));
      let _231_v9;
      _231_v9 = _dafny.Set.fromElements(_230_v8);
      let _hi1 = (((_226_v4).contains(new BigNumber((_220_v0).length))) ? ((_226_v4).get(new BigNumber((_220_v0).length))) : (new BigNumber((_231_v9).length)));
      for (let _232_i2 = _module.__default.safeDivisionInt((_this).f5, new BigNumber((_module.__default.fm4(new BigNumber((_dafny.Seq.of(_224_v2)).length), globalState)).length)); _232_i2.isLessThan(_hi1); _232_i2 = _232_i2.plus(_dafny.ONE)) {
        let _233_v10;
        _233_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_228_v6));
        let _234_v11;
        _234_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,false);
        _225_v3 = (_225_v3).update((_module.D0.create_DC2((_233_v10)[_module.__default.safeIndex(p0, new BigNumber((_233_v10).length))], (_this).f6)).dtor_cf6, new BigNumber((_234_v11).length));
        let _235_v12;
        _235_v12 = _dafny.Map.Empty.slice().updateUnsafe(_232_i2,_220_v0);
        let _236_v13;
        _236_v13 = new _dafny.CodePoint('b'.codePointAt(0));
        let _237_v14;
        _237_v14 = _module.D0.create_DC1(_dafny.Seq.update(_220_v0, _module.__default.safeIndex(_this.f10, new BigNumber((_220_v0).length)), _236_v13), (_this).f5, _225_v3, _228_v6);
        let _238_v15;
        _238_v15 = _dafny.Seq.of(_220_v0, _220_v0);
        let _239_v16;
        _239_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(652),_225_v3);
        let _240_v17;
        let _nw41 = Array((new BigNumber(24)).toNumber());
        _nw41[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_220_v0, _dafny.Seq.UnicodeFromString("md"));
        _nw41[(_dafny.ONE).toNumber()] = (((_235_v12).contains(p0)) ? ((_235_v12).get(p0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(70)), function (_241_i3) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })));
        _nw41[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_220_v0, _module.__default.safeIndex(p0, new BigNumber((_220_v0).length)), _236_v13), _220_v0);
        _nw41[(new BigNumber(3)).toNumber()] = _dafny.Seq.update((_237_v14).dtor_cf1, _module.__default.safeIndex((_this).f11, new BigNumber(((_237_v14).dtor_cf1).length)), _236_v13);
        _nw41[(new BigNumber(4)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("o");
        _nw41[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("th");
        _nw41[(new BigNumber(7)).toNumber()] = (_238_v15)[_module.__default.safeIndex(_this.f10, new BigNumber((_238_v15).length))];
        _nw41[(new BigNumber(8)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(9)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(10)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(11)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(12)).toNumber()] = _module.__default.fm5(_225_v3, _239_v16, globalState);
        _nw41[(new BigNumber(13)).toNumber()] = (((_235_v12).contains((_this).f11)) ? ((_235_v12).get((_this).f11)) : (_220_v0));
        _nw41[(new BigNumber(14)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("chm"), _220_v0);
        _nw41[(new BigNumber(16)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(17)).toNumber()] = _220_v0;
        _nw41[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_220_v0, _module.__default.safeIndex((_this).f11, new BigNumber((_220_v0).length)), _236_v13);
        _nw41[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _220_v0);
        _nw41[(new BigNumber(20)).toNumber()] = (_module.D0.create_DC1(_dafny.Seq.UnicodeFromString("ylxgswe"), _module.__default.fm6((_this).f5, (_this).f5, (_this).f6, (_this).f5, globalState), _225_v3, _228_v6)).dtor_cf1;
        _nw41[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_220_v0, _220_v0);
        _nw41[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_220_v0, _dafny.Seq.UnicodeFromString("soi"));
        _nw41[(new BigNumber(23)).toNumber()] = _220_v0;
        _240_v17 = _nw41;
        let _index31 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_240_v17).length));
        (_240_v17)[_index31] = _220_v0;
        let _242_v18;
        _242_v18 = false;
        let _243_v19;
        _243_v19 = _dafny.Seq.of((_this).fm0((_this).fm0(true, globalState), globalState), _242_v18);
        let _244_v20;
        _244_v20 = _dafny.Seq.of(_243_v19);
        let _index32 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_240_v17).length));
        let _rhs17 = (_232_i2).multipliedBy((_this.f10).multipliedBy(p0));
        let _rhs18 = p0;
        let _rhs19 = _220_v0;
        let _rhs20 = (_this.f10).isLessThan(new BigNumber((_dafny.Seq.Concat(_244_v20, _244_v20)).length));
        let _lhs12 = _this;
        let _lhs13 = _this;
        let _lhs14 = _240_v17;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_240_v17).length));
        _lhs12.f10 = _rhs17;
        _lhs13.f10 = _rhs18;
        _lhs14[_lhs15] = _rhs19;
        _242_v18 = _rhs20;
        (_this).f10 = _232_i2;
        (_this).f10 = (_dafny.ZERO).minus(((_this).f11).minus(p0));
      }
      let _245_v22;
      _245_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(_225_v3, function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(111), new BigNumber(211))) {
          let _246_v21 = _compr_8;
          if (((new BigNumber(111)).isLessThanOrEqualTo(_246_v21)) && ((_246_v21).isLessThan(new BigNumber(211)))) {
            _coll8.push([(_246_v21).multipliedBy((_this).f11),_225_v3]);
          }
        }
        return _coll8;
      }(), globalState),(_this).f11);
      (_this).f10 = new BigNumber((_245_v22).length);
      let _hi2 = _this.f10;
      for (let _247_i4 = (_this).f5; _247_i4.isLessThan(_hi2); _247_i4 = _247_i4.plus(_dafny.ONE)) {
        let _248_v23;
        let _nw42 = new _module.C0();
        _nw42.__ctor(_module.__default.fm6((_dafny.ZERO).minus(new BigNumber((_220_v0).length)), _247_i4, (_this).f6, (_this).f5, globalState), (_this).f11);
        _248_v23 = _nw42;
        if ((((_this).f6) ? ((_this).f6) : (((_this).f6) === ((_this).f6)))) {
          let _249_v24;
          let _250_v25;
          let _251_v26;
          let _252_v27;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector6 = _module.__default.m0(_this.f10, (_228_v6)[_module.__default.safeIndex((_248_v23).f13, new BigNumber((_228_v6).length))], globalState);
          _out24 = _outcollector6[0];
          _out25 = _outcollector6[1];
          _out26 = _outcollector6[2];
          _out27 = _outcollector6[3];
          _249_v24 = _out24;
          _250_v25 = _out25;
          _251_v26 = _out26;
          _252_v27 = _out27;
          let _253_v28;
          _253_v28 = _dafny.MultiSet.fromElements((_this).f6);
          let _254_v29;
          _254_v29 = _dafny.Map.Empty.slice().updateUnsafe(_253_v28,(_dafny.ZERO).minus((_248_v23).f12));
          let _255_v30;
          _255_v30 = _dafny.MultiSet.fromElements((_248_v23).f12, p0, new BigNumber((_249_v24).length), _this.f10, new BigNumber(612));
          let _256_v31;
          _256_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_253_v28);
          let _257_v32;
          _257_v32 = _module.D2.create_DC8(_255_v30);
          let _258_v33;
          _258_v33 = _dafny.Map.Empty.slice().updateUnsafe(_224_v2,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)));
          let _259_v34;
          let _nw43 = Array((new BigNumber(15)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = _255_v30;
          _nw43[(_dafny.ONE).toNumber()] = (_255_v30).Difference(_255_v30);
          _nw43[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray(_228_v6);
          _nw43[(new BigNumber(3)).toNumber()] = _255_v30;
          _nw43[(new BigNumber(4)).toNumber()] = _255_v30;
          _nw43[(new BigNumber(5)).toNumber()] = _255_v30;
          _nw43[(new BigNumber(6)).toNumber()] = _255_v30;
          _nw43[(new BigNumber(7)).toNumber()] = (_255_v30).Union(_255_v30);
          _nw43[(new BigNumber(8)).toNumber()] = _255_v30;
          _nw43[(new BigNumber(9)).toNumber()] = (_255_v30).Difference(_255_v30);
          _nw43[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), ((_260_v23) => function (_261_i5) {
            return (_260_v23).f13;
          })(_248_v23)), (_module.__default.fm9((_248_v23).f13, globalState)).dtor_cf4));
          _nw43[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_228_v6);
          _nw43[(new BigNumber(12)).toNumber()] = _255_v30;
          _nw43[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_228_v6, _dafny.Seq.of(new BigNumber(((((_256_v31).contains(_247_i4)) ? ((_256_v31).get(_247_i4)) : (_253_v28))).cardinality()), (_248_v23).f13, new BigNumber((_220_v0).length))));
          _nw43[(new BigNumber(14)).toNumber()] = ((_257_v32).dtor_cf22).update((((_258_v33).contains(_224_v2)) ? ((_258_v33).get(_224_v2)) : (new BigNumber((_225_v3).length))), _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(457)), function (_262_i6) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)));
          _259_v34 = _nw43;
          let _index33 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_259_v34).length));
          (_259_v34)[_index33] = _255_v30;
          let _263_v35;
          _263_v35 = true;
          let _index34 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_224_v2).length));
          (_224_v2)[_index34] = _263_v35;
          let _264_v36;
          _264_v36 = _dafny.Seq.of(true, (_this).f6, false);
          let _265_v37;
          _265_v37 = _module.D0.create_DC1(_249_v24, _252_v27, _225_v3, _dafny.Seq.of((_dafny.ZERO).minus((_248_v23).f13)));
          let _266_v38;
          _266_v38 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index35 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_259_v34).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_224_v2).length));
          let _rhs21 = _254_v29;
          let _rhs22 = _255_v30;
          let _rhs23 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_264_v36, _dafny.Seq.of((_this).f6, !((_this).f6))), _dafny.Seq.of(_263_v35));
          let _rhs24 = (_248_v23).fm8(_263_v35, (_265_v37).dtor_cf1, _266_v38, globalState);
          let _lhs16 = _259_v34;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_259_v34).length));
          let _lhs18 = _224_v2;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_224_v2).length));
          _254_v29 = _rhs21;
          _lhs16[_lhs17] = _rhs22;
          _263_v35 = _rhs23;
          _lhs18[_lhs19] = _rhs24;
          _249_v24 = _249_v24;
          let _index37 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_259_v34).length));
          (_259_v34)[_index37] = _255_v30;
          let _index38 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_224_v2).length));
          let _rhs25 = _module.__default.safeModuloInt((_this).f5, (_this).f11);
          let _rhs26 = _dafny.Seq.IsProperPrefixOf(_249_v24, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), ((_267_v38) => function (_268_i7) {
            return _267_v38;
          })(_266_v38)), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(125)), function (_269_i8) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), _module.__default.safeIndex((_248_v23).f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(125)), function (_270_i8) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          })).length)), _266_v38)));
          let _rhs27 = _250_v25;
          let _rhs28 = _252_v27;
          let _rhs29 = _263_v35;
          let _lhs20 = _224_v2;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_224_v2).length));
          let _lhs22 = _this;
          _252_v27 = _rhs25;
          _lhs20[_lhs21] = _rhs26;
          _250_v25 = _rhs27;
          _lhs22.f10 = _rhs28;
          _263_v35 = _rhs29;
        } else {
          let _271_v39;
          _271_v39 = new _dafny.CodePoint('w'.codePointAt(0));
          _271_v39 = _271_v39;
          (_this).f10 = ((!((_this).f6)) ? (new BigNumber((_220_v0).length)) : ((_248_v23).f12));
          let _272_v40;
          let _nw44 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _272_v40 = _nw44;
          let _273_v41;
          _273_v41 = _dafny.MultiSet.fromElements((_this).f6);
          let _274_v42;
          _274_v42 = _dafny.Seq.of((_this).f6, (_this).f6);
          let _275_v43;
          _275_v43 = _dafny.MultiSet.fromElements(new BigNumber((_220_v0).length), p0, _this.f10, new BigNumber((_274_v42).length));
          let _index39 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_272_v40).length));
          (_272_v40)[_index39] = (_273_v41).Difference(_module.__default.fm10(_275_v43, (_248_v23).f13, globalState));
          let _index40 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_272_v40).length));
          (_272_v40)[_index40] = _273_v41;
          _271_v39 = _271_v39;
          let _276_v44;
          _276_v44 = true;
          _276_v44 = (new BigNumber((_dafny.Seq.Concat(_220_v0, _dafny.Seq.UnicodeFromString("qchhltun"))).length)).isLessThan((_dafny.ZERO).minus((_247_i4).plus(_this.f10)));
        }
        if ((_this).f6) {
          let _277_v45;
          let _nw45 = new _module.C0();
          _nw45.__ctor(new BigNumber(401), (((_this).f6) ? (_247_i4) : (new BigNumber(-2))));
          _277_v45 = _nw45;
          (_this).f10 = (_dafny.ZERO).minus((_this).f11);
          let _278_v46;
          let _nw46 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _278_v46 = _nw46;
          let _279_v47;
          _279_v47 = _dafny.Set.fromElements(_278_v46, _278_v46, _278_v46);
          let _index41 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_278_v46).length));
          (_278_v46)[_index41] = (_this).f5;
          let _280_v48;
          _280_v48 = _dafny.Map.Empty.slice().updateUnsafe((_277_v45).f12,_220_v0);
          let _index42 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_278_v46).length));
          let _rhs30 = (_module.__default.safeModuloInt(new BigNumber(478), (_277_v45).f12)).plus(new BigNumber(((_280_v48).Merge(_280_v48)).length));
          let _rhs31 = _279_v47;
          let _rhs32 = _dafny.Seq.UnicodeFromString("ihvp");
          let _rhs33 = _module.__default.safeModuloInt((_277_v45).f13, (_248_v23).f12);
          let _lhs23 = _this;
          let _lhs24 = _278_v46;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_278_v46).length));
          _lhs23.f10 = _rhs30;
          _279_v47 = _rhs31;
          _220_v0 = _rhs32;
          _lhs24[_lhs25] = _rhs33;
          _224_v2 = _224_v2;
          let _281_v49;
          let _nw47 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _281_v49 = _nw47;
          let _282_v50;
          _282_v50 = _dafny.Seq.of((_this).f6);
          let _rhs34 = (((new BigNumber(115)).isLessThanOrEqualTo((_278_v46)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_278_v46).length))])) ? (_281_v49) : (_281_v49));
          let _rhs35 = _dafny.Seq.Concat(_dafny.Seq.Concat(_282_v50, _282_v50), _282_v50);
          _281_v49 = _rhs34;
          _282_v50 = _rhs35;
        } else {
          let _283_v51;
          let _init4 = ((_284_v23) => function (_285_i9) {
            return _module.__default.safeDivisionInt(_285_i9, (_284_v23).f12);
          })(_248_v23);
          let _nw48 = Array((new BigNumber(3)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw48.length); _i0_4++) {
            _nw48[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _283_v51 = _nw48;
          let _index43 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_283_v51).length));
          (_283_v51)[_index43] = p0;
          let _286_v52;
          _286_v52 = _dafny.Seq.of(!((_this).f6));
          let _287_v53;
          _287_v53 = new _dafny.CodePoint('e'.codePointAt(0));
          let _288_v54;
          _288_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_226_v4).length),(_this).f6);
          let _289_v55;
          _289_v55 = _dafny.Map.Empty.slice().updateUnsafe((_248_v23).f13,_225_v3);
          let _index44 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_283_v51).length));
          (_283_v51)[_index44] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_220_v0, _220_v0), _module.__default.safeIndex(new BigNumber((_286_v52).length), new BigNumber((_dafny.Seq.Concat(_220_v0, _220_v0)).length)), _module.__default.fm11(_287_v53, true, (_this).f6, globalState)), _dafny.Seq.Concat(_module.__default.fm5((_225_v3).update((_this).f6, new BigNumber((_288_v54).length)), _289_v55, globalState), _220_v0))).length);
          (_this).f10 = (_248_v23).f12;
          let _index45 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_224_v2).length));
          (_224_v2)[_index45] = !((_this).f6) || ((_this).f6);
          let _290_v56;
          _290_v56 = true;
          let _291_v57;
          let _nw49 = Array((new BigNumber(25)).toNumber()).fill([]);
          _291_v57 = _nw49;
          let _index46 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_291_v57).length));
          (_291_v57)[_index46] = _283_v51;
          let _292_v58;
          _292_v58 = _dafny.MultiSet.fromElements(_290_v56, (_this).f6);
          let _293_v59;
          _293_v59 = _dafny.MultiSet.fromElements((_248_v23).f12);
          let _index47 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_224_v2).length));
          let _index48 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_291_v57).length));
          let _rhs36 = ((_this).f6) === ((_module.__default.fm10(_293_v59, new BigNumber(100), globalState)).IsProperSubsetOf(_292_v58));
          let _rhs37 = _290_v56;
          let _rhs38 = _283_v51;
          let _lhs26 = _224_v2;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_224_v2).length));
          let _lhs28 = _291_v57;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_291_v57).length));
          _lhs26[_lhs27] = _rhs36;
          _290_v56 = _rhs37;
          _lhs28[_lhs29] = _rhs38;
          let _294_v60;
          _294_v60 = _dafny.Map.Empty.slice().updateUnsafe(_290_v56,false);
          _287_v53 = ((!((new BigNumber((_294_v60).length)).isLessThanOrEqualTo(new BigNumber(294)))) ? (_287_v53) : (_287_v53));
          _290_v56 = !(_dafny.Seq.IsProperPrefixOf(_286_v52, _dafny.Seq.of((_286_v52)[_module.__default.safeIndex((_248_v23).f13, new BigNumber((_286_v52).length))])));
        }
        let _295_v61;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_module.__default.safeModuloInt(new BigNumber(950), (_this).f5), p0);
        _295_v61 = _nw50;
      }
      let _296_v62;
      _296_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      (_this).f10 = (new BigNumber((_296_v62).length)).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_228_v6, _module.__default.safeIndex(p0, new BigNumber((_228_v6).length)), _this.f10), _dafny.Seq.Create(_module.__default.abs(new BigNumber(101)), function (_297_i10) {
        return new BigNumber(711);
      }))).length));
      let _298_v63;
      _298_v63 = false;
      _298_v63 = !((_this).f6);
      return;
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _299_v0;
      _299_v0 = _dafny.Seq.of(true, (_this).f6, (_this).fm0((_this).f6, globalState));
      let _300_v1;
      _300_v1 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),_299_v0);
      _300_v1 = (_300_v1).update(!(false), _299_v0);
      let _hi3 = (_this).f5;
      for (let _301_i0 = p3; _301_i0.isLessThan(_hi3); _301_i0 = _301_i0.plus(_dafny.ONE)) {
        let _302_v2;
        _302_v2 = _dafny.MultiSet.fromElements(p2);
        let _303_v3;
        _303_v3 = _dafny.Seq.of(new BigNumber((_302_v2).cardinality()), _module.__default.fm6(p3, (_this).f5, (_this).f6, new BigNumber(-771), globalState));
        let _304_v4;
        _304_v4 = _dafny.MultiSet.fromElements(_this.f10, (_303_v3)[_module.__default.safeIndex(p2, new BigNumber((_303_v3).length))], new BigNumber(929));
        let _305_v5;
        _305_v5 = _dafny.Seq.of(_303_v3);
        let _306_v6;
        let _nw51 = new _module.C0();
        _nw51.__ctor((_this).f5, (((_304_v4).contains(new BigNumber((_305_v5).length))) ? ((_304_v4).get(new BigNumber((_305_v5).length))) : (_this.f10)));
        _306_v6 = _nw51;
        let _307_v7;
        _307_v7 = _module.D0.create_DC2(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_303_v3), (_this).f6);
        let _308_v8;
        _308_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of(p3));
        let _309_v9;
        _309_v9 = _dafny.MultiSet.fromElements(_307_v7, (((_this).f6) ? (_307_v7) : (function (_pat_let7_0) {
          return function (_310_dt__update__tmp_h0) {
            return function (_pat_let8_0) {
              return function (_311_dt__update_hcf6_h0) {
                return _module.D0.create_DC2((_310_dt__update__tmp_h0).dtor_cf5, _311_dt__update_hcf6_h0);
              }(_pat_let8_0);
            }((_this).f6);
          }(_pat_let7_0);
        }(_module.D0.create_DC2(_308_v8, false)))));
        _309_v9 = _309_v9;
        let _312_v10;
        _312_v10 = false;
        _312_v10 = _312_v10;
        let _313_v11;
        _313_v11 = _module.D0.create_DC0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(455)), ((_314_p3) => function (_315_i1) {
  return _314_p3;
})(p3)));
        let _source6 = _313_v11;
        if (_source6.is_DC1) {
          let _316___mcc_h0 = (_source6).cf1;
          let _317___mcc_h1 = (_source6).cf2;
          let _318___mcc_h2 = (_source6).cf3;
          let _319___mcc_h3 = (_source6).cf4;
          let _320_cf4 = _319___mcc_h3;
          let _321_cf3 = _318___mcc_h2;
          let _322_cf2 = _317___mcc_h1;
          let _323_cf1 = _316___mcc_h0;
          let _324_v12;
          let _nw52 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _324_v12 = _nw52;
          _324_v12 = _324_v12;
          let _rhs39 = (_this).f6;
          let _rhs40 = _320_cf4;
          let _rhs41 = _312_v10;
          _312_v10 = _rhs39;
          _303_v3 = _rhs40;
          _312_v10 = _rhs41;
          let _325_v13;
          _325_v13 = _dafny.MultiSet.fromElements(_312_v10);
          let _326_v14;
          let _nw53 = new _module.C0();
          _nw53.__ctor(_module.__default.fm6(_this.f10, (((_325_v13).contains(false)) ? ((_325_v13).get(false)) : (p2)), _312_v10, (_306_v6).f13, globalState), (_306_v6).f12);
          _326_v14 = _nw53;
          let _index49 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_324_v12).length));
          (_324_v12)[_index49] = (_this).f5;
          let _index50 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_324_v12).length));
          (_324_v12)[_index50] = ((_this).f11).multipliedBy((_306_v6).f13);
        } else if (_source6.is_DC2) {
          let _327___mcc_h4 = (_source6).cf5;
          let _328___mcc_h5 = (_source6).cf6;
          let _329_cf6 = _328___mcc_h5;
          let _330_cf5 = _327___mcc_h4;
          let _331_v15;
          _331_v15 = _dafny.Set.fromElements(_329_cf6);
          let _332_v16;
          _332_v16 = _dafny.Set.fromElements(_this.f10, p3, (_306_v6).f12, new BigNumber((_331_v15).length), _301_i0);
          let _333_v17;
          let _nw54 = Array((new BigNumber(11)).toNumber());
          _nw54[(_dafny.ZERO).toNumber()] = new BigNumber((_332_v16).length);
          _nw54[(_dafny.ONE).toNumber()] = p3;
          _nw54[(new BigNumber(2)).toNumber()] = new BigNumber((_303_v3).length);
          _nw54[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw54[(new BigNumber(4)).toNumber()] = (_306_v6).f12;
          _nw54[(new BigNumber(5)).toNumber()] = (_306_v6).f13;
          _nw54[(new BigNumber(6)).toNumber()] = (_this.f10).multipliedBy((_306_v6).f12);
          _nw54[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("wecoqytg")).length);
          _nw54[(new BigNumber(8)).toNumber()] = (_this).f11;
          _nw54[(new BigNumber(9)).toNumber()] = (_306_v6).f12;
          _nw54[(new BigNumber(10)).toNumber()] = _this.f10;
          _333_v17 = _nw54;
          _333_v17 = _333_v17;
          let _334_v18;
          let _init5 = ((_335_v10) => function (_336_i2) {
            return _335_v10;
          })(_312_v10);
          let _nw55 = Array((new BigNumber(21)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw55.length); _i0_5++) {
            _nw55[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _334_v18 = _nw55;
          let _337_v19;
          let _nw56 = Array((new BigNumber(29)).toNumber());
          _nw56[(_dafny.ZERO).toNumber()] = _334_v18;
          _nw56[(_dafny.ONE).toNumber()] = _334_v18;
          _nw56[(new BigNumber(2)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(3)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(4)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(5)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(6)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(7)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(8)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(9)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(10)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(11)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(12)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(13)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(14)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(15)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(16)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(17)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(18)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(19)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(20)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(21)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(22)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(23)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(24)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(25)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(26)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(27)).toNumber()] = _334_v18;
          _nw56[(new BigNumber(28)).toNumber()] = _334_v18;
          _337_v19 = _nw56;
          let _index51 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_337_v19).length));
          (_337_v19)[_index51] = _334_v18;
          let _index52 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_337_v19).length));
          (_337_v19)[_index52] = _334_v18;
          let _338_v20;
          _338_v20 = _dafny.Seq.of((_337_v19)[_module.__default.safeIndex(new BigNumber(898), new BigNumber((_337_v19).length))]);
          let _339_v21;
          _339_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_306_v6).f12)).length),_dafny.Seq.update(_338_v20, _module.__default.safeIndex((_this).f11, new BigNumber((_338_v20).length)), _334_v18));
          let _340_v22;
          let _341_v23;
          let _342_v24;
          let _343_v25;
          let _out28;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector7 = _module.__default.m0((_this.f10).minus(new BigNumber(((((_339_v21).contains((_306_v6).f13)) ? ((_339_v21).get((_306_v6).f13)) : (_338_v20))).length)), new BigNumber(425), globalState);
          _out28 = _outcollector7[0];
          _out29 = _outcollector7[1];
          _out30 = _outcollector7[2];
          _out31 = _outcollector7[3];
          _340_v22 = _out28;
          _341_v23 = _out29;
          _342_v24 = _out30;
          _343_v25 = _out31;
          _312_v10 = (_this).fm2(_dafny.Set.fromElements((_this).f6, _312_v10, _329_cf6, _329_cf6, _312_v10), (_dafny.ZERO).minus(_342_v24), new BigNumber((_340_v22).length), globalState);
        } else {
          let _344___mcc_h6 = (_source6).cf0;
          let _345_cf0 = _344___mcc_h6;
          let _346_v26;
          _346_v26 = _dafny.Map.Empty.slice().updateUnsafe(_306_v6,_312_v10);
          let _347_v27;
          _347_v27 = _dafny.Set.fromElements(_346_v26, _346_v26);
          let _348_v28;
          _348_v28 = new _dafny.CodePoint('n'.codePointAt(0));
          let _349_v29;
          _349_v29 = _dafny.Seq.UnicodeFromString("xcfy");
          let _350_v30;
          _350_v30 = _module.D1.create_DC6((_this).f6, new BigNumber((_347_v27).length), (_306_v6).f12, !_dafny.Seq.contains(_dafny.Seq.update(_349_v29, _module.__default.safeIndex((_306_v6).f13, new BigNumber((_349_v29).length)), _348_v28), _348_v28));
          _350_v30 = _350_v30;
          _312_v10 = _312_v10;
          let _351_v31;
          _351_v31 = _dafny.MultiSet.fromElements(false);
          let _352_v32;
          _352_v32 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6) || (_312_v10),(_351_v31).IsProperSubsetOf(_351_v31));
          _352_v32 = (_352_v32).update(_312_v10, _312_v10);
          (_this).f10 = p2;
        }
      }
      let _353_v33;
      _353_v33 = _dafny.Seq.of(p2, _this.f10, _this.f10);
      let _354_v34;
      _354_v34 = _dafny.Set.fromElements((_this).f11);
      let _355_v35;
      _355_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      let _356_v36;
      let _init6 = ((_357_p3) => function (_358_i3) {
        return _module.__default.safeModuloInt(_358_i3, _357_p3);
      })(p3);
      let _nw57 = Array((new BigNumber(5)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw57.length); _i0_6++) {
        _nw57[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _356_v36 = _nw57;
      let _359_v37;
      _359_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,_356_v36);
      let _360_v38;
      _360_v38 = _module.D3.create_DC11(_356_v36);
      _353_v33 = _module.__default.fm12((new BigNumber((_354_v34).length)).multipliedBy(_this.f10), !((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_355_v35).length),(((_359_v37).contains(_this.f10)) ? ((_359_v37).get(_this.f10)) : ((_360_v38).dtor_cf27)))).length)).isEqualTo(new BigNumber(431))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-214)), function (_361_i4) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }), globalState);
      let _362_v39;
      _362_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_dafny.Seq.UnicodeFromString("vbenawu")).length));
      let _363_v40;
      _363_v40 = _dafny.MultiSet.fromElements(p2, (_this).f5);
      let _364_v41;
      _364_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_362_v39).length),(_363_v40).Intersect(_363_v40));
      let _365_v42;
      _365_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_353_v33);
      let _366_v43;
      _366_v43 = _module.D0.create_DC2(_365_v42, (_this).f6);
      let _367_v44;
      _367_v44 = _dafny.Set.fromElements(!(true));
      _364_v41 = (_364_v41).update((((_366_v43).dtor_cf6) ? ((_this).f11) : (_module.__default.fm6(new BigNumber((_367_v44).length), p3, false, new BigNumber(393), globalState))), (_363_v40).update((_this).f5, _module.__default.abs(new BigNumber(515))));
      let _368_v45;
      _368_v45 = new _dafny.CodePoint('i'.codePointAt(0));
      let _369_v46;
      _369_v46 = _dafny.Map.Empty.slice().updateUnsafe(_368_v45,p3);
      let _370_v47;
      _370_v47 = _module.D4.create_DC13(_369_v46);
      let _pat_let_tv7 = _369_v46;
      let _371_v48;
      let _nw58 = new _module.C0();
      _nw58.__ctor((_this).f11, ((!(false)) ? (new BigNumber(((function (_pat_let9_0) {
        return function (_372_dt__update__tmp_h1) {
          return function (_pat_let10_0) {
            return function (_373_dt__update_hcf30_h0) {
              return _module.D4.create_DC13(_373_dt__update_hcf30_h0);
            }(_pat_let10_0);
          }(_pat_let_tv7);
        }(_pat_let9_0);
      }(_370_v47)).dtor_cf30).length)) : (p2)));
      _371_v48 = _nw58;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_356_v36).length))) {
        let _374_i5 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_374_i5)) && ((_374_i5).isLessThan(new BigNumber((_356_v36).length))))) {
          (_356_v36)[(_374_i5)] = (_374_i5).plus(_module.__default.safeModuloInt((_this).f11, (_371_v48).f12));
        }
      }
      return;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f9) {
      let _this = this;
      (_this).f9 = f9;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true))).length))).isLessThan(_this.f9);
    };
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      (_this).f9 = _this.f9;
      let _375_v0;
      _375_v0 = false;
      let _376_i0;
      _376_i0 = _dafny.ZERO;
      L1: {
        while ((_375_v0) || (_375_v0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_376_i0)) {
              break L1;
            }
            _376_i0 = (_376_i0).plus(_dafny.ONE);
            let _377_v1;
            let _nw59 = new _module.C1();
            _nw59.__ctor(_this.f9, new BigNumber((_dafny.Seq.of(_375_v0, _375_v0, _375_v0, _375_v0)).length), _this.f9, _375_v0);
            _377_v1 = _nw59;
            let _378_v2;
            _378_v2 = _dafny.Set.fromElements(((_375_v0) ? (_377_v1) : (_377_v1)));
            let _379_v3;
            _379_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(_377_v1));
            let _380_v4;
            _380_v4 = _dafny.Seq.of(_378_v2);
            _378_v2 = ((((_379_v3).contains(p0)) ? ((_379_v3).get(p0)) : ((_380_v4)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_380_v4).length))]))).Difference((_378_v2).Difference(_378_v2));
            let _381_v5;
            _381_v5 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_377_v1).f5, new BigNumber(613)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), function (_382_i1) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            }));
            _381_v5 = (_381_v5).update((_377_v1).f5, (((_377_v1).f6) ? (_dafny.Seq.UnicodeFromString("byj")) : (p0)));
            let _383_v6;
            let _nw60 = Array((new BigNumber(2)).toNumber()).fill(false);
            _383_v6 = _nw60;
            let _index53 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length));
            (_383_v6)[_index53] = _375_v0;
            let _index54 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length));
            (_383_v6)[_index54] = !(!(((_377_v1).f5).isEqualTo(_module.__default.fm6((_377_v1).f5, _this.f9, _375_v0, (_377_v1).f5, globalState))));
            let _384_v7;
            _384_v7 = _dafny.Map.Empty.slice().updateUnsafe(_375_v0,_this.f9);
            let _385_v8;
            _385_v8 = _dafny.Map.Empty.slice().updateUnsafe((_383_v6)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length))],_384_v7);
            let _386_v9;
            _386_v9 = _dafny.Seq.of((_377_v1).f5);
            let _387_v10;
            _387_v10 = _module.D0.create_DC1(p0, _this.f9, _384_v7, _386_v9);
            let _388_v11;
            _388_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_384_v7);
            let _389_v12;
            _389_v12 = _dafny.MultiSet.fromElements((_377_v1).f5);
            let _390_v13;
            _390_v13 = _module.D1.create_DC6(_375_v0, new BigNumber((_389_v12).cardinality()), (_386_v9)[_module.__default.safeIndex((_377_v1).f5, new BigNumber((_386_v9).length))], (_377_v1).f6);
            let _391_v14;
            _391_v14 = _dafny.Seq.of((_383_v6)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length))], (_383_v6)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length))], (_383_v6)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length))], (_377_v1).f6, (_390_v13).dtor_cf20);
            let _392_v15;
            _392_v15 = _dafny.MultiSet.fromElements(new BigNumber((_391_v14).length), (_377_v1).f5, _this.f9, (_377_v1).f5, _this.f9);
            let _393_v16;
            _393_v16 = _module.D2.create_DC8(_392_v15);
            let _394_v17;
            let _nw61 = Array((new BigNumber(19)).toNumber());
            _nw61[(_dafny.ZERO).toNumber()] = (((_385_v8).contains((_377_v1).f6)) ? ((_385_v8).get((_377_v1).f6)) : (_dafny.Map.Empty.slice().updateUnsafe((_383_v6)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_383_v6).length))],(_377_v1).f5)));
            _nw61[(_dafny.ONE).toNumber()] = _384_v7;
            _nw61[(new BigNumber(2)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(3)).toNumber()] = ((_387_v10).dtor_cf3).Merge(_384_v7);
            _nw61[(new BigNumber(4)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(5)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(6)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(7)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(8)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(9)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(10)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(11)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(12)).toNumber()] = (_384_v7).Merge((((_388_v11).contains((_377_v1).f5)) ? ((_388_v11).get((_377_v1).f5)) : (_384_v7)));
            _nw61[(new BigNumber(13)).toNumber()] = (_384_v7).Merge(_384_v7);
            _nw61[(new BigNumber(14)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(15)).toNumber()] = (_384_v7).update((_377_v1).f6, (_377_v1).f5);
            _nw61[(new BigNumber(16)).toNumber()] = _384_v7;
            _nw61[(new BigNumber(17)).toNumber()] = _module.__default.fm13(p0, (((_392_v15).contains((_dafny.ZERO).minus((_377_v1).f5))) ? ((_392_v15).get((_dafny.ZERO).minus((_377_v1).f5))) : (new BigNumber(144))), (_393_v16).dtor_cf22, globalState);
            _nw61[(new BigNumber(18)).toNumber()] = _384_v7;
            _394_v17 = _nw61;
            _394_v17 = _394_v17;
          }
        }
      }
      let _395_v18;
      _395_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_this.f9);
      let _396_v19;
      _396_v19 = _dafny.Seq.of(new BigNumber(-537), _module.__default.fm6(new BigNumber((_395_v18).length), new BigNumber(210), _375_v0, _this.f9, globalState));
      let _397_i2;
      _397_i2 = _dafny.ZERO;
      L2: {
        while (_dafny.areEqual(_396_v19, ((_375_v0) ? (_396_v19) : (_dafny.Seq.of(_this.f9))))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_397_i2)) {
              break L2;
            }
            _397_i2 = (_397_i2).plus(_dafny.ONE);
            let _398_v20;
            _398_v20 = _dafny.Seq.of(_375_v0);
            let _source7 = _module.__default.fm14(!_dafny.Seq.contains(_398_v20, _375_v0), globalState);
            if (_source7.is_DC4) {
              let _399___mcc_h0 = (_source7).cf8;
              let _400___mcc_h1 = (_source7).cf9;
              let _401___mcc_h2 = (_source7).cf10;
              let _402___mcc_h3 = (_source7).cf11;
              let _403___mcc_h4 = (_source7).cf12;
              let _404_cf12 = _403___mcc_h4;
              let _405_cf11 = _402___mcc_h3;
              let _406_cf10 = _401___mcc_h2;
              let _407_cf9 = _400___mcc_h1;
              let _408_cf8 = _399___mcc_h0;
              let _rhs42 = true;
              let _rhs43 = _407_cf9;
              _407_cf9 = _rhs42;
              _407_cf9 = _rhs43;
              let _409_v21;
              _409_v21 = _dafny.MultiSet.fromElements(_375_v0, _407_cf9, false, !(!(_407_cf9) || (_408_cf8)));
              _409_v21 = _409_v21;
              _407_cf9 = _408_cf8;
              let _410_v22;
              _410_v22 = _dafny.Map.Empty.slice().updateUnsafe(_408_cf8,_module.__default.fm12(_405_cf11, _375_v0, _406_cf10, globalState));
              _407_cf9 = !(_410_v22).equals((_module.__default.fm15(_405_cf11, globalState)).Merge(_410_v22));
            } else if (_source7.is_DC5) {
              let _411___mcc_h5 = (_source7).cf13;
              let _412___mcc_h6 = (_source7).cf14;
              let _413___mcc_h7 = (_source7).cf15;
              let _414___mcc_h8 = (_source7).cf16;
              let _415_cf16 = _414___mcc_h8;
              let _416_cf15 = _413___mcc_h7;
              let _417_cf14 = _412___mcc_h6;
              let _418_cf13 = _411___mcc_h5;
              _418_cf13 = false;
              _398_v20 = _module.__default.fm16(globalState);
              let _419_v23;
              _419_v23 = _dafny.MultiSet.fromElements(new BigNumber(575));
              let _420_v24;
              _420_v24 = _dafny.MultiSet.fromElements(_375_v0, false, _418_cf13, _375_v0);
              let _421_v25;
              _421_v25 = _module.D2.create_DC8((_419_v23).Union(_dafny.MultiSet.fromElements(_this.f9, (((_420_v24).contains(_375_v0)) ? ((_420_v24).get(_375_v0)) : (_415_cf16)), _this.f9)));
              let _pat_let_tv8 = _419_v23;
              let _pat_let_tv9 = _415_cf16;
              let _pat_let_tv10 = _416_cf15;
              let _pat_let_tv11 = _415_cf16;
              _421_v25 = function (_pat_let11_0) {
                return function (_424_dt__update__tmp_h1) {
                  return function (_pat_let14_0) {
                    return function (_425_dt__update_hcf22_h1) {
                      return _module.D2.create_DC8(_425_dt__update_hcf22_h1);
                    }(_pat_let14_0);
                  }(_dafny.MultiSet.fromElements(_pat_let_tv9, _pat_let_tv10, _pat_let_tv11));
                }(_pat_let11_0);
              }(function (_pat_let12_0) {
                return function (_422_dt__update__tmp_h0) {
                  return function (_pat_let13_0) {
                    return function (_423_dt__update_hcf22_h0) {
                      return _module.D2.create_DC8(_423_dt__update_hcf22_h0);
                    }(_pat_let13_0);
                  }(_pat_let_tv8);
                }(_pat_let12_0);
              }(_421_v25));
              let _426_v26;
              _426_v26 = new _dafny.CodePoint('c'.codePointAt(0));
              _426_v26 = _426_v26;
            } else if (_source7.is_DC6) {
              let _427___mcc_h9 = (_source7).cf17;
              let _428___mcc_h10 = (_source7).cf18;
              let _429___mcc_h11 = (_source7).cf19;
              let _430___mcc_h12 = (_source7).cf20;
              let _431_cf20 = _430___mcc_h12;
              let _432_cf19 = _429___mcc_h11;
              let _433_cf18 = _428___mcc_h10;
              let _434_cf17 = _427___mcc_h9;
              _375_v0 = _431_cf20;
              let _435_v27;
              let _436_v28;
              let _437_v29;
              let _438_v30;
              let _out32;
              let _out33;
              let _out34;
              let _out35;
              let _outcollector8 = _module.__default.m0(_433_cf18, _module.__default.fm6(_this.f9, _this.f9, _431_cf20, _this.f9, globalState), globalState);
              _out32 = _outcollector8[0];
              _out33 = _outcollector8[1];
              _out34 = _outcollector8[2];
              _out35 = _outcollector8[3];
              _435_v27 = _out32;
              _436_v28 = _out33;
              _437_v29 = _out34;
              _438_v30 = _out35;
              let _439_v31;
              _439_v31 = _dafny.Map.Empty.slice().updateUnsafe(_434_cf17,_375_v0);
              let _440_v32;
              _440_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(647),_439_v31);
              let _441_v33;
              let _nw62 = new _module.C1();
              _nw62.__ctor((_438_v30).multipliedBy(new BigNumber(((((_440_v32).contains(new BigNumber(984))) ? ((_440_v32).get(new BigNumber(984))) : (_439_v31))).length)), new BigNumber(365), _437_v29, _431_cf20);
              _441_v33 = _nw62;
              _441_v33 = _441_v33;
              let _442_v34;
              _442_v34 = _module.D2.create_DC9(_375_v0, (_441_v33).f11, _432_cf19);
              let _443_v35;
              _443_v35 = _dafny.Set.fromElements(_375_v0);
              let _444_v36;
              _444_v36 = _module.D1.create_DC6(_434_cf17, new BigNumber((_443_v35).length), new BigNumber((_435_v27).length), _375_v0);
              let _445_v37;
              _445_v37 = _dafny.Map.Empty.slice().updateUnsafe((_444_v36).dtor_cf17,_396_v19);
              let _446_v38;
              _446_v38 = _module.D0.create_DC2(_445_v37, _431_cf20);
              let _447_v39;
              _447_v39 = _dafny.MultiSet.fromElements(_431_cf20);
              let _448_v40;
              _448_v40 = new _dafny.CodePoint('u'.codePointAt(0));
              let _449_v41;
              let _nw63 = Array((new BigNumber(28)).toNumber());
              _nw63[(_dafny.ZERO).toNumber()] = (!(_375_v0)) === (_431_cf20);
              _nw63[(_dafny.ONE).toNumber()] = _375_v0;
              _nw63[(new BigNumber(2)).toNumber()] = (_module.D2.create_DC9((_442_v34).dtor_cf23, new BigNumber(809), new BigNumber((p0).length))).dtor_cf23;
              _nw63[(new BigNumber(3)).toNumber()] = !(_431_cf20);
              _nw63[(new BigNumber(4)).toNumber()] = _375_v0;
              _nw63[(new BigNumber(5)).toNumber()] = _375_v0;
              _nw63[(new BigNumber(6)).toNumber()] = !(_375_v0) || (_434_cf17);
              _nw63[(new BigNumber(7)).toNumber()] = _434_cf17;
              _nw63[(new BigNumber(8)).toNumber()] = !(!(!(_375_v0) || (_431_cf20)));
              _nw63[(new BigNumber(9)).toNumber()] = _431_cf20;
              _nw63[(new BigNumber(10)).toNumber()] = _434_cf17;
              _nw63[(new BigNumber(11)).toNumber()] = !(_439_v31).contains(_375_v0);
              _nw63[(new BigNumber(12)).toNumber()] = _434_cf17;
              _nw63[(new BigNumber(13)).toNumber()] = (_442_v34).dtor_cf23;
              _nw63[(new BigNumber(14)).toNumber()] = _431_cf20;
              _nw63[(new BigNumber(15)).toNumber()] = _375_v0;
              _nw63[(new BigNumber(16)).toNumber()] = _434_cf17;
              _nw63[(new BigNumber(17)).toNumber()] = _375_v0;
              _nw63[(new BigNumber(18)).toNumber()] = !((_446_v38).dtor_cf6) || (_375_v0);
              _nw63[(new BigNumber(19)).toNumber()] = _431_cf20;
              _nw63[(new BigNumber(20)).toNumber()] = (_dafny.MultiSet.fromElements(!(_431_cf20))).IsDisjointFrom(_447_v39);
              _nw63[(new BigNumber(21)).toNumber()] = _431_cf20;
              _nw63[(new BigNumber(22)).toNumber()] = _dafny.areEqual(_396_v19, _396_v19);
              _nw63[(new BigNumber(23)).toNumber()] = (_441_v33).fm2(_443_v35, (_441_v33).f11, _433_cf18, globalState);
              _nw63[(new BigNumber(24)).toNumber()] = _431_cf20;
              _nw63[(new BigNumber(25)).toNumber()] = (_this).fm1(_441_v33.f10, _448_v40, globalState);
              _nw63[(new BigNumber(26)).toNumber()] = (_398_v20)[_module.__default.safeIndex(_438_v30, new BigNumber((_398_v20).length))];
              _nw63[(new BigNumber(27)).toNumber()] = _375_v0;
              _449_v41 = _nw63;
              let _index55 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_449_v41).length));
              (_449_v41)[_index55] = _434_cf17;
              let _450_v42;
              _450_v42 = _dafny.Map.Empty.slice().updateUnsafe(_375_v0,new BigNumber((_435_v27).length));
              let _451_v43;
              _451_v43 = _dafny.Seq.of(_450_v42);
              let _452_v44;
              _452_v44 = _module.D0.create_DC1(p0, _437_v29, (_451_v43)[_module.__default.safeIndex(_441_v33.f10, new BigNumber((_451_v43).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(557)), ((_453_v30) => function (_454_i3) {
  return _453_v30;
})(_438_v30)));
              let _index56 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_449_v41).length));
              (_449_v41)[_index56] = _dafny.Seq.contains((_452_v44).dtor_cf1, _448_v40);
            } else if (_source7.is_DC3) {
              let _455___mcc_h13 = (_source7).cf7;
              let _456_cf7 = _455___mcc_h13;
              let _457_v45;
              let _init7 = function (_458_i4) {
                return ((true) ? (new _dafny.CodePoint('w'.codePointAt(0))) : (new _dafny.CodePoint('l'.codePointAt(0))));
              };
              let _nw64 = Array((new BigNumber(19)).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw64.length); _i0_7++) {
                _nw64[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _457_v45 = _nw64;
              let _459_v46;
              _459_v46 = new _dafny.CodePoint('f'.codePointAt(0));
              let _index57 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_457_v45).length));
              (_457_v45)[_index57] = _459_v46;
              let _index58 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_457_v45).length));
              (_457_v45)[_index58] = _459_v46;
              let _460_v47;
              let _nw65 = Array((new BigNumber(22)).toNumber()).fill(false);
              _460_v47 = _nw65;
              let _index59 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_460_v47).length));
              (_460_v47)[_index59] = !(_375_v0);
              let _index60 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_460_v47).length));
              (_460_v47)[_index60] = _375_v0;
              let _461_v48;
              _461_v48 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f9), _this.f9, _this.f9, _this.f9);
              let _462_v49;
              _462_v49 = _module.D2.create_DC8(_461_v48);
              let _pat_let_tv12 = _461_v48;
              _462_v49 = function (_pat_let15_0) {
                return function (_463_dt__update__tmp_h2) {
                  return function (_pat_let16_0) {
                    return function (_464_dt__update_hcf22_h2) {
                      return _module.D2.create_DC8(_464_dt__update_hcf22_h2);
                    }(_pat_let16_0);
                  }(_pat_let_tv12);
                }(_pat_let15_0);
              }(_462_v49);
              let _465_v50;
              let _nw66 = Array((new BigNumber(5)).toNumber()).fill([]);
              _465_v50 = _nw66;
              let _index61 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_465_v50).length));
              (_465_v50)[_index61] = _460_v47;
              let _index62 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_465_v50).length));
              let _nw67 = Array((new BigNumber(18)).toNumber()).fill(false);
              (_465_v50)[_index62] = _nw67;
            } else {
              let _466___mcc_h14 = (_source7).cf21;
              let _467_cf21 = _466___mcc_h14;
              let _468_v51;
              _468_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(500),_375_v0);
              let _469_v52;
              _469_v52 = _module.D1.create_DC3(_468_v51);
              let _470_v53;
              _470_v53 = _module.D1.create_DC7(_469_v52);
              _470_v53 = _470_v53;
              let _471_v54;
              let _472_v55;
              let _473_v56;
              let _474_v57;
              let _out36;
              let _out37;
              let _out38;
              let _out39;
              let _outcollector9 = _module.__default.m0(_this.f9, _this.f9, globalState);
              _out36 = _outcollector9[0];
              _out37 = _outcollector9[1];
              _out38 = _outcollector9[2];
              _out39 = _outcollector9[3];
              _471_v54 = _out36;
              _472_v55 = _out37;
              _473_v56 = _out38;
              _474_v57 = _out39;
              let _index63 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_472_v55).length));
              (_472_v55)[_index63] = new BigNumber((function () {
                let _coll9 = new _dafny.Map();
                for (const _compr_9 of _dafny.IntegerRange(new BigNumber(809), new BigNumber(186))) {
                  let _475_v58 = _compr_9;
                  if (((new BigNumber(809)).isLessThanOrEqualTo(_475_v58)) && ((_475_v58).isLessThan(new BigNumber(186)))) {
                    _coll9.push([_module.__default.safeModuloInt(_475_v58, _this.f9),_471_v54]);
                  }
                }
                return _coll9;
              }()).length);
              let _index64 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_472_v55).length));
              (_472_v55)[_index64] = _474_v57;
              let _476_v59;
              _476_v59 = _dafny.Set.fromElements(_375_v0, _375_v0);
              let _477_v60;
              _477_v60 = _module.D1.create_DC5(_375_v0, _375_v0, new BigNumber((_476_v59).length), new BigNumber(90));
              let _478_v61;
              _478_v61 = _dafny.Map.Empty.slice().updateUnsafe(_477_v60,_396_v19);
              _478_v61 = (_478_v61).update(_477_v60, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_471_v54).length)), _396_v19));
            }
            let _479_v62;
            let _nw68 = Array((_dafny.ONE).toNumber()).fill(false);
            _479_v62 = _nw68;
            let _index65 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_479_v62).length));
            (_479_v62)[_index65] = (_398_v20)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((p0).length)), new BigNumber((_398_v20).length))];
            let _index66 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_479_v62).length));
            (_479_v62)[_index66] = _375_v0;
            let _index67 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_479_v62).length));
            (_479_v62)[_index67] = ((_479_v62)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_479_v62).length))]) || ((_479_v62)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_479_v62).length))]);
            let _index68 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_479_v62).length));
            (_479_v62)[_index68] = (_this.f9).isEqualTo(_this.f9);
          }
        }
      }
      _375_v0 = (_375_v0) === (_375_v0);
      let _480_v63;
      _480_v63 = _module.D4.create_DC14();
      _480_v63 = _480_v63;
      let _481_v64;
      _481_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_375_v0);
      _481_v64 = (_481_v64).update(_this.f9, _375_v0);
      let _482_v65;
      _482_v65 = new _dafny.CodePoint('h'.codePointAt(0));
      let _483_v66;
      _483_v66 = _module.D0.create_DC0(_dafny.Seq.of(new BigNumber(-920)));
      let _484_v67;
      _484_v67 = _dafny.Map.Empty.slice().updateUnsafe(_482_v65,_483_v66);
      let _485_v68;
      _485_v68 = _dafny.Set.fromElements(_375_v0, _375_v0);
      let _486_v69;
      _486_v69 = _dafny.Map.Empty.slice().updateUnsafe(_484_v67,new BigNumber((_485_v68).length));
      let _487_v71;
      _487_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(_375_v0),_486_v69);
      let _488_v74;
      _488_v74 = _dafny.MultiSet.fromElements(_482_v65);
      let _pat_let_tv13 = _396_v19;
      let _489_v75;
      _489_v75 = _dafny.Set.fromElements(function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_488_v74).Elements) {
          let _490_v73 = _compr_10;
          if ((_488_v74).contains(_490_v73)) {
            _coll10.push([_490_v73,_483_v66]);
          }
        }
        return _coll10;
      }(), _dafny.Map.Empty.slice().updateUnsafe(_482_v65,function (_pat_let17_0) {
        return function (_491_dt__update__tmp_h3) {
          return function (_pat_let18_0) {
            return function (_492_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_492_dt__update_hcf0_h0);
            }(_pat_let18_0);
          }(_pat_let_tv13);
        }(_pat_let17_0);
      }(_483_v66)));
      let _493_v76;
      _493_v76 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_489_v75);
      let _494_v77;
      _494_v77 = _dafny.Seq.of(_486_v69, function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of ((((_487_v71).contains(_375_v0)) ? ((_487_v71).get(_375_v0)) : (function () {
          let _coll12 = new _dafny.Map();
          for (const _compr_12 of ((((_493_v76).contains(_this.f9)) ? ((_493_v76).get(_this.f9)) : (_489_v75))).Elements) {
            let _495_v72 = _compr_12;
            if (((((_493_v76).contains(_this.f9)) ? ((_493_v76).get(_this.f9)) : (_489_v75))).contains(_495_v72)) {
              _coll12.push([_495_v72,_this.f9]);
            }
          }
          return _coll12;
        }()))).Keys.Elements) {
          let _496_v70 = _compr_11;
          if (((((_487_v71).contains(_375_v0)) ? ((_487_v71).get(_375_v0)) : (function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of ((((_493_v76).contains(_this.f9)) ? ((_493_v76).get(_this.f9)) : (_489_v75))).Elements) {
              let _495_v72 = _compr_13;
              if (((((_493_v76).contains(_this.f9)) ? ((_493_v76).get(_this.f9)) : (_489_v75))).contains(_495_v72)) {
                _coll13.push([_495_v72,_this.f9]);
              }
            }
            return _coll13;
          }()))).contains(_496_v70)) {
            _coll11.push([_496_v70,_this.f9]);
          }
        }
        return _coll11;
      }(), _486_v69);
      r0 = (_494_v77)[_module.__default.safeIndex(_this.f9, new BigNumber((_494_v77).length))];
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f8, f5, f6) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return !(((_this).f5).isLessThanOrEqualTo((_this).f8));
    };
    m1(p0, globalState) {
      let _this = this;
      let _497_i0;
      _497_i0 = _dafny.ZERO;
      L3: {
        while ((_this).f6) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_497_i0)) {
              break L3;
            }
            _497_i0 = (_497_i0).plus(_dafny.ONE);
            let _498_v0;
            _498_v0 = false;
            _498_v0 = !(_498_v0);
            let _499_v1;
            _499_v1 = _dafny.MultiSet.fromElements((_this).f6, _498_v0);
            let _500_v2;
            _500_v2 = _module.D5.create_DC15(_499_v1);
            let _501_v3;
            _501_v3 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f6, (_this).f6)).cardinality()), (_this).f8, new BigNumber(((_500_v2).dtor_cf31).cardinality()), p0);
            let _502_v4;
            let _nw69 = new _module.C2();
            _nw69.__ctor(((_501_v3)[_module.__default.safeIndex((_this).f5, new BigNumber((_501_v3).length))]).plus(p0));
            _502_v4 = _nw69;
            let _503_v5;
            _503_v5 = _dafny.Set.fromElements(p0, (_this).f5, p0, (_this).f5);
            let _504_v6;
            _504_v6 = _dafny.Map.Empty.slice().updateUnsafe(_502_v4.f9,new BigNumber(408));
            let _505_v7;
            _505_v7 = _dafny.Map.Empty.slice().updateUnsafe(_498_v0,_498_v0);
            let _506_v8;
            let _nw70 = Array((new BigNumber(15)).toNumber());
            _nw70[(_dafny.ZERO).toNumber()] = _module.__default.fm6(new BigNumber(56), _502_v4.f9, _498_v0, p0, globalState);
            _nw70[(_dafny.ONE).toNumber()] = (_this).f8;
            _nw70[(new BigNumber(2)).toNumber()] = _502_v4.f9;
            _nw70[(new BigNumber(3)).toNumber()] = _502_v4.f9;
            _nw70[(new BigNumber(4)).toNumber()] = _502_v4.f9;
            _nw70[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt((_this).f8, (_this).f5);
            _nw70[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_502_v4.f9);
            _nw70[(new BigNumber(7)).toNumber()] = (_502_v4.f9).multipliedBy(new BigNumber((_503_v5).length));
            _nw70[(new BigNumber(8)).toNumber()] = _module.__default.fm6(_502_v4.f9, (_this).f8, _498_v0, (_this).f5, globalState);
            _nw70[(new BigNumber(9)).toNumber()] = (_this).f8;
            _nw70[(new BigNumber(10)).toNumber()] = (((_504_v6).contains((_this).f5)) ? ((_504_v6).get((_this).f5)) : ((_501_v3)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_501_v3).length))]));
            _nw70[(new BigNumber(11)).toNumber()] = _502_v4.f9;
            _nw70[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(_502_v4.f9, p0);
            _nw70[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(889)), ((_507_v4) => function (_508_i1) {
              return _507_v4.f9;
            })(_502_v4)), _501_v3)).length);
            _nw70[(new BigNumber(14)).toNumber()] = ((_498_v0) ? ((_dafny.ZERO).minus(new BigNumber((_505_v7).length))) : (p0));
            _506_v8 = _nw70;
            let _index69 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_506_v8).length));
            (_506_v8)[_index69] = _502_v4.f9;
            let _509_v9;
            let _nw71 = Array((new BigNumber(20)).toNumber()).fill([]);
            _509_v9 = _nw71;
            let _510_v10;
            let _nw72 = Array((new BigNumber(29)).toNumber()).fill(false);
            _510_v10 = _nw72;
            let _index70 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length));
            (_509_v9)[_index70] = _510_v10;
            let _index71 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_506_v8).length));
            let _index72 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length));
            let _rhs44 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f5));
            let _rhs45 = _510_v10;
            let _rhs46 = (_this).f5;
            let _lhs30 = _506_v8;
            let _lhs31 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_506_v8).length));
            let _lhs32 = _509_v9;
            let _lhs33 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length));
            let _lhs34 = _502_v4;
            _lhs30[_lhs31] = _rhs44;
            _lhs32[_lhs33] = _rhs45;
            _lhs34.f9 = _rhs46;
            let _arr0 = (_509_v9)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length))];
            let _index73 = _module.__default.safeIndex(new BigNumber(338), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length))]).length));
            _arr0[_index73] = (_this).f6;
            let _511_v11;
            let _nw73 = new _module.C1();
            _nw73.__ctor((_506_v8)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_506_v8).length))], (_this).f5, _502_v4.f9, (_this).f6);
            _511_v11 = _nw73;
            let _512_v12;
            _512_v12 = new _dafny.CodePoint('b'.codePointAt(0));
            let _513_v13;
            _513_v13 = _dafny.Seq.of(_512_v12);
            let _514_v14;
            _514_v14 = _module.D5.create_DC16(_511_v11, false, new BigNumber((_513_v13).length));
            let _pat_let_tv14 = _498_v0;
            let _arr1 = (_509_v9)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length))];
            let _index74 = _module.__default.safeIndex(new BigNumber(338), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length))]).length));
            let _rhs47 = (function (_pat_let19_0) {
              return function (_515_dt__update__tmp_h0) {
                return function (_pat_let20_0) {
                  return function (_516_dt__update_hcf33_h0) {
                    return _module.D5.create_DC16((_515_dt__update__tmp_h0).dtor_cf32, _516_dt__update_hcf33_h0, (_515_dt__update__tmp_h0).dtor_cf34);
                  }(_pat_let20_0);
                }(_pat_let_tv14);
              }(_pat_let19_0);
            }(_514_v14)).dtor_cf33;
            let _rhs48 = (_this).f6;
            let _lhs35 = (_509_v9)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length))];
            let _lhs36 = _module.__default.safeIndex(new BigNumber(338), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_509_v9).length))]).length));
            _498_v0 = _rhs47;
            _lhs35[_lhs36] = _rhs48;
          }
        }
      }
      if ((p0).isLessThanOrEqualTo(p0)) {
        let _517_v15;
        _517_v15 = new _dafny.CodePoint('o'.codePointAt(0));
        let _518_v16;
        _518_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,_517_v15);
        let _519_v17;
        _519_v17 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6, (_this).f6, (_this).fm0((_this).f6, globalState), (_this).f6);
        let _520_v18;
        let _nw74 = new _module.C0();
        _nw74.__ctor(new BigNumber((_519_v17).cardinality()), (_this).f8);
        _520_v18 = _nw74;
        let _521_v19;
        _521_v19 = _dafny.Map.Empty.slice().updateUnsafe(_520_v18,_517_v15);
        let _522_v20;
        _522_v20 = _dafny.MultiSet.fromElements(_518_v16, (_518_v16).update(new BigNumber((_521_v19).length), _module.__default.fm11(_517_v15, (_this).f6, (_this).f6, globalState)));
        let _523_v21;
        _523_v21 = _dafny.Seq.of(_522_v20);
        let _524_v22;
        _524_v22 = _dafny.MultiSet.fromElements((_520_v18).f13);
        let _525_v23;
        _525_v23 = _dafny.MultiSet.fromElements(new BigNumber((_524_v22).cardinality()), _module.__default.fm6((_520_v18).f12, (_this).f5, (_this).f6, (_520_v18).f12, globalState), _module.__default.fm6((_520_v18).f13, (_this).f8, (_this).f6, (_520_v18).f13, globalState), (_this).f5, (_this).f8);
        _522_v20 = (((_this).f6) ? ((_523_v21)[_module.__default.safeIndex(new BigNumber((_525_v23).cardinality()), new BigNumber((_523_v21).length))]) : (_522_v20));
        let _526_v24;
        _526_v24 = _dafny.Seq.of(_517_v15);
        let _527_v25;
        _527_v25 = _dafny.Seq.of(_526_v24);
        let _528_v26;
        _528_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,!((_this).f6));
        let _529_v27;
        _529_v27 = _dafny.Seq.of(_module.__default.fm17((((_528_v26).contains(p0)) ? ((_528_v26).get(p0)) : ((_this).f6)), globalState));
        _527_v25 = (_529_v27)[_module.__default.safeIndex(new BigNumber(((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(119), new BigNumber(115))) {
            let _530_v28 = _compr_14;
            if (((new BigNumber(119)).isLessThanOrEqualTo(_530_v28)) && ((_530_v28).isLessThan(new BigNumber(115)))) {
              _coll14.push([_module.__default.safeDivisionInt(_530_v28, (_520_v18).f13),!((_this).f6)]);
            }
          }
          return _coll14;
        }()).Merge(function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(223), new BigNumber(441))) {
            let _531_v29 = _compr_15;
            if (((new BigNumber(223)).isLessThanOrEqualTo(_531_v29)) && ((_531_v29).isLessThan(new BigNumber(441)))) {
              _coll15.push([(_531_v29).plus((_520_v18).f13),true]);
            }
          }
          return _coll15;
        }())).length), new BigNumber((_529_v27).length))];
        let _532_v30;
        let _nw75 = Array((new BigNumber(29)).toNumber());
        _532_v30 = _nw75;
        let _533_v31;
        _533_v31 = _dafny.Set.fromElements(_517_v15, _517_v15, _517_v15, _517_v15);
        let _534_v32;
        let _nw76 = new _module.C2();
        _nw76.__ctor(new BigNumber((_533_v31).length));
        _534_v32 = _nw76;
        let _index75 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_532_v30).length));
        (_532_v30)[_index75] = _534_v32;
        let _index76 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_532_v30).length));
        (_532_v30)[_index76] = _534_v32;
        let _535_v33;
        _535_v33 = true;
        _535_v33 = (_this).f6;
        _517_v15 = new _dafny.CodePoint('r'.codePointAt(0));
      } else {
        let _536_v34;
        _536_v34 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),(_this).f6);
        let _537_v35;
        _537_v35 = _dafny.Seq.of(_536_v34, _536_v34);
        let _538_v36;
        _538_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_dafny.Seq.of((_this).f6, (_this).f6)).length));
        let _539_v37;
        let _nw77 = new _module.C1();
        _nw77.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_540_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        }), _537_v35)).length), (_this).f5, new BigNumber((_538_v36).length), (((_this).f6) ? ((_this).f6) : ((_this).f6)));
        _539_v37 = _nw77;
        let _541_v38;
        let _nw78 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
        _541_v38 = _nw78;
        let _542_v39;
        _542_v39 = _dafny.Seq.of(true);
        let _543_v40;
        _543_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_542_v39);
        let _544_v41;
        _544_v41 = _module.D6.create_DC17(_542_v39);
        let _nw79 = Array((new BigNumber(28)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_542_v39, _542_v39), _module.__default.safeIndex(new BigNumber(((((_543_v40).contains(_539_v37.f10)) ? ((_543_v40).get(_539_v37.f10)) : (_542_v39))).length), new BigNumber((_dafny.Seq.Concat(_542_v39, _542_v39)).length)), (_this).f6);
        _nw79[(_dafny.ONE).toNumber()] = _542_v39;
        _nw79[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_542_v39, _542_v39);
        _nw79[(new BigNumber(3)).toNumber()] = (_module.D6.create_DC17(_542_v39)).dtor_cf35;
        _nw79[(new BigNumber(4)).toNumber()] = _dafny.Seq.of((_this).f6, (_this).f6);
        _nw79[(new BigNumber(5)).toNumber()] = _module.__default.fm16(globalState);
        _nw79[(new BigNumber(6)).toNumber()] = _module.__default.fm16(globalState);
        _nw79[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_542_v39, _542_v39);
        _nw79[(new BigNumber(8)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(9)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_542_v39, _542_v39);
        _nw79[(new BigNumber(11)).toNumber()] = (_module.__default.fm18(globalState)).dtor_cf35;
        _nw79[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(true);
        _nw79[(new BigNumber(13)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(14)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_542_v39, (_544_v41).dtor_cf35);
        _nw79[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_542_v39, _542_v39);
        _nw79[(new BigNumber(17)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(18)).toNumber()] = _dafny.Seq.of((_this).f6);
        _nw79[(new BigNumber(19)).toNumber()] = _dafny.Seq.of((_this).f6);
        _nw79[(new BigNumber(20)).toNumber()] = _dafny.Seq.of((_this).f6, (_this).f6);
        _nw79[(new BigNumber(21)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(22)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_542_v39, _542_v39), _module.__default.safeIndex(_539_v37.f10, new BigNumber((_dafny.Seq.Concat(_542_v39, _542_v39)).length)), false);
        _nw79[(new BigNumber(23)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_542_v39, _542_v39);
        _nw79[(new BigNumber(25)).toNumber()] = _542_v39;
        _nw79[(new BigNumber(26)).toNumber()] = _dafny.Seq.of(true);
        _nw79[(new BigNumber(27)).toNumber()] = _542_v39;
        _541_v38 = _nw79;
        let _545_v42;
        let _init8 = function (_546_i3) {
          return (_this).f6;
        };
        let _nw80 = Array((new BigNumber(18)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw80.length); _i0_8++) {
          _nw80[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _545_v42 = _nw80;
        let _547_v43;
        _547_v43 = _dafny.Seq.of((_this).f5);
        let _548_v44;
        _548_v44 = _dafny.Seq.of(new BigNumber((_547_v43).length));
        let _index77 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_545_v42).length));
        (_545_v42)[_index77] = _dafny.Seq.IsPrefixOf(_548_v44, _547_v43);
        let _549_v45;
        let _nw81 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _549_v45 = _nw81;
        let _550_v46;
        _550_v46 = _module.D3.create_DC11(_549_v45);
        let _index78 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_545_v42).length));
        let _rhs49 = false;
        let _rhs50 = _550_v46;
        let _lhs37 = _545_v42;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_545_v42).length));
        _lhs37[_lhs38] = _rhs49;
        _550_v46 = _rhs50;
        let _551_v47;
        _551_v47 = new _dafny.CodePoint('n'.codePointAt(0));
        _551_v47 = _551_v47;
        let _552_v48;
        _552_v48 = _dafny.MultiSet.fromElements(new BigNumber(622));
        let _553_v49;
        _553_v49 = _dafny.Seq.of(_552_v48);
        let _554_v50;
        _554_v50 = _dafny.Seq.UnicodeFromString("fmwqpjm");
        let _555_v51;
        _555_v51 = _dafny.Seq.of(_554_v50);
        let _556_v52;
        _556_v52 = _dafny.Map.Empty.slice().updateUnsafe((_553_v49)[_module.__default.safeIndex((_this).f5, new BigNumber((_553_v49).length))],(_555_v51)[_module.__default.safeIndex((_539_v37).f11, new BigNumber((_555_v51).length))]);
        let _index79 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_545_v42).length));
        (_545_v42)[_index79] = (((_539_v37).fm0((_this).f6, globalState)) ? (!((new BigNumber((_556_v52).length)).isLessThan(_539_v37.f10))) : ((_552_v48).IsProperSubsetOf(_552_v48)));
      }
      let _557_v53;
      _557_v53 = _dafny.Seq.UnicodeFromString("whxnlgjoh");
      let _hi4 = (_dafny.ZERO).minus(new BigNumber((_557_v53).length));
      for (let _558_i4 = ((_this).f5).minus(_module.__default.fm6(p0, (_this).f5, (_this).f6, p0, globalState)); _558_i4.isLessThan(_hi4); _558_i4 = _558_i4.plus(_dafny.ONE)) {
        let _559_v54;
        _559_v54 = new BigNumber(352);
        let _560_v55;
        _560_v55 = _dafny.Seq.of(false);
        let _561_v56;
        _561_v56 = _dafny.Set.fromElements((_this).f6);
        let _562_v57;
        _562_v57 = _dafny.Seq.of(new BigNumber((_560_v55).length), new BigNumber((_557_v53).length), new BigNumber((_561_v56).length));
        _559_v54 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_562_v57, _562_v57)).length), _559_v54);
        let _563_v58;
        _563_v58 = true;
        let _564_v59;
        _564_v59 = _dafny.MultiSet.fromElements(new BigNumber((_557_v53).length));
        let _565_v60;
        _565_v60 = _dafny.Map.Empty.slice().updateUnsafe(_564_v59,(_this).f6);
        _563_v58 = (((_this).f6) ? (!((new BigNumber(924)).isLessThanOrEqualTo(new BigNumber((_565_v60).length)))) : ((_563_v58) === ((_this).f6)));
        let _566_v61;
        _566_v61 = new _dafny.CodePoint('t'.codePointAt(0));
        let _567_v62;
        let _nw82 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _567_v62 = _nw82;
        let _568_v63;
        _568_v63 = _module.D6.create_DC18((_this).f6);
        let _rhs51 = (_this).fm0(((_568_v63).dtor_cf36) === ((_this).f6), globalState);
        let _rhs52 = p0;
        let _rhs53 = _566_v61;
        let _rhs54 = _567_v62;
        let _rhs55 = p0;
        _563_v58 = _rhs51;
        _559_v54 = _rhs52;
        _566_v61 = _rhs53;
        _567_v62 = _rhs54;
        _559_v54 = _rhs55;
        _563_v58 = (_this).f6;
      }
      let _569_v64;
      _569_v64 = new _dafny.CodePoint('o'.codePointAt(0));
      let _570_v65;
      let _nw83 = Array((new BigNumber(4)).toNumber());
      _nw83[(_dafny.ZERO).toNumber()] = _569_v64;
      _nw83[(_dafny.ONE).toNumber()] = _569_v64;
      _nw83[(new BigNumber(2)).toNumber()] = _569_v64;
      _nw83[(new BigNumber(3)).toNumber()] = _569_v64;
      _570_v65 = _nw83;
      let _571_v66;
      _571_v66 = _dafny.Seq.of(_570_v65);
      let _572_v67;
      _572_v67 = _dafny.MultiSet.fromElements((_this).f6, true);
      let _573_v68;
      _573_v68 = _module.D5.create_DC15(_572_v67);
      let _574_v69;
      let _init9 = function (_575_i5) {
        return _dafny.Seq.Concat(_dafny.Seq.of((_this).f6, (_this).f6), _dafny.Seq.of((_this).f6, (_this).f6));
      };
      let _nw84 = Array((new BigNumber(7)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw84.length); _i0_9++) {
        _nw84[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _574_v69 = _nw84;
      let _576_v70;
      _576_v70 = _dafny.Seq.of((_this).f6, (_this).f6);
      let _index80 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_574_v69).length));
      (_574_v69)[_index80] = _dafny.Seq.Concat(_576_v70, _576_v70);
      let _index81 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_574_v69).length));
      (_574_v69)[_index81] = _576_v70;
      let _577_v71;
      _577_v71 = false;
      let _578_v72;
      _578_v72 = _module.D1.create_DC7(_module.D1.create_DC4(_577_v71, _577_v71, _dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_579_v64) => function (_580_i6) {
  return _579_v64;
})(_569_v64)), p0, _module.__default.fm6((_this).f8, new BigNumber(-992), _577_v71, new BigNumber(918), globalState)));
      let _581_v73;
      _581_v73 = _dafny.Map.Empty.slice().updateUnsafe(_578_v72,_572_v67);
      let _pat_let_tv15 = _581_v73;
      let _pat_let_tv16 = _577_v71;
      let _pat_let_tv17 = globalState;
      let _pat_let_tv18 = _572_v67;
      let _pat_let_tv19 = _581_v73;
      let _pat_let_tv20 = _577_v71;
      let _pat_let_tv21 = globalState;
      let _index82 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_574_v69).length));
      let _index83 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_574_v69).length));
      let _rhs56 = _571_v66;
      let _rhs57 = function (_pat_let21_0) {
        return function (_582_dt__update__tmp_h1) {
          return function (_pat_let22_0) {
            return function (_583_dt__update_hcf31_h0) {
              return _module.D5.create_DC15(_583_dt__update_hcf31_h0);
            }(_pat_let22_0);
          }((((_pat_let_tv19).contains(_module.D1.create_DC7(_module.__default.fm14(_pat_let_tv20, _pat_let_tv21)))) ? ((_pat_let_tv15).get(_module.D1.create_DC7(_module.__default.fm14(_pat_let_tv16, _pat_let_tv17)))) : (_pat_let_tv18)));
        }(_pat_let21_0);
      }(_module.D5.create_DC15(_572_v67));
      let _rhs58 = _dafny.Seq.Concat(_576_v70, _576_v70);
      let _rhs59 = _module.__default.fm16(globalState);
      let _rhs60 = _577_v71;
      let _lhs39 = _574_v69;
      let _lhs40 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_574_v69).length));
      let _lhs41 = _574_v69;
      let _lhs42 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_574_v69).length));
      _571_v66 = _rhs56;
      _573_v68 = _rhs57;
      _lhs39[_lhs40] = _rhs58;
      _lhs41[_lhs42] = _rhs59;
      _577_v71 = _rhs60;
      let _584_v74;
      let _init10 = ((_585_p0) => function (_586_i8) {
        return (_586_i8).minus(_585_p0);
      })(p0);
      let _nw85 = Array((new BigNumber(25)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw85.length); _i0_10++) {
        _nw85[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _584_v74 = _nw85;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_584_v74).length))) {
        let _587_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_587_i7)) && ((_587_i7).isLessThan(new BigNumber((_584_v74).length))))) {
          (_584_v74)[(_587_i7)] = _module.__default.safeDivisionInt(_587_i7, p0);
        }
      }
      let _588_v75;
      _588_v75 = _dafny.Set.fromElements(_577_v71, (_this).f6);
      let _589_v76;
      _589_v76 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_588_v75).length)));
      let _index84 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_584_v74).length));
      (_584_v74)[_index84] = (_589_v76)[_module.__default.safeIndex(new BigNumber((_557_v53).length), new BigNumber((_589_v76).length))];
      let _index85 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_584_v74).length));
      (_584_v74)[_index85] = ((_this).f5).plus(new BigNumber(((_588_v75).Difference(_588_v75)).length));
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _module.D0.Default();
      let _590_v0;
      _590_v0 = new BigNumber(-393);
      let _591_v1;
      _591_v1 = _dafny.Set.fromElements(p2, new BigNumber(430));
      let _592_v2;
      _592_v2 = _dafny.MultiSet.fromElements(new BigNumber((_591_v1).length), _590_v0);
      let _593_v3;
      let _nw86 = new _module.C1();
      _nw86.__ctor(new BigNumber((_592_v2).cardinality()), (_this).f8, _590_v0, (_this).f6);
      _593_v3 = _nw86;
      let _594_v4;
      _594_v4 = _module.D5.create_DC16(_593_v3, (_this).f6, p0);
      let _595_v5;
      let _nw87 = new _module.C1();
      _nw87.__ctor(new BigNumber(-929), (_594_v4).dtor_cf34, p0, (_593_v3).f6);
      _595_v5 = _nw87;
      let _596_v6;
      _596_v6 = _dafny.Seq.of(new BigNumber((p1).length), new BigNumber(968));
      let _597_v7;
      _597_v7 = _dafny.Set.fromElements((_this).f6, (_this).f6, (_595_v5).f6, (_this).f6);
      let _598_v8;
      _598_v8 = _dafny.Seq.of((_595_v5).f6, true);
      let _599_v9;
      _599_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_597_v7).length),new BigNumber((_598_v8).length));
      let _600_v10;
      _600_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(11),(_593_v3).f6);
      let _601_v11;
      _601_v11 = _dafny.Map.Empty.slice().updateUnsafe((_595_v5).f6,(((_600_v10).contains((_595_v5).f5)) ? ((_600_v10).get((_595_v5).f5)) : ((_595_v5).f6)));
      let _602_v12;
      let _nw88 = Array((new BigNumber(9)).toNumber());
      _nw88[(_dafny.ZERO).toNumber()] = p0;
      _nw88[(_dafny.ONE).toNumber()] = new BigNumber((_596_v6).length);
      _nw88[(new BigNumber(2)).toNumber()] = _590_v0;
      _nw88[(new BigNumber(3)).toNumber()] = new BigNumber(((_599_v9).Merge(_599_v9)).length);
      _nw88[(new BigNumber(4)).toNumber()] = (new BigNumber(665)).plus(new BigNumber(780));
      _nw88[(new BigNumber(5)).toNumber()] = (_595_v5).f5;
      _nw88[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt((_593_v3).f5, (_this).f5);
      _nw88[(new BigNumber(7)).toNumber()] = (_593_v3).f5;
      _nw88[(new BigNumber(8)).toNumber()] = _module.__default.fm6(p0, new BigNumber((_601_v11).length), (_595_v5).f6, _module.__default.fm6(_module.__default.fm6((_593_v3).f5, (_this).f8, (_593_v3).f6, (_595_v5).f5, globalState), _590_v0, p3, p0, globalState), globalState);
      _602_v12 = _nw88;
      let _index86 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
      (_602_v12)[_index86] = _module.__default.safeModuloInt((_this).f8, (_593_v3).f5);
      let _603_v13;
      _603_v13 = _module.D1.create_DC6(((_593_v3).f6) || ((_this).f6), (_dafny.ZERO).minus((_593_v3).f5), new BigNumber((_dafny.Seq.UnicodeFromString("a")).length), (_595_v5).f6);
      let _604_v14;
      _604_v14 = _dafny.MultiSet.fromElements(!((_593_v3).f6));
      let _605_v15;
      _605_v15 = _dafny.Map.Empty.slice().updateUnsafe(!((_593_v3).f6),_595_v5);
      let _606_v16;
      _606_v16 = _dafny.Map.Empty.slice().updateUnsafe(p3,_596_v6);
      let _index87 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
      let _rhs61 = _module.__default.fm6(new BigNumber((_module.__default.fm19(new BigNumber((_604_v14).cardinality()), globalState)).length), (_dafny.ZERO).minus((new BigNumber((p1).length)).multipliedBy(p0)), p3, new BigNumber((_dafny.Seq.UnicodeFromString("doyklvo")).length), globalState);
      let _rhs62 = (((_605_v15).contains((_this).f6)) ? ((_605_v15).get((_this).f6)) : (_593_v3));
      let _rhs63 = _module.__default.fm6((_593_v3).f5, new BigNumber((_dafny.Seq.Concat(_596_v6, (((_606_v16).contains((_593_v3).f6)) ? ((_606_v16).get((_593_v3).f6)) : (_596_v6)))).length), (_593_v3).fm0(p3, globalState), _module.__default.safeModuloInt(_590_v0, p2), globalState);
      let _rhs64 = (((_this).fm0(p3, globalState)) ? (_603_v13) : (_module.__default.fm20((_dafny.ZERO).minus(new BigNumber((p1).length)), _604_v14, (_this).f6, (_this).f6, globalState)));
      let _rhs65 = _593_v3;
      let _lhs43 = _602_v12;
      let _lhs44 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
      _590_v0 = _rhs61;
      _595_v5 = _rhs62;
      _lhs43[_lhs44] = _rhs63;
      _603_v13 = _rhs64;
      _595_v5 = _rhs65;
      if ((((_601_v11).contains((_593_v3).f6)) ? ((_601_v11).get((_593_v3).f6)) : ((_this).fm0((_595_v5).f6, globalState)))) {
        let _607_v17;
        let _nw89 = new _module.C0();
        _nw89.__ctor(p2, new BigNumber((_597_v7).length));
        _607_v17 = _nw89;
        if ((_604_v14).IsSubsetOf(_604_v14)) {
          let _index88 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          (_602_v12)[_index88] = (_607_v17).f12;
          _590_v0 = _module.__default.fm6((_590_v0).multipliedBy((_595_v5).f5), (_dafny.ZERO).minus(p2), (_593_v3).f6, ((_607_v17).f13).multipliedBy(p2), globalState);
          let _608_v18;
          _608_v18 = _dafny.Map.Empty.slice().updateUnsafe((_607_v17).f13,_dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), function (_609_i0) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }));
          let _610_v19;
          let _nw90 = Array((new BigNumber(17)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = (((_608_v18).contains((_607_v17).f12)) ? ((_608_v18).get((_607_v17).f12)) : (p1));
          _nw90[(_dafny.ONE).toNumber()] = p1;
          _nw90[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("xvmci");
          _nw90[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(962)), function (_611_i1) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          });
          _nw90[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(908)), function (_612_i2) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          });
          _nw90[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("ea");
          _nw90[(new BigNumber(6)).toNumber()] = p1;
          _nw90[(new BigNumber(7)).toNumber()] = p1;
          _nw90[(new BigNumber(8)).toNumber()] = p1;
          _nw90[(new BigNumber(9)).toNumber()] = p1;
          _nw90[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(852)), function (_613_i3) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          });
          _nw90[(new BigNumber(11)).toNumber()] = p1;
          _nw90[(new BigNumber(12)).toNumber()] = p1;
          _nw90[(new BigNumber(13)).toNumber()] = p1;
          _nw90[(new BigNumber(14)).toNumber()] = p1;
          _nw90[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), function (_614_i4) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          });
          _nw90[(new BigNumber(16)).toNumber()] = p1;
          _610_v19 = _nw90;
          let _615_v20;
          _615_v20 = _dafny.Map.Empty.slice().updateUnsafe((_595_v5).f5,_610_v19);
          _615_v20 = (_615_v20).update((_dafny.ZERO).minus((_this).f8), _610_v19);
          let _index89 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_610_v19).length));
          (_610_v19)[_index89] = p1;
          let _616_v21;
          let _nw91 = Array((new BigNumber(28)).toNumber());
          _616_v21 = _nw91;
          let _617_v22;
          _617_v22 = new _dafny.CodePoint('g'.codePointAt(0));
          let _618_v23;
          _618_v23 = _617_v22;
          let _619_v24;
          let _nw92 = Array((new BigNumber(22)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = (_593_v3).f6;
          _nw92[(_dafny.ONE).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(2)).toNumber()] = (_this).f6;
          _nw92[(new BigNumber(3)).toNumber()] = !((_593_v3).f6) || (!((((_601_v11).contains((_593_v3).f6)) ? ((_601_v11).get((_593_v3).f6)) : ((_593_v3).f6))));
          _nw92[(new BigNumber(4)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(5)).toNumber()] = !((_dafny.ZERO).minus((_595_v5).f5)).isEqualTo((_607_v17).f12);
          _nw92[(new BigNumber(6)).toNumber()] = (((_593_v3).f6) ? (!((_593_v3).f6)) : ((_593_v3).f6));
          _nw92[(new BigNumber(7)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(8)).toNumber()] = !((_595_v5).f6) || ((_593_v3).f6);
          _nw92[(new BigNumber(9)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(10)).toNumber()] = true;
          _nw92[(new BigNumber(11)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(12)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(13)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(14)).toNumber()] = ((_595_v5).f6) && ((_this).f6);
          _nw92[(new BigNumber(15)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(16)).toNumber()] = (_595_v5).fm0((_595_v5).f6, globalState);
          _nw92[(new BigNumber(17)).toNumber()] = (_593_v3).f6;
          _nw92[(new BigNumber(18)).toNumber()] = (_593_v3).fm0((_595_v5).f6, globalState);
          _nw92[(new BigNumber(19)).toNumber()] = p3;
          _nw92[(new BigNumber(20)).toNumber()] = (_595_v5).f6;
          _nw92[(new BigNumber(21)).toNumber()] = !((((_601_v11).contains((_607_v17).fm8(p3, _dafny.Seq.UnicodeFromString("vjkgdpeu"), (_618_v23), globalState))) ? ((_601_v11).get((_607_v17).fm8(p3, _dafny.Seq.UnicodeFromString("vjkgdpeu"), (_618_v23), globalState))) : ((_593_v3).f6)));
          _619_v24 = _nw92;
          let _620_v25;
          _620_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(129),_619_v24);
          let _index90 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_610_v19).length));
          let _rhs66 = _dafny.Seq.Concat(p1, p1);
          let _rhs67 = new BigNumber(481);
          let _rhs68 = _616_v21;
          let _rhs69 = (((_620_v25).contains(_590_v0)) ? ((_620_v25).get(_590_v0)) : (_619_v24));
          let _lhs45 = _610_v19;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_610_v19).length));
          _lhs45[_lhs46] = _rhs66;
          _590_v0 = _rhs67;
          _616_v21 = _rhs68;
          _619_v24 = _rhs69;
          let _621_v26;
          _621_v26 = _module.D6.create_DC17(_598_v8);
          let _pat_let_tv22 = _598_v8;
          let _pat_let_tv23 = _598_v8;
          _621_v26 = function (_pat_let23_0) {
            return function (_622_dt__update__tmp_h0) {
              return function (_pat_let24_0) {
                return function (_623_dt__update_hcf35_h0) {
                  return _module.D6.create_DC17(_623_dt__update_hcf35_h0);
                }(_pat_let24_0);
              }(_dafny.Seq.Concat(_pat_let_tv22, _pat_let_tv23));
            }(_pat_let23_0);
          }(_621_v26);
        } else {
          let _624_v27;
          _624_v27 = _module.D1.create_DC4((_595_v5).f6, p3, _dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("ahcbwma")), (_595_v5).f5, _module.__default.safeDivisionInt(new BigNumber((_600_v10).length), (_607_v17).f12));
          _624_v27 = _624_v27;
          let _625_v29;
          _625_v29 = _dafny.Seq.of(_591_v1, function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(626), new BigNumber(843))) {
              let _626_v28 = _compr_16;
              if (((new BigNumber(626)).isLessThanOrEqualTo(_626_v28)) && ((_626_v28).isLessThan(new BigNumber(843)))) {
                _coll16.add((_626_v28).plus((_595_v5).f5));
              }
            }
            return _coll16;
          }(), _591_v1);
          let _627_v31;
          let _nw93 = Array((new BigNumber(20)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = (_591_v1).Difference(_591_v1);
          _nw93[(_dafny.ONE).toNumber()] = _591_v1;
          _nw93[(new BigNumber(2)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(3)).toNumber()] = _module.__default.fm21(globalState);
          _nw93[(new BigNumber(4)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(5)).toNumber()] = (_625_v29)[_module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_625_v29).length))];
          _nw93[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements((_593_v3).f5)).Union(_591_v1);
          _nw93[(new BigNumber(7)).toNumber()] = (_591_v1).Difference(_dafny.Set.fromElements(new BigNumber(-219)));
          _nw93[(new BigNumber(8)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(9)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(10)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(11)).toNumber()] = function () {
            let _coll17 = new _dafny.Set();
            for (const _compr_17 of _dafny.IntegerRange(new BigNumber(-558), new BigNumber(-764))) {
              let _628_v30 = _compr_17;
              if (((new BigNumber(-558)).isLessThanOrEqualTo(_628_v30)) && ((_628_v30).isLessThan(new BigNumber(-764)))) {
                _coll17.add(_module.__default.safeDivisionInt(_628_v30, (_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))]));
              }
            }
            return _coll17;
          }();
          _nw93[(new BigNumber(12)).toNumber()] = (_591_v1).Union(_591_v1);
          _nw93[(new BigNumber(13)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(14)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(15)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(16)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(17)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(18)).toNumber()] = _591_v1;
          _nw93[(new BigNumber(19)).toNumber()] = _591_v1;
          _627_v31 = _nw93;
          let _index91 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_627_v31).length));
          (_627_v31)[_index91] = _591_v1;
          let _index92 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_627_v31).length));
          (_627_v31)[_index92] = (_591_v1).Intersect((_591_v1).Union(_591_v1));
          let _629_v32;
          let _init11 = ((_630_v3) => function (_631_i5) {
            return (_630_v3).f6;
          })(_593_v3);
          let _nw94 = Array((new BigNumber(5)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw94.length); _i0_11++) {
            _nw94[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _629_v32 = _nw94;
          let _index93 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_629_v32).length));
          (_629_v32)[_index93] = (_593_v3).f6;
          let _index94 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_629_v32).length));
          (_629_v32)[_index94] = !(false) || (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-432)), function (_632_i6) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), p1));
          let _index95 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_629_v32).length));
          let _index96 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _rhs70 = !(false) || ((_629_v32)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_629_v32).length))]);
          let _rhs71 = (_module.__default.safeModuloInt(p0, new BigNumber(672))).plus((_593_v3).f5);
          let _lhs47 = _629_v32;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_629_v32).length));
          let _lhs49 = _602_v12;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          _lhs47[_lhs48] = _rhs70;
          _lhs49[_lhs50] = _rhs71;
          let _633_v33;
          _633_v33 = _module.D3.create_DC11(_602_v12);
          _633_v33 = _633_v33;
        }
        let _634_v34;
        let _nw95 = Array((new BigNumber(9)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = (_593_v3).f6;
        _nw95[(_dafny.ONE).toNumber()] = (_593_v3).f6;
        _nw95[(new BigNumber(2)).toNumber()] = p3;
        _nw95[(new BigNumber(3)).toNumber()] = p3;
        _nw95[(new BigNumber(4)).toNumber()] = (_593_v3).f6;
        _nw95[(new BigNumber(5)).toNumber()] = (_593_v3).f6;
        _nw95[(new BigNumber(6)).toNumber()] = (_593_v3).f6;
        _nw95[(new BigNumber(7)).toNumber()] = false;
        _nw95[(new BigNumber(8)).toNumber()] = (_595_v5).f6;
        _634_v34 = _nw95;
        let _index97 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_634_v34).length));
        (_634_v34)[_index97] = (_595_v5).f6;
        let _index98 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_634_v34).length));
        (_634_v34)[_index98] = p3;
        let _index99 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_634_v34).length));
        (_634_v34)[_index99] = (_595_v5).f6;
        let _index100 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
        (_602_v12)[_index100] = (_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))];
      } else {
        let _635_v35;
        _635_v35 = _dafny.Map.Empty.slice().updateUnsafe(_596_v6,(_593_v3).f6);
        let _636_v36;
        let _nw96 = Array((new BigNumber(23)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = true;
        _nw96[(_dafny.ONE).toNumber()] = (_595_v5).f6;
        _nw96[(new BigNumber(2)).toNumber()] = (_593_v3).f6;
        _nw96[(new BigNumber(3)).toNumber()] = (_591_v1).IsDisjointFrom(_591_v1);
        _nw96[(new BigNumber(4)).toNumber()] = (_595_v5).f6;
        _nw96[(new BigNumber(5)).toNumber()] = true;
        _nw96[(new BigNumber(6)).toNumber()] = (_593_v3).f6;
        _nw96[(new BigNumber(7)).toNumber()] = (_595_v5).f6;
        _nw96[(new BigNumber(8)).toNumber()] = (_593_v3).f6;
        _nw96[(new BigNumber(9)).toNumber()] = p3;
        _nw96[(new BigNumber(10)).toNumber()] = (_595_v5).fm0((_593_v3).f6, globalState);
        _nw96[(new BigNumber(11)).toNumber()] = ((_this).f8).isLessThan(_590_v0);
        _nw96[(new BigNumber(12)).toNumber()] = ((p3) ? ((_593_v3).f6) : (p3));
        _nw96[(new BigNumber(13)).toNumber()] = (_595_v5).f6;
        _nw96[(new BigNumber(14)).toNumber()] = (_595_v5).f6;
        _nw96[(new BigNumber(15)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(16)).toNumber()] = (((_635_v35).contains(_dafny.Seq.of(p2, (_595_v5).f5, (_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))], (_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))]))) ? ((_635_v35).get(_dafny.Seq.of(p2, (_595_v5).f5, (_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))], (_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))]))) : ((_this).f6));
        _nw96[(new BigNumber(17)).toNumber()] = true;
        _nw96[(new BigNumber(18)).toNumber()] = (_593_v3).f6;
        _nw96[(new BigNumber(19)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(20)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(21)).toNumber()] = !((_595_v5).f6) || (p3);
        _nw96[(new BigNumber(22)).toNumber()] = !((_598_v8)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).fm0((_593_v3).f6, globalState), (_this).f6)).cardinality())), new BigNumber((_598_v8).length))]) || ((_593_v3).f6);
        _636_v36 = _nw96;
        let _index101 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
        (_636_v36)[_index101] = p3;
        let _index102 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
        (_636_v36)[_index102] = p3;
        let _637_v37;
        _637_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,_636_v36);
        _637_v37 = (_637_v37).update((_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))], _636_v36);
        _598_v8 = _module.__default.fm16(globalState);
        let _index103 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
        (_636_v36)[_index103] = (_595_v5).f6;
        let _638_v38;
        _638_v38 = _dafny.Seq.of(p1);
        let _639_v39;
        _639_v39 = _module.D1.create_DC4((_590_v0).isLessThan(p0), (_593_v3).fm0(p3, globalState), p1, ((_593_v3).f5).minus(new BigNumber(((_638_v38)[_module.__default.safeIndex((_dafny.ZERO).minus((_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))]), new BigNumber((_638_v38).length))]).length)), (_593_v3).f5);
        let _source8 = _639_v39;
        if (_source8.is_DC4) {
          let _640___mcc_h0 = (_source8).cf8;
          let _641___mcc_h1 = (_source8).cf9;
          let _642___mcc_h2 = (_source8).cf10;
          let _643___mcc_h3 = (_source8).cf11;
          let _644___mcc_h4 = (_source8).cf12;
          let _645_cf12 = _644___mcc_h4;
          let _646_cf11 = _643___mcc_h3;
          let _647_cf10 = _642___mcc_h2;
          let _648_cf9 = _641___mcc_h1;
          let _649_cf8 = _640___mcc_h0;
          let _650_v40;
          _650_v40 = _dafny.Map.Empty.slice().updateUnsafe(_648_cf9,(_593_v3).f5);
          let _651_v41;
          _651_v41 = _dafny.Map.Empty.slice().updateUnsafe(_601_v11,!((_dafny.ZERO).minus(_module.__default.fm6(new BigNumber((_650_v40).length), (_593_v3).f5, p3, new BigNumber((_dafny.Seq.UnicodeFromString("htbyp")).length), globalState))).isEqualTo((_593_v3).f5));
          _651_v41 = _651_v41;
          let _652_v42;
          _652_v42 = new _dafny.CodePoint('e'.codePointAt(0));
          _652_v42 = _652_v42;
          let _653_v43;
          let _nw97 = Array((new BigNumber(29)).toNumber());
          _653_v43 = _nw97;
          let _index104 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_653_v43).length));
          (_653_v43)[_index104] = _595_v5;
          let _index105 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_653_v43).length));
          (_653_v43)[_index105] = (((_605_v15).contains(p3)) ? ((_605_v15).get(p3)) : (_595_v5));
          _649_cf8 = (_593_v3).f6;
        } else if (_source8.is_DC5) {
          let _654___mcc_h5 = (_source8).cf13;
          let _655___mcc_h6 = (_source8).cf14;
          let _656___mcc_h7 = (_source8).cf15;
          let _657___mcc_h8 = (_source8).cf16;
          let _658_cf16 = _657___mcc_h8;
          let _659_cf15 = _656___mcc_h7;
          let _660_cf14 = _655___mcc_h6;
          let _661_cf13 = _654___mcc_h5;
          _659_cf15 = (_595_v5).f5;
          let _662_v44;
          _662_v44 = new _dafny.CodePoint('e'.codePointAt(0));
          let _rhs72 = _602_v12;
          let _rhs73 = (_module.D1.create_DC4((_595_v5).f6, (((_601_v11).contains(false)) ? ((_601_v11).get(false)) : (p3)), p1, p0, (_595_v5).f5)).dtor_cf8;
          let _rhs74 = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("vhrqm"), _662_v44);
          _602_v12 = _rhs72;
          _660_cf14 = _rhs73;
          _661_cf13 = _rhs74;
          _658_cf16 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), ((_663_v3) => function (_664_i7) {
            return (_663_v3).f5;
          })(_593_v3))).length);
          _660_cf14 = (_595_v5).f6;
        } else if (_source8.is_DC6) {
          let _665___mcc_h9 = (_source8).cf17;
          let _666___mcc_h10 = (_source8).cf18;
          let _667___mcc_h11 = (_source8).cf19;
          let _668___mcc_h12 = (_source8).cf20;
          let _669_cf20 = _668___mcc_h12;
          let _670_cf19 = _667___mcc_h11;
          let _671_cf18 = _666___mcc_h10;
          let _672_cf17 = _665___mcc_h9;
          let _673_v45;
          let _nw98 = new _module.C2();
          _nw98.__ctor((_593_v3).f5);
          _673_v45 = _nw98;
          let _674_v46;
          _674_v46 = _dafny.Map.Empty.slice().updateUnsafe(_602_v12,_673_v45);
          let _675_v47;
          _675_v47 = _dafny.Seq.of(_602_v12);
          let _676_v48;
          _676_v48 = _dafny.Map.Empty.slice().updateUnsafe((((_674_v46).contains((_675_v47)[_module.__default.safeIndex((_593_v3).f5, new BigNumber((_675_v47).length))])) ? ((_674_v46).get((_675_v47)[_module.__default.safeIndex((_593_v3).f5, new BigNumber((_675_v47).length))])) : (_673_v45)),_dafny.Set.fromElements((_593_v3).f5, p2));
          _676_v48 = _676_v48;
          _636_v36 = _636_v36;
          let _677_v49;
          _677_v49 = new _dafny.CodePoint('k'.codePointAt(0));
          _669_cf20 = !((_673_v45).fm1(p0, _677_v49, globalState));
          let _678_v50;
          let _nw99 = Array((new BigNumber(27)).toNumber()).fill([]);
          _678_v50 = _nw99;
          let _679_v52;
          _679_v52 = _dafny.Seq.of(_599_v9);
          let _680_v53;
          let _nw100 = Array((new BigNumber(15)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = _599_v9;
          _nw100[(_dafny.ONE).toNumber()] = _599_v9;
          _nw100[(new BigNumber(2)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(3)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(4)).toNumber()] = function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_599_v9).Keys.Elements) {
              let _681_v51 = _compr_18;
              if ((_599_v9).contains(_681_v51)) {
                _coll18.push([(_681_v51).multipliedBy(_670_cf19),_670_cf19]);
              }
            }
            return _coll18;
          }();
          _nw100[(new BigNumber(5)).toNumber()] = (_599_v9).update((_593_v3).f5, (_595_v5).f5);
          _nw100[(new BigNumber(6)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(7)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(8)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(9)).toNumber()] = (_679_v52)[_module.__default.safeIndex(new BigNumber(-264), new BigNumber((_679_v52).length))];
          _nw100[(new BigNumber(10)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(11)).toNumber()] = (_599_v9).update(new BigNumber(967), (_this).f8);
          _nw100[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(515),_590_v0);
          _nw100[(new BigNumber(13)).toNumber()] = _599_v9;
          _nw100[(new BigNumber(14)).toNumber()] = _599_v9;
          _680_v53 = _nw100;
          let _index106 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_678_v50).length));
          (_678_v50)[_index106] = _680_v53;
          let _682_v54;
          _682_v54 = _dafny.Seq.of(_680_v53, _680_v53, _680_v53, _680_v53);
          let _index107 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_678_v50).length));
          (_678_v50)[_index107] = (_682_v54)[_module.__default.safeIndex(new BigNumber(-573), new BigNumber((_682_v54).length))];
        } else if (_source8.is_DC3) {
          let _683___mcc_h13 = (_source8).cf7;
          let _684_cf7 = _683___mcc_h13;
          let _index108 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
          (_636_v36)[_index108] = (_595_v5).f6;
          let _index109 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _index110 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _rhs75 = (_595_v5).f5;
          let _rhs76 = (_593_v3).f5;
          let _lhs51 = _602_v12;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _lhs53 = _602_v12;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          _lhs51[_lhs52] = _rhs75;
          _lhs53[_lhs54] = _rhs76;
          let _685_v55;
          _685_v55 = _module.D3.create_DC11(_602_v12);
          let _686_v56;
          _686_v56 = _dafny.Map.Empty.slice().updateUnsafe(_685_v55,(_595_v5).f5);
          _686_v56 = (_686_v56).update((((_this).f6) ? (_685_v55) : (_module.D3.create_DC11(_602_v12))), (_593_v3).f5);
          let _687_v57;
          _687_v57 = _dafny.Map.Empty.slice().updateUnsafe((_595_v5).f5,p1);
          _687_v57 = (_687_v57).update(((_595_v5).f5).plus(p0), p1);
        } else {
          let _688___mcc_h14 = (_source8).cf21;
          let _689_cf21 = _688___mcc_h14;
          let _index111 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
          (_636_v36)[_index111] = (_593_v3).f6;
          _590_v0 = p2;
          let _690_v58;
          let _nw101 = Array((new BigNumber(9)).toNumber()).fill([]);
          _690_v58 = _nw101;
          let _691_v59;
          _691_v59 = _dafny.Seq.of(_636_v36, _636_v36);
          let _692_v60;
          _692_v60 = _dafny.Map.Empty.slice().updateUnsafe(true,(_603_v13).dtor_cf19);
          let _index112 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
          let _index113 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _index114 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _rhs77 = !((true) === (!_dafny.Seq.contains(_691_v59, _636_v36)));
          let _rhs78 = _690_v58;
          let _rhs79 = _module.__default.safeModuloInt(new BigNumber(((_604_v14).Difference(_604_v14)).cardinality()), new BigNumber((_692_v60).length));
          let _rhs80 = _module.__default.fm6(_590_v0, ((_593_v3).f5).plus(p2), true, new BigNumber(197), globalState);
          let _lhs55 = _636_v36;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
          let _lhs57 = _602_v12;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          let _lhs59 = _602_v12;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length));
          _lhs55[_lhs56] = _rhs77;
          _690_v58 = _rhs78;
          _lhs57[_lhs58] = _rhs79;
          _lhs59[_lhs60] = _rhs80;
          let _index115 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_636_v36).length));
          (_636_v36)[_index115] = true;
        }
      }
      let _693_v61;
      let _nw102 = new _module.C2();
      _nw102.__ctor(((_602_v12)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_602_v12).length))]).multipliedBy((_dafny.ZERO).minus(_590_v0)));
      _693_v61 = _nw102;
      let _694_i8;
      _694_i8 = _dafny.ZERO;
      L4: {
        while (!((_595_v5).f6)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_694_i8)) {
              break L4;
            }
            _694_i8 = (_694_i8).plus(_dafny.ONE);
            let _695_v62;
            let _nw103 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _695_v62 = _nw103;
            let _696_v63;
            _696_v63 = _dafny.Map.Empty.slice().updateUnsafe(((!(false)) ? (p1) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(441)), function (_697_i9) {
              return new _dafny.CodePoint('b'.codePointAt(0));
            }))),_695_v62);
            _696_v63 = (_696_v63).update(_dafny.Seq.Concat(p1, p1), _695_v62);
            let _698_v64;
            _698_v64 = _dafny.Map.Empty.slice().updateUnsafe((_593_v3).f6,(_593_v3).f5);
            let _699_v65;
            _699_v65 = _dafny.Seq.of(_596_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), ((_700_p2) => function (_701_i10) {
              return _700_p2;
            })(p2)));
            _698_v64 = (_698_v64).update(_dafny.Seq.IsPrefixOf((_699_v65)[_module.__default.safeIndex((_595_v5).f5, new BigNumber((_699_v65).length))], _596_v6), (_dafny.ZERO).minus((_595_v5).f5));
            let _702_v66;
            let _nw104 = new _module.C2();
            _nw104.__ctor(_module.__default.safeModuloInt(new BigNumber(585), (_dafny.ZERO).minus((_this).f8)));
            _702_v66 = _nw104;
            let _index116 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_695_v62).length));
            (_695_v62)[_index116] = p1;
            let _703_v67;
            _703_v67 = true;
            let _704_v68;
            _704_v68 = _dafny.Map.Empty.slice().updateUnsafe((_595_v5).f5,_698_v64);
            let _705_v69;
            _705_v69 = _dafny.Seq.of(_module.__default.fm5(_dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_dafny.Seq.of((_595_v5).f5)).length)), _704_v68, globalState), p1);
            let _index117 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_695_v62).length));
            let _rhs81 = _dafny.Seq.Concat(p1, (_705_v69)[_module.__default.safeIndex(new BigNumber((_598_v8).length), new BigNumber((_705_v69).length))]);
            let _rhs82 = (_this).f5;
            let _rhs83 = p3;
            let _rhs84 = !(((_593_v3).f5).isLessThanOrEqualTo((_593_v3).f5)) || ((_this).f6);
            let _rhs85 = !(_693_v61.f9).isEqualTo((_595_v5).f5);
            let _lhs61 = _695_v62;
            let _lhs62 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_695_v62).length));
            let _lhs63 = _693_v61;
            _lhs61[_lhs62] = _rhs81;
            _lhs63.f9 = _rhs82;
            _703_v67 = _rhs83;
            _703_v67 = _rhs84;
            _703_v67 = _rhs85;
          }
        }
      }
      let _706_v70;
      _706_v70 = _module.D3.create_DC11(_602_v12);
      let _707_v71;
      _707_v71 = _dafny.Seq.of(_706_v70);
      let _708_v72;
      _708_v72 = _dafny.MultiSet.fromElements(_706_v70);
      let _709_v73;
      let _nw105 = Array((new BigNumber(6)).toNumber());
      _nw105[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_707_v71, _707_v71));
      _nw105[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_706_v70)).Difference(_708_v72);
      _nw105[(new BigNumber(2)).toNumber()] = _708_v72;
      _nw105[(new BigNumber(3)).toNumber()] = _708_v72;
      _nw105[(new BigNumber(4)).toNumber()] = _708_v72;
      _nw105[(new BigNumber(5)).toNumber()] = (_708_v72).Difference(_708_v72);
      _709_v73 = _nw105;
      _709_v73 = (_709_v73);
      let _710_v74;
      let _init12 = ((_711_v9) => function (_712_i11) {
        return _711_v9;
      })(_599_v9);
      let _nw106 = Array((new BigNumber(2)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw106.length); _i0_12++) {
        _nw106[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _710_v74 = _nw106;
      _710_v74 = _710_v74;
      let _713_v75;
      let _nw107 = Array((new BigNumber(21)).toNumber()).fill([]);
      _713_v75 = _nw107;
      let _714_v76;
      _714_v76 = _dafny.Seq.of(_713_v75, _713_v75);
      r0 = (_714_v76)[_module.__default.safeIndex(_590_v0, new BigNumber((_714_v76).length))];
      let _715_v77;
      _715_v77 = _module.D0.create_DC0(_596_v6);
      r1 = _715_v77;
      return [r0, r1];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this.f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f7, f5, f6) {
      let _this = this;
      (_this).f7 = f7;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      let _source9 = _module.D0.create_DC0(_dafny.Seq.of(_this.f7));
      if (_source9.is_DC1) {
        let _716___mcc_h0 = (_source9).cf1;
        let _717___mcc_h1 = (_source9).cf2;
        let _718___mcc_h2 = (_source9).cf3;
        let _719___mcc_h3 = (_source9).cf4;
        let _720_cf4 = _719___mcc_h3;
        let _721_cf3 = _718___mcc_h2;
        let _722_cf2 = _717___mcc_h1;
        let _723_cf1 = _716___mcc_h0;
        return true;
      } else if (_source9.is_DC2) {
        let _724___mcc_h4 = (_source9).cf5;
        let _725___mcc_h5 = (_source9).cf6;
        let _726_cf6 = _725___mcc_h5;
        let _727_cf5 = _724___mcc_h4;
        return (_dafny.Set.fromElements((_this).f6)).contains((_this).f6);
      } else {
        let _728___mcc_h6 = (_source9).cf0;
        let _729_cf0 = _728___mcc_h6;
        return !((_this).f6);
      }
    };
    m1(p0, globalState) {
      let _this = this;
      let _730_v0;
      _730_v0 = _dafny.Seq.of((_this).f6, (_this).f6);
      let _731_v1;
      _731_v1 = _dafny.Seq.of(_this.f7, p0, p0, new BigNumber((_730_v0).length), p0);
      let _732_v2;
      _732_v2 = _module.D0.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(false,_731_v1), (_this).f6);
      let _733_v3;
      _733_v3 = _dafny.Set.fromElements(_this.f7, _this.f7);
      let _734_v4;
      _734_v4 = _module.D1.create_DC6((_this).f6, _this.f7, new BigNumber((_733_v3).length), false);
      let _735_v5;
      let _nw108 = new _module.C0();
      _nw108.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("fcehnr")).length), (_734_v4).dtor_cf19);
      _735_v5 = _nw108;
      let _736_v6;
      _736_v6 = _dafny.Map.Empty.slice().updateUnsafe(_735_v5,(_735_v5).f12);
      let _737_v7;
      _737_v7 = _dafny.MultiSet.fromElements(_736_v6, _736_v6);
      let _738_v8;
      _738_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-102),_this.f7);
      let _739_v9;
      let _nw109 = new _module.C3();
      _nw109.__ctor((_735_v5).f12, (_dafny.ZERO).minus(new BigNumber((_738_v8).length)), (_this).f6);
      _739_v9 = _nw109;
      let _740_v10;
      let _nw110 = new _module.C3();
      _nw110.__ctor(new BigNumber((_737_v7).cardinality()), new BigNumber((_dafny.Set.fromElements(_739_v9, _739_v9)).length), (_this).f6);
      _740_v10 = _nw110;
      let _741_v11;
      _741_v11 = _dafny.Set.fromElements(_740_v10);
      let _742_v12;
      _742_v12 = _dafny.Seq.UnicodeFromString("nhstj");
      let _743_v13;
      let _nw111 = Array((new BigNumber(17)).toNumber());
      _nw111[(_dafny.ZERO).toNumber()] = (_732_v2).dtor_cf6;
      _nw111[(_dafny.ONE).toNumber()] = (p0).isLessThan(new BigNumber(56));
      _nw111[(new BigNumber(2)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(3)).toNumber()] = _dafny.areEqual(_731_v1, _731_v1);
      _nw111[(new BigNumber(4)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(5)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(6)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(7)).toNumber()] = !_dafny.Seq.contains(_731_v1, _this.f7);
      _nw111[(new BigNumber(8)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(9)).toNumber()] = !((_this).f6);
      _nw111[(new BigNumber(10)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(11)).toNumber()] = (_this).f6;
      _nw111[(new BigNumber(12)).toNumber()] = (_741_v11).IsDisjointFrom(_741_v11);
      _nw111[(new BigNumber(13)).toNumber()] = _dafny.areEqual(_742_v12, _742_v12);
      _nw111[(new BigNumber(14)).toNumber()] = ((_this).f6) && ((_this).f6);
      _nw111[(new BigNumber(15)).toNumber()] = (_740_v10).f6;
      _nw111[(new BigNumber(16)).toNumber()] = (_this).f6;
      _743_v13 = _nw111;
      let _index118 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
      (_743_v13)[_index118] = (_740_v10).f6;
      let _index119 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
      (_743_v13)[_index119] = (_740_v10).f6;
      let _index120 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
      (_743_v13)[_index120] = (_740_v10).f6;
      if ((_this).f6) {
        let _744_v14;
        let _nw112 = Array((new BigNumber(17)).toNumber()).fill(_dafny.MultiSet.Empty);
        _744_v14 = _nw112;
        let _745_v15;
        _745_v15 = _744_v14;
        let _index121 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
        let _index122 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
        let _rhs86 = false;
        let _rhs87 = _745_v15;
        let _rhs88 = !((_740_v10).f6);
        let _rhs89 = _742_v12;
        let _rhs90 = (_739_v9).f8;
        let _lhs64 = _743_v13;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
        let _lhs66 = _743_v13;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
        let _lhs68 = _this;
        _lhs64[_lhs65] = _rhs86;
        _745_v15 = _rhs87;
        _lhs66[_lhs67] = _rhs88;
        _742_v12 = _rhs89;
        _lhs68.f7 = _rhs90;
        let _index123 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length));
        (_743_v13)[_index123] = (((_743_v13)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length))]) ? ((_743_v13)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length))]) : ((_740_v10).f6));
        let _746_v16;
        _746_v16 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jijr"));
        _746_v16 = (_746_v16).Difference(_746_v16);
        let _747_v17;
        let _748_v18;
        let _749_v19;
        let _750_v20;
        let _out40;
        let _out41;
        let _out42;
        let _out43;
        let _outcollector10 = _module.__default.m0(((_735_v5).f12).multipliedBy(_this.f7), new BigNumber(253), globalState);
        _out40 = _outcollector10[0];
        _out41 = _outcollector10[1];
        _out42 = _outcollector10[2];
        _out43 = _outcollector10[3];
        _747_v17 = _out40;
        _748_v18 = _out41;
        _749_v19 = _out42;
        _750_v20 = _out43;
        _749_v19 = _this.f7;
      } else {
        let _751_v21;
        _751_v21 = _module.D5.create_DC16(_740_v10, (_740_v10).f6, (_dafny.ZERO).minus((_dafny.ZERO).minus((_740_v10).f5)));
        let _752_v22;
        _752_v22 = _dafny.Map.Empty.slice().updateUnsafe((_751_v21).dtor_cf33,(_731_v1)[_module.__default.safeIndex(_this.f7, new BigNumber((_731_v1).length))]);
        let _753_v23;
        _753_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_752_v22).length),true);
        let _754_v24;
        let _nw113 = new _module.C3();
        _nw113.__ctor((_739_v9).f8, new BigNumber((_753_v23).length), false);
        _754_v24 = _nw113;
        let _755_v25;
        let _nw114 = new _module.C3();
        _nw114.__ctor((_739_v9).f8, (_this.f7).multipliedBy((((_752_v22).contains((_743_v13)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length))])) ? ((_752_v22).get((_743_v13)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length))])) : (new BigNumber((_dafny.MultiSet.fromElements((_743_v13)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length))])).cardinality())))), ((_754_v24).f8).isLessThan((_754_v24).f8));
        _755_v25 = _nw114;
        let _756_v26;
        _756_v26 = new _dafny.CodePoint('e'.codePointAt(0));
        _742_v12 = _dafny.Seq.Concat(_dafny.Seq.Concat(_742_v12, _742_v12), _dafny.Seq.update(_742_v12, _module.__default.safeIndex(_module.__default.fm6(new BigNumber(-546), (_735_v5).f12, (_740_v10).f6, (_735_v5).f13, globalState), new BigNumber((_742_v12).length)), _756_v26));
        let _757_v27;
        let _nw115 = new _module.C0();
        _nw115.__ctor((_755_v25).f8, (_740_v10).f5);
        _757_v27 = _nw115;
        (_this).f7 = _module.__default.fm6((_757_v27).f12, _this.f7, (_740_v10).f6, (_this).f5, globalState);
      }
      let _758_i0;
      _758_i0 = _dafny.ZERO;
      L5: {
        while ((_740_v10).f6) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_758_i0)) {
              break L5;
            }
            _758_i0 = (_758_i0).plus(_dafny.ONE);
            let _759_v28;
            let _init13 = ((_760_v0) => function (_761_i1) {
              return _module.D5.create_DC15(_dafny.MultiSet.FromArray(_760_v0));
            })(_730_v0);
            let _nw116 = Array((new BigNumber(29)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw116.length); _i0_13++) {
              _nw116[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _759_v28 = _nw116;
            let _762_v29;
            _762_v29 = _dafny.MultiSet.fromElements((_740_v10).f6);
            let _763_v30;
            _763_v30 = _module.D5.create_DC15((_762_v29).update((_this).f6, _module.__default.abs(new BigNumber((_742_v12).length))));
            let _pat_let_tv24 = _735_v5;
            let _pat_let_tv25 = _735_v5;
            let _pat_let_tv26 = globalState;
            let _index124 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_759_v28).length));
            (_759_v28)[_index124] = function (_pat_let25_0) {
              return function (_764_dt__update__tmp_h0) {
                return function (_pat_let26_0) {
                  return function (_765_dt__update_hcf31_h0) {
                    return _module.D5.create_DC15(_765_dt__update_hcf31_h0);
                  }(_pat_let26_0);
                }(_module.__default.fm10(_dafny.MultiSet.fromElements((_pat_let_tv24).f12, (_pat_let_tv25).f12, new BigNumber(-863)), (_this).f5, _pat_let_tv26));
              }(_pat_let25_0);
            }(_763_v30);
            let _766_v31;
            let _nw117 = new _module.C1();
            _nw117.__ctor(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)))).length), (_this).f5, (_739_v9).f8, false);
            _766_v31 = _nw117;
            let _767_v32;
            _767_v32 = _dafny.Map.Empty.slice().updateUnsafe((_740_v10).f6,_766_v31);
            let _768_v33;
            _768_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f7);
            let _index125 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_759_v28).length));
            (_759_v28)[_index125] = _module.__default.fm22(new BigNumber(((_767_v32).Merge(_767_v32)).length), _742_v12, ((_766_v31).f11).minus((_module.D0.create_DC1(_742_v12, (_739_v9).f8, _768_v33, _731_v1)).dtor_cf2), _dafny.Seq.Create(_module.__default.abs(new BigNumber(126)), function (_769_i2) {
              return new _dafny.CodePoint('o'.codePointAt(0));
            }), globalState);
            let _770_v34;
            _770_v34 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_735_v5).f12), new BigNumber((_dafny.Set.fromElements((_this).f5)).length));
            let _771_v35;
            _771_v35 = _module.D2.create_DC8(_770_v34);
            _771_v35 = _module.__default.fm23(_this.f7, (new BigNumber((_742_v12).length)).isLessThan(new BigNumber((_731_v1).length)), new BigNumber(237), ((_735_v5).f12).plus(new BigNumber((_742_v12).length)), globalState);
            let _772_v36;
            _772_v36 = _dafny.Map.Empty.slice().updateUnsafe((_739_v9).f8,!_dafny.areEqual(_742_v12, _742_v12));
            let _773_v37;
            _773_v37 = _module.D5.create_DC16(_740_v10, (_this).f6, (_735_v5).f13);
            _772_v36 = (_772_v36).update((_773_v37).dtor_cf34, (_this).f6);
            let _rhs91 = (_dafny.ZERO).minus((_735_v5).f12);
            let _rhs92 = (_dafny.Map.Empty.slice().updateUnsafe((_740_v10).f5,(_740_v10).f6)).update(((_740_v10).f5).multipliedBy((_735_v5).f13), (_743_v13)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_743_v13).length))]);
            let _lhs69 = _this;
            _lhs69.f7 = _rhs91;
            _772_v36 = _rhs92;
          }
        }
      }
      _742_v12 = _742_v12;
      let _774_v38;
      let _nw118 = Array((new BigNumber(5)).toNumber());
      _nw118[(_dafny.ZERO).toNumber()] = (_this).f5;
      _nw118[(_dafny.ONE).toNumber()] = _this.f7;
      _nw118[(new BigNumber(2)).toNumber()] = (_735_v5).f12;
      _nw118[(new BigNumber(3)).toNumber()] = (_740_v10).f5;
      _nw118[(new BigNumber(4)).toNumber()] = (_this).f5;
      _774_v38 = _nw118;
      let _775_v39;
      _775_v39 = _dafny.MultiSet.fromElements(_774_v38);
      let _776_v40;
      let _nw119 = Array((new BigNumber(7)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = _775_v39;
      _nw119[(_dafny.ONE).toNumber()] = (_775_v39).Intersect(_775_v39);
      _nw119[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_774_v38, _774_v38, _774_v38, _774_v38, _774_v38);
      _nw119[(new BigNumber(3)).toNumber()] = (_775_v39).Union(_dafny.MultiSet.fromElements(_774_v38));
      _nw119[(new BigNumber(4)).toNumber()] = _775_v39;
      _nw119[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_774_v38, _774_v38);
      _nw119[(new BigNumber(6)).toNumber()] = _775_v39;
      _776_v40 = _nw119;
      let _index126 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_776_v40).length));
      (_776_v40)[_index126] = _775_v39;
      let _index127 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_776_v40).length));
      (_776_v40)[_index127] = _775_v39;
      return;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
